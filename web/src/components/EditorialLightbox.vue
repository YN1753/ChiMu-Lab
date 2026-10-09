<script setup lang="ts">
import { watch, onBeforeUnmount } from 'vue'
import { X, Camera, MapPin, Calendar } from 'lucide-vue-next'

const props = defineProps<{
  isOpen: boolean
  imageUrl: string
  title?: string
  meta?: string
  location?: string
  date?: string
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === 'Escape' && props.isOpen) {
    emit('close')
  }
}

watch(
  () => props.isOpen,
  (open) => {
    if (open) {
      window.addEventListener('keydown', handleKeydown)
      document.body.style.overflow = 'hidden'
    } else {
      window.removeEventListener('keydown', handleKeydown)
      document.body.style.overflow = ''
    }
  },
  { immediate: true }
)

onBeforeUnmount(() => {
  window.removeEventListener('keydown', handleKeydown)
  document.body.style.overflow = ''
})
</script>

<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 z-[100] flex flex-col justify-between p-4 sm:p-8 bg-[#121110]/95 backdrop-blur-md text-[#ece9e1] animate-in fade-in duration-200 select-none cursor-default"
    @click.self="emit('close')"
  >
    <!-- 顶部暗房灯箱控制栏 -->
    <header class="w-full max-w-6xl mx-auto flex items-center justify-between py-2 border-b border-white/10 text-xs font-mono-archive text-[#98958c]">
      <div class="flex items-center gap-3">
        <span class="text-[#d8764b] font-medium tracking-widest">// 暗房底片灯箱</span>
        <span class="hidden sm:inline opacity-60">· DARKROOM LIGHTBOX</span>
      </div>

      <button
        @click="emit('close')"
        class="inline-flex items-center gap-1.5 px-3 py-1 text-xs text-[#ece9e1] hover:text-[#d8764b] border border-white/10 hover:border-white/30 rounded-xs transition-colors cursor-pointer"
        title="按 ESC 或点击关闭"
      >
        <span>ESC 关闭</span>
        <X class="w-3.5 h-3.5" />
      </button>
    </header>

    <!-- 中央装裱大图视口 -->
    <main
      class="flex-grow flex items-center justify-center p-2 sm:p-6 overflow-hidden"
      @click.self="emit('close')"
    >
      <div class="relative max-w-5xl max-h-[76vh] flex flex-col items-center">
        <img
          :src="imageUrl"
          :alt="title || '底片装裱'"
          class="max-w-full max-h-[76vh] object-contain rounded-xs shadow-[0_25px_60px_rgba(0,0,0,0.6)] ring-1 ring-white/10 select-none"
        />
      </div>
    </main>

    <!-- 底部出版元信息装裱条 (Passe-partout Strip) -->
    <footer class="w-full max-w-6xl mx-auto py-3 border-t border-white/10 flex flex-wrap items-center justify-between gap-4 text-xs font-mono-archive">
      <div class="space-y-1">
        <h4 v-if="title" class="font-serif-editorial text-base sm:text-lg text-[#ece9e1] font-normal">
          {{ title }}
        </h4>
        <div class="flex flex-wrap items-center gap-x-4 gap-y-1 text-[#98958c]">
          <span v-if="date" class="inline-flex items-center gap-1">
            <Calendar class="w-3 h-3 text-[#d8764b]" />
            <span>{{ date }}</span>
          </span>
          <span v-if="location" class="inline-flex items-center gap-1">
            <MapPin class="w-3 h-3 text-[#d8764b]" />
            <span>{{ location }}</span>
          </span>
          <span v-if="meta" class="inline-flex items-center gap-1 text-[#ece9e1]/80">
            <Camera class="w-3 h-3 text-[#d8764b]" />
            <span>{{ meta }}</span>
          </span>
        </div>
      </div>

      <div class="text-[#78756d] text-[11px] hidden sm:block">
        35mm ARCHIVE · LIFE > PROJECTS
      </div>
    </footer>
  </div>
</template>
