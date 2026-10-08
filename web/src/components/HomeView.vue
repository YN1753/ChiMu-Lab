<script setup lang="ts">
import { ref, computed } from 'vue'
import type { LifeEntry, NowStatus, ArchiveCategory } from '../types'
import LifeActivityMap from './LifeActivityMap.vue'
import LifeStream from './LifeStream.vue'
import { ArrowRight } from 'lucide-vue-next'

const props = defineProps<{
  entries: LifeEntry[]
  now?: NowStatus | null
  currentCategory?: ArchiveCategory
}>()

const emit = defineEmits<{
  (e: 'navigate', view: string): void
  (e: 'changeCategory', category: ArchiveCategory): void
}>()

const selectedDate = ref<string>('')

const handleDateSelect = (date: string) => {
  selectedDate.value = date
  if (date) {
    const el = document.getElementById('timeline-section')
    if (el) {
      el.scrollIntoView({ behavior: 'smooth' })
    }
  }
}

const currentCategoryInfo = computed(() => {
  switch (props.currentCategory) {
    case 'daily':
      return {
        label: '日常',
        title: '日常 // 生活时间流',
        desc: '现实生活足迹：出门、吃饭、旅行、手冲、行迹与身边小事',
      }
    case 'thought':
      return {
        label: '想法',
        title: '想法 // 生活时间流',
        desc: '思想与文字：随笔、感悟、计划与碎片化思索',
      }
    case 'project':
      return {
        label: '项目',
        title: '项目 // 生活时间流',
        desc: '创造与探索：ArchCanvas、Go 底层工具、开源服务与个人创作',
      }
    case 'collection':
      return {
        label: '收藏',
        title: '收藏 // 生活时间流',
        desc: '喜爱与长伴：唱片、书籍、游戏、设备与值得保留之物',
      }
    default:
      return {
        label: '全部记录',
        title: '全部记录 // 生活时间流',
        desc: '完整生活时间线：日常、想法、项目与收藏自然相融',
      }
  }
})
</script>

