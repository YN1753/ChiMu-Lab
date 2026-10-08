<script setup lang="ts">
import { computed } from 'vue'
import type { Profile, Stats } from '../types'
import { ArrowRight, Code2, Cpu, Zap, Compass } from 'lucide-vue-next'

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
  <section id="hero" class="relative pt-16 pb-28 cinematic-canvas">
    <div class="max-w-6xl mx-auto px-5 sm:px-8 relative text-left">
      <!-- 隐喻式分镜刻度条 (Scene Header) -->
      <div class="flex items-center justify-between text-[11px] font-mono text-[#8c8f9b] uppercase tracking-widest pb-6 border-b border-[#e8e6df] mb-12">
        <div class="flex items-center gap-2">
          <span class="w-2 h-2 rounded-full bg-[#d97706]/70"></span>
          <span>Act I · Prologue & Monograph</span>
        </div>
        <div class="hidden sm:flex items-center gap-4 text-[#716e64]">
          <span>Location: Hangzhou 30.27°N</span>
          <span>•</span>
          <span>Tone: Warm Alabaster</span>
          <span>•</span>
          <span>Shutter: 180°</span>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 items-center">
        <!-- 左侧文本海报排版 -->
        <div class="lg:col-span-7 space-y-8">
          <!-- 身份导语 -->
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-[#edeae1] text-[#716e64] text-xs font-mono">
            <Compass class="w-3.5 h-3.5 text-[#d97706]" />
            <span>迟暮的数字工坊 · ChiMu-Lab 2026</span>
          </div>

          <!-- 电影海报式大标题 -->
          <div class="space-y-4">
            <h1 class="text-4xl sm:text-5xl lg:text-6xl font-medium tracking-tight text-[#14151a] font-serif-cinematic leading-[1.18]">
              写有确定性的代码，
              <br />
              <span class="italic text-[#92400e]">
                做可自洽的工程。
              </span>
            </h1>
            <p class="text-base sm:text-lg text-[#525662] leading-relaxed max-w-2xl pt-2 font-normal">
              你好，我是 <strong class="text-[#14151a] font-semibold">迟暮 (ChiMu)</strong>。专注于 Go 高性能后端架构、现代 Web 交互与底层系统工程。在数字世界构建如胶片般具备时间厚度的实验工坊，用严谨的算法切片与架构推导，沉淀真实的工程价值。
            </p>
          </div>

          <!-- 专长标签 -->
          <div class="pt-2">
            <div class="text-[11px] font-mono uppercase tracking-wider text-[#8c8f9b] mb-3.5 flex items-center gap-1.5">
              <Cpu class="w-3.5 h-3.5 text-[#d97706]" />
              <span>Technical Repertoire · 核心专长与架构栈</span>
            </div>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="skill in skillList"
                :key="skill"
                class="px-3 py-1 rounded-lg bg-white border border-[#e8e6df] text-xs font-mono text-[#525662] hover:border-[#d97706]/40 hover:text-[#92400e] hover:bg-[#faf9f5] transition-all shadow-2xs"
              >
                {{ skill }}
              </span>
            </div>
          </div>

          <!-- 交互操作按钮 -->
          <div class="flex flex-wrap items-center gap-4 pt-4">
            <button
              @click="scrollToProjects"
              class="px-6 py-3.5 rounded-xl bg-[#14151a] hover:bg-[#252834] text-white font-medium text-xs flex items-center gap-2.5 shadow-md hover:shadow-lg transition-all transform hover:-translate-y-0.5 cursor-pointer"
            >
              <span>浏览作品分镜展厅</span>
              <ArrowRight class="w-4 h-4 text-slate-300" />
            </button>

            <a
              href="/interview"
              class="px-5 py-3.5 rounded-xl bg-white hover:bg-[#faf9f5] text-[#14151a] font-medium text-xs border border-[#e2ded4] hover:border-[#14151a]/20 flex items-center gap-2.5 transition-all shadow-2xs"
            >
              <Code2 class="w-4 h-4 text-[#d97706]" />
              <span>今日 · 面试练习站</span>
              <span class="text-[10px] px-1.5 py-0.5 rounded bg-[#f3efe4] text-[#92400e] font-mono">
                Live
              </span>
            </a>
          </div>
        </div>

        <!-- 右侧：精装画册质感系统 HUD 切片卡片 -->
        <div class="lg:col-span-5">
          <div class="film-card p-7 sm:p-8 bg-white border border-[#e8e6df] shadow-xl text-left">
            <!-- 头部 -->
            <div class="flex items-center justify-between pb-4 border-b border-[#f0ede6]">
              <div class="flex items-center gap-2.5">
                <div class="w-2 h-2 rounded-full bg-[#d97706]"></div>
                <span class="text-xs font-mono text-[#14151a] font-semibold tracking-wide">
                  SYSTEM TELEMETRY
                </span>
              </div>
              <span class="text-[11px] font-mono text-[#0d766e] bg-[#f0fdf4] px-2 py-0.5 rounded border border-[#bbf7d0] flex items-center gap-1">
                <Zap class="w-3 h-3 text-[#0d766e]" />
                Operational
              </span>
            </div>

            <!-- 数据四宫格 -->
            <div class="grid grid-cols-2 gap-3.5 my-6">
              <div class="p-3.5 rounded-xl bg-[#fbfbfa] border border-[#f0ede6]">
                <div class="text-[11px] font-mono text-[#8c8f9b]">RUNTIME</div>
                <div class="text-sm font-bold text-[#14151a] font-mono mt-1">
                  {{ stats?.go_version || 'go1.27' }}
                </div>
                <div class="text-[10px] text-[#8c8f9b] font-mono mt-0.5">
                  Linux AMD64 CVM
                </div>
              </div>

              <div class="p-3.5 rounded-xl bg-[#fbfbfa] border border-[#f0ede6]">
                <div class="text-[11px] font-mono text-[#8c8f9b]">MEMORY RSS</div>
                <div class="text-sm font-bold text-[#0d766e] font-mono mt-1">
                  {{ stats?.memory_alloc_mb ? `${stats.memory_alloc_mb} MB` : '1.76 MB' }}
                </div>
                <div class="text-[10px] text-[#8c8f9b] font-mono mt-0.5">
                  Whisper-quiet Footprint
                </div>
              </div>

              <div class="p-3.5 rounded-xl bg-[#fbfbfa] border border-[#f0ede6]">
                <div class="text-[11px] font-mono text-[#8c8f9b]">DATA ENGINE</div>
                <div class="text-xs font-bold text-[#14151a] font-mono mt-1 truncate">
                  Pure-Go SQLite
                </div>
                <div class="text-[10px] text-[#8c8f9b] font-mono mt-0.5">
                  CGO-Free · 0.2ms Latency
                </div>
              </div>

              <div class="p-3.5 rounded-xl bg-[#fbfbfa] border border-[#f0ede6]">
                <div class="text-[11px] font-mono text-[#8c8f9b]">SECURITY</div>
                <div class="text-xs font-bold text-[#14151a] font-mono mt-1">
                  TLS 1.3 / Port 443
                </div>
                <div class="text-[10px] text-[#8c8f9b] font-mono mt-0.5">
                  TrustAsia Active
                </div>
              </div>
            </div>

            <!-- 底部分镜元数据注记 -->
            <div class="p-3.5 rounded-xl bg-[#faf9f5] border border-[#e8e6df] font-mono text-[11px] text-[#525662] space-y-1.5">
              <div class="flex items-center justify-between text-[#8c8f9b] pb-1 border-b border-[#ece8df]">
                <span>PROD STATUS // VERIFIED</span>
                <span class="text-[#0d766e]">PASS</span>
              </div>
              <div class="leading-relaxed">
                <div>> Domain: <span class="text-[#14151a] font-semibold">codeactivityhub.top</span></div>
                <div>> License: <span class="text-[#92400e] font-semibold">浙ICP备2026081664号</span></div>
                <div>> Goroutines: <span class="text-[#14151a]">{{ stats?.goroutines || 4 }} concurrent</span></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
