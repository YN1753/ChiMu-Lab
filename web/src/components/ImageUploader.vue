<script setup lang="ts">
import { ref, onUnmounted } from 'vue'
import { Plus, X, RotateCcw, AlertCircle } from 'lucide-vue-next'
import { presignUpload, uploadDirectToR2, completeUpload, cleanupUploads } from '../utils/api'
import type { Attachment } from '../types'

const props = defineProps<{
  initialAttachments?: Attachment[]
  storageConfigured: boolean
}>()

const emit = defineEmits<{
  (e: 'update:attachments', ids: number[]): void
}>()

interface UploadItem {
  id: string
  file?: File
  name: string
  size: number
  previewUrl: string
  progress: number
  status: 'idle' | 'uploading' | 'completed' | 'error'
  errorMsg?: string
  attachmentId?: number
  objectKey?: string
}

const items = ref<UploadItem[]>(
  (props.initialAttachments || []).map((att) => ({
    id: `existing-${att.id}`,
    name: att.original_filename || 'image',
    size: att.size || 0,
    previewUrl: att.url,
    progress: 100,
    status: 'completed',
    attachmentId: att.id,
    objectKey: att.object_key,
  }))
)

const fileInputRef = ref<HTMLInputElement | null>(null)
const isDragging = ref(false)

const triggerSelect = () => {
  fileInputRef.value?.click()
}

const notifyChange = () => {
  const ids = items.value
    .filter((it) => it.status === 'completed' && it.attachmentId)
    .map((it) => it.attachmentId!)
  emit('update:attachments', ids)
}

const handleFileProcess = async (file: File) => {
  // 单文件 20MB 限制
  if (file.size > 20 * 1024 * 1024) {
    alert(`文件 ${file.name} 超过 20MB 限制`)
    return
  }

  const previewUrl = URL.createObjectURL(file)
  const item: UploadItem = {
    id: `${Date.now()}-${Math.random().toString(36).substring(2, 7)}`,
    file,
    name: file.name,
    size: file.size,
    previewUrl,
    progress: 0,
    status: 'idle',
  }

  items.value.push(item)
  await startUpload(item)
}

const startUpload = async (item: UploadItem) => {
  if (!item.file) return
  if (!props.storageConfigured) {
    item.status = 'error'
    item.errorMsg = 'R2 存储未配置'
    return
  }

  try {
    item.status = 'uploading'
    item.progress = 5
    item.errorMsg = undefined

    // 1. 请求预签名 PUT URL
    const presign = await presignUpload(item.file.name, item.file.type || 'image/jpeg', item.file.size)
    item.objectKey = presign.object_key
    item.progress = 20

    // 2. 直传 R2
    await uploadDirectToR2(presign.upload_url, item.file, (pct) => {
      // 映射进度 20% -> 90%
      item.progress = Math.min(90, Math.max(20, Math.round(20 + pct * 0.7)))
    })

    item.progress = 95

    // 3. 确认上传完成，落库 Attachment
    const att = await completeUpload({
      upload_id: presign.upload_id,
      object_key: presign.object_key,
      original_filename: item.file.name,
      mime_type: item.file.type || 'image/jpeg',
      size: item.file.size,
    })

    item.attachmentId = att.id
    item.status = 'completed'
    item.progress = 100
    notifyChange()
  } catch (err: any) {
    item.status = 'error'
    item.errorMsg = err.message || '上传失败'
  }
}

const onFilesSelected = (e: Event) => {
  const target = e.target as HTMLInputElement
  if (!target.files) return
  Array.from(target.files).forEach((f) => handleFileProcess(f))
  target.value = ''
}

const onDrop = (e: DragEvent) => {
  isDragging.value = false
  if (!e.dataTransfer?.files) return
  Array.from(e.dataTransfer.files).forEach((f) => handleFileProcess(f))
}

const removeItem = (item: UploadItem) => {
  if (item.objectKey && !item.id.startsWith('existing-')) {
    cleanupUploads([item.objectKey])
  }
  if (item.previewUrl.startsWith('blob:')) {
    URL.revokeObjectURL(item.previewUrl)
  }
  items.value = items.value.filter((it) => it.id !== item.id)
  notifyChange()
}

