<script setup lang="ts">
import { ref, computed } from 'vue'
import type { LifeEntry } from '../types'

const props = defineProps<{
  entries: LifeEntry[]
}>()

const selectedYear = ref('2026')

// 获取所有出现的年份
const years = computed(() => {
  const set = new Set<string>()
  props.entries.forEach(e => {
    if (e.year) set.add(e.year)
    else set.add('2026')
  })
  return Array.from(set).sort((a, b) => b.localeCompare(a))
})

// 当前年份的所有条目
const yearEntries = computed(() => {
  return props.entries.filter(e => (e.year || '2026') === selectedYear.value)
})

// 轻量统计
const yearStats = computed(() => {
  const counts: Record<string, number> = {}
  yearEntries.value.forEach(e => {
    counts[e.type] = (counts[e.type] || 0) + 1
  })
  return counts
})

// 按月份分组
const monthGroups = computed(() => {
  const map = new Map<string, LifeEntry[]>()
  const monthOrder = ['12', '11', '10', '09', '08', '07', '06', '05', '04', '03', '02', '01']
  
  monthOrder.forEach(m => {
    const list = yearEntries.value.filter(e => e.month === m || e.month === parseInt(m, 10).toString())
    if (list.length > 0) {
      map.set(m, list)
    }
  })
  return map
})

const getMonthFull = (m: string) => {
  const names: Record<string, string> = {
    '10': 'OCTOBER',
    '09': 'SEPTEMBER',
    '08': 'AUGUST',
    '07': 'JULY',
    '06': 'JUNE',
    '05': 'MAY',
    '04': 'APRIL',
    '03': 'MARCH',
    '02': 'FEBRUARY',
    '01': 'JANUARY',
  }
  return names[m] || `MONTH ${m}`
}
</script>

<template>
  <div class="max-w-4xl mx-auto px-5 sm:px-8 py-20 sm:py-28 text-left">
    
    <!-- 标头 -->
    <header class="pb-10 border-b border-[var(--border-subtle)] mb-12 space-y-3">
      <h1 class="font-serif-editorial text-4xl sm:text-5xl font-normal text-[var(--ink-primary)]">
        ARCHIVE
      </h1>
      <p class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
        Chronological index of memories, moments and things made.
      </p>
    </header>

    <!-- 年份切换（文本式，非按钮 Card） -->
    <div class="flex items-center gap-6 mb-8 font-mono-archive text-sm">
      <button
        v-for="y in years"
        :key="y"
        @click="selectedYear = y"
        class="transition-colors cursor-pointer py-1"
        :class="selectedYear === y 
          ? 'text-[var(--ink-primary)] font-bold border-b-2 border-[var(--ink-primary)]' 
          : 'text-[var(--ink-muted)] hover:text-[var(--ink-primary)]'"
      >
        {{ y }}
      </button>
    </div>

    <!-- 极轻量文本统计（非 Dashboard，纯文字注脚） -->
    <div class="font-mono-archive text-xs text-[var(--ink-muted)] pb-8 mb-12 border-b border-[var(--border-subtle)]">
      <span>{{ selectedYear }} // </span>
      <span class="text-[var(--ink-secondary)] font-medium">{{ yearEntries.length }} entries</span>
      <span class="mx-2">·</span>
      <span v-for="(count, type) in yearStats" :key="type" class="mr-3">
        {{ count }} {{ type }}s
      </span>
    </div>

    <!-- 按月份沉浸式浏览时间线清单 -->
    <div class="space-y-16">
      <section
        v-for="[month, items] in monthGroups"
        :key="month"
        class="space-y-6"
      >
        <div class="font-mono-archive text-xs tracking-widest text-[var(--ink-muted)] uppercase border-b border-[var(--border-subtle)] pb-2 flex justify-between items-center">
          <span class="text-[var(--ink-secondary)] font-semibold">{{ getMonthFull(month) }}</span>
          <span>{{ items.length }} records</span>
        </div>

        <div class="divide-y divide-[var(--border-subtle)]/60">
          <div
            v-for="item in items"
            :key="item.id"
            class="py-4 flex flex-col sm:flex-row sm:items-baseline justify-between gap-2 group hover:bg-[var(--bg-subtle)]/40 px-2 rounded-sm transition-colors"
          >
            <div class="flex items-baseline gap-4">
              <span class="font-mono-archive text-xs text-[var(--ink-muted)] shrink-0 w-12">
                {{ item.month }}.{{ item.day }}
              </span>
              <span class="font-serif-editorial text-base sm:text-lg text-[var(--ink-primary)] group-hover:text-[var(--accent-warm)] transition-colors">
                {{ item.title }}
              </span>
            </div>

            <div class="flex items-center gap-3 text-xs font-mono-archive text-[var(--ink-muted)] shrink-0 pl-16 sm:pl-0">
              <span v-if="item.location" class="hidden md:inline">{{ item.location }}</span>
              <span class="text-[10px] uppercase tracking-wider text-[var(--ink-muted)]">#{{ item.type }}</span>
            </div>
          </div>
        </div>
      </section>
    </div>

  </div>
</template>
