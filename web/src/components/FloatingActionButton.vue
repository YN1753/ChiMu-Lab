<script setup lang="ts">
import { Plus } from 'lucide-vue-next'

const emit = defineEmits<{
  (e: 'click'): void
}>()
</script>

<template>
  <aside aria-label="快捷操作" class="fixed right-6 bottom-6 sm:right-8 sm:bottom-8 z-40">
    <button
      @click="emit('click')"
      class="w-12 h-12 sm:w-13 sm:h-13 rounded-full bg-[var(--ink-primary)] text-[var(--bg-archive)] shadow-[0_6px_24px_rgba(0,0,0,0.18)] hover:scale-105 active:scale-95 transition-all duration-200 flex items-center justify-center cursor-pointer group"
      title="记录一点什么 (快捷新增)"
      aria-label="新增生活记录"
    >
      <Plus class="w-5 h-5 group-hover:rotate-90 transition-transform duration-200" />
    </button>
  </aside>
</template>
