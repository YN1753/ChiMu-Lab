<script setup lang="ts">
import { ArrowUpRight } from 'lucide-vue-next'
</script>

<template>
  <footer class="border-t border-[var(--border-subtle)] bg-[var(--bg-archive)] py-16 text-left transition-colors duration-300">
    <div class="max-w-5xl mx-auto px-5 sm:px-8 space-y-8 font-mono-archive text-xs text-[var(--ink-muted)]">
      
      <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-6 pb-8 border-b border-[var(--border-subtle)]">
        <div>
          <span class="text-[var(--ink-primary)] font-semibold tracking-wider font-serif-editorial text-sm">
            迟暮的生活档案馆
          </span>
          <p class="mt-1 text-[11px] text-[var(--ink-secondary)]">
            记录做过的事、看过的风景、思考过的念头，和不想遗忘的瞬间。
          </p>
        </div>

        <!-- 规范工信部备案号 (严格合规) -->
        <div class="flex items-center gap-3">
          <a
            href="https://beian.miit.gov.cn"
            target="_blank"
            rel="noopener noreferrer"
            class="hover:text-[var(--ink-primary)] hover-underline transition-colors py-1 inline-flex items-center gap-1 text-[var(--ink-secondary)]"
          >
            <span>浙ICP备2026081664号</span>
            <ArrowUpRight class="w-3 h-3 opacity-60" />
          </a>
        </div>
      </div>

      <!-- 版权与定位 -->
      <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 text-[11px]">
        <div>
          © 2024–2026 迟暮 (ChiMu). 时间刻度下的真实生活。
        </div>
        <div class="flex items-center gap-4">
          <a href="https://github.com/YN1753" target="_blank" class="hover:text-[var(--ink-primary)]">
            github/YN1753
          </a>
          <span>·</span>
          <span>中国 · 杭州</span>
        </div>
      </div>

    </div>
  </footer>
</template>
