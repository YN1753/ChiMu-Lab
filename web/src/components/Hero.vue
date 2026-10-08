<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { ArrowDown, Sparkles, MapPin, Coffee, BookOpen, Music, Film, Compass, Heart, Terminal } from 'lucide-vue-next'
import { audio } from '../utils/audio'

// 迟暮当下生活便签状态
const todayNotes = [
  { label: '所在坐标', val: '中国 · 杭州 (西湖区 / 余杭)', icon: MapPin },
  { label: '当前时令', val: '寒露微凉 · 桂花初落满青石板', icon: Sparkles },
  { label: '今日风味', val: '手冲埃塞俄比亚古吉花魁 (92°C 细水萃取)', icon: Coffee },
  { label: '正在循环', val: '坂本龙一 · 《async》/《Merry Christmas Mr. Lawrence》', icon: Music },
  { label: '案头在读', val: '罗伯特·波西格 · 《禅与摩托车维修艺术》', icon: BookOpen },
  { label: '随身镜头', val: 'Contax T2 · 38mm f/2.8 · Kodak Portra 400', icon: Film },
  { label: '闲暇造物', val: 'SUSE 协会业务后端 & Go 调度状态机', icon: Terminal },
]

// 实时流动时间
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

const scrollTo = (selector: string) => {
  audio.playTink()
  const el = document.querySelector(selector)
  if (el) el.scrollIntoView({ behavior: 'smooth' })
}
</script>

