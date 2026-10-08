<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { ArrowUpRight, ChevronDown, CloudRain, Menu, X } from 'lucide-vue-next'
import { audio } from '../utils/audio'
import type { ArchiveCategory } from '../types'

const props = defineProps<{
  currentView: string
  currentCategory: ArchiveCategory
}>()

const emit = defineEmits<{
  (e: 'navigate', view: string): void
  (e: 'changeCategory', category: ArchiveCategory): void
}>()

const mobileMenuOpen = ref(false)
const isRainPlaying = ref(false)
const isDropdownOpen = ref(false)
const dropdownRef = ref<HTMLElement | null>(null)

// 导航栏目（极简克制，非开发者作品集栏目）
const navItems = [
  { key: 'home', label: '首页' },
  { key: 'archive', label: '归档' },
  { key: 'now', label: '当下' },
  { key: 'about', label: '关于' },
]

// 核心生活档案五重视图
const categoryOptions: { key: ArchiveCategory; label: string }[] = [
  { key: 'all', label: '全部记录' },
  { key: 'daily', label: '日常' },
  { key: 'thought', label: '想法' },
  { key: 'project', label: '项目' },
  { key: 'collection', label: '收藏' },
]

const currentCategoryLabel = computed(() => {
  const match = categoryOptions.find(o => o.key === props.currentCategory)
  return match ? match.label : '全部记录'
})

const setView = (view: string) => {
  audio.playTink()
  emit('navigate', view)
  mobileMenuOpen.value = false
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const selectCategory = (cat: ArchiveCategory) => {
  audio.playTink()
  emit('changeCategory', cat)
  isDropdownOpen.value = false
  mobileMenuOpen.value = false
  if (props.currentView !== 'home') {
    emit('navigate', 'home')
  }
}

const toggleRain = () => {
  const active = audio.toggleRain()
  isRainPlaying.value = active
}

// 点击外部自动关闭下拉浮层
const handleDocumentClick = (e: MouseEvent) => {
  if (dropdownRef.value && !dropdownRef.value.contains(e.target as Node)) {
    isDropdownOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleDocumentClick)
})

onUnmounted(() => {
  document.removeEventListener('click', handleDocumentClick)
})
</script>

<template>
  <header class="sticky top-0 z-50 w-full bg-[var(--bg-archive)]/92 backdrop-blur-md border-b border-[var(--border-subtle)] transition-colors duration-300">
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

      <!-- 右上角：二级视图菜单（全部记录⌄） + 雨声 + GitHub -->
      <div class="hidden sm:flex items-center gap-4 text-xs font-mono-archive text-[var(--ink-muted)]">
        
        <!-- 二级视图切换入口（全部记录⌄）：生活档案核心视图切换 -->
        <div class="relative" ref="dropdownRef">
          <button
            @click.stop="isDropdownOpen = !isDropdownOpen"
            class="flex items-center gap-1 text-xs tracking-wider transition-colors cursor-pointer py-1 px-2.5 rounded-sm border border-[var(--border-subtle)] hover:border-[var(--border-divider)] bg-[var(--bg-archive)]"
            :class="[
              isDropdownOpen || currentCategory !== 'all' 
                ? 'text-[var(--ink-primary)] font-medium border-[var(--border-divider)]' 
                : 'text-[var(--ink-secondary)] hover:text-[var(--ink-primary)]'
            ]"
            title="切换生活档案观察视角"
          >
            <span class="font-serif-editorial text-[13px]">{{ currentCategoryLabel }}</span>
            <ChevronDown class="w-3 h-3 opacity-60 transition-transform duration-200" :class="{ 'rotate-180': isDropdownOpen }" />
          </button>

          <!-- 简洁、精致、克制的二级浮层菜单 (Editorial 风格，零卡片堆叠) -->
          <div
            v-if="isDropdownOpen"
            class="absolute right-0 top-full mt-2 w-36 bg-[var(--bg-archive)] border border-[var(--border-subtle)] shadow-[0_4px_24px_rgba(0,0,0,0.06)] py-2 z-50 rounded-xs animate-in fade-in duration-150"
            @click.stop
          >
            <div class="px-3 pb-1.5 mb-1 text-[10px] font-mono-archive tracking-widest text-[var(--ink-muted)] border-b border-[var(--border-subtle)]/60">
              查看 // VIEW
            </div>

            <div class="space-y-0.5">
              <button
                v-for="opt in categoryOptions"
                :key="opt.key"
                @click="selectCategory(opt.key)"
                class="w-full text-left px-3 py-1.5 text-xs font-serif-editorial transition-colors flex items-center justify-between cursor-pointer group"
                :class="currentCategory === opt.key 
                  ? 'text-[var(--ink-primary)] font-semibold' 
                  : 'text-[var(--ink-secondary)] hover:text-[var(--ink-primary)]'"
              >
                <span>{{ opt.label }}</span>
                <!-- 克制的小圆点指示当前选中 -->
                <span
                  v-if="currentCategory === opt.key"
                  class="w-1.5 h-1.5 rounded-full bg-[var(--ink-primary)] shrink-0"
                />
              </button>
            </div>
          </div>
        </div>

        <span class="text-[var(--border-divider)]">/</span>

        <!-- 自然白噪音 -->
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

        <!-- 外部代码仓 -->
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
      class="md:hidden border-b border-[var(--border-subtle)] bg-[var(--bg-archive)] px-6 py-6 space-y-6"
    >
      <!-- 视图分类选择 -->
      <div class="space-y-2">
        <span class="text-[11px] font-mono-archive tracking-widest text-[var(--ink-muted)] block">
          查看视角 //
        </span>
        <div class="grid grid-cols-2 gap-2 text-sm font-serif-editorial">
          <button
            v-for="opt in categoryOptions"
            :key="opt.key"
            @click="selectCategory(opt.key)"
            class="text-left py-1.5 px-2.5 rounded-xs border text-xs flex items-center justify-between"
            :class="currentCategory === opt.key 
              ? 'border-[var(--ink-primary)] text-[var(--ink-primary)] font-bold' 
              : 'border-[var(--border-subtle)] text-[var(--ink-muted)]'"
          >
            <span>{{ opt.label }}</span>
            <span v-if="currentCategory === opt.key" class="w-1.5 h-1.5 rounded-full bg-[var(--ink-primary)]"></span>
          </button>
        </div>
      </div>

      <!-- 栏目导航 -->
      <div class="space-y-2 pt-2 border-t border-[var(--border-subtle)]">
        <span class="text-[11px] font-mono-archive tracking-widest text-[var(--ink-muted)] block">
          页面导航 //
        </span>
        <div class="space-y-2 text-sm">
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
