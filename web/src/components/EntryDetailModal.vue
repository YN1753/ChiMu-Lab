<script setup lang="ts">
import { ref, computed } from 'vue'
import { X, Edit2, Trash2, ArrowUpRight, Check, AlertTriangle } from 'lucide-vue-next'
import ImageUploader from './ImageUploader.vue'
import { updateLifeEntry, deleteLifeEntry } from '../utils/api'
import type { LifeEntry, Project } from '../types'

const props = defineProps<{
  entry: LifeEntry | null
  storageConfigured: boolean
  projects?: Project[]
}>()

const emit = defineEmits<{
  (e: 'close'): void
  (e: 'updated', entry: LifeEntry): void
  (e: 'deleted', id: number): void
}>()

const isEditing = ref(false)
const showDeleteConfirm = ref(false)
const isSubmitting = ref(false)
const errorMsg = ref<string | null>(null)

// 编辑表单字段
const editTitle = ref('')
const editContent = ref('')
const editLocation = ref('')
const editOccurredAt = ref('')
const editProjectId = ref<number | null>(null)
const editAttachmentIds = ref<number[]>([])

const startEdit = () => {
  if (!props.entry) return
  isEditing.value = true
  editTitle.value = props.entry.title || ''
  editContent.value = props.entry.content || ''
  editLocation.value = props.entry.location || ''
  editProjectId.value = props.entry.project_id || null
  editAttachmentIds.value = (props.entry.attachments || []).map((a) => a.id)

  if (props.entry.occurred_at) {
    const d = new Date(props.entry.occurred_at)
    if (!isNaN(d.getTime())) {
      const pad = (n: number) => String(n).padStart(2, '0')
      editOccurredAt.value = `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())}T${pad(d.getHours())}:${pad(d.getMinutes())}`
    }
  } else if (props.entry.date) {
    editOccurredAt.value = `${props.entry.date.replace(/\./g, '-')}T${props.entry.time || '12:00'}`
  }
}

const cancelEdit = () => {
  isEditing.value = false
  errorMsg.value = null
}

const handleSaveEdit = async () => {
  if (!props.entry) return
  if (!editContent.value.trim()) {
    errorMsg.value = '记录内容不能为空'
    return
  }

  try {
    isSubmitting.value = true
    const updated = await updateLifeEntry(props.entry.id, {
      type: props.entry.type,
      title: editTitle.value.trim(),
      content: editContent.value.trim(),
      location: editLocation.value.trim(),
      project_id: editProjectId.value,
      occurred_at: editOccurredAt.value,
      attachment_ids: editAttachmentIds.value,
    })

    emit('updated', updated)
    isEditing.value = false
    errorMsg.value = null
  } catch (err: any) {
    errorMsg.value = err.message || '更新失败'
  } finally {
    isSubmitting.value = false
  }
}

const handleDelete = async () => {
  if (!props.entry) return
  try {
    isSubmitting.value = true
    await deleteLifeEntry(props.entry.id)
    emit('deleted', props.entry.id)
    emit('close')
  } catch (err: any) {
    alert(err.message || '删除失败')
  } finally {
    isSubmitting.value = false
    showDeleteConfirm.value = false
  }
}

const displayImages = computed(() => {
  if (!props.entry) return []
  if (props.entry.attachments && props.entry.attachments.length > 0) {
    return props.entry.attachments.map((a) => a.url)
  }
  if (props.entry.images) {
    return props.entry.images.split(',').filter(Boolean)
  }
  return []
})

const getKindLabel = (type: string) => {
  const map: Record<string, string> = {
    life: '生活',
    daily: '日常',
    photo: '照片',
    thought: '想法',
    project: '项目',
    collection: '收藏',
    coffee: '咖啡',
    music: '听音',
    book: '书摘',
    place: '行迹',
    game: '游戏',
    purchase: '好物',
  }
  return map[type] || '记录'
}
</script>

