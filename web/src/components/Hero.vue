<script setup lang="ts">
import { computed } from 'vue'
import type { Profile, Stats } from '../types'
import { Sparkles, ArrowRight, Code2, MapPin, Terminal, Cpu } from 'lucide-vue-next'

const props = defineProps<{
  profile: Profile | null
  stats: Stats | null
}>()

const skillList = computed(() => {
  if (!props.profile?.skills) {
    return ['Go', 'Vue 3', 'TypeScript', 'Docker', 'Linux', 'SQLite', 'Gin', 'Tailwind CSS']
  }
  return props.profile.skills.split(',').map(s => s.trim())
})

const scrollToProjects = () => {
  const el = document.querySelector('#projects')
  if (el) el.scrollIntoView({ behavior: 'smooth' })
}
</script>

<template>
  <section id="hero" class="relative pt-12 pb-20 overflow-hidden">
    <!-- 背景极客网格装饰 -->
    <div class="absolute inset-0 bg-[linear-gradient(to_right,#1e293b15_1px,transparent_1px),linear-gradient(to_bottom,#1e293b15_1px,transparent_1px)] bg-[size:4rem_4rem] [mask-image:radial-gradient(ellipse_60%_50%_at_50%_0%,#000_70%,transparent_100%)] pointer-events-none"></div>

    <div class="max-w-6xl mx-auto px-4 sm:px-6 relative">
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 items-center">
        <!-- 左侧文本区域 -->
        <div class="lg:col-span-7 space-y-6 text-left">
          <!-- 标签徽章 -->
          <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-slate-900 border border-cyan-500/30 text-cyan-400 text-xs font-mono shadow-sm">
            <Sparkles class="w-3.5 h-3.5" />
            <span>迟暮的个人极客空间 · Code Activity Hub</span>
          </div>

          <!-- 主标题 -->
          <div class="space-y-2">
            <h1 class="text-4xl sm:text-5xl lg:text-6xl font-extrabold tracking-tight text-white leading-tight">
              构建优雅坚固的
              <br />
              <span class="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 via-teal-300 to-blue-500">
                工程实验与代码工坊
              </span>
            </h1>
            <p class="text-lg text-slate-300 leading-relaxed max-w-2xl pt-2">
              {{ profile?.bio || '追求极致简洁与工程美学。探索 Go 高性能后端、现代化前端架构与自动化技术，沉淀有价值的代码与思考。' }}
            </p>
          </div>

          <!-- 技能徽章墙 -->
          <div class="pt-2">
            <p class="text-xs font-mono uppercase tracking-wider text-slate-400 mb-3 flex items-center gap-1.5">
              <Cpu class="w-3.5 h-3.5 text-cyan-400" />
              Core Tech Stack · 核心技术栈
            </p>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="skill in skillList"
                :key="skill"
                class="px-3 py-1 rounded-md bg-slate-900/90 border border-slate-700/80 text-xs font-mono text-slate-200 hover:border-cyan-500/50 hover:text-cyan-300 transition-colors shadow-sm"
              >
                {{ skill }}
              </span>
            </div>
          </div>

          <!-- 动作按钮组 -->
          <div class="flex flex-wrap items-center gap-4 pt-4">
            <button
              @click="scrollToProjects"
              class="px-6 py-3 rounded-xl bg-gradient-to-r from-cyan-500 to-blue-600 hover:from-cyan-400 hover:to-blue-500 text-white font-semibold text-sm flex items-center gap-2 shadow-lg shadow-cyan-500/25 transition-all transform hover:-translate-y-0.5 cursor-pointer"
            >
              <span>浏览实验室展厅</span>
              <ArrowRight class="w-4 h-4" />
            </button>

            <a
              href="/interview"
              class="px-6 py-3 rounded-xl bg-slate-900 hover:bg-slate-800 text-slate-200 font-medium text-sm border border-slate-700 hover:border-slate-600 flex items-center gap-2 transition-all shadow-sm"
            >
              <Code2 class="w-4 h-4 text-cyan-400" />
              <span>今日 · 面试练习站</span>
              <span class="text-xs px-1.5 py-0.5 rounded bg-cyan-950 text-cyan-400 font-mono">Live</span>
            </a>
          </div>
        </div>

        <!-- 右侧个人名片卡片 -->
        <div class="lg:col-span-5">
          <div class="glass-panel rounded-2xl p-6 sm:p-8 glow-card border border-white/10 shadow-2xl relative overflow-hidden">
            <!-- 顶部装饰光圈 -->
            <div class="absolute -top-12 -right-12 w-40 h-40 bg-cyan-500/10 rounded-full blur-2xl pointer-events-none"></div>

            <div class="flex items-start gap-5">
              <div class="w-16 h-16 rounded-2xl bg-gradient-to-tr from-cyan-600 to-blue-600 p-0.5 shadow-lg shadow-cyan-500/20 shrink-0">
                <div class="w-full h-full rounded-[14px] bg-slate-900 flex items-center justify-center overflow-hidden">
                  <span class="text-2xl font-bold text-cyan-400 font-mono">迟暮</span>
                </div>
              </div>
              <div>
                <h3 class="text-xl font-bold text-white tracking-tight flex items-center gap-2">
                  {{ profile?.name || '迟暮' }}
                  <span class="text-xs font-mono px-2 py-0.5 rounded-full bg-emerald-950 text-emerald-400 border border-emerald-800">
                    Online
                  </span>
                </h3>
                <p class="text-sm text-cyan-400 font-medium mt-0.5">
                  {{ profile?.title || '全栈工程师 / Gopher / 独立开发者' }}
                </p>
                <div class="flex items-center gap-2 text-xs text-slate-400 mt-2 font-mono">
                  <MapPin class="w-3.5 h-3.5 text-slate-400" />
                  <span>{{ profile?.location || '中国 · 杭州' }}</span>
                </div>
              </div>
            </div>

            <!-- 数据指示面板 -->
            <div class="grid grid-cols-3 gap-3 my-6 py-4 px-3 rounded-xl bg-slate-950/60 border border-slate-800/80">
              <div class="text-center">
                <div class="text-2xl font-extrabold text-white font-mono">
                  {{ stats?.total_projects || 4 }}
                </div>
                <div class="text-[11px] text-slate-400 mt-0.5">开源与实验项目</div>
              </div>
              <div class="text-center border-x border-slate-800">
                <div class="text-2xl font-extrabold text-cyan-400 font-mono">
                  {{ stats?.active_days || 180 }}+
                </div>
                <div class="text-[11px] text-slate-400 mt-0.5">打卡积累天数</div>
              </div>
              <div class="text-center">
                <div class="text-2xl font-extrabold text-blue-400 font-mono">
                  100%
                </div>
                <div class="text-[11px] text-slate-400 mt-0.5">纯粹极客精神</div>
              </div>
            </div>

            <!-- 终端简况 -->
            <div class="rounded-lg bg-black/60 p-3.5 font-mono text-xs text-slate-400 space-y-1.5 border border-slate-800/60">
              <div class="flex items-center gap-2 text-slate-500">
                <Terminal class="w-3.5 h-3.5 text-emerald-400" />
                <span>chimu@codeactivityhub ~ % status</span>
              </div>
              <p class="text-slate-300">
                > Domain: <span class="text-cyan-400">codeactivityhub.top</span>
              </p>
              <p class="text-slate-300">
                > Core Stack: <span class="text-yellow-400">Go 1.27 + Vue 3 + SQLite</span>
              </p>
              <p class="text-slate-300">
                > Security: <span class="text-emerald-400">TLS 1.3 / 443 SSL Active</span>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
