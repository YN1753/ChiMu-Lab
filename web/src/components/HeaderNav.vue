<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import {
  ArrowUpRight,
  ChevronDown,
  CloudRain,
  Menu,
  X,
  Plus,
} from 'lucide-vue-next'
import { audio } from '../utils/audio'
import type { ArchiveCategory } from '../types'

const props = defineProps<{
  currentView: string
  currentCategory: ArchiveCategory
}>()

const emit = defineEmits<{
  (e: 'navigate', view: string): void
  (e: 'changeCategory', category: ArchiveCategory): void
  (e: 'openAdd'): void
  (e: 'openStats'): void
  (e: 'openStorage'): void
}>()

const mobileMenuOpen = ref(false)
const isRainPlaying = ref(false)
const isDropdownOpen = ref(false)
const dropdownRef = ref<HTMLElement | null>(null)

// 页面导航（极简克制，中文编排）
const navItems = [
  { key: 'home', label: '首页' },
  { key: 'archive', label: '归档' },
  { key: 'now', label: '当下' },
  { key: 'about', label: '关于' },
]

// 我的生活功能视图选项
interface LifeViewOption {
  key: string
  label: string
  desc: string
  action: () => void
}

const lifeViewOptions = computed<LifeViewOption[]>(() => [
  {
    key: 'timeline',
    label: '时间线',
    desc: '完整生活时间流',
    action: () => {
      emit('changeCategory', 'all')
      emit('navigate', 'home')
    },
  },
  {
    key: 'transactions',
    label: '记账',
    desc: '日常收支与月度记录',
    action: () => {
      emit('navigate', 'transactions')
    },
  },
  {
    key: 'projects',
    label: '项目',
    desc: '造物、代码与长期项目',
    action: () => {
      emit('navigate', 'projects')
    },
  },
  {
    key: 'collection',
    label: '收藏',
    desc: '唱片、书籍、游戏与好物',
    action: () => {
      emit('changeCategory', 'collection')
      emit('navigate', 'home')
    },
  },
  {
    key: 'photo',
    label: '照片',
    desc: '定格光影与胶卷瞬间',
    action: () => {
      emit('changeCategory', 'photo')
      emit('navigate', 'home')
    },
  },
  {
    key: 'stats',
    label: '统计',
    desc: '生活刻度与聚合指标',
    action: () => {
      emit('openStats')
    },
  },
  {
    key: 'storage',
    label: '存储设置',
    desc: 'Cloudflare R2 连接状态',
    action: () => {
      emit('openStorage')
    },
  },
])

