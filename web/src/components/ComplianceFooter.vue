<script setup lang="ts">
import type { SiteConfig } from '../types'
import { Shield, ExternalLink } from 'lucide-vue-next'

defineProps<{
  config: SiteConfig | null
}>()
</script>

<template>
  <footer class="border-t border-[#e8e6df] bg-[#f2efe9] text-[#525662] py-16">
    <div class="max-w-6xl mx-auto px-5 sm:px-8 text-left">
      <div class="flex flex-col md:flex-row items-start md:items-center justify-between gap-8 pb-12 border-b border-[#e2ded4]">
        <!-- 左侧 Logo 与描述 -->
        <div>
          <div class="flex items-center gap-2.5">
            <div class="w-7 h-7 rounded-lg bg-[#14151a] flex items-center justify-center text-white text-xs font-serif-cinematic font-bold">
              迟
            </div>
            <span class="text-sm font-bold text-[#14151a] font-serif-cinematic tracking-wide">
              ChiMu-Lab · 迟暮实验室
            </span>
          </div>
          <p class="text-xs text-[#8c8f9b] font-mono mt-2.5">
            Code Activity Hub · Powered by Go 1.27 + Vue 3 + Pure-Go SQLite
          </p>
        </div>

        <!-- 中间/右侧：合规备案专区（严密合规审核要求） -->
        <div class="flex flex-wrap items-center gap-4 text-xs font-mono">
          <!-- 工信部 ICP 备案号 -->
          <a
            :href="config?.icp_link || 'https://beian.miit.gov.cn'"
            target="_blank"
            rel="noopener noreferrer"
            class="flex items-center gap-2 text-[#14151a] hover:text-[#92400e] transition-all py-2 px-4 rounded-xl bg-white border border-[#dedcd5] hover:border-[#d97706]/50 shadow-2xs"
          >
            <Shield class="w-3.5 h-3.5 text-[#d97706]" />
            <span class="font-semibold">{{ config?.icp_number || '浙ICP备2026081664号' }}</span>
            <ExternalLink class="w-3 h-3 text-[#8c8f9b]" />
          </a>

          <!-- 全国公安机关联网备案号预留 -->
          <a
            href="http://www.beian.gov.cn/portal/registerSystemInfo"
            target="_blank"
            rel="noopener noreferrer"
            class="flex items-center gap-2 text-[#716e64] hover:text-[#14151a] transition-all py-2 px-3.5 rounded-xl bg-[#edeae1] border border-[#dedcd5]"
          >
            <span class="text-xs">👮‍♂️</span>
            <span>{{ config?.police_number || '全国公安联网备案审核中' }}</span>
          </a>

          <!-- 节点状态 -->
          <div class="flex items-center gap-1.5 text-[#716e64] text-[11px] py-2 px-3 rounded-xl bg-white/60 border border-[#e2ded4]">
            <span class="w-1.5 h-1.5 rounded-full bg-[#0d766e]"></span>
            <span>Hangzhou CVM · TLS 1.3 Active</span>
          </div>
        </div>
      </div>

      <!-- 底部版权与胶片演职员表感 -->
      <div class="pt-8 flex flex-col sm:flex-row items-center justify-between gap-4 text-xs text-[#8c8f9b] font-mono text-center sm:text-left">
        <div>
          © 2024–2026 ChiMu-Lab (迟暮实验室). All Rights Reserved.
        </div>
        <div class="flex items-center gap-5">
          <a href="https://github.com/YN1753" target="_blank" class="hover:text-[#14151a] transition-colors">
            GitHub @YN1753
          </a>
          <span class="text-[#dedcd5]">|</span>
          <a href="/interview" class="hover:text-[#92400e] transition-colors">
            今日 · 面试练习站
          </a>
          <span class="text-[#dedcd5]">|</span>
          <span>Hangzhou, CN</span>
        </div>
      </div>
    </div>
  </footer>
</template>
