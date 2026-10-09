<script setup lang="ts">
import { ref, computed } from 'vue'
import { X, ArrowLeft, Check } from 'lucide-vue-next'
import ImageUploader from './ImageUploader.vue'
import { createLifeEntry, createTransaction, createProject } from '../utils/api'
import type { LifeEntry, Transaction, Project } from '../types'

const props = defineProps<{
  isOpen: boolean
  storageConfigured: boolean
  projects?: Project[]
}>()

const emit = defineEmits<{
  (e: 'close'): void
  (e: 'entryCreated', entry: LifeEntry): void
  (e: 'transactionCreated', tx: Transaction): void
  (e: 'projectCreated', proj: Project): void
}>()

type EntryKind = 'life' | 'thought' | 'photo' | 'project' | 'collection' | 'transaction'

const step = ref<'select' | 'form'>('select')
const selectedKind = ref<EntryKind>('life')
const isSubmitting = ref(false)
const errorMsg = ref<string | null>(null)

// 基础表单状态
const title = ref('')
const content = ref('')
const location = ref('')
const occurredAt = ref(getTodayDateTimeString())
const selectedProjectId = ref<number | null>(null)
const attachmentIds = ref<number[]>([])

// 记账专用状态
const txAmount = ref<string>('')
const txType = ref<'expense' | 'income'>('expense')
const txCategory = ref<string>('餐饮')
const txPaymentMethod = ref<string>('微信支付')
const txNote = ref<string>('')

const txCategories = ['餐饮', '交通', '购物', '娱乐', '学习', '住房', '医疗', '其他']
const paymentMethods = ['微信支付', '支付宝', '银行卡', '现金', '其他']

const kindOptions: { key: EntryKind; label: string; desc: string }[] = [
  { key: 'life', label: '生活记录', desc: '出门、吃饭、旅行、手冲或今天发生的小事' },
  { key: 'thought', label: '想法', desc: '随笔、突发灵感、人生感悟与文字碎片' },
  { key: 'photo', label: '照片', desc: '胶卷底片、光影瞬间与视觉记忆' },
  { key: 'project', label: '项目', desc: '亲手创造的工具、设计实验与长线造物' },
  { key: 'collection', label: '收藏', desc: '喜欢的音乐唱片、书籍、游戏与长伴物件' },
  { key: 'transaction', label: '记账', desc: '记录一笔收支，把握日常真实开销' },
]

function getTodayDateTimeString(): string {
  const now = new Date()
  const pad = (n: number) => String(n).padStart(2, '0')
  return `${now.getFullYear()}-${pad(now.getMonth() + 1)}-${pad(now.getDate())}T${pad(now.getHours())}:${pad(now.getMinutes())}`
}

const selectKind = (k: EntryKind) => {
  selectedKind.value = k
  step.value = 'form'
  errorMsg.value = null
}

const resetForm = () => {
  step.value = 'select'
  selectedKind.value = 'life'
  title.value = ''
  content.value = ''
  location.value = ''
  occurredAt.value = getTodayDateTimeString()
  selectedProjectId.value = null
  attachmentIds.value = []
  txAmount.value = ''
  txType.value = 'expense'
  txCategory.value = '餐饮'
  txPaymentMethod.value = '微信支付'
  txNote.value = ''
  errorMsg.value = null
  isSubmitting.value = false
}

const handleClose = () => {
  resetForm()
  emit('close')
}

const kindTitle = computed(() => {
  const match = kindOptions.find((o) => o.key === selectedKind.value)
  return match ? match.label : '记录'
})

const contentPlaceholder = computed(() => {
  switch (selectedKind.value) {
    case 'thought':
      return '写下此刻脑海里的念头、感触或片段……'
    case 'photo':
      return '为这些照片写几句注脚或拍摄时的心境……'
    case 'project':
      return '记录这个项目的进展、思考、踩坑或新的突破……'
    case 'collection':
      return '记录这张唱片、这本书、这款游戏或这件好物吸引你的地方……'
    default:
      return '记录今天真实发生的一件事、迎面的微风、味道或琐碎日常……'
  }
})

