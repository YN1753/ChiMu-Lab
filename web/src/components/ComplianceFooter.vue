<script setup lang="ts">
import type { SiteConfig } from '../types'
import { Shield, ExternalLink, Terminal } from 'lucide-vue-next'

defineProps<{
  config: SiteConfig | null
}>()
</script>

<template>
  <footer class="mt-20 border-t border-white/[0.08] bg-[#050609] text-slate-400 py-14">
    <div class="max-w-6xl mx-auto px-4 sm:px-6">
      <div class="flex flex-col md:flex-row items-center justify-between gap-8 pb-10 border-b border-white/[0.06]">
        <!-- 左侧 Logo 与描述 -->
        <div class="text-center md:text-left">
          <div class="flex items-center justify-center md:justify-start gap-2.5">
            <div class="w-7 h-7 rounded-lg bg-cyan-500/10 border border-cyan-500/30 flex items-center justify-center text-cyan-400">
              <Terminal class="w-3.5 h-3.5" />
            </div>
            <span class="text-sm font-bold text-white tracking-wide">
              ChiMu-Lab · 迟暮实验室
            </span>
          </div>
          <p class="text-xs text-slate-500 font-mono mt-2">
            Code Activity Hub · Powered by Go 1.27 + Vue 3 + Pure SQLite
          </p>
        </div>

        <!-- 中间/右侧合规备案专区（审核核心要求） -->
        <div class="flex flex-wrap items-center justify-center gap-4 text-xs font-mono">
          <!-- 工信部 ICP 备案号 -->
          <a
            :href="config?.icp_link || 'https://beian.miit.gov.cn'"
            target="_blank"
            rel="noopener noreferrer"
            class="flex items-center gap-2 text-slate-300 hover:text-cyan-300 transition-all py-1.5 px-3.5 rounded-lg bg-white/[0.03] border border-white/[0.08] hover:border-cyan-500/40 shadow-sm"
          >
            <Shield class="w-3.5 h-3.5 text-cyan-400" />
            <span class="font-semibold">{{ config?.icp_number || '浙ICP备2026081664号' }}</span>
            <ExternalLink class="w-3 h-3 text-slate-500" />
          </a>

          <!-- 全国公安机关联网备案号预留 -->
          <a
            href="http://www.beian.gov.cn/portal/registerSystemInfo"
            target="_blank"
            rel="noopener noreferrer"
            class="flex items-center gap-2 text-slate-400 hover:text-slate-200 transition-all py-1.5 px-3.5 rounded-lg bg-white/[0.02] border border-white/[0.06] hover:border-white/20"
          >
            <span class="text-xs">👮</span>
            <span>{{ config?.police_number || '全国公安联网备案审核中' }}</span>
          </a>

          <!-- 节点状态 -->
          <div class="flex items-center gap-1.5 text-slate-500 text-[11px] py-1.5 px-2.5 rounded-lg bg-black/40 border border-white/[0.04]">
            <span class="w-1.5 h-1.5 rounded-full bg-emerald-400"></span>
            <span>Tencent Cloud CVM · TLS 1.3 Active</span>
          </div>
        </div>
      </div>

      <!-- 底部版权 -->
      <div class="pt-8 flex flex-col sm:flex-row items-center justify-between gap-4 text-xs text-slate-500 font-mono text-center sm:text-left">
        <div>
          © 2024–2026 ChiMu-Lab (迟暮实验室). All Rights Reserved.
        </div>
        <div class="flex items-center gap-4">
          <a href="https://github.com/YN1753" target="_blank" class="hover:text-slate-300 transition-colors">
            GitHub @YN1753
          </a>
          <span class="text-slate-700">|</span>
          <a href="/interview" class="hover:text-cyan-400 transition-colors">
            今日 · 面试练习站
          </a>
        </div>
      </div>
    </div>
  </footer>
</template>
