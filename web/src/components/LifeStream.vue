<script setup lang="ts">
import { ref, computed } from 'vue'
import type { LifeEntry } from '../types'
import { ArrowUpRight } from 'lucide-vue-next'

const props = defineProps<{
  entries: LifeEntry[]
  limit?: number
  showFilters?: boolean
  dateFilter?: string
}>()

const emit = defineEmits<{
  (e: 'clearDateFilter'): void
}>()

const activeFilter = ref<string>('all')

const filters = [
  { key: 'all', label: '全部' },
  { key: 'thought', label: '随想' },
  { key: 'photo', label: '胶卷摄影' },
  { key: 'coffee', label: '手冲咖啡' },
  { key: 'music', label: '听音' },
  { key: 'book', label: '书摘' },
  { key: 'place', label: '行迹' },
  { key: 'project', label: '造物' },
  { key: 'game', label: '游戏' },
  { key: 'purchase', label: '好物' },
]

const filteredEntries = computed(() => {
  let list = props.entries
  if (props.dateFilter) {
    const target = props.dateFilter.replace(/-/g, '.').trim()
    list = list.filter(e => e.date.replace(/-/g, '.').trim() === target)
  }
  if (activeFilter.value !== 'all') {
    list = list.filter(e => e.type === activeFilter.value)
  }
  if (props.limit && props.limit > 0 && !props.dateFilter) {
    return list.slice(0, props.limit)
  }
  return list
})

const getTypeName = (type: string) => {
  const map: Record<string, string> = {
    thought: '随想',
    photo: '胶卷摄影',
    coffee: '手冲咖啡',
    music: '听音',
    book: '书摘',
    place: '行迹',
    project: '造物',
    game: '游戏',
    purchase: '日常好物',
    gear: '日常装备',
    moment: '日常',
  }
  return map[type] || type
}
</script>

<template>
  <div class="w-full text-left">
    
    <!-- 克制的文字筛选栏 -->
    <div v-if="showFilters" class="flex flex-wrap items-center gap-x-5 gap-y-2 mb-14 text-xs font-mono-archive text-[var(--ink-muted)] border-b border-[var(--border-subtle)] pb-4">
      <span class="text-[var(--ink-secondary)] mr-2">分类 //</span>
      <button
        v-for="f in filters"
        :key="f.key"
        @click="activeFilter = f.key"
        class="transition-colors cursor-pointer py-1 hover:text-[var(--ink-primary)]"
        :class="{ 'text-[var(--ink-primary)] font-semibold border-b border-[var(--ink-primary)]': activeFilter === f.key }"
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
        [显示全部记录]
      </button>
    </div>

    <!-- 空状态 -->
    <div v-if="filteredEntries.length === 0" class="py-16 text-center text-sm font-serif-editorial text-[var(--ink-muted)]">
      该筛选条件下暂无生活记录。
    </div>

    <!-- 时间流列表：无 Card UI，纯粹的独立出版物杂志排版，无任何 emoji -->
    <div v-else class="divide-y divide-[var(--border-subtle)]">
      
      <article
        v-for="entry in filteredEntries"
        :key="entry.id"
        class="py-12 sm:py-16 first:pt-4"
      >
        <div class="grid grid-cols-1 md:grid-cols-12 gap-6 md:gap-12 items-start">
          
          <!-- 左侧：时间锚点与类型 -->
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

            <!-- 极简类型文本标识 -->
            <div class="pt-2">
              <span class="inline-block text-[11px] tracking-wider text-[var(--ink-muted)] border-b border-[var(--border-subtle)] pb-0.5">
                · {{ getTypeName(entry.type) }}
              </span>
            </div>

            <!-- 地点注脚 (无 emoji) -->
            <div v-if="entry.location" class="text-xs text-[var(--ink-muted)] pt-1 font-normal">
              地点：{{ entry.location }}
            </div>
          </div>

          <!-- 右侧：有机排版的内容主体 -->
          <div class="md:col-span-9 space-y-5">
            
            <!-- 标题 -->
            <h3 class="text-xl sm:text-2xl font-serif-editorial font-medium tracking-tight text-[var(--ink-primary)] leading-snug">
              {{ entry.title }}
            </h3>

            <!-- 大幅摄影排版 -->
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

            <!-- 项目/仓库链接 -->
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
