import type {
  LifeEntry,
  Attachment,
  Transaction,
  TransactionSummary,
  StorageStatus,
  GlobalStats,
  Project,
} from '../types'

const API_BASE = '/api/v1'
const ADMIN_TOKEN_KEY = 'chimu_admin_token'

export function getAdminToken(): string {
  try {
    return localStorage.getItem(ADMIN_TOKEN_KEY) || ''
  } catch {
    return ''
  }
}

export function setAdminToken(token: string): void {
  try {
    localStorage.setItem(ADMIN_TOKEN_KEY, token.trim())
  } catch {}
}

export function removeAdminToken(): void {
  try {
    localStorage.removeItem(ADMIN_TOKEN_KEY)
  } catch {}
}

// 通用请求处理
async function request<T>(url: string, options?: RequestInit): Promise<T> {
  const token = getAdminToken()
  const authHeaders: Record<string, string> = {}
  if (token) {
    authHeaders['Authorization'] = `Bearer ${token}`
  }

  const res = await fetch(url, {
    headers: {
      'Content-Type': 'application/json',
      ...authHeaders,
      ...options?.headers,
    },
    ...options,
  })

  if (!res.ok) {
    let msg = `HTTP error ${res.status}`
    try {
      const errJson = await res.json()
      if (errJson.message) msg = errJson.message
    } catch {
      // ignore
    }
    throw new Error(msg)
  }

  const json = await res.json()
  // 兼容标准结构 { code: 200, message: "success", data: ... }
  if (json && typeof json === 'object' && 'code' in json && 'data' in json) {
    return json.data as T
  }
  return json as T
}

// 1. Life Entries
export async function fetchEntries(params?: {
  page?: number
  page_size?: number
  type?: string
  category?: string
  year?: string
}): Promise<LifeEntry[]> {
  const query = new URLSearchParams()
  if (params?.page) query.set('page', String(params.page))
  if (params?.page_size) query.set('page_size', String(params.page_size))
  if (params?.type && params.type !== 'all') query.set('type', params.type)
  if (params?.category && params.category !== 'all') query.set('category', params.category)
  if (params?.year && params.year !== 'all') query.set('year', params.year)

  const url = `${API_BASE}/entries?${query.toString()}`
  try {
    const data = await request<{ entries: LifeEntry[]; total: number }>(url)
    return data.entries || []
  } catch {
    // 降级兼容 /api/entries
    return request<LifeEntry[]>(`/api/entries?${query.toString()}`)
  }
}

export async function fetchEntryById(id: number): Promise<LifeEntry> {
  return request<LifeEntry>(`${API_BASE}/entries/${id}`)
}

export async function createLifeEntry(data: {
  type: string
  title?: string
  content: string
  occurred_at?: string
  location?: string
  project_id?: number | null
  attachment_ids?: number[]
}): Promise<LifeEntry> {
  return request<LifeEntry>(`${API_BASE}/entries`, {
    method: 'POST',
    body: JSON.stringify(data),
  })
}