const setView = (view: string) => {
  audio.playTink()
  emit('navigate', view)
  mobileMenuOpen.value = false
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const handleLifeViewClick = (opt: LifeViewOption) => {
  audio.playTink()
  opt.action()
  isDropdownOpen.value = false
  mobileMenuOpen.value = false
}

const handleOpenAdd = () => {
  audio.playTink()
  emit('openAdd')
  mobileMenuOpen.value = false
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

      <!-- 右上角：[＋ 记录] + [我的生活⌄] + 雨声 + GitHub -->
      <div class="hidden sm:flex items-center gap-3 text-xs font-mono-archive text-[var(--ink-muted)]">
        
        <!-- 全局「＋」新增入口 (负责添加东西) -->
        <button
          @click="handleOpenAdd"
          class="flex items-center gap-1 text-xs py-1 px-2.5 rounded-xs border border-[var(--ink-primary)] bg-[var(--ink-primary)] text-[var(--bg-archive)] hover:opacity-90 transition-opacity cursor-pointer shadow-xs"
          title="记录一点什么 (生活记录 / 想法 / 照片 / 项目 / 收藏 / 记账)"
        >
          <Plus class="w-3.5 h-3.5" />
          <span class="font-serif-editorial text-[13px]">记录</span>
        </button>

        <!-- 我的生活⌄ 二级功能视图切换入口 (负责查看东西) -->
        <div class="relative" ref="dropdownRef">
          <button
            @click.stop="isDropdownOpen = !isDropdownOpen"
            class="flex items-center gap-1.5 text-xs tracking-wider transition-colors cursor-pointer py-1 px-2.5 rounded-xs border border-[var(--border-subtle)] hover:border-[var(--border-divider)] bg-[var(--bg-archive)]"
            :class="[
              isDropdownOpen || currentView === 'transactions' || currentCategory !== 'all'
                ? 'text-[var(--ink-primary)] font-medium border-[var(--border-divider)]' 
                : 'text-[var(--ink-secondary)] hover:text-[var(--ink-primary)]'
            ]"
            title="查看我的生活维度"
          >
            <span class="font-serif-editorial text-[13px]">我的生活</span>
            <ChevronDown class="w-3 h-3 opacity-60 transition-transform duration-200" :class="{ 'rotate-180': isDropdownOpen }" />
          </button>

          <!-- 简洁、精致、克制的二级浮层菜单 (Editorial 风格) -->
          <div
            v-if="isDropdownOpen"
            class="absolute right-0 top-full mt-2 w-48 bg-[var(--bg-archive)] border border-[var(--border-subtle)] shadow-[0_8px_30px_rgba(0,0,0,0.08)] py-2 z-50 rounded-xs animate-in fade-in duration-150"
            @click.stop
          >
            <div class="px-3 pb-1.5 mb-1 text-[10px] font-mono-archive tracking-widest text-[var(--ink-muted)] border-b border-[var(--border-subtle)]/60">
              生活视图 // VIEWS
            </div>

            <div class="space-y-0.5">
              <button
                v-for="opt in lifeViewOptions"
                :key="opt.key"
                @click="handleLifeViewClick(opt)"
                class="w-full text-left px-3 py-2 text-xs font-serif-editorial transition-colors flex items-center justify-between cursor-pointer group hover:bg-[var(--bg-subtle)]/60"
                :class="[
                  (opt.key === 'transactions' && currentView === 'transactions') ||
                  (opt.key === 'projects' && currentView === 'projects') ||
                  (opt.key === 'collection' && currentCategory === 'collection') ||
                  (opt.key === 'photo' && currentCategory === 'photo') ||
                  (opt.key === 'timeline' && currentView === 'home' && currentCategory === 'all')
                    ? 'text-[var(--ink-primary)] font-semibold' 
                    : 'text-[var(--ink-secondary)] hover:text-[var(--ink-primary)]'
                ]"
              >
                <div>
                  <div class="text-xs">{{ opt.label }}</div>
                  <div class="text-[10px] text-[var(--ink-muted)] font-mono-archive mt-0.5">{{ opt.desc }}</div>
                </div>
                <span
                  v-if="
                    (opt.key === 'transactions' && currentView === 'transactions') ||
                    (opt.key === 'projects' && currentView === 'projects') ||
                    (opt.key === 'collection' && currentCategory === 'collection') ||
                    (opt.key === 'photo' && currentCategory === 'photo') ||
                    (opt.key === 'timeline' && currentView === 'home' && currentCategory === 'all')
                  "
                  class="w-1.5 h-1.5 rounded-full bg-[var(--ink-primary)] shrink-0 ml-2"
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

      <!-- 移动端右上角：＋ 与 菜单按钮 -->
      <div class="flex items-center gap-2 md:hidden">
        <button
          @click="handleOpenAdd"
          class="p-1.5 bg-[var(--ink-primary)] text-[var(--bg-archive)] rounded-xs cursor-pointer shadow-xs"
          title="新建记录"
        >
          <Plus class="w-4 h-4" />
        </button>
        <button
          @click="mobileMenuOpen = !mobileMenuOpen"
          class="p-1.5 text-[var(--ink-primary)] cursor-pointer"
        >
          <Menu v-if="!mobileMenuOpen" class="w-5 h-5" />
          <X v-else class="w-5 h-5" />
        </button>
      </div>

    </div>

    <!-- 移动端展开抽屉 -->
    <div
      v-if="mobileMenuOpen"
      class="md:hidden border-b border-[var(--border-subtle)] bg-[var(--bg-archive)] px-6 py-6 space-y-6"
    >
      <!-- 快捷新增 -->
      <div>
        <button
          @click="handleOpenAdd"
          class="w-full py-2.5 px-3 bg-[var(--ink-primary)] text-[var(--bg-archive)] rounded-xs text-xs font-mono-archive flex items-center justify-center gap-2 cursor-pointer"
        >
          <Plus class="w-4 h-4" />
          <span class="font-serif-editorial text-sm">记录一点什么</span>
        </button>
      </div>

      <!-- 我的生活视图 -->
      <div class="space-y-2">
        <span class="text-[11px] font-mono-archive tracking-widest text-[var(--ink-muted)] block">
          我的生活 //
        </span>
        <div class="grid grid-cols-2 gap-2 text-sm font-serif-editorial">
          <button
            v-for="opt in lifeViewOptions"
            :key="opt.key"
            @click="handleLifeViewClick(opt)"
            class="text-left py-2 px-2.5 rounded-xs border text-xs flex items-center justify-between"
            :class="[
              (opt.key === 'transactions' && currentView === 'transactions') ||
              (opt.key === 'projects' && currentView === 'projects') ||
              (opt.key === 'collection' && currentCategory === 'collection') ||
              (opt.key === 'photo' && currentCategory === 'photo') ||
              (opt.key === 'timeline' && currentView === 'home' && currentCategory === 'all')
                ? 'border-[var(--ink-primary)] text-[var(--ink-primary)] font-bold' 
                : 'border-[var(--border-subtle)] text-[var(--ink-muted)]'
            ]"
          >
            <span>{{ opt.label }}</span>
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
        <button @click="toggleRain" class="flex items-center gap-1.5 cursor-pointer">
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
