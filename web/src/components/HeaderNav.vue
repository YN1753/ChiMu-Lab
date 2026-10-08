<script setup lang="ts">
import { ref } from 'vue'
import { ArrowUpRight, CloudRain, Menu, X } from 'lucide-vue-next'
import { audio } from '../utils/audio'

defineProps<{
  currentView: string
}>()

const emit = defineEmits<{
  (e: 'navigate', view: string): void
}>()

const mobileMenuOpen = ref(false)
const isRainPlaying = ref(false)

const navItems = [
  { key: 'home', label: '首页' },
  { key: 'life', label: '生活' },
  { key: 'archive', label: '归档' },
  { key: 'now', label: '当下' },
  { key: 'projects', label: '造物' },
  { key: 'about', label: '关于' },
]

const setView = (view: string) => {
  audio.playTink()
  emit('navigate', view)
  mobileMenuOpen.value = false
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const toggleRain = () => {
  const active = audio.toggleRain()
  isRainPlaying.value = active
}
</script>

<template>
  <header class="sticky top-0 z-50 w-full bg-[var(--bg-archive)]/90 backdrop-blur-md border-b border-[var(--border-subtle)] transition-colors duration-300">
    <div class="max-w-5xl mx-auto px-5 sm:px-8 h-16 flex items-center justify-between">
      
      <!-- 极简中文标识：迟暮 · 生活档案 -->
      <button
        @click="setView('home')"
        class="text-left group cursor-pointer"
      >
        <span class="font-serif-editorial text-xl font-medium tracking-wider text-[var(--ink-primary)]">
          迟暮
        </span>
        <span class="text-xs text-[var(--ink-muted)] ml-2.5 hidden sm:inline tracking-wider font-normal">
          / 生活档案
        </span>
      </button>

      <!-- 桌面端克制纯中文导航 -->
      <nav class="hidden md:flex items-center gap-7">
        <button
          v-for="item in navItems"
          :key="item.key"
          @click="setView(item.key)"
          class="text-sm tracking-wider transition-colors cursor-pointer py-1"
          :class="currentView === item.key 
            ? 'text-[var(--ink-primary)] font-semibold border-b border-[var(--ink-primary)]' 
            : 'text-[var(--ink-muted)] hover:text-[var(--ink-primary)]'"
        >
          {{ item.label }}
        </button>
      </nav>

      <!-- 辅助极简挂件：细雨白噪音 + GitHub -->
      <div class="hidden sm:flex items-center gap-4 text-xs font-mono-archive text-[var(--ink-muted)]">
        <button
          @click="toggleRain"
          class="flex items-center gap-1.5 transition-colors cursor-pointer hover:text-[var(--ink-primary)]"
          :class="{ 'text-[var(--accent-moss)]': isRainPlaying }"
          title="自然雨声白噪音"
        >
          <CloudRain class="w-3.5 h-3.5" :class="{ 'animate-pulse': isRainPlaying }" />
          <span class="text-xs">{{ isRainPlaying ? '雨声 · 开' : '雨声' }}</span>
        </button>

        <span class="text-[var(--border-divider)]">/</span>

        <a
          href="https://github.com/YN1753"
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center gap-1 hover:text-[var(--ink-primary)] transition-colors"
        >
          <span>GitHub</span>
          <ArrowUpRight class="w-3 h-3" />
        </a>
      </div>

      <!-- 移动端按钮 -->
      <button
        @click="mobileMenuOpen = !mobileMenuOpen"
        class="md:hidden p-1.5 text-[var(--ink-primary)]"
      >
        <Menu v-if="!mobileMenuOpen" class="w-5 h-5" />
        <X v-else class="w-5 h-5" />
      </button>

    </div>

    <!-- 移动端展开 -->
    <div
      v-if="mobileMenuOpen"
      class="md:hidden border-b border-[var(--border-subtle)] bg-[var(--bg-archive)] px-6 py-6 space-y-4"
    >
      <div class="space-y-3 text-base">
        <button
          v-for="item in navItems"
          :key="item.key"
          @click="setView(item.key)"
          class="block w-full text-left py-1 text-[var(--ink-primary)]"
          :class="{ 'font-bold': currentView === item.key }"
        >
          {{ item.label }}
        </button>
      </div>

      <div class="pt-4 border-t border-[var(--border-subtle)] flex items-center justify-between text-xs font-mono-archive text-[var(--ink-muted)]">
        <button @click="toggleRain" class="flex items-center gap-1.5">
          <CloudRain class="w-3.5 h-3.5" />
          <span>{{ isRainPlaying ? '雨声：开' : '雨声：关' }}</span>
        </button>
        <a href="https://github.com/YN1753" target="_blank" class="flex items-center gap-1">
          <span>GitHub</span>
          <ArrowUpRight class="w-3 h-3" />
        </a>
      </div>
    </div>
  </header>
</template>