export async function updateLifeEntry(
  id: number,
  data: {
    type?: string
    title?: string
    content?: string
    occurred_at?: string
    location?: string
    project_id?: number | null
    attachment_ids?: number[]
  }
): Promise<LifeEntry> {
  return request<LifeEntry>(`${API_BASE}/entries/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
  })
}

export async function deleteLifeEntry(id: number): Promise<boolean> {
  await request(`${API_BASE}/entries/${id}`, {
    method: 'DELETE',
  })
  return true
}

// 2. R2 Presigned Upload
export interface PresignResult {
  upload_id: string
  object_key: string
  upload_url: string
  expires_in: number
}

export async function presignUpload(
  filename: string,
  contentType: string,
  size: number
): Promise<PresignResult> {
  return request<PresignResult>(`${API_BASE}/uploads/presign`, {
    method: 'POST',
    body: JSON.stringify({
      filename,
      content_type: contentType,
      size,
    }),
  })
}

// 客户端使用 XMLHttpRequest 直传 R2 预签名 URL，带进度回调
export function uploadDirectToR2(
  uploadUrl: string,
  file: File,
  onProgress?: (percent: number) => void
): Promise<void> {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()
    xhr.open('PUT', uploadUrl, true)
    xhr.setRequestHeader('Content-Type', file.type || 'application/octet-stream')

    if (xhr.upload && onProgress) {
      xhr.upload.onprogress = (e) => {
        if (e.lengthComputable) {
          const percent = Math.round((e.loaded / e.total) * 100)
          onProgress(percent)
        }
      }
    }

    xhr.onload = () => {
      if (xhr.status >= 200 && xhr.status < 300) {
        resolve()
      } else {
        reject(new Error(`上传至存储桶失败 (HTTP ${xhr.status})`))
      }
    }

    xhr.onerror = () => {
      reject(new Error('网络中断或跨域限制，直传 R2 失败'))
    }

    xhr.send(file)
  })
}

export async function completeUpload(data: {
  upload_id: string
  object_key: string
  original_filename: string
  mime_type: string
  size: number
  width?: number
  height?: number
}): Promise<Attachment> {
  return request<Attachment>(`${API_BASE}/uploads/complete`, {
    method: 'POST',
    body: JSON.stringify(data),
  })
}

export async function cleanupUploads(objectKeys: string[]): Promise<void> {
  if (objectKeys.length === 0) return
  try {
    await request(`${API_BASE}/uploads/cleanup`, {
      method: 'POST',
      body: JSON.stringify({ object_keys: objectKeys }),
    })
  } catch {
    // 忽略静默清理失败
  }
}

export async function fetchStorageStatus(): Promise<StorageStatus> {
  return request<StorageStatus>(`${API_BASE}/storage/status`)
}

// 3. Transactions (记账)
export async function fetchTransactions(params?: {
  year?: number
  month?: number
}): Promise<Transaction[]> {
  const query = new URLSearchParams()
  if (params?.year) query.set('year', String(params.year))
  if (params?.month) query.set('month', String(params.month))
  return request<Transaction[]>(`${API_BASE}/transactions?${query.toString()}`)
}

export async function fetchTransactionSummary(params?: {
  year?: number
  month?: number
}): Promise<TransactionSummary> {
  const query = new URLSearchParams()
  if (params?.year) query.set('year', String(params.year))
  if (params?.month) query.set('month', String(params.month))
  return request<TransactionSummary>(`${API_BASE}/transactions/summary?${query.toString()}`)
}

export async function createTransaction(data: {
  amount: number
  type: 'income' | 'expense'
  category: string
  title: string
  note?: string
  payment_method?: string
  occurred_at?: string
}): Promise<Transaction> {
  return request<Transaction>(`${API_BASE}/transactions`, {
    method: 'POST',
    body: JSON.stringify(data),
  })
}

export async function updateTransaction(
  id: number,
  data: Partial<Transaction>
): Promise<Transaction> {
  return request<Transaction>(`${API_BASE}/transactions/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
  })
}

export async function deleteTransaction(id: number): Promise<boolean> {
  await request(`${API_BASE}/transactions/${id}`, {
    method: 'DELETE',
  })
  return true
}

// 4. Projects
export async function fetchProjects(): Promise<Project[]> {
  try {
    return await request<Project[]>(`${API_BASE}/projects`)
  } catch {
    return request<Project[]>('/api/projects')
  }
}

export async function createProject(data: {
  name: string
  description?: string
  content?: string
  category?: string
  github_url?: string
}): Promise<Project> {
  return request<Project>(`${API_BASE}/projects`, {
    method: 'POST',
    body: JSON.stringify(data),
  })
}

// 5. Global Stats
export async function fetchGlobalStats(): Promise<GlobalStats> {
  return request<GlobalStats>(`${API_BASE}/stats`)
}