<template>
  <div class="max-w-5xl mx-auto px-5 sm:px-8 text-left">
    
    <!-- 第一屏：极度克制、优雅的开篇封面 (纯中文典雅排版) -->
    <section class="pt-20 sm:pt-28 pb-16 sm:pb-20 border-b border-[var(--border-subtle)]">
      <div class="max-w-3xl space-y-8">
        
        <div class="space-y-3">
          <h1 class="font-serif-editorial text-4xl sm:text-6xl lg:text-7xl font-normal tracking-tight text-[var(--ink-primary)] leading-[1.08]">
            迟暮
          </h1>
          <p class="font-serif-editorial text-sm sm:text-base tracking-widest text-[var(--ink-muted)]">
            一个属于我自己的数字生活档案馆。
          </p>
        </div>

        <div class="text-xl sm:text-2xl text-[var(--ink-secondary)] font-serif-editorial leading-relaxed max-w-xl pt-2 space-y-1">
          <p>我做过的事，</p>
          <p>我看过的风景，</p>
          <p>我思考过的念头，</p>
          <p>和我不想遗忘的瞬间。</p>
        </div>

        <div class="pt-4 font-mono-archive text-xs text-[var(--ink-muted)] tracking-wider">
          <span>杭州 · 2026年10月09日</span>
        </div>

      </div>
    </section>

    <!-- 首页第一视觉焦点：2026 生活刻度热力图 (Life Activity / Life in 2026) -->
    <LifeActivityMap
      :entries="entries"
      :selectedDate="selectedDate"
      @selectDate="handleDateSelect"
    />

    <!-- 当下步调 // NOW (轻量自然的人性化剪影，突出“我作为一个人”) -->
    <section v-if="now" class="py-12 border-b border-[var(--border-subtle)]">
      <div class="flex flex-col sm:flex-row sm:items-baseline justify-between gap-4 pb-6">
        <div class="flex items-center gap-3">
          <h2 class="font-serif-editorial text-xl sm:text-2xl font-medium tracking-tight text-[var(--ink-primary)]">
            当下步调
          </h2>
          <span class="font-mono-archive text-[11px] uppercase tracking-widest text-[var(--ink-muted)]">
            // NOW
          </span>
        </div>
        <button
          @click="emit('navigate', 'now')"
          class="text-xs font-mono-archive text-[var(--ink-muted)] hover:text-[var(--ink-primary)] hover-underline cursor-pointer flex items-center gap-1 self-start sm:self-auto"
        >
          <span>查看完整当下状态</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </button>
      </div>

      <!-- 纯排版、无卡片、有呼吸感的状态网格 -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 text-left font-serif-editorial">
        <div class="space-y-1.5">
          <span class="font-mono-archive text-[11px] tracking-wider text-[var(--ink-muted)] block">// 正在构建</span>
          <p class="text-base text-[var(--ink-primary)] leading-snug">{{ now.building }}</p>
        </div>
        <div class="space-y-1.5">
          <span class="font-mono-archive text-[11px] tracking-wider text-[var(--ink-muted)] block">// 正在听与读</span>
          <p class="text-base text-[var(--ink-primary)] leading-snug">{{ now.listening }} · {{ now.reading }}</p>
        </div>
        <div class="space-y-1.5">
          <span class="font-mono-archive text-[11px] tracking-wider text-[var(--ink-muted)] block">// 所思所想</span>
          <p class="text-base text-[var(--ink-secondary)] leading-snug italic">“{{ now.thinking }}”</p>
        </div>
      </div>
    </section>

    <!-- 首页第二核心：生活时间流 (Timeline / Life Stream) -->
    <section id="timeline-section" class="pt-16 sm:pt-20 pb-28">
      
      <!-- 动态视图标题与状态 -->
      <div class="flex flex-col sm:flex-row sm:items-baseline justify-between gap-4 pb-8 mb-6 border-b border-[var(--border-subtle)]">
        <div>
          <div class="flex items-center gap-3">
            <h2 class="text-base sm:text-lg font-semibold tracking-wider text-[var(--ink-primary)] font-serif-editorial">
              {{ currentCategoryInfo.title }}
            </h2>
            <span class="text-xs font-mono-archive text-[var(--ink-muted)]">
              （当前视图 {{ entries.length }} 条）
            </span>
          </div>
          <p class="text-xs font-mono-archive text-[var(--ink-muted)] mt-1">
            {{ currentCategoryInfo.desc }}
          </p>
        </div>

        <div class="flex items-center gap-4 text-xs font-mono-archive">
          <!-- 切换回全景 -->
          <button
            v-if="currentCategory && currentCategory !== 'all'"
            @click="emit('changeCategory', 'all')"
            class="text-[var(--accent-warm)] hover:text-[var(--ink-primary)] hover-underline cursor-pointer"
          >
            [查看全部记录]
          </button>

          <button
            @click="emit('navigate', 'archive')"
            class="text-[var(--ink-muted)] hover:text-[var(--ink-primary)] hover-underline cursor-pointer flex items-center gap-1"
          >
            <span>年度归档</span>
            <ArrowRight class="w-3.5 h-3.5" />
          </button>
        </div>
      </div>

      <!-- 时间流列表 (支持右上角 5 重视角切换与刻度聚焦) -->
      <LifeStream
        :entries="entries"
        :limit="currentCategory === 'all' ? 8 : 0"
        :showFilters="false"
        :dateFilter="selectedDate"
        :currentCategory="currentCategory"
        @clearDateFilter="selectedDate = ''"
        @changeCategory="emit('changeCategory', $event)"
      />

      <!-- 底部探索更多 -->
      <div class="pt-16 mt-8 border-t border-[var(--border-subtle)] flex items-center justify-between">
        <span class="text-xs font-mono-archive text-[var(--ink-muted)]">
          {{ currentCategoryInfo.label }} · 真实生活记录
        </span>
        <button
          @click="emit('navigate', 'archive')"
          class="inline-flex items-center gap-2 text-xs text-[var(--ink-primary)] hover-underline cursor-pointer font-mono-archive"
        >
          <span>按年归档一览</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </button>
      </div>

    </section>

  </div>
</template>
