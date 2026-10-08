<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Github, Menu, X, ArrowUpRight, Volume2, VolumeX, CloudRain, Sun, Moon, Sparkles } from 'lucide-vue-next'
import { audio } from '../utils/audio'

const mobileMenuOpen = ref(false)
const currentTheme = ref<'alabaster' | 'amber' | 'noir'>('alabaster')
const isSoundOn = ref(true)
const isRainPlaying = ref(false)

const navLinks = [
  { label: '01 / 生活方方面面', en: 'LIFE CHRONICLES', href: '#life' },
  { label: '02 / 造物与手艺', en: 'THE CRAFT & CODE', href: '#craft' },
  { label: '03 / 岁月与足迹', en: 'TEMPORAL TRACE', href: '#activity' },
]

const toggleTheme = (theme: 'alabaster' | 'amber' | 'noir') => {
  currentTheme.value = theme
  document.body.classList.remove('theme-alabaster', 'theme-amber', 'theme-noir')
  if (theme !== 'alabaster') {
    document.body.classList.add(`theme-${theme}`)
  }
  localStorage.setItem('chimu-theme', theme)
  audio.playShutter()
}

const toggleSound = () => {
  isSoundOn.value = !isSoundOn.value
  audio.soundEnabled = isSoundOn.value
  if (isSoundOn.value) {
    audio.playTink()
  }
}

const toggleRain = () => {
  const active = audio.toggleRain()
  isRainPlaying.value = active
  if (active) {
    audio.playShutter()
  }
}

const scrollToSection = (href: string) => {
  mobileMenuOpen.value = false
  audio.playTink()
  const target = document.querySelector(href)
  if (target) {
    target.scrollIntoView({ behavior: 'smooth' })
  }
}

onMounted(() => {
  const savedTheme = localStorage.getItem('chimu-theme') as 'alabaster' | 'amber' | 'noir' | null
  if (savedTheme && ['alabaster', 'amber', 'noir'].includes(savedTheme)) {
    toggleTheme(savedTheme)
  }
})
</script>

