<script setup lang="ts">
import { ref, computed } from 'vue'
import type { Activity } from '../types'
import { Activity as ActivityIcon, GitCommit, Award, BookOpen, Rocket, Flame, GitFork, BarChart3 } from 'lucide-vue-next'

const props = defineProps<{
  activities: Activity[]
}>()

const activeHoverDay = ref<{ date: string; count: number } | null>(null)

// 真实月份标签（跨越 12 个月）
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
      
      // 基于伪随机产生真实开发活跃分布
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
      return 'text-purple-400 bg-purple-950/80 border-purple-850'
    case 'milestone':
      return 'text-amber-400 bg-amber-950/80 border-amber-850'
    case 'study':
      return 'text-emerald-400 bg-emerald-950/80 border-emerald-850'
    default:
      return 'text-cyan-400 bg-cyan-950/80 border-cyan-850'
  }
}
</script>

<template>
  <section id="activity" class="py-20 border-t border-white/[0.06] relative">
    <div class="max-w-6xl mx-auto px-4 sm:px-6">
      <div class="mb-12 text-left">
        <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-white/[0.04] border border-white/[0.08] text-cyan-400 text-xs font-mono mb-3">
          <ActivityIcon class="w-3.5 h-3.5" />
          <span>Code Activity Hub · 动态矩阵</span>
        </div>
        <h2 class="text-3xl sm:text-4xl font-extrabold tracking-tight text-white">
          代码动态与工程热力
        </h2>
        <p class="text-slate-400 text-sm sm:text-base mt-2 max-w-xl">
          保持日常开发节奏。记录每一行确定性提交、架构重构与算法演练。
        </p>
      </div>

      <!-- 🌟 热力图与语言分布 Bento 卡片 -->
      <div class="bento-card p-6 sm:p-8 bg-[#0a0d16] border border-white/[0.08] mb-8">
        <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-6">
          <div class="flex items-center gap-3">
            <div class="p-2.5 rounded-xl bg-emerald-950/80 border border-emerald-800/80 text-emerald-400">
              <Flame class="w-5 h-5" />
            </div>
            <div>
              <h3 class="text-base font-bold text-white flex items-center gap-2">
                年度贡献与提交热力图
                <span class="text-xs font-mono text-emerald-400 bg-emerald-950 px-2 py-0.5 rounded border border-emerald-800">
                  52 周沉淀
                </span>
              </h3>
              <p class="text-xs text-slate-400 font-mono mt-0.5">
                {{ activeHoverDay ? `${activeHoverDay.date} : ${activeHoverDay.count} 次有效提交` : '共计 642 次代码提交与工程构建' }}
              </p>
            </div>
          </div>

          <!-- 图例 -->
          <div class="flex items-center gap-2 text-xs font-mono text-slate-400">
            <span class="text-[11px]">Less</span>
            <span class="w-3 h-3 rounded-[3px] bg-slate-800/60"></span>
            <span class="w-3 h-3 rounded-[3px] bg-emerald-950 border border-emerald-850"></span>
            <span class="w-3 h-3 rounded-[3px] bg-emerald-800"></span>
            <span class="w-3 h-3 rounded-[3px] bg-emerald-600"></span>
            <span class="w-3 h-3 rounded-[3px] bg-emerald-400 shadow-sm shadow-emerald-400/50"></span>
            <span class="text-[11px]">More</span>
          </div>
        </div>

        <!-- 月份标题行 -->
        <div class="overflow-x-auto pb-4 scrollbar-thin">
          <div class="min-w-[820px]">
            <div class="flex justify-between text-[10px] font-mono text-slate-500 mb-2 pl-6 pr-2">
              <span v-for="(m, idx) in monthLabels" :key="idx">{{ m }}</span>
            </div>

            <!-- 热力格子主体 -->
            <div class="flex gap-1.5 items-start">
              <!-- 星期标签 -->
              <div class="flex flex-col gap-1.5 text-[9px] font-mono text-slate-500 pt-1 pr-1.5">
                <span>周一</span>
                <span class="opacity-0">周二</span>
                <span>周三</span>
                <span class="opacity-0">周四</span>
                <span>周五</span>
                <span class="opacity-0">周六</span>
                <span class="opacity-0">周日</span>
              </div>

              <!-- 52周格子 -->
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
                      'w-3.5 h-3.5 rounded-[3px] transition-all cursor-pointer transform hover:scale-125',
                      day.level === 0 ? 'bg-slate-800/50 hover:bg-slate-700' : '',
                      day.level === 1 ? 'bg-emerald-950 border border-emerald-900/60 hover:bg-emerald-800' : '',
                      day.level === 2 ? 'bg-emerald-800 hover:bg-emerald-700' : '',
                      day.level === 3 ? 'bg-emerald-600 hover:bg-emerald-500' : '',
                      day.level === 4 ? 'bg-emerald-400 shadow-sm shadow-emerald-400/40 hover:bg-emerald-300' : '',
                    ]"
                  ></div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 语言与技术栈分布比例条 (Language Distribution) -->
        <div class="mt-6 pt-5 border-t border-white/[0.06]">
          <div class="flex items-center justify-between text-xs font-mono text-slate-400 mb-2">
            <span class="flex items-center gap-1.5">
              <BarChart3 class="w-3.5 h-3.5 text-cyan-400" />
              Language Breakdown · 代码成分
            </span>
            <span class="text-slate-500">64% Go Dominant</span>
          </div>

          <div class="h-2 w-full rounded-full bg-slate-800 overflow-hidden flex">
            <div class="h-full bg-cyan-400 w-[64%]" title="Go: 64%"></div>
            <div class="h-full bg-emerald-400 w-[22%]" title="Vue/TypeScript: 22%"></div>
            <div class="h-full bg-amber-400 w-[8%]" title="Python/Shell: 8%"></div>
            <div class="h-full bg-purple-400 w-[6%]" title="SQL/SQLite: 6%"></div>
          </div>

          <div class="flex flex-wrap items-center gap-5 mt-3 text-xs font-mono text-slate-400">
            <span class="flex items-center gap-1.5">
              <span class="w-2 h-2 rounded-full bg-cyan-400"></span> Go 64.0%
            </span>
            <span class="flex items-center gap-1.5">
              <span class="w-2 h-2 rounded-full bg-emerald-400"></span> Vue / TypeScript 22.0%
            </span>
            <span class="flex items-center gap-1.5">
              <span class="w-2 h-2 rounded-full bg-amber-400"></span> Python & Shell 8.0%
            </span>
            <span class="flex items-center gap-1.5">
              <span class="w-2 h-2 rounded-full bg-purple-400"></span> SQLite & SQL 6.0%
            </span>
          </div>
        </div>
      </div>

      <!-- 动态脉络 Timeline -->
      <div class="text-left">
        <h3 class="text-sm font-mono uppercase tracking-wider text-slate-400 mb-4 flex items-center gap-2">
          <GitFork class="w-4 h-4 text-cyan-400" />
          Recent Engineering Changelog · 近期动态
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div
            v-for="act in activities"
            :key="act.id"
            class="bento-card p-5 border border-white/[0.06] hover:border-slate-700 transition-all text-left flex gap-4 items-start"
          >
            <div :class="['p-2 rounded-xl border shrink-0', getActivityColor(act.type)]">
              <component :is="getActivityIcon(act.type)" class="w-4 h-4" />
            </div>

            <div class="flex-1 min-w-0">
              <div class="flex items-center justify-between gap-2">
                <span class="text-xs font-mono text-cyan-400 truncate">
                  {{ act.repo_name }}
                </span>
                <span class="text-[11px] font-mono text-slate-500 shrink-0">
                  {{ act.date }}
                </span>
              </div>
              <h4 class="text-sm font-semibold text-white mt-1 mb-1 truncate">
                {{ act.title }}
              </h4>
              <p class="text-xs text-slate-300 leading-relaxed">
                {{ act.description }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
