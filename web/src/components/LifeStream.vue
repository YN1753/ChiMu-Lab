<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import type { LifeEntry, ArchiveCategory } from '../types'
import { getEntryCategory, entryMatchesCategory } from '../types'
import { ArrowUpRight, Maximize2 } from 'lucide-vue-next'
import EditorialLightbox from './EditorialLightbox.vue'

const props = defineProps<{
  entries: LifeEntry[]
  limit?: number
  showFilters?: boolean
  dateFilter?: string
  currentCategory?: ArchiveCategory
}>()

const emit = defineEmits<{
  (e: 'clearDateFilter'): void
  (e: 'changeCategory', category: ArchiveCategory): void
  (e: 'selectEntry', entry: LifeEntry): void
  (e: 'openAdd'): void
}>()

const activeCategory = ref<ArchiveCategory>(props.currentCategory || 'all')

// 暗房灯箱状态
const isLightboxOpen = ref(false)
const lightboxPayload = ref({
  imageUrl: '',
  title: '',
  meta: '',
  location: '',
  date: '',
})

const openLightbox = (entry: LifeEntry, customUrl?: string) => {
  lightboxPayload.value = {
    imageUrl: customUrl || entry.images || '',
    title: entry.title || '',
    meta: entry.meta || '',
    location: entry.location || '',
    date: formatEntryDate(entry),
  }
  isLightboxOpen.value = true
}

watch(
  () => props.currentCategory,
  (newCat) => {
    if (newCat) {
      activeCategory.value = newCat
    }
  }
)

// 视图过滤分类
const categoryFilters: { key: ArchiveCategory; label: string }[] = [
  { key: 'all', label: '全部记录' },
  { key: 'daily', label: '日常' },
  { key: 'thought', label: '想法' },
  { key: 'photo', label: '照片' },
  { key: 'transaction', label: '记账' },
  { key: 'project', label: '项目' },
  { key: 'collection', label: '收藏' },
]

const filteredEntries = computed(() => {
  let list = props.entries

  // 1. 日期筛选
  if (props.dateFilter) {
    const target = props.dateFilter.replace(/-/g, '.').trim()
    list = list.filter((e) => (e.date ? e.date.replace(/-/g, '.').trim() === target : true))
  }

  // 2. 类别筛选
  if (activeCategory.value !== 'all') {
    list = list.filter((e) => entryMatchesCategory(e, activeCategory.value))
  }

  // 3. 数量限制
  if (props.limit && props.limit > 0 && !props.dateFilter && activeCategory.value === 'all') {
    return list.slice(0, props.limit)
  }

  return list
})

const handleCategoryClick = (cat: ArchiveCategory) => {
  activeCategory.value = cat
  emit('changeCategory', cat)
}

// 归一化条目显示归属
const getCategoryName = (entry: LifeEntry) => {
  const cat = getEntryCategory(entry)
  const map: Record<ArchiveCategory, string> = {
    all: '全部记录',
    daily: '日常',
    thought: '想法',
    project: '项目',
    collection: '收藏',
    photo: '照片',
    transaction: '记账',
  }
  return map[cat] || '日常'
}

const formatEntryDate = (entry: LifeEntry) => {
  if (entry.month && entry.day) {
    return `${entry.month}月${entry.day}日`
  }
  if (entry.occurred_at) {
    const d = new Date(entry.occurred_at)
    if (!isNaN(d.getTime())) {
      return `${d.getMonth() + 1}月${d.getDate()}日`
    }
  }
  return entry.date || '近期'
}

const formatEntryTime = (entry: LifeEntry) => {
  if (entry.time) return entry.time
  if (entry.occurred_at) {
    const d = new Date(entry.occurred_at)
    if (!isNaN(d.getTime())) {
      const pad = (n: number) => String(n).padStart(2, '0')
      return `${pad(d.getHours())}:${pad(d.getMinutes())}`
    }
  }
  return ''
}

const formatMoney = (cents: number): string => {
  return (cents / 100).toLocaleString('zh-CN', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  })
}

