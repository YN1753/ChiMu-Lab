<script setup lang="ts">
import { X, Server, CheckCircle2, AlertCircle } from 'lucide-vue-next'
import type { StorageStatus } from '../types'

defineProps<{
  isOpen: boolean
  status: StorageStatus | null
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()
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
          <Server class="w-4 h-4 text-[var(--ink-secondary)]" />
          <h2 class="font-serif-editorial text-xl font-medium text-[var(--ink-primary)]">
            对象存储设置 // STORAGE
          </h2>
        </div>
        <button
          @click="emit('close')"
          class="p-1 text-[var(--ink-muted)] hover:text-[var(--ink-primary)] cursor-pointer"
        >
          <X class="w-4 h-4" />
        </button>
      </div>

      <div class="space-y-4 font-serif-editorial">
        <div class="flex items-center justify-between py-2 border-b border-[var(--border-subtle)]/60 text-sm">
          <span class="text-[var(--ink-muted)]">存储提供方</span>
          <span class="text-[var(--ink-primary)] font-mono-archive text-xs">Cloudflare R2</span>
        </div>

        <div class="flex items-center justify-between py-2 border-b border-[var(--border-subtle)]/60 text-sm">
          <span class="text-[var(--ink-muted)]">连接状态</span>
          <div class="flex items-center gap-1.5 font-mono-archive text-xs">
            <span
              v-if="status?.configured"
              class="text-emerald-700 flex items-center gap-1"
            >
              <CheckCircle2 class="w-3.5 h-3.5" />
              <span>已连接</span>
            </span>
            <span
              v-else
              class="text-[var(--accent-warm)] flex items-center gap-1"
            >
              <AlertCircle class="w-3.5 h-3.5" />
              <span>未配置</span>
            </span>
          </div>
        </div>

        <div class="flex items-center justify-between py-2 border-b border-[var(--border-subtle)]/60 text-sm">
          <span class="text-[var(--ink-muted)]">存储桶 (Bucket)</span>
          <span class="text-[var(--ink-primary)] font-mono-archive text-xs">
            {{ status?.bucket || '—' }}
          </span>
        </div>

        <div class="flex items-center justify-between py-2 border-b border-[var(--border-subtle)]/60 text-sm">
          <span class="text-[var(--ink-muted)]">公网域名 (Domain)</span>
          <span class="text-[var(--ink-primary)] font-mono-archive text-xs truncate max-w-[200px]">
            {{ status?.public_domain || '—' }}
          </span>
        </div>
      </div>

      <div
        v-if="!status?.configured"
        class="p-3 bg-[var(--bg-subtle)] border border-[var(--border-subtle)] rounded-xs text-xs font-serif-editorial text-[var(--ink-secondary)] leading-relaxed space-y-1"
      >
        <p class="font-medium text-[var(--ink-primary)]">R2 存储未配置：</p>
        <p>媒体照片上传暂不可用。如需开启直传，请在服务器环境变量或 <code>.env</code> 中配置 R2 凭据。</p>
        <p class="text-[11px] font-mono-archive text-[var(--ink-muted)] pt-1">
          注：文字生活记录、随笔与记账功能无需 R2 即可 100% 正常使用。
        </p>
      </div>

      <div class="pt-2 flex justify-end">
        <button
          @click="emit('close')"
          class="px-4 py-1.5 text-xs font-mono-archive bg-[var(--ink-primary)] text-[var(--bg-archive)] rounded-xs cursor-pointer hover:opacity-90"
        >
          确定
        </button>
      </div>
    </div>
  </div>
</template>
