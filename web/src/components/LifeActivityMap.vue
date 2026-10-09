<script setup lang="ts">
import { ref, computed } from 'vue'
import type { LifeEntry } from '../types'

const props = defineProps<{
  entries: LifeEntry[]
  selectedDate?: string
}>()

const emit = defineEmits<{
  (e: 'selectDate', date: string): void
}>()

const hoveredDay = ref<{
  date: string
  dayOfWeek: string
  entries: LifeEntry[]
  count: number
} | null>(null)

// 规范化日期格式辅助函数：确保 "2026.10.09" 或 "2026-10-09" 统一匹配
const normalizeDate = (d: string) => {
  return d.replace(/-/g, '.').trim()
}

// 建立日期映射表
const entryMap = computed(() => {
  const map = new Map<string, LifeEntry[]>()
  props.entries.forEach(e => {
    const rawDate = e.date || (e.occurred_at ? e.occurred_at.slice(0, 10) : '')
    if (!rawDate) return
    const key = normalizeDate(rawDate)
    const list = map.get(key) || []
    list.push(e)
    map.set(key, list)
  })
  return map
})

// 周几名称
const weekDays = ['一', '二', '三', '四', '五', '六', '日']
const weekDayLabels = ['一', '', '三', '', '五', '', '日']

interface DayCell {
  date: string
  displayDate: string
  dayOfWeekStr: string
  month: number
  inYear: boolean
  isToday: boolean
  isFuture: boolean
  entries: LifeEntry[]
  count: number
  level: number
}

// 生成 2026 年完整 52~53 周的刻度网格
const calendarWeeks = computed(() => {
  const weeks: DayCell[][] = []
  
  // 2026年1月1日是周四，当前周的周一为 2025年12月29日；补齐到2026年最后一周周日 (2027-01-03)
  const startDate = new Date(2025, 11, 29) // 2025-12-29 (周一)
  const finalDate = new Date(2027, 0, 3)

  const realNow = new Date()
  const pad = (n: number) => String(n).padStart(2, '0')
  const todayStr = `${realNow.getFullYear()}.${pad(realNow.getMonth() + 1)}.${pad(realNow.getDate())}`

  let current = new Date(startDate)
  let currentWeek: DayCell[] = []

  while (current <= finalDate) {
    const y = current.getFullYear()
    const m = current.getMonth() + 1
    const d = current.getDate()
    
    const mStr = m < 10 ? `0${m}` : `${m}`
    const dStr = d < 10 ? `0${d}` : `${d}`
    const dateStr = `${y}.${mStr}.${dStr}`
    
    const dayOfWeek = (current.getDay() + 6) % 7 // 0=周一, 6=周日
    const inYear = y === 2026
    const isToday = dateStr === todayStr
    const isFuture = dateStr > todayStr

    const matchedEntries = entryMap.value.get(dateStr) || []
    const count = matchedEntries.length

    let level = 0
    if (count === 1) level = 1
    else if (count === 2) level = 2
    else if (count >= 3) level = 3

    currentWeek.push({
      date: dateStr,
      displayDate: `${m}月${d}日`,
      dayOfWeekStr: `周${weekDays[dayOfWeek]}`,
      month: m,
      inYear,
      isToday,
      isFuture,
      entries: matchedEntries,
      count,
      level,
    })

    if (currentWeek.length === 7) {
      weeks.push(currentWeek)
      currentWeek = []
    }

    current.setDate(current.getDate() + 1)
  }

  return weeks
})

// 计算月份标签在横坐标的位置
const monthLabels = computed(() => {
  const months: { label: string; weekIndex: number }[] = []
  let lastMonth = -1

  calendarWeeks.value.forEach((week, wIdx) => {
    // 以本周第一天判断月份切换
    const firstDay = week.find(d => d.inYear)
    if (firstDay && firstDay.month !== lastMonth) {
      lastMonth = firstDay.month
      const chineseMonths = ['1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月', '11月', '12月']
      months.push({
        label: chineseMonths[lastMonth - 1],
        weekIndex: wIdx,
      })
    }
  })

  return months
})

// 活跃总统计
const totalEntriesCount = computed(() => {
  return props.entries.length
})

const activeDaysCount = computed(() => {
  return entryMap.value.size
})

const onCellClick = (cell: DayCell) => {
  if (!cell.inYear || cell.count === 0) return
  if (props.selectedDate === cell.date) {
    emit('selectDate', '')
  } else {
    emit('selectDate', cell.date)
  }
}
</script>