<template>
  <header class="sticky top-0 z-50 w-full border-b border-[var(--border-color)] bg-[var(--bg-page)]/90 backdrop-blur-xl transition-colors duration-400">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-18 flex items-center justify-between">
      
      <!-- 品牌标识：迟暮的生活自留地 -->
      <a href="#hero" @click="audio.playShutter()" class="flex items-center gap-3.5 group">
        <div class="relative w-9 h-9 rounded-xl bg-[var(--ink-primary)] flex items-center justify-center text-[var(--bg-page)] font-serif-cinematic text-base font-semibold shadow-sm group-hover:scale-105 transition-transform duration-300">
          <span>暮</span>
          <span class="absolute -top-1 -right-1 w-2.5 h-2.5 rounded-full bg-[var(--accent-amber)] border-2 border-[var(--bg-page)]"></span>
        </div>
        <div class="text-left">
          <div class="flex items-center gap-2">
            <span class="font-bold text-base tracking-tight text-[var(--ink-primary)] font-serif-cinematic">
              迟暮 · ChiMu
            </span>
            <span class="text-[10px] px-2 py-0.5 rounded-full bg-[var(--bg-surface-subtle)] text-[var(--ink-secondary)] font-mono border border-[var(--border-color)]">
              生活志与自留地
            </span>
          </div>
          <p class="text-[11px] text-[var(--ink-muted)] font-mono tracking-wider -mt-0.5 hidden sm:block">
            记录生活的方方面面 · 杭州 (30.27° N, 120.15° E)
          </p>
        </div>
      </a>

      <!-- 中央生活章节导航 -->
      <nav class="hidden md:flex items-center gap-1 bg-[var(--bg-surface-subtle)]/80 border border-[var(--border-color)] rounded-full px-4 py-1.5 backdrop-blur-sm shadow-2xs">
        <button
          v-for="item in navLinks"
          :key="item.href"
          @click="scrollToSection(item.href)"
          class="group text-xs font-medium text-[var(--ink-secondary)] hover:text-[var(--ink-primary)] px-3.5 py-1.5 rounded-full hover:bg-[var(--bg-surface)] transition-all cursor-pointer font-mono flex items-center gap-1.5"
        >
          <span>{{ item.label }}</span>
          <span class="text-[10px] text-[var(--ink-muted)] opacity-60 group-hover:opacity-100">{{ item.en }}</span>
        </button>
      </nav>

      <!-- 电影控制台：滤镜调色 + 自然白噪音 + 触感 + 个人主页 -->
      <div class="hidden lg:flex items-center gap-3">
        
        <!-- 自然微雨白噪音 -->
        <button
          @click="toggleRain"
          class="flex items-center gap-1.5 px-2.5 py-1.5 rounded-full border border-[var(--border-color)] bg-[var(--bg-surface)] text-xs font-mono transition-all hover:border-[var(--accent-teal)]"
          :class="isRainPlaying ? 'text-[var(--accent-teal)] border-[var(--accent-teal)] shadow-xs' : 'text-[var(--ink-secondary)]'"
          title="切换秋夜微雨自然白噪音"
        >
          <CloudRain class="w-3.5 h-3.5" :class="{ 'animate-bounce': isRainPlaying }" />
          <span class="text-[11px]">{{ isRainPlaying ? 'Rain 432Hz' : 'Rain' }}</span>
          <span v-if="isRainPlaying" class="flex gap-0.5 items-end h-3 ml-0.5">
            <span class="w-0.5 h-1.5 bg-[var(--accent-teal)] animate-pulse"></span>
            <span class="w-0.5 h-3 bg-[var(--accent-teal)] animate-pulse delay-75"></span>
            <span class="w-0.5 h-2 bg-[var(--accent-teal)] animate-pulse delay-150"></span>
          </span>
        </button>

        <!-- 快门触感开关 -->
        <button
          @click="toggleSound"
          class="p-2 rounded-full border border-[var(--border-color)] bg-[var(--bg-surface)] text-[var(--ink-secondary)] hover:text-[var(--ink-primary)] transition-all"
          :title="isSoundOn ? '微音效已开启（机械快门音）' : '微音效已静音'"
        >
          <Volume2 v-if="isSoundOn" class="w-3.5 h-3.5" />
          <VolumeX v-else class="w-3.5 h-3.5 text-[var(--ink-muted)]" />
        </button>

        <!-- 电影调色 LUT 切换器 -->
        <div class="flex items-center bg-[var(--bg-surface-subtle)] p-0.5 rounded-full border border-[var(--border-color)]">
          <button
            @click="toggleTheme('alabaster')"
            class="px-2.5 py-1 rounded-full text-[11px] font-mono transition-all"
            :class="currentTheme === 'alabaster' ? 'bg-[var(--bg-surface)] text-[var(--ink-primary)] font-semibold shadow-2xs' : 'text-[var(--ink-muted)] hover:text-[var(--ink-secondary)]'"
            title="温润暖骨白"
          >
            <Sun class="w-3 h-3 inline mr-1 text-[var(--accent-amber)]" />暖白
          </button>
          <button
            @click="toggleTheme('amber')"
            class="px-2.5 py-1 rounded-full text-[11px] font-mono transition-all"
            :class="currentTheme === 'amber' ? 'bg-[var(--bg-surface)] text-[var(--ink-primary)] font-semibold shadow-2xs' : 'text-[var(--ink-muted)] hover:text-[var(--ink-secondary)]'"
            title="琥珀薄暮（Sunset Amber）"
          >
            <Sparkles class="w-3 h-3 inline mr-1 text-amber-600" />黄昏
          </button>
          <button
            @click="toggleTheme('noir')"
            class="px-2.5 py-1 rounded-full text-[11px] font-mono transition-all"
            :class="currentTheme === 'noir' ? 'bg-[var(--bg-surface)] text-[var(--ink-primary)] font-semibold shadow-2xs' : 'text-[var(--ink-muted)] hover:text-[var(--ink-secondary)]'"
            title="胶片暗房（Midnight Darkroom）"
          >
            <Moon class="w-3 h-3 inline mr-1 text-teal-400" />暗房
          </button>
        </div>

        <!-- GitHub 外部主页 -->
        <a
          href="https://github.com/YN1753"
          target="_blank"
          rel="noopener noreferrer"
          @click="audio.playShutter()"
          class="flex items-center gap-1.5 px-3 py-1.5 rounded-full bg-[var(--ink-primary)] hover:opacity-90 text-[var(--bg-page)] text-xs font-mono font-medium transition-all shadow-sm group"
        >
          <Github class="w-3.5 h-3.5" />
          <span>YN1753</span>
          <ArrowUpRight class="w-3 h-3 opacity-60 group-hover:opacity-100 transition-opacity" />
        </a>
      </div>

      <!-- 移动端汉堡按钮 -->
      <div class="flex items-center gap-2 lg:hidden">
        <button
          @click="toggleRain"
          class="p-2 rounded-full border border-[var(--border-color)] bg-[var(--bg-surface)] text-xs text-[var(--ink-secondary)]"
          :class="{ 'text-[var(--accent-teal)]': isRainPlaying }"
        >
          <CloudRain class="w-4 h-4" />
        </button>
        <button
          @click="mobileMenuOpen = !mobileMenuOpen"
          class="p-2 rounded-lg text-[var(--ink-secondary)] hover:text-[var(--ink-primary)] hover:bg-black/5"
        >
          <Menu v-if="!mobileMenuOpen" class="w-5 h-5" />
          <X v-else class="w-5 h-5" />
        </button>
      </div>

    </div>

    <!-- 移动端抽屉 -->
    <div v-if="mobileMenuOpen" class="lg:hidden border-b border-[var(--border-color)] bg-[var(--bg-page)] px-6 py-5 space-y-4 shadow-lg">
      <div class="flex items-center justify-between pb-3 border-b border-[var(--border-color)]">
        <span class="text-xs font-mono text-[var(--ink-muted)]">电影滤镜 LUT</span>
        <div class="flex items-center gap-1.5">
          <button @click="toggleTheme('alabaster')" class="px-2.5 py-1 rounded text-xs border border-[var(--border-color)] bg-[var(--bg-surface)]">暖白</button>
          <button @click="toggleTheme('amber')" class="px-2.5 py-1 rounded text-xs border border-[var(--border-color)] bg-[var(--bg-surface)]">黄昏</button>
          <button @click="toggleTheme('noir')" class="px-2.5 py-1 rounded text-xs border border-[var(--border-color)] bg-[var(--bg-surface)]">暗房</button>
        </div>
      </div>

      <div class="space-y-2">
        <button
          v-for="item in navLinks"
          :key="item.href"
          @click="scrollToSection(item.href)"
          class="block w-full text-left py-2 text-sm font-medium text-[var(--ink-secondary)] hover:text-[var(--ink-primary)] font-mono"
        >
          {{ item.label }} · {{ item.en }}
        </button>
      </div>

      <div class="pt-3 border-t border-[var(--border-color)] flex items-center justify-between">
        <a
          href="https://github.com/YN1753"
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center gap-2 text-xs font-mono text-[var(--ink-primary)]"
        >
          <Github class="w-4 h-4" />
          <span>github.com/YN1753</span>
        </a>
      </div>
    </div>
  </header>
</template>
