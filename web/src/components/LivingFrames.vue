<script setup lang="ts">
import { ref, computed } from 'vue'
import type { LifeMoment } from '../types'
import { Camera, Heart, Sparkles, MapPin, Cloud, RotateCw, Coffee, Music, Film, Compass } from 'lucide-vue-next'
import { audio } from '../utils/audio'

const props = defineProps<{
  moments: LifeMoment[]
}>()

const activeCategory = ref<string>('all')
const flippedCardIds = ref<number[]>([])

const categories = [
  { key: 'all', label: '全部切片 · ALL', icon: Sparkles },
  { key: 'thought', label: '随笔与思考 · THOUGHTS', icon: Compass },
  { key: 'film', label: '胶卷定格 · FILM', icon: Film },
  { key: 'coffee', label: '手冲手记 · COFFEE', icon: Coffee },
  { key: 'reading', label: '阅读与音乐 · SOUND', icon: Music },
]

const filteredMoments = computed(() => {
  if (activeCategory.value === 'all') return props.moments
  return props.moments.filter(m => m.category === activeCategory.value)
})

const setCategory = (cat: string) => {
  activeCategory.value = cat
  audio.playTink()
}

// 翻转拍立得卡片 (3D 翻转)
const toggleFlip = (id: number) => {
  if (flippedCardIds.value.includes(id)) {
    flippedCardIds.value = flippedCardIds.value.filter(i => i !== id)
  } else {
    flippedCardIds.value.push(id)
  }
  audio.playPaperFlip()
}

// 互动点赞 / 盖章
const likeMoment = async (moment: LifeMoment, event: Event) => {
  event.stopPropagation()
  moment.likes++
  audio.playShutter()
  try {
    await fetch(`/api/moments/${moment.id}/like`, { method: 'POST' })
  } catch {
    // 乐观更新
  }
}
</script>