<template>
  <section class="py-12 border-b border-[var(--border-subtle)] text-left select-none">
    
    <!-- 标头与副标：自然出版物美学，拒绝 SaaS 卡片 -->
    <div class="flex flex-col sm:flex-row sm:items-end justify-between gap-4 pb-6">
      <div class="space-y-1.5">
        <div class="flex items-center gap-3">
          <h2 class="font-serif-editorial text-2xl sm:text-3xl font-medium tracking-tight text-[var(--ink-primary)]">
            2026 生活刻度
          </h2>
          <span class="font-mono-archive text-[11px] uppercase tracking-widest text-[var(--ink-muted)]">
            // LIFE IN 2026
          </span>
        </div>
        <p class="font-serif-editorial text-xs sm:text-sm text-[var(--ink-secondary)]">
          时间的刻度与日常印记。每一格，都是真实度过的一天。
        </p>
      </div>

      <!-- 右侧轻量统计指标 -->
      <div class="font-mono-archive text-xs text-[var(--ink-muted)] flex items-center gap-4">
        <span>已记录 <strong class="text-[var(--ink-primary)] font-medium">{{ totalEntriesCount }}</strong> 个生活瞬间</span>
        <span>·</span>
        <span>覆盖 <strong class="text-[var(--ink-primary)] font-medium">{{ activeDaysCount }}</strong> 天</span>
      </div>
    </div>

    <!-- 热力图本体容器 (支持横向自然滑动) -->
    <div class="overflow-x-auto pb-3 pt-2 scrollbar-none">
      <div class="min-w-[760px] max-w-full">
        
        <!-- 月份标头 -->
        <div class="relative h-5 mb-2 font-mono-archive text-[11px] text-[var(--ink-muted)]">
          <span
            v-for="m in monthLabels"
            :key="m.label"
            class="absolute top-0 transform"
            :style="{ left: `${m.weekIndex * 14.5 + 24}px` }"
          >
            {{ m.label }}
          </span>
        </div>

        <!-- 网格主体：左侧周几 + 右侧 53 列 -->
        <div class="flex gap-2 items-start">
          
          <!-- 周几列 -->
          <div class="flex flex-col gap-[3.5px] pt-[1px] font-mono-archive text-[10px] text-[var(--ink-muted)] shrink-0 w-4">
            <span v-for="(lbl, idx) in weekDayLabels" :key="idx" class="h-[11px] leading-[11px] block">
              {{ lbl }}
            </span>
          </div>

          <!-- 53周刻度方格 -->
          <div class="flex gap-[3.5px]">
            <div
              v-for="(week, wIdx) in calendarWeeks"
              :key="wIdx"
              class="flex flex-col gap-[3.5px]"
            >
              <button
                v-for="(day, dIdx) in week"
                :key="dIdx"
                @mouseenter="hoveredDay = { date: day.date, dayOfWeek: day.dayOfWeekStr, entries: day.entries, count: day.count }"
                @mouseleave="hoveredDay = null"
                @click="onCellClick(day)"
                :disabled="!day.inYear || day.count === 0"
                class="w-[11px] h-[11px] rounded-[1.5px] transition-all duration-150 relative cursor-default"
                :class="[
                  // 不属于2026年
                  !day.inYear ? 'opacity-0 pointer-events-none' : '',
                  // 未来日期微弱留白
                  day.inYear && day.isFuture ? 'bg-[var(--border-subtle)]/30' : '',
                  // 过去日期但无记录
                  day.inYear && !day.isFuture && day.level === 0 ? 'bg-[var(--border-subtle)]/60 hover:bg-[var(--border-divider)]' : '',
                  // 等级 1: 柔和苔绿 / 暖墨色
                  day.level === 1 ? 'bg-[#5b7a66] hover:opacity-85 cursor-pointer shadow-xs' : '',
                  // 等级 2: 中度苔绿
                  day.level === 2 ? 'bg-[#3b5946] hover:opacity-85 cursor-pointer shadow-xs' : '',
                  // 等级 3+: 浓郁沉稳墨绿
                  day.level >= 3 ? 'bg-[#22392b] hover:opacity-85 cursor-pointer shadow-xs' : '',
                  // 今日高亮圆圈
                  day.isToday ? 'outline-1 outline-[var(--ink-primary)] outline-offset-1' : '',
                  // 选中当前日期
                  selectedDate === day.date ? 'ring-2 ring-[var(--accent-warm)] ring-offset-1 scale-125 z-10' : '',
                ]"
                :title="day.inYear ? `${day.date} (${day.dayOfWeekStr}): ${day.count} 条记录` : ''"
              />
            </div>
          </div>

        </div>

      </div>
    </div>

    <!-- 底部状态与图例解释 (纯文字排版，无 AI 卡片与渐变) -->
    <div class="pt-4 flex flex-col sm:flex-row sm:items-center justify-between gap-3 text-xs font-mono-archive text-[var(--ink-muted)]">
      
      <!-- 动态悬浮/选中提示 -->
      <div class="flex items-center gap-3">
        <template v-if="selectedDate">
          <span class="text-[var(--accent-warm)] font-medium">
            正在查看 {{ selectedDate }} 的生活印记
          </span>
          <button
            @click="emit('selectDate', '')"
            class="text-[var(--ink-secondary)] hover:text-[var(--ink-primary)] hover-underline cursor-pointer"
          >
            [清除筛选]
          </button>
        </template>
        <template v-else-if="hoveredDay && hoveredDay.count > 0">
          <span class="text-[var(--ink-primary)]">
            {{ hoveredDay.date }} ({{ hoveredDay.dayOfWeek }}) · {{ hoveredDay.count }} 条生活印记：
            <span class="font-serif-editorial text-[var(--ink-secondary)]">
              {{ hoveredDay.entries.map(e => e.title).join('、') }}
            </span>
          </span>
        </template>
        <template v-else>
          <span>点击有记录的日期，可直接跳转/筛选当天时间线</span>
        </template>
      </div>

      <!-- 刻度图例 (纯中文无 emoji) -->
      <div class="flex items-center gap-2 text-[11px] self-end sm:self-auto">
        <span>留白</span>
        <span class="w-2.5 h-2.5 rounded-[1px] bg-[var(--border-subtle)]/60 inline-block"></span>
        <span class="w-2.5 h-2.5 rounded-[1px] bg-[#5b7a66] inline-block"></span>
        <span class="w-2.5 h-2.5 rounded-[1px] bg-[#3b5946] inline-block"></span>
        <span class="w-2.5 h-2.5 rounded-[1px] bg-[#22392b] inline-block"></span>
        <span>饱满</span>
      </div>

    </div>

  </section>
</template>
