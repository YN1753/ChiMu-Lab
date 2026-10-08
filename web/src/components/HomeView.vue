<script setup lang="ts">
import type { LifeEntry } from '../types'
import LifeStream from './LifeStream.vue'
import { ArrowRight } from 'lucide-vue-next'

defineProps<{
  entries: LifeEntry[]
}>()

const emit = defineEmits<{
  (e: 'navigate', view: string): void
}>()
</script>

<template>
  <div class="max-w-5xl mx-auto px-5 sm:px-8 text-left">
    
    <!-- 第一屏：极度克制、优雅的开篇封面 (The Monograph Opening) -->
    <section class="pt-20 sm:pt-28 pb-20 sm:pb-28 border-b border-[var(--border-subtle)]">
      <div class="max-w-3xl space-y-8">
        
        <div class="space-y-3">
          <h1 class="font-serif-editorial text-4xl sm:text-6xl lg:text-7xl font-normal tracking-tight text-[var(--ink-primary)] leading-[1.08]">
            CHIMU
          </h1>
          <p class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
            A personal archive.
          </p>
        </div>

        <div class="text-xl sm:text-2xl text-[var(--ink-secondary)] font-serif-editorial leading-relaxed max-w-xl italic pt-2">
          <p>things I've done,</p>
          <p>things I've seen,</p>
          <p>things I've thought about,</p>
          <p>and things I don't want to forget.</p>
        </div>

        <div class="pt-6 font-mono-archive text-xs text-[var(--ink-muted)] tracking-wider">
          <span>Hangzhou · 2026.10.09</span>
        </div>

      </div>
    </section>

    <!-- 下方直接切入最近的生活时间线 (RECENTLY) -->
    <section class="pt-16 sm:pt-20 pb-28">
      
      <div class="flex items-center justify-between pb-8 mb-4 border-b border-[var(--border-subtle)]">
        <h2 class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-secondary)] font-semibold">
          RECENTLY // 最近的印记
        </h2>
        <button
          @click="emit('navigate', 'life')"
          class="text-xs font-mono-archive text-[var(--ink-muted)] hover:text-[var(--ink-primary)] hover-underline cursor-pointer flex items-center gap-1"
        >
          <span>All life entries</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </button>
      </div>

      <!-- 时间流列表 -->
      <LifeStream :entries="entries" :limit="6" :showFilters="false" />

      <!-- 底部探索更多 -->
      <div class="pt-16 mt-8 border-t border-[var(--border-subtle)] flex items-center justify-between">
        <span class="text-xs font-mono-archive text-[var(--ink-muted)]">
          End of recent entries
        </span>
        <button
          @click="emit('navigate', 'life')"
          class="inline-flex items-center gap-2 text-xs font-mono-archive text-[var(--ink-primary)] hover-underline cursor-pointer"
        >
          <span>Explore complete life stream</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </button>
      </div>

    </section>

  </div>
</template>
