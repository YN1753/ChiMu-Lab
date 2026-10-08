<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import type { LifeEntry, ArchiveCategory } from '../types'
import { getEntryCategory } from '../types'
import { ArrowUpRight } from 'lucide-vue-next'

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
}>()

const activeCategory = ref<ArchiveCategory>(props.currentCategory || 'all')

watch(() => props.currentCategory, (newCat) => {
  if (newCat) {
    activeCategory.value = newCat
  }
})

// 五重核心生活档案视图
const categoryFilters: { key: ArchiveCategory; label: string }[] = [
  { key: 'all', label: '全部记录' },
  { key: 'daily', label: '日常' },
  { key: 'thought', label: '想法' },
  { key: 'project', label: '项目' },
  { key: 'collection', label: '收藏' },
]

const filteredEntries = computed(() => {
  let list = props.entries

  // 1. 日期筛选 (来自生活刻度点击)
  if (props.dateFilter) {
    const target = props.dateFilter.replace(/-/g, '.').trim()
    list = list.filter(e => e.date.replace(/-/g, '.').trim() === target)
  }

  // 2. 核心 5 重分类视图筛选
  if (activeCategory.value !== 'all') {
    list = list.filter(e => getEntryCategory(e) === activeCategory.value)
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

// 归一化条目显示归属 (日常 / 想法 / 项目 / 收藏)
const getCategoryName = (entry: LifeEntry) => {
  const cat = getEntryCategory(entry)
  const map: Record<ArchiveCategory, string> = {
    all: '全部记录',
    daily: '日常',
    thought: '想法',
    project: '项目',
    collection: '收藏',
  }
  return map[cat] || '日常'
}
</script>

<template>
  <div class="w-full text-left">
    
    <!-- 克制的五重分类筛选栏 (无 Tag 标签系统，仅 5 种纯粹视图观察方式) -->
    <div v-if="showFilters" class="flex flex-wrap items-center gap-x-6 gap-y-2 mb-12 text-xs font-mono-archive text-[var(--ink-muted)] border-b border-[var(--border-subtle)] pb-4">
      <span class="text-[var(--ink-secondary)] mr-1">视角 //</span>
      <button
        v-for="f in categoryFilters"
        :key="f.key"
        @click="handleCategoryClick(f.key)"
        class="transition-colors cursor-pointer py-1 hover:text-[var(--ink-primary)]"
        :class="{ 'text-[var(--ink-primary)] font-semibold border-b border-[var(--ink-primary)]': activeCategory === f.key }"
      >
        {{ f.label }}
      </button>
    </div>

    <!-- 日期筛选提醒 (纯文字，无卡片) -->
    <div v-if="dateFilter" class="mb-10 pb-4 border-b border-[var(--border-subtle)] flex items-center justify-between text-xs font-mono-archive">
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
    <div v-if="filteredEntries.length === 0" class="py-16 text-center text-sm font-serif-editorial text-[var(--ink-muted)]">
      该观察视角下暂无生活记录。
    </div>

    <!-- 时间流列表：无 Card UI，纯粹的独立出版物杂志排版，无任何 emoji，无杂乱 Tag 标签 -->
    <div v-else class="divide-y divide-[var(--border-subtle)]">
      
      <article
        v-for="entry in filteredEntries"
        :key="entry.id"
        class="py-12 sm:py-16 first:pt-4"
      >
        <div class="grid grid-cols-1 md:grid-cols-12 gap-6 md:gap-12 items-start">
          
          <!-- 左侧：时间锚点与生活归属 (纯文字注脚) -->
          <div class="md:col-span-3 space-y-2">
            <div class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
              <span class="text-[var(--ink-secondary)] font-semibold block text-sm">{{ entry.year }}年</span>
              <span class="text-base text-[var(--ink-primary)] font-serif-editorial block mt-0.5">
                {{ entry.month }}月{{ entry.day }}日
              </span>
              <span v-if="entry.time" class="block text-[11px] text-[var(--ink-muted)] mt-1">
                {{ entry.time }}
              </span>
            </div>

            <!-- 极简归类：日常 · 想法 · 项目 · 收藏 -->
            <div class="pt-2">
              <span class="inline-block text-[11px] tracking-wider text-[var(--ink-muted)] border-b border-[var(--border-subtle)] pb-0.5">
                · {{ getCategoryName(entry) }}
              </span>
            </div>

            <!-- 地点注脚 -->
            <div v-if="entry.location" class="text-xs text-[var(--ink-muted)] pt-1 font-normal font-serif-editorial">
              {{ entry.location }}
            </div>
          </div>

          <!-- 右侧：有机排版的内容主体 -->
          <div class="md:col-span-9 space-y-5">
            
            <!-- 标题 -->
            <h3 class="text-xl sm:text-2xl font-serif-editorial font-medium tracking-tight text-[var(--ink-primary)] leading-snug">
              {{ entry.title }}
            </h3>

            <!-- 大幅摄影表现 (摄影是内容的表达形式，容纳在日常记录中) -->
            <div v-if="entry.images" class="pt-2">
              <div class="relative w-full overflow-hidden bg-[var(--bg-subtle)] rounded-sm">
                <img
                  :src="entry.images"
                  :alt="entry.title"
                  class="w-full max-h-[640px] object-cover editorial-image"
                  loading="lazy"
                />
              </div>
              <div v-if="entry.meta" class="mt-2 text-xs font-mono-archive text-[var(--ink-muted)] flex items-center justify-between">
                <span>{{ entry.meta }}</span>
                <span v-if="entry.location" class="hidden sm:inline">{{ entry.location }}</span>
              </div>
            </div>

            <!-- 正文叙事 -->
            <p class="text-base sm:text-lg text-[var(--ink-secondary)] leading-relaxed font-normal whitespace-pre-line max-w-2xl font-serif-editorial">
              {{ entry.content }}
            </p>

            <!-- 附加参数与说明 (无大图时的元信息) -->
            <div v-if="entry.meta && !entry.images" class="pt-2 flex items-center gap-3 text-xs font-mono-archive text-[var(--ink-muted)]">
              <span>— {{ entry.meta }}</span>
            </div>

            <!-- 项目/造物链接 (克制细线，无按钮卡片) -->
            <div v-if="entry.link" class="pt-2">
              <a
                :href="entry.link"
                target="_blank"
                rel="noopener noreferrer"
                class="inline-flex items-center gap-1.5 text-xs text-[var(--ink-primary)] hover-underline pb-0.5 font-mono-archive"
              >
                <span>查看项目与代码仓库</span>
                <ArrowUpRight class="w-3.5 h-3.5" />
              </a>
            </div>

          </div>

        </div>
      </article>

    </div>

  </div>
</template>
