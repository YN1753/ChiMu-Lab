<script setup lang="ts">
import { X, BarChart3 } from 'lucide-vue-next'
import type { GlobalStats } from '../types'

defineProps<{
  isOpen: boolean
  stats: GlobalStats | null
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const formatMoney = (cents: number): string => {
  return (cents / 100).toLocaleString('zh-CN', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  })
}
</script>

<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/40 backdrop-blur-xs animate-in fade-in duration-200"
    @click.self="emit('close')"
  >
    <div
      class="bg-[var(--bg-archive)] border border-[var(--border-subtle)] shadow-[0_16px_40px_rgba(0,0,0,0.1)] w-full max-w-md rounded-xs p-6 sm:p-8 text-left space-y-6"
    >
      <div class="flex items-center justify-between pb-4 border-b border-[var(--border-subtle)]">
        <div class="flex items-center gap-2.5">
          <BarChart3 class="w-4 h-4 text-[var(--ink-secondary)]" />
          <h2 class="font-serif-editorial text-xl font-medium text-[var(--ink-primary)]">
            生活刻度统计 // STATS
          </h2>
        </div>
        <button
          @click="emit('close')"
          class="p-1 text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer"
        >
          <X class="w-4 h-4" />
        </button>
      </div>

      <!-- 核心指标网格 -->
      <div class="grid grid-cols-2 gap-6 font-serif-editorial">
        <div class="space-y-1 p-3 bg-[var(--bg-subtle)]/50 rounded-xs border border-[var(--border-subtle)]">
          <span class="font-mono-archive text-[11px] text-[var(--ink-muted)] block">生活印记总数</span>
          <div class="text-3xl font-medium tracking-tight text-[var(--ink-primary)]">
            {{ stats?.total_entries ?? 0 }}
          </div>
        </div>

        <div class="space-y-1 p-3 bg-[var(--bg-subtle)]/50 rounded-xs border border-[var(--border-subtle)]">
          <span class="font-mono-archive text-[11px] text-[var(--ink-muted)] block">胶卷与照片</span>
          <div class="text-3xl font-medium tracking-tight text-[var(--ink-primary)]">
            {{ stats?.total_photos ?? 0 }}
          </div>
        </div>

        <div class="space-y-1 p-3 bg-[var(--bg-subtle)]/50 rounded-xs border border-[var(--border-subtle)]">
          <span class="font-mono-archive text-[11px] text-[var(--ink-muted)] block">造物与项目</span>
          <div class="text-3xl font-medium tracking-tight text-[var(--ink-primary)]">
            {{ stats?.total_projects ?? 0 }}
          </div>
        </div>

        <div class="space-y-1 p-3 bg-[var(--bg-subtle)]/50 rounded-xs border border-[var(--border-subtle)]">
          <span class="font-mono-archive text-[11px] text-[var(--ink-muted)] block">本月实际支出</span>
          <div class="text-2xl font-medium tracking-tight text-[var(--accent-warm)]">
            ¥{{ formatMoney(stats?.month_expense ?? 0) }}
          </div>
        </div>
      </div>

      <p class="text-xs font-serif-editorial text-[var(--ink-muted)] text-center pt-2">
        每一个数字，都是真实流淌度过的时间。
      </p>

      <div class="pt-2 flex justify-end">
        <button
          @click="emit('close')"
          class="px-4 py-1.5 text-xs font-mono-archive bg-[var(--ink-primary)] text-[var(--bg-archive)] rounded-xs cursor-pointer hover:opacity-90"
        >
          关闭
        </button>
      </div>
    </div>
  </div>
</template>
