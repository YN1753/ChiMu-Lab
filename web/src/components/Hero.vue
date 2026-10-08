<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import type { Profile, Stats } from '../types'
import { ArrowDown, Cpu, Sparkles, Terminal, ArrowUpRight, RefreshCw } from 'lucide-vue-next'
import { audio } from '../utils/audio'

const props = defineProps<{
  profile: Profile | null
  stats: Stats | null
}>()

// 迟暮当下状态轮盘
const currentStates = [
  { icon: '⚙️', text: '正在调优 suseoaa 的并发通道与事务', tag: 'SUSE-OAA' },
  { icon: '☕', text: '手冲一杯浅烘耶加雪菲 · 92°C 细水闷蒸', tag: 'HandDrip' },
  { icon: '📐', text: '在写 ArchCanvas 的 AST 架构剪枝生成', tag: 'ArchCanvas' },
  { icon: '🚲', text: '秋夜在西湖边骑行，感受微凉晚风', tag: 'Hangzhou' },
  { icon: '🎧', text: '在听坂本龙一《async》，终端里静默编译', tag: 'BGM' },
]

const currentStateIndex = ref(0)
const cycleState = () => {
  currentStateIndex.value = (currentStateIndex.value + 1) % currentStates.length
  audio.playShutter()
}

// 实时时间流
const currentTime = ref('')
let timer: number | null = null

const updateTime = () => {
  const d = new Date()
  currentTime.value = d.toLocaleTimeString('zh-CN', { hour12: false })
}

onMounted(() => {
  updateTime()
  timer = window.setInterval(updateTime, 1000)
})

onUnmounted(() => {
  if (timer) clearInterval(timer)
})

const skillList = computed(() => {
  if (!props.profile?.skills) {
    return ['Go', 'Wails', 'Vue 3', 'TypeScript', 'Docker', 'Linux', 'SQLite', 'Gin', 'Tailwind']
  }
  return props.profile.skills.split(',').map(s => s.trim())
})

const scrollTo = (selector: string) => {
  audio.playTink()
  const el = document.querySelector(selector)
  if (el) el.scrollIntoView({ behavior: 'smooth' })
}
</script>

