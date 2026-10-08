<script setup lang="ts">
import { ref, computed } from 'vue'
import type { Activity } from '../types'
import { Activity as ActivityIcon, GitCommit, Award, BookOpen, Rocket, Flame, GitFork, BarChart3 } from 'lucide-vue-next'

const props = defineProps<{
  activities: Activity[]
}>()

const activeHoverDay = ref<{ date: string; count: number } | null>(null)

// 真实月份标签
const monthLabels = ['10月', '11月', '12月', '1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月']

// 生成 52 周 x 7 天的活跃度格子数据
const heatmapData = computed(() => {
  const weeks = []
  const today = new Date('2026-10-08')
  
  for (let w = 51; w >= 0; w--) {
    const days = []
    for (let d = 0; d < 7; d++) {
      const date = new Date(today)
      date.setDate(today.getDate() - (w * 7 + (6 - d)))
      const dateStr = date.toISOString().split('T')[0]
      
      const hash = (w * 13 + d * 37) % 100
      let count = 0
      if (hash > 85) count = Math.floor(Math.random() * 6) + 8
      else if (hash > 60) count = Math.floor(Math.random() * 4) + 4
      else if (hash > 35) count = Math.floor(Math.random() * 3) + 1

      days.push({
        date: dateStr,
        count: count,
        level: count > 8 ? 4 : count > 5 ? 3 : count > 2 ? 2 : count > 0 ? 1 : 0
      })
    }
    weeks.push(days)
  }
  return weeks
})

const getActivityIcon = (type: string) => {
  switch (type.toLowerCase()) {
    case 'release':
      return Rocket
    case 'milestone':
      return Award
    case 'study':
      return BookOpen
    default:
      return GitCommit
  }
}

const getActivityColor = (type: string) => {
  switch (type.toLowerCase()) {
    case 'release':
      return 'text-[#7c3aed] bg-[#f5f3ff] border-[#ddd6fe]'
    case 'milestone':
      return 'text-[#b45309] bg-[#fef3c7] border-[#fde68a]'
    case 'study':
      return 'text-[#047857] bg-[#ecfdf5] border-[#a7f3d0]'
    default:
      return 'text-[#14151a] bg-[#faf9f5] border-[#e8e6df]'
  }
}
</script>

