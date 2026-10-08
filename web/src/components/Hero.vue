<script setup lang="ts">
import { computed } from 'vue'
import type { Profile, Stats } from '../types'
import { ArrowRight, Code2, Cpu, Terminal, Zap, Shield, Database, HardDrive, Check } from 'lucide-vue-next'

const props = defineProps<{
  profile: Profile | null
  stats: Stats | null
}>()

const skillList = computed(() => {
  if (!props.profile?.skills) {
    return ['Go (Core)', 'Goroutine', 'Vue 3', 'TypeScript', 'Docker', 'Linux', 'SQLite', 'Gin', 'Tailwind', 'Git']
  }
  return props.profile.skills.split(',').map(s => s.trim())
})

const scrollToProjects = () => {
  const el = document.querySelector('#projects')
  if (el) el.scrollIntoView({ behavior: 'smooth' })
}
</script>

<template>
  <section id="hero" class="relative pt-12 pb-24 overflow-hidden bg-ambient-grid">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 relative">
      <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 items-center">
        <!-- 左侧个人叙事与定位 -->
        <div class="lg:col-span-7 space-y-7 text-left">
          <!-- 极客身份徽章 -->
          <div class="inline-flex items-center gap-2.5 px-3.5 py-1.5 rounded-full bg-white/[0.04] border border-white/[0.08] text-xs font-mono text-cyan-300 backdrop-blur-md">
            <span class="w-1.5 h-1.5 rounded-full bg-cyan-400"></span>
            <span>ChiMu-Lab · 迟暮的个人极客工坊</span>
            <span class="text-slate-600">/</span>
            <span class="text-slate-400">codeactivityhub.top</span>
          </div>

          <!-- 主标题：有态度、有力量、有美感 -->
          <div class="space-y-3">
            <h1 class="text-4xl sm:text-5xl lg:text-6xl font-black tracking-tight text-white leading-[1.15]">
              写有确定性的代码，
              <br />
              <span class="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 via-sky-300 to-indigo-400">
                做可自洽的工程。
              </span>
            </h1>
            <p class="text-base sm:text-lg text-slate-300/90 leading-relaxed max-w-2xl pt-2 font-normal">
              你好，我是 <strong class="text-white font-semibold">迟暮 (ChiMu)</strong>。专注于 Go 高性能后端架构、现代 Web 交互与自动化工程探索。在这里，拒绝空壳模板，用真实的算法演练、系统探针与扎实架构构建自己的数字空间。
            </p>
          </div>

          <!-- 核心技能栈徽章墙 -->
          <div class="pt-1">
            <div class="text-[11px] font-mono uppercase tracking-wider text-slate-500 mb-3 flex items-center gap-1.5">
              <Cpu class="w-3.5 h-3.5 text-cyan-400" />
              <span>Core Specialization · 专长与技术沉淀</span>
            </div>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="skill in skillList"
                :key="skill"
                class="px-3 py-1 rounded-lg bg-white/[0.03] border border-white/[0.08] text-xs font-mono text-slate-300 hover:border-cyan-500/50 hover:text-cyan-300 hover:bg-white/[0.06] transition-all shadow-sm"
              >
                {{ skill }}
              </span>
            </div>
          </div>

          <!-- 交互操作按钮 -->
          <div class="flex flex-wrap items-center gap-4 pt-3">
            <button
              @click="scrollToProjects"
              class="px-6 py-3 rounded-xl bg-gradient-to-r from-cyan-500 to-blue-600 hover:from-cyan-400 hover:to-blue-500 text-white font-semibold text-sm flex items-center gap-2 shadow-lg shadow-cyan-500/20 hover:shadow-cyan-500/35 transition-all transform hover:-translate-y-0.5 cursor-pointer"
            >
              <span>进入实验室 Bento 展厅</span>
              <ArrowRight class="w-4 h-4" />
            </button>

            <a
              href="/interview"
              class="px-5 py-3 rounded-xl bg-white/[0.04] hover:bg-white/[0.08] text-slate-200 font-medium text-sm border border-white/[0.1] hover:border-white/20 flex items-center gap-2.5 transition-all backdrop-blur-md"
            >
              <Code2 class="w-4 h-4 text-cyan-400" />
              <span>今日 · 面试练习站</span>
              <span class="text-[10px] px-1.5 py-0.5 rounded bg-emerald-950/80 text-emerald-400 border border-emerald-800 font-mono">
                Live
              </span>
            </a>
          </div>
        </div>

        <!-- 右侧：实时运行监控 HUD (Live System Dashboard) -->
        <div class="lg:col-span-5">
          <div class="bento-card p-6 shadow-2xl text-left border border-white/[0.08] bg-[#0c0f18]/90">
            <!-- 头部控制栏 -->
            <div class="flex items-center justify-between pb-4 border-b border-white/[0.06]">
              <div class="flex items-center gap-2">
                <span class="w-2.5 h-2.5 rounded-full bg-rose-500/80"></span>
                <span class="w-2.5 h-2.5 rounded-full bg-amber-500/80"></span>
                <span class="w-2.5 h-2.5 rounded-full bg-emerald-500/80"></span>
                <span class="text-xs font-mono text-slate-400 ml-2">node@chimu-cvm ~ hud</span>
              </div>
              <span class="text-[11px] font-mono text-emerald-400 bg-emerald-950/80 px-2 py-0.5 rounded border border-emerald-800/80 flex items-center gap-1">
                <Zap class="w-3 h-3 text-emerald-400" />
                Active
              </span>
            </div>

            <!-- 核心数据网格 -->
            <div class="grid grid-cols-2 gap-3.5 my-5">
              <div class="p-3 rounded-xl bg-white/[0.02] border border-white/[0.05]">
                <div class="text-[11px] font-mono text-slate-400 flex items-center gap-1.5">
                  <Terminal class="w-3.5 h-3.5 text-cyan-400" />
                  <span>Go Runtime</span>
                </div>
                <div class="text-base font-bold text-white font-mono mt-1">
                  {{ stats?.go_version || 'go1.27' }}
                </div>
                <div class="text-[10px] text-slate-500 font-mono mt-0.5">
                  Linux AMD64 CVM
                </div>
              </div>

              <div class="p-3 rounded-xl bg-white/[0.02] border border-white/[0.05]">
                <div class="text-[11px] font-mono text-slate-400 flex items-center gap-1.5">
                  <HardDrive class="w-3.5 h-3.5 text-emerald-400" />
                  <span>Memory RSS</span>
                </div>
                <div class="text-base font-bold text-emerald-400 font-mono mt-1">
                  {{ stats?.memory_alloc_mb ? `${stats.memory_alloc_mb} MB` : '3.8 MB' }}
                </div>
                <div class="text-[10px] text-slate-500 font-mono mt-0.5">
                  Ultra-light Footprint
                </div>
              </div>

              <div class="p-3 rounded-xl bg-white/[0.02] border border-white/[0.05]">
                <div class="text-[11px] font-mono text-slate-400 flex items-center gap-1.5">
                  <Database class="w-3.5 h-3.5 text-indigo-400" />
                  <span>Database</span>
                </div>
                <div class="text-sm font-bold text-white font-mono mt-1 truncate">
                  Pure-Go SQLite
                </div>
                <div class="text-[10px] text-slate-500 font-mono mt-0.5">
                  Zero CGO · 0.2ms Latency
                </div>
              </div>

              <div class="p-3 rounded-xl bg-white/[0.02] border border-white/[0.05]">
                <div class="text-[11px] font-mono text-slate-400 flex items-center gap-1.5">
                  <Shield class="w-3.5 h-3.5 text-amber-400" />
                  <span>Security & TLS</span>
                </div>
                <div class="text-sm font-bold text-white font-mono mt-1">
                  TLS 1.3 / Port 443
                </div>
                <div class="text-[10px] text-slate-500 font-mono mt-0.5">
                  TrustAsia Validated
                </div>
              </div>
            </div>

            <!-- 终端命令与状态摘要 -->
            <div class="p-3.5 rounded-xl bg-black/60 border border-white/[0.06] font-mono text-xs space-y-1.5">
              <div class="flex items-center justify-between text-slate-500 text-[11px]">
                <span>$ chimu-health --verbose</span>
                <span class="text-emerald-400 flex items-center gap-1">
                  <Check class="w-3 h-3" /> All Systems Pass
                </span>
              </div>
              <div class="text-slate-400 text-[11px] leading-relaxed">
                <div>> Host: <span class="text-cyan-400">codeactivityhub.top</span></div>
                <div>> IPC: <span class="text-amber-400">浙ICP备2026081664号</span> (Approved)</div>
                <div>> Goroutines: <span class="text-purple-400">{{ stats?.goroutines || 6 }} active</span></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
