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
      <div class="flex items-center gap-3">
        <h1 class="font-serif-editorial text-4xl sm:text-5xl font-normal text-[var(--ink-primary)]">
          此刻
        </h1>
        <span class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
          // NOW
        </span>
      </div>
      <p class="font-serif-editorial text-sm sm:text-base text-[var(--ink-secondary)]">
        定格我此时此刻的生活步调、手头做的事情与心中所想。
      </p>
    </header>

    <!-- 「此刻的我」清晰排版：绝无任何 Card UI，纯粹的出版物间距排版，零 emoji -->
    <div class="space-y-12 max-w-2xl font-serif-editorial">
      
      <!-- 正在构建 -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
          // 正在构建 · BUILDING
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.building || '生活自留地 · 真正记录生活的数字档案馆' }}
        </p>
      </section>

      <!-- 在学在看 -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
          // 在学在看 · LEARNING
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.learning || '编译器底层机制、旁轴胶片冲洗、手冲浅烘水温曲线' }}
        </p>
      </section>

      <!-- 在玩游戏 -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
          // 在玩游戏 · PLAYING
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.playing || '黑神话：悟空、塞尔达传说、偶尔和朋友 Steam 连麦放空' }}
        </p>
      </section>

      <!-- 在听唱片 -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
          // 在听唱片 · LISTENING
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.listening || '坂本龙一 - 《async》/《BTTB》' }}
        </p>
      </section>

      <!-- 在读书籍 -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
          // 在读书籍 · READING
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.reading || '罗伯特·M·波西格 - 《禅与摩托车维修艺术》' }}
        </p>
      </section>

      <!-- 所思所想 -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
          // 所思所想 · THINKING
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.thinking || '如何在喧嚣的信息时代，认认真真过好具体的生活，保持安宁与自洽。' }}
        </p>
      </section>

      <!-- 常伴物件 -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
          // 常伴物件 · USING
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.using || 'MacBook Pro 14, Contax T2, 纯白陶瓷滤杯, Midori MD 方格手账' }}
        </p>
      </section>

      <!-- 身处坐标 -->
      <section class="space-y-2">
        <div class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
          // 身处坐标 · LOCATION
        </div>
        <p class="text-xl sm:text-2xl text-[var(--ink-primary)] leading-relaxed">
          {{ now?.location || '中国 · 杭州 (西湖区 / 余杭)' }}
        </p>
      </section>

      <!-- 页脚注释 -->
      <footer class="pt-8 border-t border-[var(--border-subtle)] font-mono-archive text-xs text-[var(--ink-muted)] flex flex-col sm:flex-row sm:items-center justify-between gap-2">
        <span>源自 Derek Sivers 倡导的「Now Page」理念</span>
        <span>最近更新：{{ now?.last_updated || '2026.10.09' }}</span>
      </footer>

    </div>

  </div>
</template>
