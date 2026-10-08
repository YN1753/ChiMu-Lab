<script setup lang="ts">
import { ref } from 'vue'
import { ArrowUpRight, CloudRain, Menu, X } from 'lucide-vue-next'
import { audio } from '../utils/audio'

const props = defineProps<{
  currentView: string
}>()

const emit = defineEmits<{
  (e: 'navigate', view: string): void
}>()

const mobileMenuOpen = ref(false)
const isRainPlaying = ref(false)

const navItems = [
  { key: 'home', label: 'HOME' },
  { key: 'life', label: 'LIFE' },
  { key: 'archive', label: 'ARCHIVE' },
  { key: 'now', label: 'NOW' },
  { key: 'projects', label: 'PROJECTS' },
  { key: 'about', label: 'ABOUT' },
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
      
      <!-- 极简文字标识 -->
      <button
        @click="setView('home')"
        class="text-left group cursor-pointer"
      >
        <span class="font-serif-editorial text-lg tracking-widest text-[var(--ink-primary)] font-medium">
          CHIMU
        </span>
        <span class="text-[11px] font-mono-archive text-[var(--ink-muted)] ml-2.5 hidden sm:inline tracking-wider">
          / archive
        </span>
      </button>

      <!-- 桌面端克制导航 -->
      <nav class="hidden md:flex items-center gap-7">
        <button
          v-for="item in navItems"
          :key="item.key"
          @click="setView(item.key)"
          class="text-xs font-mono-archive tracking-wider transition-colors cursor-pointer py-1"
          :class="currentView === item.key 
            ? 'text-[var(--ink-primary)] font-semibold border-b border-[var(--ink-primary)]' 
            : 'text-[var(--ink-muted)] hover:text-[var(--ink-primary)]'"
        >
          {{ item.label }}
        </button>
      </nav>

      <!-- 辅助极简小挂件：自然雨声 + 小 GitHub 链接 -->
      <div class="hidden sm:flex items-center gap-4 text-xs font-mono-archive text-[var(--ink-muted)]">
        <button
          @click="toggleRain"
          class="flex items-center gap-1.5 transition-colors cursor-pointer hover:text-[var(--ink-primary)]"
          :class="{ 'text-[var(--accent-moss)]': isRainPlaying }"
          title="自然雨声白噪音"
        >
          <CloudRain class="w-3.5 h-3.5" :class="{ 'animate-pulse': isRainPlaying }" />
          <span class="text-[11px]">{{ isRainPlaying ? 'rain on' : 'rain' }}</span>
        </button>

        <span class="text-[var(--border-divider)]">/</span>

        <a
          href="https://github.com/YN1753"
          target="_blank"
          rel="noopener noreferrer"
          class="flex items-center gap-1 hover:text-[var(--ink-primary)] transition-colors"
        >
          <span>github</span>
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
      <div class="space-y-3 font-mono-archive text-sm">
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
          <span>{{ isRainPlaying ? 'rain: on' : 'rain: off' }}</span>
        </button>
        <a href="https://github.com/YN1753" target="_blank" class="flex items-center gap-1">
          <span>github.com/YN1753</span>
          <ArrowUpRight class="w-3 h-3" />
        </a>
      </div>
    </div>
  </header>
</template>
