<script setup lang="ts">
import type { SiteConfig } from '../types'
import { Shield, ExternalLink, Heart, Terminal } from 'lucide-vue-next'

defineProps<{
  config: SiteConfig | null
}>()
</script>

<template>
  <footer class="mt-20 border-t border-white/10 bg-slate-950/90 text-slate-400 py-12">
    <div class="max-w-6xl mx-auto px-4 sm:px-6">
      <div class="flex flex-col md:flex-row items-center justify-between gap-6 pb-8 border-b border-white/5">
        <!-- 左侧 Logo & Slogan -->
        <div class="flex items-center gap-3">
          <div class="w-8 h-8 rounded-lg bg-cyan-600/30 border border-cyan-500/40 flex items-center justify-center text-cyan-400">
            <Terminal class="w-4 h-4" />
          </div>
          <div>
            <span class="text-sm font-bold text-white tracking-wide">
              ChiMu-Lab 迟暮实验室
            </span>
            <p class="text-xs text-slate-500 font-mono">
              Code Activity Hub · Powered by Go 1.27 + Vue 3 + SQLite
            </p>
          </div>
        </div>

        <!-- 备案合规专区（审核通过关键） -->
        <div class="flex flex-wrap items-center justify-center gap-6 text-xs font-mono">
          <!-- 工信部 ICP 备案号 -->
          <a
            :href="config?.icp_link || 'https://beian.miit.gov.cn'"
            target="_blank"
            rel="noopener noreferrer"
            class="flex items-center gap-1.5 text-slate-300 hover:text-cyan-400 transition-colors py-1 px-2.5 rounded bg-slate-900 border border-slate-800 hover:border-cyan-500/40"
          >
            <Shield class="w-3.5 h-3.5 text-cyan-400" />
            <span>{{ config?.icp_number || '苏ICP备2024000000号-1' }}</span>
            <ExternalLink class="w-3 h-3 text-slate-500" />
          </a>

          <!-- 全国公安联网备案号预留 -->
          <a
            href="http://www.beian.gov.cn/portal/registerSystemInfo"
            target="_blank"
            rel="noopener noreferrer"
            class="flex items-center gap-1.5 text-slate-400 hover:text-slate-200 transition-colors py-1 px-2.5 rounded bg-slate-900/60 border border-slate-800"
          >
            <span class="text-xs">👮‍♂️</span>
            <span>{{ config?.police_number || '全国公安联网备案审核中' }}</span>
          </a>
        </div>
      </div>

      <!-- 底部版权与免责声明 -->
      <div class="pt-6 flex flex-col sm:flex-row items-center justify-between gap-4 text-xs text-slate-500 font-mono text-center sm:text-left">
        <div>
          © 2024–2026 ChiMu-Lab (迟暮实验室). All Rights Reserved.
        </div>
        <div class="flex items-center gap-2">
          <span>Designed with</span>
          <Heart class="w-3.5 h-3.5 text-rose-500 inline fill-rose-500/20" />
          <span>for Engineering Excellence</span>
        </div>
      </div>
    </div>
  </footer>
</template>
