<script setup lang="ts">
import { computed } from 'vue'
import type { Activity } from '../types'
import { Activity as ActivityIcon, GitCommit, Award, BookOpen, Rocket, Flame, GitFork } from 'lucide-vue-next'

const props = defineProps<{
  activities: Activity[]
}>()

// 模拟生成 52 周 x 7 天的活跃度格子数据
const heatmapWeeks = computed(() => {
  const weeks = []
  for (let w = 0; w < 40; w++) {
    const days = []
    for (let d = 0; d < 7; d++) {
      // 随机权重产生自然的活跃点状分布
      const rand = Math.random()
      let level = 0
      if (rand > 0.8) level = 3
      else if (rand > 0.55) level = 2
      else if (rand > 0.35) level = 1
      days.push(level)
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
      return 'text-purple-400 bg-purple-950/80 border-purple-800'
    case 'milestone':
      return 'text-amber-400 bg-amber-950/80 border-amber-800'
    case 'study':
      return 'text-emerald-400 bg-emerald-950/80 border-emerald-800'
    default:
      return 'text-cyan-400 bg-cyan-950/80 border-cyan-800'
  }
}
</script>

<template>
  <section id="activity" class="py-16 border-t border-white/5 relative">
    <div class="max-w-6xl mx-auto px-4 sm:px-6">
      <div class="mb-10 text-left">
        <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-slate-900 border border-slate-800 text-cyan-400 text-xs font-mono mb-2">
          <ActivityIcon class="w-3.5 h-3.5" />
          <span>Code Activity Hub · 动态追踪</span>
        </div>
        <h2 class="text-3xl font-bold tracking-tight text-white">
          代码沉淀与活跃热力
        </h2>
        <p class="text-slate-400 text-sm mt-1 max-w-xl">
          记录日常工程开发、算法演练与架构里程碑。代码不断流，思考不停歇。
        </p>
      </div>

      <!-- 热力图面板 -->
      <div class="glass-panel rounded-2xl p-6 glow-card border border-white/10 mb-8 overflow-hidden">
        <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-6">
          <div class="flex items-center gap-3">
            <div class="p-2 rounded-lg bg-slate-900 border border-slate-800 text-emerald-400">
              <Flame class="w-5 h-5" />
            </div>
            <div>
              <h3 class="text-base font-bold text-white">
                年度贡献与打卡热力图
              </h3>
              <p class="text-xs text-slate-400 font-mono">
                500+ Contributions in the last year
              </p>
            </div>
          </div>

          <!-- 图例 -->
          <div class="flex items-center gap-2 text-xs font-mono text-slate-400">
            <span>Less</span>
            <span class="w-3 h-3 rounded-[2px] bg-slate-800"></span>
            <span class="w-3 h-3 rounded-[2px] bg-emerald-900"></span>
            <span class="w-3 h-3 rounded-[2px] bg-emerald-600"></span>
            <span class="w-3 h-3 rounded-[2px] bg-emerald-400"></span>
            <span>More</span>
          </div>
        </div>

        <!-- 热力方格滑动容器 -->
        <div class="overflow-x-auto pb-2 scrollbar-thin">
          <div class="inline-flex gap-1.5 min-w-[700px]">
            <div
              v-for="(week, wIdx) in heatmapWeeks"
              :key="wIdx"
              class="flex flex-col gap-1.5"
            >
              <div
                v-for="(day, dIdx) in week"
                :key="dIdx"
                :class="[
                  'w-3.5 h-3.5 rounded-[3px] transition-colors',
                  day === 0 ? 'bg-slate-800/80 hover:bg-slate-700' : '',
                  day === 1 ? 'bg-emerald-950 border border-emerald-800/50 hover:bg-emerald-800' : '',
                  day === 2 ? 'bg-emerald-700 hover:bg-emerald-600' : '',
                  day === 3 ? 'bg-emerald-400 shadow-sm shadow-emerald-400/30 hover:bg-emerald-300' : '',
                ]"
                :title="`第 ${wIdx + 1} 周 活跃等级: ${day}`"
              ></div>
            </div>
          </div>
        </div>
      </div>

      <!-- 时间线 Timeline -->
      <div class="space-y-4">
        <h3 class="text-sm font-mono uppercase tracking-wider text-slate-400 mb-4 flex items-center gap-2">
          <GitFork class="w-4 h-4 text-cyan-400" />
          Recent Activity Stream · 近期动态脉络
        </h3>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div
            v-for="act in activities"
            :key="act.id"
            class="glass-panel rounded-xl p-4.5 border border-white/5 hover:border-slate-700 transition-all text-left flex gap-4 items-start"
          >
            <div :class="['p-2 rounded-lg border shrink-0', getActivityColor(act.type)]">
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
              <h4 class="text-sm font-semibold text-white mt-0.5 mb-1 truncate">
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
