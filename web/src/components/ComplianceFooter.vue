<script setup lang="ts">
import type { SiteConfig } from '../types'
import { Shield, ExternalLink, Heart } from 'lucide-vue-next'
import { audio } from '../utils/audio'

defineProps<{
  config: SiteConfig | null
}>()
</script>

<template>
  <footer class="border-t border-[var(--border-color)] bg-[var(--bg-surface-subtle)] text-[var(--ink-secondary)] py-16 transition-colors duration-400">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-left">
      <div class="flex flex-col md:flex-row items-start md:items-center justify-between gap-8 pb-12 border-b border-[var(--border-color)]">
        
        <!-- 左侧 Logo 与工坊自白 -->
        <div>
          <div class="flex items-center gap-2.5">
            <div class="w-8 h-8 rounded-lg bg-[var(--ink-primary)] flex items-center justify-center text-[var(--bg-page)] text-xs font-serif-cinematic font-bold">
              暮
            </div>
            <span class="text-base font-bold text-[var(--ink-primary)] font-serif-cinematic tracking-wide">
              ChiMu-Lab · 迟暮的数字工坊
            </span>
          </div>
          <p class="text-xs text-[var(--ink-muted)] font-mono mt-2.5">
            Personal Workbench & Living Vignettes · Go 1.27 + Vue 3 + Pure SQLite
          </p>
        </div>

        <!-- 中间/右侧：合规备案专区（工信部 ICP 规范） -->
        <div class="flex flex-wrap items-center gap-4 text-xs font-mono">
          <!-- 工信部 ICP 备案号 -->
          <a
            :href="config?.icp_link || 'https://beian.miit.gov.cn'"
            target="_blank"
            rel="noopener noreferrer"
            @click="audio.playShutter()"
            class="flex items-center gap-2 text-[var(--ink-primary)] hover:text-[var(--accent-amber)] transition-all py-2 px-4 rounded-xl bg-[var(--bg-surface)] border border-[var(--border-color)] hover:border-[var(--accent-amber)] shadow-2xs"
          >
            <Shield class="w-3.5 h-3.5 text-[var(--accent-amber)]" />
            <span class="font-semibold">{{ config?.icp_number || '浙ICP备2026081664号' }}</span>
            <ExternalLink class="w-3 h-3 text-[var(--ink-muted)]" />
          </a>

          <!-- 全国公安机关联网备案号预留 -->
          <a
            href="http://www.beian.gov.cn/portal/registerSystemInfo"
            target="_blank"
            rel="noopener noreferrer"
            class="flex items-center gap-2 text-[var(--ink-secondary)] hover:text-[var(--ink-primary)] transition-all py-2 px-3.5 rounded-xl bg-[var(--bg-surface-subtle)] border border-[var(--border-color)]"
          >
            <span class="text-xs">👮‍♂️</span>
            <span>{{ config?.police_number || '全国公安联网备案审核中' }}</span>
          </a>

          <!-- 节点状态 -->
          <div class="flex items-center gap-1.5 text-[var(--ink-muted)] text-[11px] py-2 px-3 rounded-xl bg-[var(--bg-surface)] border border-[var(--border-color)]">
            <span class="w-1.5 h-1.5 rounded-full bg-[var(--accent-teal)]"></span>
            <span>Hangzhou CVM · TLS 1.3 Active</span>
          </div>
        </div>

      </div>

      <!-- 底部版权与胶片演职员表感 -->
      <div class="pt-8 flex flex-col sm:flex-row items-center justify-between gap-4 text-xs text-[var(--ink-muted)] font-mono text-center sm:text-left">
        <div class="flex items-center gap-1.5">
          <span>© 2024–2026 ChiMu (迟暮). Crafted with</span>
          <Heart class="w-3 h-3 text-rose-500 fill-rose-500 inline" />
          <span>in Hangzhou.</span>
        </div>
        
        <div class="flex items-center gap-5">
          <a href="https://github.com/YN1753" target="_blank" class="hover:text-[var(--ink-primary)] transition-colors">
            GitHub @YN1753
          </a>
          <span class="text-[var(--border-color)]">|</span>
          <span>Film LUT: Alabaster / Dusk / Noir</span>
          <span class="text-[var(--border-color)]">|</span>
          <span>Hangzhou, CN</span>
        </div>
      </div>
    </div>
  </footer>
</template>