<template>
  <div
    v-if="entry"
    class="fixed inset-0 z-50 flex items-center justify-center p-4 sm:p-6 bg-black/45 backdrop-blur-xs animate-in fade-in duration-200"
    @click.self="emit('close')"
  >
    <div
      class="bg-[var(--bg-archive)] border border-[var(--border-subtle)] shadow-[0_20px_60px_rgba(0,0,0,0.15)] w-full max-w-2xl max-h-[92vh] overflow-y-auto rounded-xs p-6 sm:p-10 text-left relative transition-all"
    >
      <!-- 顶部关闭与操作工具栏 -->
      <div class="flex items-center justify-between pb-6 border-b border-[var(--border-subtle)] mb-8">
        <div class="flex items-center gap-2 font-mono-archive text-xs text-[var(--ink-muted)]">
          <span class="text-[var(--ink-primary)] font-medium">
            {{ entry.year || '2026' }}.{{ entry.month || '10' }}.{{ entry.day || '09' }}
          </span>
          <span v-if="entry.time">· {{ entry.time }}</span>
          <span>·</span>
          <span class="text-[var(--ink-secondary)]">{{ getKindLabel(entry.type) }}</span>
        </div>

        <div class="flex items-center gap-3">
          <template v-if="!isEditing">
            <button
              @click="startEdit"
              class="text-xs font-mono-archive text-[var(--ink-muted)] hover:text-[var(--ink-primary)] transition-colors cursor-pointer flex items-center gap-1 py-1 px-2 border border-transparent hover:border-[var(--border-subtle)] rounded-xs"
              title="编辑记录"
            >
              <Edit2 class="w-3.5 h-3.5" />
              <span>编辑</span>
            </button>
            <button
              @click="showDeleteConfirm = true"
              class="text-xs font-mono-archive text-red-600/70 hover:text-red-700 transition-colors cursor-pointer flex items-center gap-1 py-1 px-2 border border-transparent hover:border-red-200 rounded-xs"
              title="删除记录"
            >
              <Trash2 class="w-3.5 h-3.5" />
              <span>删除</span>
            </button>
          </template>

          <button
            @click="emit('close')"
            class="p-1 text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer transition-colors"
          >
            <X class="w-5 h-5" />
          </button>
        </div>
      </div>

      <!-- 删除确认浮层 -->
      <div
        v-if="showDeleteConfirm"
        class="mb-6 p-4 bg-red-50/70 border border-red-200 rounded-xs space-y-3 font-serif-editorial text-sm"
      >
        <div class="flex items-center gap-2 text-red-800 font-medium">
          <AlertTriangle class="w-4 h-4 text-red-600" />
          <span>确定删除这条记录？</span>
        </div>
        <p class="text-xs text-red-700 leading-relaxed font-mono-archive">
          删除后将同时清理关联的照片与存储对象，该操作不可撤销。
        </p>
        <div class="flex items-center justify-end gap-3 pt-1">
          <button
            @click="showDeleteConfirm = false"
            class="px-3 py-1 text-xs font-mono-archive text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer"
          >
            取消
          </button>
          <button
            @click="handleDelete"
            :disabled="isSubmitting"
            class="px-3.5 py-1 text-xs font-mono-archive bg-red-600 text-white hover:bg-red-700 rounded-xs transition-colors cursor-pointer disabled:opacity-50"
          >
            {{ isSubmitting ? '正在删除...' : '确认删除' }}
          </button>
        </div>
      </div>

      <!-- 编辑模式 -->
      <form v-if="isEditing" @submit.prevent="handleSaveEdit" class="space-y-6">
        <div v-if="errorMsg" class="p-3 bg-red-50 text-red-700 text-xs rounded-xs border border-red-200">
          {{ errorMsg }}
        </div>

        <div class="space-y-1.5">
          <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">标题</label>
          <input
            v-model="editTitle"
            type="text"
            class="w-full px-3 py-2 text-base font-serif-editorial bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
          />
        </div>

        <div class="space-y-1.5">
          <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">正文内容 *</label>
          <textarea
            v-model="editContent"
            rows="6"
            class="w-full px-3 py-2 text-base font-serif-editorial leading-relaxed bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none resize-y"
          ></textarea>
        </div>

        <!-- 附件管理 -->
        <ImageUploader
          :initialAttachments="entry.attachments"
          :storageConfigured="storageConfigured"
          @update:attachments="editAttachmentIds = $event"
        />

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="space-y-1.5">
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">发生时间</label>
            <input
              v-model="editOccurredAt"
              type="datetime-local"
              class="w-full px-3 py-2 text-xs font-mono-archive bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
            />
          </div>

          <div class="space-y-1.5">
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">地点</label>
            <input
              v-model="editLocation"
              type="text"
              class="w-full px-3 py-2 text-xs font-serif-editorial bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
            />
          </div>
        </div>

        <!-- 项目关联 -->
        <div v-if="projects && projects.length > 0" class="space-y-1.5">
          <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">关联造物项目</label>
          <select
            v-model="editProjectId"
            class="w-full px-3 py-2 text-xs font-serif-editorial bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
          >
            <option :value="null">不关联项目</option>
            <option v-for="p in projects" :key="p.id" :value="p.id">{{ p.title }}</option>
          </select>
        </div>

        <div class="pt-6 border-t border-[var(--border-subtle)] flex items-center justify-end gap-3">
          <button
            type="button"
            @click="cancelEdit"
            class="px-4 py-2 text-xs font-mono-archive text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer"
          >
            取消
          </button>
          <button
            type="submit"
            :disabled="isSubmitting"
            class="px-5 py-2 text-xs font-mono-archive bg-[var(--ink-primary)] text-[var(--bg-archive)] hover:opacity-90 rounded-xs transition-opacity cursor-pointer disabled:opacity-50 flex items-center gap-1.5"
          >
            <Check v-if="!isSubmitting" class="w-3.5 h-3.5" />
            <span>{{ isSubmitting ? '保存中...' : '保存修改' }}</span>
          </button>
        </div>
      </form>

      <!-- 浏览模式 (Editorial 典雅展示) -->
      <article v-else class="space-y-8 font-serif-editorial">
        <!-- 标题 -->
        <h1 v-if="entry.title" class="text-2xl sm:text-3xl font-medium tracking-tight text-[var(--ink-primary)] leading-snug">
          {{ entry.title }}
        </h1>

        <!-- 大尺寸图片画廊 -->
        <div v-if="displayImages.length > 0" class="space-y-4">
          <div
            v-for="(imgUrl, idx) in displayImages"
            :key="idx"
            class="overflow-hidden rounded-xs bg-[var(--bg-subtle)] border border-[var(--border-subtle)]"
          >
            <img
              :src="imgUrl"
              :alt="entry.title"
              class="w-full max-h-[680px] object-cover editorial-image"
              loading="lazy"
            />
          </div>
        </div>

        <!-- 正文叙事 -->
        <div class="text-lg text-[var(--ink-secondary)] leading-relaxed whitespace-pre-line font-normal">
          {{ entry.content }}
        </div>

        <!-- 注脚信息：地点、关联项目与元信息 -->
        <div class="pt-6 border-t border-[var(--border-subtle)] space-y-2 text-xs font-mono-archive text-[var(--ink-muted)]">
          <div v-if="entry.location" class="flex items-center gap-2">
            <span>坐标：</span>
            <span class="text-[var(--ink-primary)] font-serif-editorial">{{ entry.location }}</span>
          </div>

          <div v-if="entry.project" class="flex items-center gap-2">
            <span>关联造物：</span>
            <span class="text-[var(--ink-primary)]">{{ entry.project.title }}</span>
          </div>

          <div v-if="entry.meta" class="flex items-center gap-2">
            <span>注脚：</span>
            <span>{{ entry.meta }}</span>
          </div>

          <div v-if="entry.link" class="pt-2">
            <a
              :href="entry.link"
              target="_blank"
              rel="noopener noreferrer"
              class="inline-flex items-center gap-1 text-[var(--ink-primary)] hover-underline"
            >
              <span>查看关联外部链接</span>
              <ArrowUpRight class="w-3 h-3" />
            </a>
          </div>
        </div>
      </article>

    </div>
  </div>
</template>
