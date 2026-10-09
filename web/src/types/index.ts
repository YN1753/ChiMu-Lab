export type EntryType =
  | 'life'
  | 'thought'
  | 'photo'
  | 'moment'
  | 'place'
  | 'music'
  | 'book'
  | 'game'
  | 'project'
  | 'purchase'
  | 'coffee'

export type ArchiveCategory =
  | 'all'
  | 'daily'
  | 'thought'
  | 'project'
  | 'collection'
  | 'photo'
  | 'transaction'

export const getEntryCategory = (entry: LifeEntry): ArchiveCategory => {
  const t = (entry.type || '').toLowerCase()
  if (t === 'transaction' || entry.is_transaction) {
    return 'transaction'
  }
  if (t === 'thought' || t === 'idea' || t === 'note' || t === 'essay') {
    return 'thought'
  }
  if (t === 'project' || t === 'code' || t === 'craft') {
    return 'project'
  }
  if (t === 'music' || t === 'book' || t === 'game' || t === 'purchase' || t === 'gear' || t === 'movie' || t === 'device') {
    return 'collection'
  }
  if (t === 'photo') {
    return 'photo'
  }
  return 'daily'
}

export const entryMatchesCategory = (entry: LifeEntry, category: ArchiveCategory): boolean => {
  if (category === 'all') return true
  const cat = getEntryCategory(entry)
  if (category === 'transaction') {
    return cat === 'transaction' || entry.type === 'transaction' || Boolean(entry.is_transaction)
  }
  if (category === 'photo') {
    return cat === 'photo' || Boolean(entry.images) || Boolean(entry.attachments && entry.attachments.length > 0)
  }
  if (category === 'daily') {
    return cat === 'daily' || cat === 'photo'
  }
  return cat === category
}

export const getCategoryDisplayName = (cat: ArchiveCategory): string => {
  const map: Record<ArchiveCategory, string> = {
    all: '全部记录',
    daily: '日常',
    thought: '想法',
    project: '项目',
    collection: '收藏',
    photo: '照片',
    transaction: '记账',
  }
  return map[cat] || '全部记录'
}

export interface Attachment {
  id: number
  entry_id?: number
  object_key: string
  original_filename: string
  mime_type: string
  size: number
  width?: number
  height?: number
  url: string
  created_at?: string
}

export interface LifeEntry {
  id: number
  type: EntryType | string
  title: string
  content: string
  occurred_at?: string
  location?: string
  project_id?: number
  project?: Project
  attachments?: Attachment[]

  // 兼容老版本
  date?: string
  year?: string
  month?: string
  day?: string
  time?: string
  images?: string
  meta?: string
  tags?: string
  link?: string
  related_project?: string
  featured?: boolean

  // 记账支持
  amount?: number
  tx_type?: 'income' | 'expense'
  category?: string
  payment_method?: string
  is_transaction?: boolean

  created_at?: string
  updated_at?: string
}

export interface Transaction {
  id: number
  amount: number // 单位：分 (例如 1250 表示 12.50 元)
  type: 'income' | 'expense'
  category: string // 餐饮, 交通, 购物, 娱乐, 学习, 住房, 医疗, 其他
  title: string
  note?: string
  payment_method?: string
  occurred_at: string
  created_at?: string
}

export interface TransactionSummary {
  year: number
  month: number
  total_expense: number // 分
  total_income: number  // 分
  balance: number       // 分
  category_breakdown: Record<string, number>
  count: number
}

export interface StorageStatus {
  provider: string
  configured: boolean
  bucket: string
  public_domain: string
  endpoint: string
}

export interface GlobalStats {
  total_entries: number
  total_photos: number
  total_projects: number
  month_expense: number
  current_year: number
  current_month: number
}

export interface NowStatus {
  id?: number
  building: string
  learning: string
  playing: string
  listening: string
  reading: string
  thinking: string
  using: string
  location: string
  last_updated: string
}

export interface Project {
  id: number
  name?: string
  title: string
  subtitle?: string
  description: string
  content?: string
  story?: string
  category: string
  tags: string
  demo_url?: string
  github_url: string
  status: string
  featured?: boolean
  order?: number
  started_at?: string
}

export interface YearArchiveStat {
  year: string
  total_count: number
  type_counts: Record<string, number>
  months: string[]
}

export interface SiteConfig {
  site_name: string
  site_desc: string
  domain: string
  icp_number: string
  icp_link: string
}
