<script setup lang="ts">
import type { NowStatus } from '../types'

defineProps<{
  now: NowStatus | null
}>()
</script>

<template>
  <div class="max-w-4xl mx-auto px-5 sm:px-8 py-20 sm:py-28 text-left">
    
    <!-- 标头 -->
    <header class="pb-12 border-b border-[var(--border-subtle)] mb-14 space-y-3">
      <h1 class="font-serif-editorial text-4xl sm:text-5xl font-normal text-[var(--ink-primary)]">
        NOW
      </h1>
      <p class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
        A snapshot of where I am and what I'm doing right now.
      </p>
    </header>

    <!-- 「此刻的我」清晰排版：绝无任何 Card UI，纯粹的出版物间距排版 -->
    <div class="space-y-12 max-w-2xl font-serif-editorial">
      
      <!-- Building -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
          // Currently Building
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.building || '生活自留地 · 真正记录生活的数字档案馆' }}
        </p>
      </section>

      <!-- Learning -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
          // Learning & Exploring
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.learning || '编译器底层机制、旁轴胶片冲洗、手冲浅烘水温曲线' }}
        </p>
      </section>

      <!-- Playing -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
          // Playing
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.playing || '黑神话：悟空、塞尔达传说、偶尔和朋友 Steam 连麦放空' }}
        </p>
      </section>

      <!-- Listening -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
          // Listening
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.listening || '坂本龙一 - 《async》/《BTTB》' }}
        </p>
      </section>

      <!-- Reading -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
          // Reading
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.reading || '罗伯特·M·波西格 - 《禅与摩托车维修艺术》' }}
        </p>
      </section>

      <!-- Thinking about -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
          // Thinking about
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.thinking || '如何在喧嚣的信息时代，认认真真过好具体的生活，保持安宁与自洽。' }}
        </p>
      </section>

      <!-- Using -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
          // Using
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.using || 'MacBook Pro 14, Contax T2, 纯白陶瓷滤杯, Midori MD 方格手账' }}
        </p>
      </section>

      <!-- Location -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
          // Location
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.location || '中国 · 杭州 (西湖区 / 余杭)' }}
        </p>
      </section>

      <!-- Last updated -->
      <footer class="pt-8 border-t border-[var(--border-subtle)] font-mono-archive text-xs text-[var(--ink-muted)] flex items-center justify-between">
        <span>This is a "Now page", inspired by Derek Sivers.</span>
        <span>Last updated: {{ now?.last_updated || '2026.10.09' }}</span>
      </footer>

    </div>

  </div>
</template>
