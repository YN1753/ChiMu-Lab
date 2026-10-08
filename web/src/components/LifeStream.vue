<script setup lang="ts">
import { ref, computed } from 'vue'
import type { LifeEntry } from '../types'
import { ArrowUpRight } from 'lucide-vue-next'

const props = defineProps<{
  entries: LifeEntry[]
  limit?: number
  showFilters?: boolean
}>()

const activeFilter = ref<string>('all')

const filters = [
  { key: 'all', label: 'ALL' },
  { key: 'thought', label: 'THOUGHTS' },
  { key: 'photo', label: 'PHOTOGRAPHY' },
  { key: 'music', label: 'MUSIC' },
  { key: 'book', label: 'BOOKS' },
  { key: 'place', label: 'PLACES' },
  { key: 'project', label: 'PROJECTS' },
  { key: 'game', label: 'GAMES' },
]

const filteredEntries = computed(() => {
  let list = props.entries
  if (activeFilter.value !== 'all') {
    list = list.filter(e => e.type === activeFilter.value)
  }
  if (props.limit && props.limit > 0) {
    return list.slice(0, props.limit)
  }
  return list
})

const getMonthName = (monthStr: string) => {
  const m = parseInt(monthStr, 10)
  const names = ['JAN', 'FEB', 'MAR', 'APR', 'MAY', 'JUN', 'JUL', 'AUG', 'SEP', 'OCT', 'NOV', 'DEC']
  return names[m - 1] || monthStr
}
</script>

<template>
  <div class="w-full text-left">
    
    <!-- 克制的文字筛选栏 (仅在显式开启时展示) -->
    <div v-if="showFilters" class="flex flex-wrap items-center gap-x-5 gap-y-2 mb-14 text-xs font-mono-archive text-[var(--ink-muted)] border-b border-[var(--border-subtle)] pb-4">
      <span class="text-[var(--ink-secondary)] mr-2">FILTER //</span>
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

    <!-- 时间流列表：绝无千篇一律的 Card UI，纯粹的独立出版物杂志排版 -->
    <div class="divide-y divide-[var(--border-subtle)]">
      
      <article
        v-for="entry in filteredEntries"
        :key="entry.id"
        class="py-12 sm:py-16 first:pt-4"
      >
        <div class="grid grid-cols-1 md:grid-cols-12 gap-6 md:gap-12 items-start">
          
          <!-- 左侧：时间与类型锚点 (Typographic Anchor) -->
          <div class="md:col-span-3 space-y-2">
            <div class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
              <span class="text-[var(--ink-secondary)] font-semibold block text-sm">{{ entry.year }}</span>
              <span class="text-base text-[var(--ink-primary)] font-serif-editorial block mt-0.5">
                {{ getMonthName(entry.month) }} {{ entry.day }}
              </span>
              <span v-if="entry.time" class="block text-[11px] text-[var(--ink-muted)] mt-1">
                {{ entry.time }}
              </span>
            </div>

            <!-- 极简类型标识 -->
            <div class="pt-2">
              <span class="inline-block text-[10px] font-mono-archive uppercase tracking-widest text-[var(--ink-muted)] border-b border-[var(--border-subtle)] pb-0.5">
                # {{ entry.type }}
              </span>
            </div>

            <!-- 地点注脚 -->
            <div v-if="entry.location" class="text-xs font-mono-archive text-[var(--ink-muted)] pt-1">
              📍 {{ entry.location }}
            </div>
          </div>

          <!-- 右侧：有机自然排版的内容主体 (根据内容形态自适应，非固定 Card) -->
          <div class="md:col-span-9 space-y-5">
            
            <!-- 标题 -->
            <h3 class="text-xl sm:text-2xl font-serif-editorial font-medium tracking-tight text-[var(--ink-primary)] leading-snug">
              {{ entry.title }}
            </h3>

            <!-- 大幅摄影排版 (Large Editorial Photo) -->
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

            <!-- 正文叙事 (宽松行距与呼吸感) -->
            <p class="text-base sm:text-lg text-[var(--ink-secondary)] leading-relaxed font-normal whitespace-pre-line max-w-2xl">
              {{ entry.content }}
            </p>

            <!-- 附加信息与链接 (音乐专辑 / 项目仓库 / 书籍作者 / 随笔参数) -->
            <div v-if="entry.meta && !entry.images" class="pt-2 flex items-center gap-3 text-xs font-mono-archive text-[var(--ink-muted)]">
              <span>— {{ entry.meta }}</span>
            </div>

            <div v-if="entry.link" class="pt-2">
              <a
                :href="entry.link"
                target="_blank"
                rel="noopener noreferrer"
                class="inline-flex items-center gap-1.5 text-xs font-mono-archive text-[var(--ink-primary)] hover-underline pb-0.5"
              >
                <span>View project / repository</span>
                <ArrowUpRight class="w-3.5 h-3.5" />
              </a>
            </div>

          </div>

        </div>
      </article>

    </div>

  </div>
</template>