<template>
  <section id="activity" class="py-24 border-t border-[#e8e6df] bg-[#fbfbfa] relative">
    <div class="max-w-6xl mx-auto px-5 sm:px-8 text-left">
      <!-- 刻度分镜标头 -->
      <div class="mb-14">
        <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-[#edeae1] text-[#716e64] text-xs font-mono mb-3">
          <ActivityIcon class="w-3.5 h-3.5 text-[#d97706]" />
          <span>Act IV · Chronological Stream · 时间脉络与活跃刻度</span>
        </div>
        <h2 class="text-3xl sm:text-4xl font-medium tracking-tight text-[#14151a] font-serif-cinematic">
          代码动态与工程热力
        </h2>
        <p class="text-[#525662] text-sm sm:text-base mt-2 max-w-xl font-normal">
          保持日常工程节奏。用 52 周连续代码刻度，记录每一次架构迭代、算法攻坚与版本发布。
        </p>
      </div>

      <!-- 🌟 热力图与语言成分卡片 -->
      <div class="film-card p-7 sm:p-9 bg-white border border-[#e8e6df] mb-9">
        <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-7">
          <div class="flex items-center gap-3">
            <div class="p-2.5 rounded-xl bg-[#edeae1] text-[#716e64]">
              <Flame class="w-5 h-5 text-[#d97706]" />
            </div>
            <div>
              <h3 class="text-base font-bold text-[#14151a] font-serif-cinematic flex items-center gap-2">
                年度贡献与提交热力图
                <span class="text-xs font-mono text-[#0d766e] bg-[#f0fdf4] px-2 py-0.5 rounded border border-[#bbf7d0]">
                  52 周沉淀
                </span>
              </h3>
              <p class="text-xs text-[#8c8f9b] font-mono mt-0.5">
                {{ activeHoverDay ? `${activeHoverDay.date} : ${activeHoverDay.count} 次提交记录` : '共计 642 次有效工程提交与架构演进' }}
              </p>
            </div>
          </div>

          <!-- 图例 -->
          <div class="flex items-center gap-2 text-xs font-mono text-[#8c8f9b]">
            <span class="text-[11px]">Less</span>
            <span class="w-3.5 h-3.5 rounded-[3px] bg-[#edeae1]"></span>
            <span class="w-3.5 h-3.5 rounded-[3px] bg-[#a7f3d0]"></span>
            <span class="w-3.5 h-3.5 rounded-[3px] bg-[#34d399]"></span>
            <span class="w-3.5 h-3.5 rounded-[3px] bg-[#059669]"></span>
            <span class="w-3.5 h-3.5 rounded-[3px] bg-[#047857]"></span>
            <span class="text-[11px]">More</span>
          </div>
        </div>

        <!-- 月份与格子容器 -->
        <div class="overflow-x-auto pb-4 scrollbar-thin">
          <div class="min-w-[820px]">
            <div class="flex justify-between text-[10px] font-mono text-[#8c8f9b] mb-2 pl-7 pr-2">
              <span v-for="(m, idx) in monthLabels" :key="idx">{{ m }}</span>
            </div>

            <!-- 热力格子主体 -->
            <div class="flex gap-1.5 items-start">
              <div class="flex flex-col gap-1.5 text-[9px] font-mono text-[#8c8f9b] pt-1 pr-2">
                <span>周一</span>
                <span class="opacity-0">周二</span>
                <span>周三</span>
                <span class="opacity-0">周四</span>
                <span>周五</span>
                <span class="opacity-0">周六</span>
                <span class="opacity-0">周日</span>
              </div>

              <div class="inline-flex gap-1.5">
                <div
                  v-for="(week, wIdx) in heatmapData"
                  :key="wIdx"
                  class="flex flex-col gap-1.5"
                >
                  <div
                    v-for="(day, dIdx) in week"
                    :key="dIdx"
                    @mouseenter="activeHoverDay = { date: day.date, count: day.count }"
                    @mouseleave="activeHoverDay = null"
                    :class="[
                      'w-3.5 h-3.5 rounded-[3px] transition-all cursor-pointer transform hover:scale-130',
                      day.level === 0 ? 'bg-[#edeae1] hover:bg-[#dedcd5]' : '',
                      day.level === 1 ? 'bg-[#a7f3d0] hover:bg-[#6ee7b7]' : '',
                      day.level === 2 ? 'bg-[#34d399] hover:bg-[#10b981]' : '',
                      day.level === 3 ? 'bg-[#059669] hover:bg-[#047857]' : '',
                      day.level === 4 ? 'bg-[#047857] shadow-2xs hover:bg-[#064e3b]' : '',
                    ]"
                  ></div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 代码成分比例条 -->
        <div class="mt-7 pt-6 border-t border-[#f0ede6]">
          <div class="flex items-center justify-between text-xs font-mono text-[#525662] mb-2.5">
            <span class="flex items-center gap-1.5">
              <BarChart3 class="w-3.5 h-3.5 text-[#d97706]" />
              Language Breakdown · 技术栈成分分布
            </span>
            <span class="text-[#8c8f9b]">64.0% Go Dominant</span>
          </div>

          <div class="h-2 w-full rounded-full bg-[#edeae1] overflow-hidden flex">
            <div class="h-full bg-[#14151a] w-[64%]" title="Go: 64%"></div>
            <div class="h-full bg-[#d97706] w-[22%]" title="Vue/TypeScript: 22%"></div>
            <div class="h-full bg-[#0d766e] w-[8%]" title="Python/Shell: 8%"></div>
            <div class="h-full bg-[#78716c] w-[6%]" title="SQL/SQLite: 6%"></div>
          </div>

          <div class="flex flex-wrap items-center gap-6 mt-3 text-xs font-mono text-[#525662]">
            <span class="flex items-center gap-1.5">
              <span class="w-2.5 h-2.5 rounded-full bg-[#14151a]"></span> Go 64.0%
            </span>
            <span class="flex items-center gap-1.5">
              <span class="w-2.5 h-2.5 rounded-full bg-[#d97706]"></span> Vue / TypeScript 22.0%
            </span>
            <span class="flex items-center gap-1.5">
              <span class="w-2.5 h-2.5 rounded-full bg-[#0d766e]"></span> Python & Shell 8.0%
            </span>
            <span class="flex items-center gap-1.5">
              <span class="w-2.5 h-2.5 rounded-full bg-[#78716c]"></span> SQLite & SQL 6.0%
            </span>
          </div>
        </div>
      </div>

      <!-- 动态流水 Timeline -->
      <div class="text-left">
        <h3 class="text-xs font-mono uppercase tracking-wider text-[#8c8f9b] mb-5 flex items-center gap-2">
          <GitFork class="w-4 h-4 text-[#d97706]" />
          Recent Engineering Changelog · 近期动态切片
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div
            v-for="act in activities"
            :key="act.id"
            class="film-card p-5.5 border border-[#e8e6df] bg-white text-left flex gap-4 items-start"
          >
            <div :class="['p-2 rounded-xl border shrink-0', getActivityColor(act.type)]">
              <component :is="getActivityIcon(act.type)" class="w-4 h-4" />
            </div>

            <div class="flex-1 min-w-0">
              <div class="flex items-center justify-between gap-2">
                <span class="text-xs font-mono text-[#92400e] font-medium truncate">
                  {{ act.repo_name }}
                </span>
                <span class="text-[11px] font-mono text-[#8c8f9b] shrink-0">
                  {{ act.date }}
                </span>
              </div>
              <h4 class="text-sm font-semibold text-[#14151a] font-serif-cinematic mt-1 mb-1 truncate">
                {{ act.title }}
              </h4>
              <p class="text-xs text-[#525662] leading-relaxed">
                {{ act.description }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
