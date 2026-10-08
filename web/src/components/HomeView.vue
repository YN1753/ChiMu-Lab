<script setup lang="ts">
import { ref } from 'vue'
import type { LifeEntry, NowStatus } from '../types'
import LifeActivityMap from './LifeActivityMap.vue'
import LifeStream from './LifeStream.vue'
import { ArrowRight } from 'lucide-vue-next'

defineProps<{
  entries: LifeEntry[]
  now?: NowStatus | null
}>()

const emit = defineEmits<{
  (e: 'navigate', view: string): void
}>()

const selectedDate = ref<string>('')

const handleDateSelect = (date: string) => {
  selectedDate.value = date
  if (date) {
    const el = document.getElementById('timeline-section')
    if (el) {
      el.scrollIntoView({ behavior: 'smooth' })
    }
  }
}
</script>

<template>
  <div class="max-w-5xl mx-auto px-5 sm:px-8 text-left">
    
    <!-- 第一屏：极度克制、优雅的开篇封面 (纯中文典雅排版) -->
    <section class="pt-20 sm:pt-28 pb-16 sm:pb-20 border-b border-[var(--border-subtle)]">
      <div class="max-w-3xl space-y-8">
        
        <div class="space-y-3">
          <h1 class="font-serif-editorial text-4xl sm:text-6xl lg:text-7xl font-normal tracking-tight text-[var(--ink-primary)] leading-[1.08]">
            迟暮
          </h1>
          <p class="font-serif-editorial text-sm sm:text-base tracking-widest text-[var(--ink-muted)]">
            一个属于我自己的数字生活档案馆。
          </p>
        </div>

        <div class="text-xl sm:text-2xl text-[var(--ink-secondary)] font-serif-editorial leading-relaxed max-w-xl pt-2 space-y-1">
          <p>我做过的事，</p>
          <p>我看过的风景，</p>
          <p>我思考过的念头，</p>
          <p>和我不想遗忘的瞬间。</p>
        </div>

        <div class="pt-4 font-mono-archive text-xs text-[var(--ink-muted)] tracking-wider">
          <span>杭州 · 2026年10月09日</span>
        </div>

      </div>
    </section>

    <!-- 首页第一视觉焦点：2026 生活刻度热力图 (Life Activity / Life in 2026) -->
    <LifeActivityMap
      :entries="entries"
      :selectedDate="selectedDate"
      @selectDate="handleDateSelect"
    />

    <!-- 当下步调 // NOW (轻量自然的人性化剪影，突出“我作为一个人”) -->
    <section v-if="now" class="py-12 border-b border-[var(--border-subtle)]">
      <div class="flex flex-col sm:flex-row sm:items-baseline justify-between gap-4 pb-6">
        <div class="flex items-center gap-3">
          <h2 class="font-serif-editorial text-xl sm:text-2xl font-medium tracking-tight text-[var(--ink-primary)]">
            当下步调
          </h2>
          <span class="font-mono-archive text-[11px] uppercase tracking-widest text-[var(--ink-muted)]">
            // NOW
          </span>
        </div>
        <button
          @click="emit('navigate', 'now')"
          class="text-xs font-mono-archive text-[var(--ink-muted)] hover:text-[var(--ink-primary)] hover-underline cursor-pointer flex items-center gap-1 self-start sm:self-auto"
        >
          <span>查看完整当下状态</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </button>
      </div>

      <!-- 纯排版、无卡片、有呼吸感的状态网格 -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 text-left font-serif-editorial">
        <div class="space-y-1.5">
          <span class="font-mono-archive text-[11px] tracking-wider text-[var(--ink-muted)] block">// 正在构建</span>
          <p class="text-base text-[var(--ink-primary)] leading-snug">{{ now.building }}</p>
        </div>
        <div class="space-y-1.5">
          <span class="font-mono-archive text-[11px] tracking-wider text-[var(--ink-muted)] block">// 正在听与读</span>
          <p class="text-base text-[var(--ink-primary)] leading-snug">{{ now.listening }} · {{ now.reading }}</p>
        </div>
        <div class="space-y-1.5">
          <span class="font-mono-archive text-[11px] tracking-wider text-[var(--ink-muted)] block">// 所思所想</span>
          <p class="text-base text-[var(--ink-secondary)] leading-snug italic">“{{ now.thinking }}”</p>
        </div>
      </div>
    </section>

    <!-- 首页第二核心：生活时间流 (Timeline / Life Stream) -->
    <section id="timeline-section" class="pt-16 sm:pt-20 pb-28">
      
      <div class="flex items-center justify-between pb-8 mb-6 border-b border-[var(--border-subtle)]">
        <div>
          <h2 class="text-base sm:text-lg font-semibold tracking-wider text-[var(--ink-primary)] font-serif-editorial">
            近况印记 // 生活时间流
          </h2>
          <p class="text-xs font-mono-archive text-[var(--ink-muted)] mt-1">
            随想、摄影、听音、造物、手冲与行迹自然相融
          </p>
        </div>

        <button
          @click="emit('navigate', 'life')"
          class="text-xs text-[var(--ink-muted)] hover:text-[var(--ink-primary)] hover-underline cursor-pointer flex items-center gap-1 font-mono-archive"
        >
          <span>生活完整档案</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </button>
      </div>

      <!-- 时间流列表 (支持点击热力图聚焦某一天) -->
      <LifeStream
        :entries="entries"
        :limit="7"
        :showFilters="false"
        :dateFilter="selectedDate"
        @clearDateFilter="selectedDate = ''"
      />

      <!-- 底部探索更多 -->
      <div class="pt-16 mt-8 border-t border-[var(--border-subtle)] flex items-center justify-between">
        <span class="text-xs font-mono-archive text-[var(--ink-muted)]">
          以上为近期的生活刻度
        </span>
        <button
          @click="emit('navigate', 'life')"
          class="inline-flex items-center gap-2 text-xs text-[var(--ink-primary)] hover-underline cursor-pointer font-mono-archive"
        >
          <span>翻阅全部生活档案</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </button>
      </div>

    </section>

  </div>
</template>