const handleSubmit = async () => {
  errorMsg.value = null

  if (selectedKind.value === 'transaction') {
    // 提交记账
    const amtFloat = parseFloat(txAmount.value)
    if (isNaN(amtFloat) || amtFloat <= 0) {
      errorMsg.value = '请输入有效的金额'
      return
    }
    const amountInCents = Math.round(amtFloat * 100)

    try {
      isSubmitting.value = true
      const createdTx = await createTransaction({
        amount: amountInCents,
        type: txType.value,
        category: txCategory.value,
        title: title.value.trim() || txCategory.value,
        note: txNote.value.trim(),
        payment_method: txPaymentMethod.value,
        occurred_at: occurredAt.value,
      })
      emit('transactionCreated', createdTx)
      handleClose()
    } catch (err: any) {
      errorMsg.value = err.message || '记账失败'
    } finally {
      isSubmitting.value = false
    }
    return
  }

  // 提交 Life Entry
  if (!content.value.trim() && selectedKind.value !== 'photo') {
    errorMsg.value = '请填写记录内容'
    return
  }
  if (selectedKind.value === 'photo' && attachmentIds.value.length === 0 && !content.value.trim()) {
    errorMsg.value = '请上传至少一张照片或填写文字记录'
    return
  }

  try {
    isSubmitting.value = true

    // 如果是新建项目类型，同时允许创建基础 Project
    let projId = selectedProjectId.value
    if (selectedKind.value === 'project' && title.value.trim() && !projId) {
      try {
        const newProj = await createProject({
          name: title.value.trim(),
          description: content.value.trim().slice(0, 100),
          content: content.value.trim(),
        })
        projId = newProj.id
        emit('projectCreated', newProj)
      } catch {
        // ignore
      }
    }

    const created = await createLifeEntry({
      type: selectedKind.value,
      title: title.value.trim(),
      content: content.value.trim() || (selectedKind.value === 'photo' ? '定格的照片。' : ''),
      occurred_at: occurredAt.value,
      location: location.value.trim(),
      project_id: projId,
      attachment_ids: attachmentIds.value,
    })

    emit('entryCreated', created)
    handleClose()
  } catch (err: any) {
    errorMsg.value = err.message || '保存失败'
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/40 backdrop-blur-xs animate-in fade-in duration-200"
    @click.self="handleClose"
  >
    <div
      class="bg-[var(--bg-archive)] border border-[var(--border-subtle)] shadow-[0_16px_50px_rgba(0,0,0,0.12)] w-full max-w-xl max-h-[90vh] overflow-y-auto rounded-xs p-6 sm:p-8 text-left transition-all duration-200 relative"
    >
      <!-- 顶部轻量标头 -->
      <div class="flex items-center justify-between pb-6 border-b border-[var(--border-subtle)] mb-6">
        <div class="flex items-center gap-3">
          <button
            v-if="step === 'form'"
            @click="step = 'select'"
            class="text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer p-1 -ml-1 transition-colors"
            title="返回类型选择"
          >
            <ArrowLeft class="w-4 h-4" />
          </button>
          <div>
            <h2 class="font-serif-editorial text-xl sm:text-2xl font-medium tracking-tight text-[var(--ink-primary)]">
              {{ step === 'select' ? '记录一点什么' : `新建 · ${kindTitle}` }}
            </h2>
            <p class="font-mono-archive text-[11px] text-[var(--ink-muted)] mt-0.5">
              {{ step === 'select' ? '选择你想记录的生活形态' : '沉淀真实的生活记忆' }}
            </p>
          </div>
        </div>

        <button
          @click="handleClose"
          class="p-1 text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer transition-colors"
          title="关闭"
        >
          <X class="w-5 h-5" />
        </button>
      </div>

      <!-- 第一步：轻量选择菜单 -->
      <div v-if="step === 'select'" class="space-y-3 py-2">
        <button
          v-for="opt in kindOptions"
          :key="opt.key"
          @click="selectKind(opt.key)"
          class="w-full text-left p-4 rounded-xs border border-[var(--border-subtle)] hover:border-[var(--ink-secondary)] hover:bg-[var(--bg-subtle)]/40 transition-all cursor-pointer group flex items-start justify-between"
        >
          <div class="space-y-1">
            <span class="font-serif-editorial text-base text-[var(--ink-primary)] group-hover:text-[var(--accent-warm)] transition-colors block font-medium">
              {{ opt.label }}
            </span>
            <p class="text-xs text-[var(--ink-muted)] font-serif-editorial">
              {{ opt.desc }}
            </p>
          </div>
          <span class="text-xs font-mono-archive text-[var(--ink-muted)] opacity-40 group-hover:opacity-100 transition-opacity">
            →
          </span>
        </button>
      </div>

      <!-- 第二步：克制清晰的表单 -->
      <form v-else @submit.prevent="handleSubmit" class="space-y-6">
        
        <!-- 错误提示 -->
        <div v-if="errorMsg" class="p-3 bg-red-50 text-red-700 text-xs rounded-xs border border-red-200">
          {{ errorMsg }}
        </div>

        <!-- 记账表单分支 -->
        <template v-if="selectedKind === 'transaction'">
          <!-- 金额与类型 -->
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div class="space-y-1.5">
              <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">金额 (¥) *</label>
              <div class="relative">
                <span class="absolute left-3 top-2.5 text-sm font-mono-archive text-[var(--ink-muted)]">¥</span>
                <input
                  v-model="txAmount"
                  type="number"
                  step="0.01"
                  min="0.01"
                  required
                  placeholder="0.00"
                  class="w-full pl-7 pr-3 py-2 text-lg font-mono-archive bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
                  autofocus
                />
              </div>
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">收支类型</label>
              <div class="flex gap-2 pt-0.5">
                <button
                  type="button"
                  @click="txType = 'expense'"
                  class="flex-1 py-2 text-xs font-serif-editorial rounded-xs border transition-colors cursor-pointer"
                  :class="txType === 'expense' ? 'border-[var(--ink-primary)] font-bold text-[var(--ink-primary)]' : 'border-[var(--border-subtle)] text-[var(--ink-muted)]'"
                >
                  支出
                </button>
                <button
                  type="button"
                  @click="txType = 'income'"
                  class="flex-1 py-2 text-xs font-serif-editorial rounded-xs border transition-colors cursor-pointer"
                  :class="txType === 'income' ? 'border-[var(--ink-primary)] font-bold text-[var(--ink-primary)]' : 'border-[var(--border-subtle)] text-[var(--ink-muted)]'"
                >
                  收入
                </button>
              </div>
            </div>
          </div>

          <!-- 分类标签选择 -->
          <div class="space-y-1.5">
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">消费分类</label>
            <div class="flex flex-wrap gap-2">
              <button
                v-for="cat in txCategories"
                :key="cat"
                type="button"
                @click="txCategory = cat"
                class="px-2.5 py-1 text-xs font-serif-editorial rounded-xs border transition-colors cursor-pointer"
                :class="txCategory === cat ? 'border-[var(--ink-primary)] font-medium text-[var(--ink-primary)] bg-[var(--bg-subtle)]' : 'border-[var(--border-subtle)] text-[var(--ink-muted)]'"
              >
                {{ cat }}
              </button>
            </div>
          </div>

          <!-- 账目名称与支付方式 -->
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div class="space-y-1.5">
              <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">事项名称</label>
              <input
                v-model="title"
                type="text"
                placeholder="例如：午饭、咖啡豆、键盘"
                class="w-full px-3 py-2 text-sm bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none font-serif-editorial"
              />
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">支付渠道</label>
              <select
                v-model="txPaymentMethod"
                class="w-full px-3 py-2 text-sm bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none font-serif-editorial"
              >
                <option v-for="m in paymentMethods" :key="m" :value="m">{{ m }}</option>
              </select>
            </div>
          </div>

          <!-- 备注 -->
          <div class="space-y-1.5">
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">备注说明 (可选)</label>
            <input
              v-model="txNote"
              type="text"
              placeholder="添加细节或心情批注……"
              class="w-full px-3 py-2 text-sm bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none font-serif-editorial"
            />
          </div>
        </template>

        <!-- 普通 Life Entry 表单分支 (生活 / 想法 / 照片 / 项目 / 收藏) -->
        <template v-else>
          <!-- 标题 (可选) -->
          <div class="space-y-1.5">
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">
              标题 (可选)
            </label>
            <input
              v-model="title"
              type="text"
              :placeholder="selectedKind === 'project' ? '项目名称 (例如: ArchCanvas)' : '给这条记录一个简短标题……'"
              class="w-full px-3 py-2 text-base font-serif-editorial bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
            />
          </div>

          <!-- 内容主体 (必填) -->
          <div class="space-y-1.5">
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">
              记录内容 *
            </label>
            <textarea
              v-model="content"
              rows="5"
              :placeholder="contentPlaceholder"
              class="w-full px-3 py-2 text-base font-serif-editorial leading-relaxed bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none resize-y"
            ></textarea>
          </div>

          <!-- 照片上传组件 -->
          <ImageUploader
            :storageConfigured="storageConfigured"
            @update:attachments="attachmentIds = $event"
          />

          <!-- 发生时间与地点 -->
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div class="space-y-1.5">
              <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">发生时间</label>
              <input
                v-model="occurredAt"
                type="datetime-local"
                class="w-full px-3 py-2 text-xs font-mono-archive bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
              />
            </div>

            <div class="space-y-1.5">
              <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">地点 (可选)</label>
              <input
                v-model="location"
                type="text"
                placeholder="例如：杭州 · 满觉陇"
                class="w-full px-3 py-2 text-xs font-serif-editorial bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
              />
            </div>
          </div>

          <!-- 项目关联 (可选) -->
          <div v-if="projects && projects.length > 0 && selectedKind !== 'project'" class="space-y-1.5">
            <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">关联造物项目 (可选)</label>
            <select
              v-model="selectedProjectId"
              class="w-full px-3 py-2 text-xs font-serif-editorial bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
            >
              <option :value="null">不关联项目</option>
              <option v-for="p in projects" :key="p.id" :value="p.id">{{ p.title }}</option>
            </select>
          </div>
        </template>

        <!-- 发生时间 (记账时置底) -->
        <div v-if="selectedKind === 'transaction'" class="space-y-1.5">
          <label class="block text-xs font-mono-archive text-[var(--ink-muted)]">记录时间</label>
          <input
            v-model="occurredAt"
            type="datetime-local"
            class="w-full px-3 py-2 text-xs font-mono-archive bg-transparent border border-[var(--border-subtle)] focus:border-[var(--ink-primary)] rounded-xs outline-none"
          />
        </div>

        <!-- 操作按钮 -->
        <div class="pt-6 border-t border-[var(--border-subtle)] flex items-center justify-end gap-3">
          <button
            type="button"
            @click="handleClose"
            class="px-4 py-2 text-xs font-mono-archive text-[var(--ink-muted)] hover:text-[var(--ink-primary)] transition-colors cursor-pointer"
          >
            取消
          </button>
          <button
            type="submit"
            :disabled="isSubmitting"
            class="px-5 py-2 text-xs font-mono-archive bg-[var(--ink-primary)] text-[var(--bg-archive)] hover:opacity-90 rounded-xs transition-opacity cursor-pointer disabled:opacity-50 flex items-center gap-1.5"
          >
            <Check v-if="!isSubmitting" class="w-3.5 h-3.5" />
            <span>{{ isSubmitting ? '保存中...' : '记录到生活档案' }}</span>
          </button>
        </div>

      </form>
    </div>
  </div>
</template>