<template>
  <section id="hero" class="relative pt-12 pb-24 cinematic-canvas transition-colors duration-400">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative text-left">
      
      <!-- 胶片场次打板横标 (Film Slate Bar) -->
      <div class="flex flex-wrap items-center justify-between text-[11px] font-mono text-[var(--ink-muted)] tracking-wider pb-5 border-b border-[var(--border-color)] mb-12 gap-3">
        <div class="flex items-center gap-3">
          <span class="inline-block w-2 h-2 rounded-full bg-[var(--accent-amber)] animate-pulse"></span>
          <span class="text-[var(--ink-primary)] font-semibold">SCENE: CHIMU-ATELIER</span>
          <span class="text-[var(--border-hover)]">/</span>
          <span>TAKE: 2026.FALL</span>
          <span class="text-[var(--border-hover)]">/</span>
          <span>HANGZHOU (30.27° N, 120.15° E)</span>
        </div>
        
        <div class="flex items-center gap-4">
          <span class="px-2 py-0.5 rounded bg-[var(--bg-surface)] border border-[var(--border-color)] text-[var(--accent-teal)]">
            UTC+8 {{ currentTime }}
          </span>
          <span class="hidden sm:inline">KODAK WARM TONE</span>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 lg:gap-16 items-start">
        
        <!-- 左侧：非对称电影级排版与自白 -->
        <div class="lg:col-span-7 space-y-8">
          
          <!-- 迟暮此刻正在做什么（互动轮盘） -->
          <div
            @click="cycleState"
            class="group inline-flex items-center gap-3 px-3.5 py-1.5 rounded-full bg-[var(--bg-surface)] border border-[var(--border-color)] hover:border-[var(--accent-amber)] cursor-pointer shadow-2xs transition-all duration-300"
            title="点击切换迟暮的当下状态"
          >
            <span class="text-sm">{{ currentStates[currentStateIndex].icon }}</span>
            <span class="text-xs text-[var(--ink-primary)] font-medium font-mono">
              {{ currentStates[currentStateIndex].text }}
            </span>
            <span class="text-[10px] px-1.5 py-0.5 rounded bg-[var(--bg-surface-subtle)] text-[var(--ink-muted)] font-mono border border-[var(--border-color)] group-hover:text-[var(--accent-amber)] flex items-center gap-1">
              <RefreshCw class="w-2.5 h-2.5 group-hover:rotate-180 transition-transform duration-500" />
              <span>切频</span>
            </span>
          </div>

          <!-- 电影海报主标题 -->
          <div class="space-y-4">
            <h1 class="text-4xl sm:text-5xl lg:text-6xl font-medium tracking-tight text-[var(--ink-primary)] font-serif-cinematic leading-[1.18]">
              造有骨肉的工程，
              <br />
              <span class="italic text-[var(--accent-amber)]">
                过有体温的生活。
              </span>
            </h1>
            
            <p class="text-base sm:text-lg text-[var(--ink-secondary)] leading-relaxed max-w-2xl pt-2 font-normal">
              我是 <strong class="text-[var(--ink-primary)] font-semibold">迟暮 (ChiMu)</strong>。这里不是对外宣讲的简历橱窗，而是我自己的私人暗房与数字工坊。左手记录真实写下的代码架构、底层调度与踩坑；右手收集秋天的桂花、夜雨的咖啡、胶卷底片与吉光片羽。
            </p>
          </div>

          <!-- 双轨切入交互按钮 -->
          <div class="flex flex-wrap items-center gap-4 pt-2">
            <button
              @click="scrollTo('#projects')"
              class="px-5 py-3 rounded-xl bg-[var(--ink-primary)] hover:opacity-90 text-[var(--bg-page)] text-xs font-mono font-medium transition-all shadow-sm flex items-center gap-2 group cursor-pointer"
            >
              <Terminal class="w-4 h-4 text-[var(--accent-amber)]" />
              <span>01 / 探访造物工坊 (The Craft)</span>
              <ArrowDown class="w-3.5 h-3.5 group-hover:translate-y-0.5 transition-transform" />
            </button>

            <button
              @click="scrollTo('#moments')"
              class="px-5 py-3 rounded-xl bg-[var(--bg-surface)] hover:bg-[var(--bg-surface-subtle)] border border-[var(--border-color)] text-[var(--ink-primary)] text-xs font-mono font-medium transition-all shadow-2xs flex items-center gap-2 group cursor-pointer"
            >
              <Sparkles class="w-4 h-4 text-[var(--accent-amber)]" />
              <span>02 / 翻阅生活切片 (Life Frames)</span>
              <ArrowDown class="w-3.5 h-3.5 group-hover:translate-y-0.5 transition-transform" />
            </button>
          </div>

          <!-- 常备技术手艺 -->
          <div class="pt-4 border-t border-[var(--border-color)]">
            <div class="text-[11px] font-mono uppercase tracking-wider text-[var(--ink-muted)] mb-3 flex items-center gap-2">
              <Cpu class="w-3.5 h-3.5 text-[var(--accent-amber)]" />
              <span>常用手艺 · Personal Tech Stack</span>
            </div>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="skill in skillList"
                :key="skill"
                class="px-2.5 py-1 rounded-lg bg-[var(--bg-surface)] border border-[var(--border-color)] text-xs font-mono text-[var(--ink-secondary)] hover:border-[var(--accent-amber)] hover:text-[var(--ink-primary)] transition-all shadow-2xs"
              >
                {{ skill }}
              </span>
            </div>
          </div>

        </div>

        <!-- 右侧：物理感工坊档案卡 (Atelier Dossier Specimen) -->
        <div class="lg:col-span-5">
          <div class="film-card p-6 sm:p-7 relative overflow-hidden">
            
            <!-- 胶片孔与档案戳 -->
            <div class="flex items-center justify-between pb-5 border-b border-[var(--border-color)] mb-6">
              <div class="flex items-center gap-2">
                <span class="w-2.5 h-2.5 rounded-full bg-[var(--accent-teal)]"></span>
                <span class="text-xs font-mono font-semibold text-[var(--ink-primary)] tracking-wide">
                  ATELIER SPECIMEN · 01
                </span>
              </div>
              <span class="text-[10px] font-mono px-2 py-0.5 rounded bg-[var(--bg-surface-subtle)] text-[var(--ink-muted)] border border-[var(--border-color)]">
                AUTHENTIC
              </span>
            </div>

            <!-- 当前首位置顶作品重点索引 -->
            <div class="mb-6 p-4 rounded-xl bg-[var(--bg-surface-subtle)] border border-[var(--border-color)] relative group">
              <div class="text-[10px] font-mono text-[var(--accent-amber)] font-medium mb-1 flex items-center gap-1.5">
                <span>★ 置顶核心仓库 #1</span>
                <span>·</span>
                <span>suse-edu-cn 组织项目</span>
              </div>
              <h3 class="text-base font-serif-cinematic font-semibold text-[var(--ink-primary)] group-hover:text-[var(--accent-amber)] transition-colors">
                SUSE-OAA-BACKEND
              </h3>
              <p class="text-xs text-[var(--ink-secondary)] mt-1.5 line-clamp-2">
                四川轻化工大学开放原子开源协会业务后端服务体系。Go + Gin 高性能微架构。
              </p>
              <div class="mt-3 flex items-center justify-between pt-2 border-t border-[var(--border-color)]/60 text-[11px] font-mono">
                <span class="text-[var(--ink-muted)]">Campus OpenSource</span>
                <a
                  href="https://github.com/suse-edu-cn/SUSE-OAA-BACKEND"
                  target="_blank"
                  rel="noopener noreferrer"
                  @click="audio.playShutter()"
                  class="text-[var(--ink-primary)] hover:text-[var(--accent-amber)] inline-flex items-center gap-1 font-medium"
                >
                  <span>检视仓库</span>
                  <ArrowUpRight class="w-3 h-3" />
                </a>
              </div>
            </div>

            <!-- 运行时遥测指标 -->
            <div class="space-y-3 font-mono text-xs">
              <div class="flex justify-between items-center py-1.5 border-b border-[var(--border-color)]/50">
                <span class="text-[var(--ink-muted)]">运行时架构</span>
                <span class="text-[var(--ink-primary)] font-medium">{{ stats?.go_version || 'Go 1.27 (Linux)' }}</span>
              </div>
              <div class="flex justify-between items-center py-1.5 border-b border-[var(--border-color)]/50">
                <span class="text-[var(--ink-muted)]">活动协程数</span>
                <span class="text-[var(--accent-teal)] font-medium">{{ stats?.goroutines || 8 }} Goroutines</span>
              </div>
              <div class="flex justify-between items-center py-1.5 border-b border-[var(--border-color)]/50">
                <span class="text-[var(--ink-muted)]">数据库引擎</span>
                <span class="text-[var(--ink-primary)] font-medium">Pure-Go SQLite</span>
              </div>
              <div class="flex justify-between items-center py-1.5 border-b border-[var(--border-color)]/50">
                <span class="text-[var(--ink-muted)]">查询响应时延</span>
                <span class="text-[var(--accent-amber)] font-medium">{{ stats?.query_latency_ms || 0.45 }} ms</span>
              </div>
              <div class="flex justify-between items-center py-1.5">
                <span class="text-[var(--ink-muted)]">服务连续存活</span>
                <span class="text-[var(--ink-primary)] font-medium">{{ stats?.uptime_hours || 24 }} 小时</span>
              </div>
            </div>

            <!-- 底部印章 -->
            <div class="mt-6 pt-4 border-t border-[var(--border-color)] flex items-center justify-between text-[11px] font-mono text-[var(--ink-muted)]">
              <span>CHIMU / ARCHIVED</span>
              <span class="text-[var(--ink-secondary)]">100% 真实源码足迹</span>
            </div>

          </div>
        </div>

      </div>

    </div>
  </section>
</template>