<template>
  <section id="hero" class="relative pt-10 pb-20 cinematic-canvas transition-colors duration-400">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative text-left">
      
      <!-- 场次与时令横标 -->
      <div class="flex flex-wrap items-center justify-between text-[11px] font-mono text-[var(--ink-muted)] tracking-wider pb-5 border-b border-[var(--border-color)] mb-12 gap-3">
        <div class="flex items-center gap-3">
          <span class="inline-block w-2 h-2 rounded-full bg-[var(--accent-amber)] animate-pulse"></span>
          <span class="text-[var(--ink-primary)] font-semibold">CHIMU'S LIFE JOURNAL</span>
          <span class="text-[var(--border-hover)]">/</span>
          <span>AUTUMN 2026</span>
          <span class="text-[var(--border-hover)]">/</span>
          <span>HANGZHOU (30.27° N, 120.15° E)</span>
        </div>
        
        <div class="flex items-center gap-4">
          <span class="px-2 py-0.5 rounded bg-[var(--bg-surface)] border border-[var(--border-color)] text-[var(--accent-teal)]">
            UTC+8 {{ currentTime }}
          </span>
          <span class="hidden sm:inline">温润暖骨白 · 胶片生活志</span>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-12 gap-12 lg:gap-16 items-start">
        
        <!-- 左侧：生活态度自白 -->
        <div class="lg:col-span-7 space-y-7">
          
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-[var(--bg-surface)] text-[var(--ink-secondary)] text-xs font-mono border border-[var(--border-color)] shadow-2xs">
            <Compass class="w-3.5 h-3.5 text-[var(--accent-amber)]" />
            <span>迟暮的私人生活自留地 · Life & Mind</span>
          </div>

          <!-- 主标题：生活是全方位的 -->
          <div class="space-y-4">
            <h1 class="text-4xl sm:text-5xl lg:text-6xl font-medium tracking-tight text-[var(--ink-primary)] font-serif-cinematic leading-[1.18]">
              认认真真生活，
              <br />
              <span class="italic text-[var(--accent-amber)]">
                安安静静记录。
              </span>
            </h1>
            
            <p class="text-base sm:text-lg text-[var(--ink-secondary)] leading-relaxed max-w-2xl pt-2 font-normal">
              我是 <strong class="text-[var(--ink-primary)] font-semibold">迟暮 (ChiMu)</strong>。这里是属于我自己的自留地，记录着<strong class="text-[var(--ink-primary)]">生活的方方面面</strong>——骑行路上迎面吹来的夜风、清晨手冲咖啡的香气、读过的书摘与听过的唱片、拍下的胶卷，以及偶尔亲手敲下的代码与造物。
            </p>
          </div>

          <!-- 双轨导航入口 -->
          <div class="flex flex-wrap items-center gap-4 pt-2">
            <button
              @click="scrollTo('#life')"
              class="px-5 py-3 rounded-xl bg-[var(--ink-primary)] hover:opacity-90 text-[var(--bg-page)] text-xs font-mono font-medium transition-all shadow-sm flex items-center gap-2 group cursor-pointer"
            >
              <Sparkles class="w-4 h-4 text-[var(--accent-amber)]" />
              <span>01 / 漫步生活全景 (Life Chronicles)</span>
              <ArrowDown class="w-3.5 h-3.5 group-hover:translate-y-0.5 transition-transform" />
            </button>

            <button
              @click="scrollTo('#craft')"
              class="px-5 py-3 rounded-xl bg-[var(--bg-surface)] hover:bg-[var(--bg-surface-subtle)] border border-[var(--border-color)] text-[var(--ink-primary)] text-xs font-mono font-medium transition-all shadow-2xs flex items-center gap-2 group cursor-pointer"
            >
              <Terminal class="w-4 h-4 text-[var(--accent-amber)]" />
              <span>02 / 探访造物工坊 (Craft & Code)</span>
              <ArrowDown class="w-3.5 h-3.5 group-hover:translate-y-0.5 transition-transform" />
            </button>
          </div>

          <!-- 生活棱镜标签 -->
          <div class="pt-4 border-t border-[var(--border-color)]">
            <div class="text-[11px] font-mono uppercase tracking-wider text-[var(--ink-muted)] mb-3 flex items-center gap-2">
              <Heart class="w-3.5 h-3.5 text-rose-500" />
              <span>生活的多面体 · Facets of Living</span>
            </div>
            <div class="flex flex-wrap gap-2 text-xs font-mono">
              <span class="px-2.5 py-1 rounded-lg bg-[var(--bg-surface)] border border-[var(--border-color)] text-[var(--ink-secondary)]">🚲 城市骑行</span>
              <span class="px-2.5 py-1 rounded-lg bg-[var(--bg-surface)] border border-[var(--border-color)] text-[var(--ink-secondary)]">☕ 手冲咖啡</span>
              <span class="px-2.5 py-1 rounded-lg bg-[var(--bg-surface)] border border-[var(--border-color)] text-[var(--ink-secondary)]">📷 胶卷摄影</span>
              <span class="px-2.5 py-1 rounded-lg bg-[var(--bg-surface)] border border-[var(--border-color)] text-[var(--ink-secondary)]">📖 书房阅读</span>
              <span class="px-2.5 py-1 rounded-lg bg-[var(--bg-surface)] border border-[var(--border-color)] text-[var(--ink-secondary)]">🎧 唱片音乐</span>
              <span class="px-2.5 py-1 rounded-lg bg-[var(--bg-surface)] border border-[var(--border-color)] text-[var(--ink-secondary)]">💭 深夜随想</span>
              <span class="px-2.5 py-1 rounded-lg bg-[var(--bg-surface)] border border-[var(--border-color)] text-[var(--ink-secondary)]">🎒 桌面好物</span>
              <span class="px-2.5 py-1 rounded-lg bg-[var(--bg-surface)] border border-[var(--border-color)] text-[var(--ink-secondary)]">⚙️ 造物手艺</span>
            </div>
          </div>

        </div>

        <!-- 右侧：生活便签板 (Today's Life Slate) -->
        <div class="lg:col-span-5">
          <div class="film-card p-6 sm:p-7 relative overflow-hidden bg-[var(--bg-surface)] border border-[var(--border-color)]">
            
            <div class="flex items-center justify-between pb-4 border-b border-[var(--border-color)] mb-5">
              <div class="flex items-center gap-2">
                <span class="w-2.5 h-2.5 rounded-full bg-[var(--accent-amber)]"></span>
                <span class="text-xs font-mono font-semibold text-[var(--ink-primary)] tracking-wide">
                  TODAY'S LIFE SLATE · 当下便签
                </span>
              </div>
              <span class="text-[10px] font-mono px-2 py-0.5 rounded bg-[var(--bg-surface-subtle)] text-[var(--ink-muted)] border border-[var(--border-color)]">
                AUTHENTIC
              </span>
            </div>

            <!-- 便签条目列表 -->
            <div class="space-y-3.5 text-xs font-mono">
              <div
                v-for="(item, i) in todayNotes"
                :key="i"
                class="p-2.5 rounded-xl bg-[var(--bg-surface-subtle)]/70 border border-[var(--border-color)]/60 flex items-start gap-3"
              >
                <component :is="item.icon" class="w-4 h-4 text-[var(--accent-amber)] shrink-0 mt-0.5" />
                <div class="flex-1 min-w-0">
                  <div class="text-[10px] text-[var(--ink-muted)]">{{ item.label }}</div>
                  <div class="text-[var(--ink-primary)] font-medium mt-0.5 truncate">{{ item.val }}</div>
                </div>
              </div>
            </div>

            <!-- 底部生活注脚 -->
            <div class="mt-5 pt-4 border-t border-[var(--border-color)] flex items-center justify-between text-[11px] font-mono text-[var(--ink-muted)]">
              <span>CHIMU'S MEMO</span>
              <span class="text-[var(--ink-secondary)]">生活是具体的，且充满温度</span>
            </div>

          </div>
        </div>

      </div>

    </div>
  </section>
</template>