<template>
  <section id="moments" class="py-24 border-t border-[var(--border-color)] bg-[var(--bg-page)] relative transition-colors duration-400">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-left">
      
      <!-- 分镜标头 -->
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-12">
        <div>
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-[var(--bg-surface)] text-[var(--ink-secondary)] text-xs font-mono mb-3 border border-[var(--border-color)] shadow-2xs">
            <Camera class="w-3.5 h-3.5 text-[var(--accent-amber)]" />
            <span>LIVING VIGNETTES · 生活切片与胶片手记</span>
          </div>
          <h2 class="text-3xl sm:text-4xl font-medium tracking-tight text-[var(--ink-primary)] font-serif-cinematic">
            代码之外：生活、咖啡与吉光片羽
          </h2>
          <p class="text-sm text-[var(--ink-secondary)] mt-2 font-mono">
            真实记录当下心绪、手冲笔记与秋日街道 · 点击卡片角落可「3D 翻转查看便签」
          </p>
        </div>

        <div class="text-xs font-mono text-[var(--ink-muted)]">
          <span>共记录 {{ moments.length }} 个真实瞬间</span>
        </div>
      </div>

      <!-- 分类滤镜胶囊 -->
      <div class="flex flex-wrap items-center gap-2 mb-12">
        <button
          v-for="cat in categories"
          :key="cat.key"
          @click="setCategory(cat.key)"
          class="px-4 py-2 rounded-full text-xs font-mono transition-all flex items-center gap-2 cursor-pointer border"
          :class="activeCategory === cat.key 
            ? 'bg-[var(--ink-primary)] text-[var(--bg-page)] border-[var(--ink-primary)] shadow-sm font-semibold' 
            : 'bg-[var(--bg-surface)] text-[var(--ink-secondary)] border-[var(--border-color)] hover:border-[var(--accent-amber)]'"
        >
          <component :is="cat.icon" class="w-3.5 h-3.5" :class="activeCategory === cat.key ? 'text-[var(--accent-amber)]' : ''" />
          <span>{{ cat.label }}</span>
        </button>
      </div>

      <!-- 拍立得 3D 卡片矩阵 -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div
          v-for="m in filteredMoments"
          :key="m.id"
          class="perspective-1000 min-h-[460px]"
        >
          <!-- 3D 翻转容器 -->
          <div
            class="relative w-full h-full transform-style-3d transition-transform duration-600 rounded-2xl"
            :class="{ 'rotate-y-180': flippedCardIds.includes(m.id) }"
          >
            
            <!-- 正面：拍立得照片与生活诗意 -->
            <div
              class="absolute inset-0 backface-hidden film-card p-5 sm:p-6 flex flex-col justify-between overflow-hidden bg-[var(--bg-surface)] border border-[var(--border-color)]"
            >
              <div>
                <!-- 顶部元数据：时间、地点、天气 -->
                <div class="flex items-center justify-between text-[11px] font-mono text-[var(--ink-muted)] pb-3 border-b border-[var(--border-color)] mb-4">
                  <div class="flex items-center gap-1.5 text-[var(--ink-secondary)]">
                    <MapPin class="w-3 h-3 text-[var(--accent-amber)]" />
                    <span>{{ m.location }}</span>
                  </div>
                  <div class="flex items-center gap-1.5">
                    <Cloud class="w-3 h-3 text-[var(--accent-teal)]" />
                    <span>{{ m.weather }}</span>
                  </div>
                </div>

                <!-- 胶片影像预览 (带画框微质感) -->
                <div class="relative w-full h-44 rounded-xl overflow-hidden bg-stone-100 mb-4 border border-[var(--border-color)] group">
                  <img
                    :src="m.image_url"
                    :alt="m.title"
                    class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"
                    loading="lazy"
                  />
                  <!-- 相机刻度水印 -->
                  <div class="absolute bottom-2 left-2 px-2 py-0.5 rounded bg-black/60 text-white text-[10px] font-mono backdrop-blur-xs">
                    {{ m.camera }}
                  </div>
                </div>

                <!-- 标题与生活正文 -->
                <h4 class="text-lg font-serif-cinematic font-semibold text-[var(--ink-primary)] mb-2">
                  {{ m.title }}
                </h4>
                <p class="text-xs text-[var(--ink-secondary)] leading-relaxed line-clamp-3">
                  {{ m.content }}
                </p>
              </div>

              <!-- 底部操作：翻转按钮与戳一戳 -->
              <div class="pt-4 mt-4 border-t border-[var(--border-color)] flex items-center justify-between">
                <button
                  @click="toggleFlip(m.id)"
                  class="text-xs font-mono text-[var(--ink-secondary)] hover:text-[var(--accent-amber)] inline-flex items-center gap-1.5 cursor-pointer transition-colors"
                >
                  <RotateCw class="w-3.5 h-3.5" />
                  <span>翻转查看手记</span>
                </button>

                <!-- 盖章/点赞 -->
                <button
                  @click="likeMoment(m, $event)"
                  class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full bg-[var(--bg-surface-subtle)] hover:bg-[var(--accent-amber)]/10 text-[var(--ink-primary)] text-xs font-mono transition-all cursor-pointer border border-[var(--border-color)] group"
                  title="为这个瞬间盖章"
                >
                  <Heart class="w-3.5 h-3.5 text-rose-500 group-hover:scale-125 transition-transform" />
                  <span>{{ m.likes }}</span>
                </button>
              </div>

            </div>

            <!-- 反面：复古明信片手写便签与印章 (3D 翻转背面) -->
            <div
              class="absolute inset-0 backface-hidden rotate-y-180 film-card p-6 flex flex-col justify-between overflow-hidden bg-[var(--bg-surface)] border-2 border-[var(--accent-amber)]/30 text-left"
            >
              <div>
                <!-- 明信片邮票与邮戳 -->
                <div class="flex items-center justify-between pb-4 border-b border-[var(--border-color)] mb-5">
                  <div class="text-[11px] font-mono text-[var(--accent-amber)] font-semibold">
                    POSTCARD FROM HANGZHOU · {{ m.date }}
                  </div>
                  <!-- 复古邮戳 -->
                  <div class="w-10 h-10 rounded-full border border-dashed border-[var(--accent-amber)] flex items-center justify-center text-[9px] font-mono text-[var(--accent-amber)] rotate-12">
                    OCT '26
                  </div>
                </div>

                <!-- 便签标题 -->
                <div class="text-xs font-mono text-[var(--ink-muted)] mb-2">
                  迟暮手记 · {{ m.mood }}状态下写下：
                </div>

                <!-- 手写体质感文字 -->
                <div class="p-4 rounded-xl bg-[var(--bg-surface-subtle)] border border-[var(--border-color)] text-xs text-[var(--ink-primary)] font-serif-cinematic leading-relaxed italic">
                  “{{ m.note }}”
                </div>

                <!-- 标签与器材 -->
                <div class="mt-4 space-y-1.5 text-[11px] font-mono text-[var(--ink-muted)]">
                  <div>镜头/参数：<span class="text-[var(--ink-secondary)]">{{ m.camera }}</span></div>
                  <div>记录时间：<span class="text-[var(--ink-secondary)]">{{ m.date }} {{ m.time }}</span></div>
                </div>
              </div>

              <!-- 翻回正面按钮 -->
              <div class="pt-4 border-t border-[var(--border-color)] flex items-center justify-between">
                <button
                  @click="toggleFlip(m.id)"
                  class="text-xs font-mono text-[var(--ink-primary)] hover:text-[var(--accent-amber)] inline-flex items-center gap-1.5 cursor-pointer font-medium"
                >
                  <RotateCw class="w-3.5 h-3.5" />
                  <span>翻回正面照片</span>
                </button>

                <span class="text-[11px] font-mono text-[var(--ink-muted)]">
                  #{{ m.category }}
                </span>
              </div>

            </div>

          </div>
        </div>
      </div>

    </div>
  </section>
</template>