// 岁月留白刻度计算 (计算两篇相邻记录之间的空白跨度)
const getEntryTimestamp = (entry: LifeEntry): number => {
  if (entry.occurred_at) {
    const t = new Date(entry.occurred_at).getTime()
    if (!isNaN(t)) return t
  }
  if (entry.date) {
    const clean = entry.date.replace(/\./g, '-')
    const t = new Date(clean).getTime()
    if (!isNaN(t)) return t
  }
  if (entry.year && entry.month && entry.day) {
    const t = new Date(`${entry.year}-${entry.month}-${entry.day}`).getTime()
    if (!isNaN(t)) return t
  }
  return 0
}

const getGapDays = (index: number): number => {
  if (index === 0) return 0
  const current = filteredEntries.value[index]
  const prev = filteredEntries.value[index - 1]
  const tCurrent = getEntryTimestamp(current)
  const tPrev = getEntryTimestamp(prev)
  if (!tCurrent || !tPrev) return 0
  const diffMs = Math.abs(tPrev - tCurrent)
  return Math.round(diffMs / (1000 * 60 * 60 * 24))
}
</script>

<template>
  <div class="w-full text-left">
    <!-- 视角分类筛选栏 (极简呼吸感) -->
    <div
      v-if="showFilters"
      class="flex flex-wrap items-center gap-x-6 gap-y-2 mb-12 text-xs font-mono-archive text-[var(--ink-muted)] border-b border-[var(--border-subtle)] pb-4"
    >
      <span class="text-[var(--ink-secondary)] mr-1">视角 //</span>
      <button
        v-for="f in categoryFilters"
        :key="f.key"
        @click="handleCategoryClick(f.key)"
        class="transition-colors cursor-pointer py-1 hover:text-[var(--ink-primary)]"
        :class="{
          'text-[var(--ink-primary)] font-semibold border-b border-[var(--ink-primary)]':
            activeCategory === f.key,
        }"
      >
        {{ f.label }}
      </button>
    </div>

    <!-- 日期筛选提醒 -->
    <div
      v-if="dateFilter"
      class="mb-10 pb-4 border-b border-[var(--border-subtle)] flex items-center justify-between text-xs font-mono-archive"
    >
      <div class="flex items-center gap-2">
        <span class="text-[var(--accent-warm)] font-medium">聚焦日期：{{ dateFilter }}</span>
        <span class="text-[var(--ink-muted)]">（共 {{ filteredEntries.length }} 条生活印记）</span>
      </div>
      <button
        @click="emit('clearDateFilter')"
        class="text-[var(--ink-secondary)] hover:text-[var(--ink-primary)] hover-underline cursor-pointer"
      >
        [显示全部日期记录]
      </button>
    </div>

    <!-- 空状态 -->
    <div v-if="filteredEntries.length === 0" class="py-20 text-center space-y-4">
      <div class="space-y-1.5 font-serif-editorial">
        <h3 class="text-lg text-[var(--ink-primary)]">
          {{
            activeCategory === 'photo'
              ? '还没有照片'
              : activeCategory === 'transaction'
              ? '这个月还没有记账'
              : '还没有记录'
          }}
        </h3>
        <p class="text-xs text-[var(--ink-muted)] leading-relaxed">
          从今天开始，<br />
          记录一些你不想忘记的事情。
        </p>
      </div>
      <button
        @click="emit('openAdd')"
        class="inline-flex items-center gap-1.5 px-4 py-2 text-xs font-mono-archive border border-[var(--border-subtle)] hover:border-[var(--ink-primary)] rounded-xs cursor-pointer transition-colors"
      >
        <span>＋ 添加第一条记录</span>
      </button>
    </div>

    <!-- 核心：纵向时间线 (Vertical Spine Timeline) -->
    <div
      v-else
      class="relative pl-6 sm:pl-10 border-l border-[var(--border-subtle)] ml-3 sm:ml-5 space-y-12 sm:space-y-16 py-2"
    >
      <template v-for="(entry, index) in filteredEntries" :key="entry.id">
        <!-- 岁月留白刻度 (跨越 >= 14 天的静默时段) -->
        <div
          v-if="getGapDays(index) >= 14"
          class="relative py-2 flex items-center gap-3 sm:gap-4 text-xs font-mono-archive text-[var(--ink-muted)] select-none -my-4 sm:-my-6"
        >
          <!-- 标尺切线点 -->
          <div class="absolute -left-[30px] sm:-left-[46px] w-2.5 h-px bg-[var(--ink-muted)]/60"></div>
          <span class="tracking-widest text-[11px] text-[var(--ink-muted)]/80">
            // 跨越 {{ getGapDays(index) }} 天的时光留白
          </span>
          <span class="flex-grow h-px bg-[var(--border-subtle)]/70 max-w-xs"></span>
        </div>

        <article
          @click="emit('selectEntry', entry)"
          class="relative group cursor-pointer transition-all"
        >
          <!-- 纵向主轨锚点圆点 (Timeline Node) -->
          <div
            class="absolute -left-[31px] sm:-left-[47px] top-1.5 flex items-center justify-center"
          >
            <div
              class="w-2.5 h-2.5 rounded-full bg-[var(--bg-archive)] border-2 border-[var(--ink-secondary)] group-hover:border-[var(--accent-warm)] group-hover:scale-125 transition-all"
              :class="{
                'border-[var(--accent-warm)] bg-[var(--accent-warm)]/20': entry.featured,
                'border-emerald-600': entry.type === 'transaction' || entry.is_transaction,
              }"
            ></div>
          </div>

          <!-- 顶部元信息行：日期 · 时间 · 类型标签 · 地点 -->
          <div class="flex items-center gap-2 text-xs font-mono-archive text-[var(--ink-muted)] mb-2.5">
            <span class="font-semibold text-[var(--ink-primary)]">
              {{ formatEntryDate(entry) }}
            </span>
            <span v-if="formatEntryTime(entry)">· {{ formatEntryTime(entry) }}</span>
            <span>·</span>
            <span
              class="text-[var(--ink-secondary)]"
              :class="{
                'text-[var(--accent-warm)] font-medium': entry.type === 'thought',
                'text-emerald-700 font-medium': entry.type === 'transaction' || entry.is_transaction,
              }"
            >
              {{ getCategoryName(entry) }}
            </span>
            <span v-if="entry.location" class="hidden sm:inline text-[var(--ink-muted)]">
              · {{ entry.location }}
            </span>
          </div>

          <!-- 多态内容排版 (Polymorphic Layouts) -->

          <!-- 1. 想法 / 随笔形态 (Thought)：免标题，舒适大字号引言体，轻盈随手记 -->
          <div v-if="entry.type === 'thought'" class="space-y-2">
            <blockquote
              class="font-serif-editorial text-lg sm:text-xl text-[var(--ink-primary)] leading-relaxed italic border-l-2 border-[var(--ink-muted)]/30 pl-4 py-1 group-hover:border-[var(--accent-warm)] transition-colors"
            >
              “{{ entry.content }}”
            </blockquote>
            <div
              v-if="entry.title && entry.title !== entry.content && entry.title !== '想法'"
              class="text-xs font-mono-archive text-[var(--ink-muted)] pl-4"
            >
              — {{ entry.title }}
            </div>
            <div v-if="entry.meta" class="text-xs font-mono-archive text-[var(--ink-muted)]/70 pl-4">
              {{ entry.meta }}
            </div>
          </div>

          <!-- 2. 记账形态 (Transaction)：单行账本微卡片，柴米油盐融入时间流 -->
          <div
            v-else-if="entry.type === 'transaction' || entry.is_transaction"
            class="inline-flex flex-wrap items-center gap-3 py-2.5 px-4 bg-[var(--bg-subtle)]/60 rounded-xs border border-[var(--border-subtle)] group-hover:border-[var(--border-divider)] transition-colors"
          >
            <span
              class="text-[11px] font-mono-archive text-[var(--ink-muted)] px-1.5 py-0.2 rounded-xs border border-[var(--border-subtle)]"
            >
              {{ entry.category || '日常' }}
            </span>
            <span class="font-serif-editorial text-base text-[var(--ink-primary)] font-medium">
              {{ entry.title }}
            </span>
            <span
              class="font-mono-archive text-base font-semibold"
              :class="entry.tx_type === 'income' ? 'text-[var(--accent-moss)]' : 'text-[var(--ink-primary)]'"
            >
              {{ entry.tx_type === 'income' ? '+' : '-' }}¥{{ formatMoney(entry.amount || 0) }}
            </span>
            <span v-if="entry.content" class="text-xs font-serif-editorial text-[var(--ink-muted)]">
              — {{ entry.content }}
            </span>
            <span
              v-if="entry.payment_method"
              class="text-[11px] font-mono-archive text-[var(--ink-muted)] hidden sm:inline"
            >
              ({{ entry.payment_method }})
            </span>
          </div>

          <!-- 3. 照片 / 胶卷形态 (Photo)：视觉画卷，大图与胶片参数 -->
          <div v-else-if="entry.type === 'photo'" class="space-y-3">
            <h3
              v-if="entry.title"
              class="text-xl sm:text-2xl font-serif-editorial font-medium tracking-tight text-[var(--ink-primary)] leading-snug group-hover:text-[var(--accent-warm)] transition-colors"
            >
              {{ entry.title }}
            </h3>
            <!-- 胶片装裱底片框 (Passe-partout) -->
            <div
              v-if="entry.images"
              class="film-frame bg-[var(--bg-subtle)]/80 p-2 sm:p-3 rounded-xs border border-[var(--border-subtle)] hover:border-[var(--border-divider)] transition-all max-w-3xl group/photo cursor-zoom-in"
              @click.stop="openLightbox(entry)"
              title="点击在暗房灯箱中查看底片"
            >
              <div class="relative overflow-hidden rounded-xs bg-[var(--bg-archive)]">
                <img
                  :src="entry.images"
                  :alt="entry.title"
                  class="w-full max-h-[580px] object-cover editorial-image group-hover/photo:scale-[1.01] transition-transform duration-500"
                  loading="lazy"
                />
                <div class="absolute bottom-2 right-2 px-2 py-1 bg-black/60 backdrop-blur-xs text-white text-[10px] font-mono-archive rounded-xs opacity-0 group-hover/photo:opacity-100 transition-opacity flex items-center gap-1">
                  <Maximize2 class="w-3 h-3" />
                  <span>暗房灯箱</span>
                </div>
              </div>
              <!-- 齿孔与冲印装裱条 -->
              <div class="pt-2 px-1 flex items-center justify-between text-[11px] font-mono-archive text-[var(--ink-muted)]">
                <span class="tracking-widest uppercase">EXP. // {{ entry.meta || 'FILM ROLL' }}</span>
                <div class="flex items-center gap-2">
                  <span v-if="entry.location">{{ entry.location }}</span>
                  <span class="opacity-60 hidden sm:inline">[底片装裱]</span>
                </div>
              </div>
            </div>
            <p
              v-if="entry.content"
              class="text-base sm:text-lg text-[var(--ink-secondary)] leading-relaxed font-normal whitespace-pre-line max-w-2xl font-serif-editorial"
            >
              {{ entry.content }}
            </p>
            <div
              v-if="entry.meta"
              class="text-xs font-mono-archive text-[var(--ink-muted)] flex items-center justify-between max-w-2xl"
            >
              <span>— {{ entry.meta }}</span>
              <span v-if="entry.location" class="hidden sm:inline">{{ entry.location }}</span>
            </div>
          </div>

          <!-- 4. 造物 / 项目形态 (Project)：手艺与故事，带源码仓库直链 -->
          <div
            v-else-if="entry.type === 'project'"
            class="space-y-3 p-4 sm:p-5 rounded-xs border border-[var(--border-subtle)] bg-[var(--bg-subtle)]/30 group-hover:border-[var(--border-divider)] transition-colors max-w-2xl"
          >
            <div class="flex items-center justify-between">
              <span class="text-[10px] font-mono-archive tracking-widest text-[var(--accent-warm)] uppercase">
                // 造物经历 · PROJECT
              </span>
              <span v-if="entry.meta" class="text-[11px] font-mono-archive text-[var(--ink-muted)]">
                {{ entry.meta }}
              </span>
            </div>
            <h3
              class="text-xl font-serif-editorial font-medium text-[var(--ink-primary)] group-hover:text-[var(--accent-warm)] transition-colors"
            >
              {{ entry.title }}
            </h3>
            <p class="text-sm sm:text-base text-[var(--ink-secondary)] leading-relaxed font-serif-editorial">
              {{ entry.content }}
            </p>
            <div v-if="entry.link" class="pt-1">
              <a
                :href="entry.link"
                target="_blank"
                rel="noopener noreferrer"
                @click.stop
                class="inline-flex items-center gap-1.5 text-xs text-[var(--ink-primary)] hover-underline font-mono-archive"
              >
                <span>查看代码与项目仓库</span>
                <ArrowUpRight class="w-3.5 h-3.5" />
              </a>
            </div>
          </div>

          <!-- 5. 默认 / 日常形态 (Daily, Coffee, Music, Place, Books, Games, Collection) -->
          <div v-else class="space-y-3">
            <h3
              v-if="entry.title"
              class="text-xl sm:text-2xl font-serif-editorial font-medium tracking-tight text-[var(--ink-primary)] leading-snug group-hover:text-[var(--accent-warm)] transition-colors"
            >
              {{ entry.title }}
            </h3>

            <div
              v-if="entry.images"
              class="relative w-full overflow-hidden bg-[var(--bg-subtle)] rounded-xs border border-[var(--border-subtle)] max-w-2xl cursor-zoom-in group/img"
              @click.stop="openLightbox(entry)"
              title="点击在暗房灯箱中放大"
            >
              <img
                :src="entry.images"
                :alt="entry.title"
                class="w-full max-h-[500px] object-cover editorial-image group-hover/img:scale-[1.01] transition-transform duration-500"
                loading="lazy"
              />
            </div>

            <p
              class="text-base sm:text-lg text-[var(--ink-secondary)] leading-relaxed font-normal whitespace-pre-line max-w-2xl font-serif-editorial"
            >
              {{ entry.content }}
            </p>

            <div
              v-if="entry.meta && !entry.images"
              class="text-xs font-mono-archive text-[var(--ink-muted)]"
            >
              — {{ entry.meta }}
            </div>

            <div v-if="entry.link" class="pt-1">
              <a
                :href="entry.link"
                target="_blank"
                rel="noopener noreferrer"
                @click.stop
                class="inline-flex items-center gap-1.5 text-xs text-[var(--ink-primary)] hover-underline font-mono-archive"
              >
                <span>相关链接</span>
                <ArrowUpRight class="w-3.5 h-3.5" />
              </a>
            </div>
          </div>

          <!-- 悬浮显露轻量操作提示 -->
          <div
            class="pt-2 text-[11px] font-mono-archive text-[var(--ink-muted)] opacity-0 group-hover:opacity-100 transition-opacity flex items-center gap-1"
          >
            <span>查看详情 / 编辑 →</span>
          </div>
        </article>
      </template>
    </div>

    <!-- 暗房底片灯箱 -->
    <EditorialLightbox
      :isOpen="isLightboxOpen"
      :imageUrl="lightboxPayload.imageUrl"
      :title="lightboxPayload.title"
      :meta="lightboxPayload.meta"
      :location="lightboxPayload.location"
      :date="lightboxPayload.date"
      @close="isLightboxOpen = false"
    />
  </div>
</template>