onUnmounted(() => {
  items.value.forEach((it) => {
    if (it.previewUrl.startsWith('blob:')) {
      URL.revokeObjectURL(it.previewUrl)
    }
  })
})
</script>

<template>
  <div class="space-y-3 select-none">
    <div class="flex items-center justify-between text-xs font-mono-archive text-[var(--ink-muted)]">
      <span>照片 / 附件</span>
      <span v-if="!storageConfigured" class="text-[var(--accent-warm)] flex items-center gap-1 text-[11px]">
        <AlertCircle class="w-3 h-3" />
        <span>R2 存储未配置，媒体直传暂不可用</span>
      </span>
      <span v-else class="text-[11px]">
        单张图片最大 20MB
      </span>
    </div>

    <!-- 缩略图列表与添加区域 -->
    <div
      class="border border-dashed rounded-xs p-4 transition-colors duration-150"
      :class="[
        isDragging ? 'border-[var(--ink-primary)] bg-[var(--bg-subtle)]' : 'border-[var(--border-subtle)] hover:border-[var(--border-divider)]',
        !storageConfigured ? 'opacity-85' : ''
      ]"
      @dragover.prevent="isDragging = true"
      @dragleave.prevent="isDragging = false"
      @drop.prevent="onDrop"
    >
      <!-- 隐藏的文件选择器 -->
      <input
        ref="fileInputRef"
        type="file"
        multiple
        accept="image/jpeg,image/png,image/webp,image/gif"
        class="hidden"
        @change="onFilesSelected"
      />

      <div class="flex flex-wrap gap-3 items-center">
        <!-- 已选/上传中的缩略图 -->
        <div
          v-for="item in items"
          :key="item.id"
          class="relative w-24 h-24 rounded-xs overflow-hidden border border-[var(--border-subtle)] group bg-[var(--bg-subtle)] shrink-0"
        >
          <img
            :src="item.previewUrl"
            :alt="item.name"
            class="w-full h-full object-cover"
          />

          <!-- 进度遮罩 -->
          <div
            v-if="item.status === 'uploading'"
            class="absolute inset-0 bg-black/50 flex flex-col items-center justify-center text-white text-[10px] font-mono-archive"
          >
            <span>{{ item.progress }}%</span>
            <div class="w-16 h-1 bg-white/30 rounded-full mt-1 overflow-hidden">
              <div
                class="h-full bg-white transition-all duration-150"
                :style="{ width: `${item.progress}%` }"
              />
            </div>
          </div>

          <!-- 失败状态 -->
          <div
            v-if="item.status === 'error'"
            class="absolute inset-0 bg-black/60 flex flex-col items-center justify-center text-white p-1 text-center"
          >
            <span class="text-[10px] text-red-300 truncate max-w-full">
              {{ item.errorMsg || '上传失败' }}
            </span>
            <button
              v-if="storageConfigured"
              type="button"
              @click.stop="startUpload(item)"
              class="mt-1 p-1 hover:bg-white/20 rounded-full cursor-pointer"
              title="重试"
            >
              <RotateCcw class="w-3.5 h-3.5 text-white" />
            </button>
          </div>

          <!-- 删除按钮 -->
          <button
            type="button"
            @click.stop="removeItem(item)"
            class="absolute top-1 right-1 w-5 h-5 bg-black/60 hover:bg-black/80 text-white rounded-full flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer"
            title="删除"
          >
            <X class="w-3 h-3" />
          </button>
        </div>

        <!-- 添加按钮 -->
        <button
          type="button"
          @click="triggerSelect"
          :disabled="!storageConfigured"
          class="w-24 h-24 border border-dashed border-[var(--border-subtle)] hover:border-[var(--ink-secondary)] rounded-xs flex flex-col items-center justify-center gap-1 text-[var(--ink-muted)] hover:text-[var(--ink-primary)] transition-colors cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed shrink-0"
        >
          <Plus class="w-4 h-4" />
          <span class="text-[10px] font-mono-archive">添加照片</span>
        </button>

        <div v-if="items.length === 0" class="text-xs text-[var(--ink-muted)] font-serif-editorial ml-2">
          点击或拖拽照片到此处
        </div>
      </div>
    </div>
  </div>
</template>
