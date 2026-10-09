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
  (e: 'selectEntry', entry: LifeEntry): void
  (e: 'openAdd'): void
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
const showHeatmap = ref(false)
</script>

<template>
  <div class="max-w-4xl mx-auto px-5 sm:px-8 text-left">
    
    <!-- 开篇封面 (紧凑、静默、优雅的自白引言) -->
    <section class="pt-12 sm:pt-16 pb-8 border-b border-[var(--border-subtle)]">
      <div class="space-y-4">
        <div class="flex items-baseline justify-between">
          <h1 class="font-serif-editorial text-3xl sm:text-5xl font-medium tracking-tight text-[var(--ink-primary)]">
            迟暮
          </h1>
          <span class="font-mono-archive text-xs text-[var(--ink-muted)]">
            杭州 · 2026.10.09
          </span>
        </div>

        <p class="font-serif-editorial text-xs sm:text-sm tracking-widest text-[var(--ink-muted)]">
          一个属于我自己的数字生活空间。
        </p>

        <p class="font-serif-editorial text-base sm:text-lg text-[var(--ink-secondary)] leading-relaxed italic max-w-2xl pt-1">
          “我做过的事，我看过的风景，我思考过的念头，和我不想遗忘的瞬间。”
        </p>

        <!-- 当下切片微横幅 (NOW) -->
        <div
          v-if="now"
          class="pt-3 border-t border-[var(--border-subtle)]/60 flex flex-wrap items-center gap-x-6 gap-y-1 text-xs font-mono-archive text-[var(--ink-muted)]"
        >
          <span class="text-[var(--ink-primary)] font-medium">此刻 //</span>
          <span>正在构建：{{ now.building }}</span>
          <span>听：{{ now.listening }}</span>
          <button
            @click="emit('navigate', 'now')"
            class="hover:text-[var(--ink-primary)] hover-underline cursor-pointer ml-auto hidden sm:inline"
          >
            完整当下状态 →
          </button>
        </div>
      </div>
    </section>

    <!-- 可选折叠的 2026 生活刻度热力图 -->
    <section class="py-4 border-b border-[var(--border-subtle)]">
      <div class="flex items-center justify-between text-xs font-mono-archive text-[var(--ink-muted)]">
        <div class="flex items-center gap-2">
          <span class="text-[var(--ink-secondary)] font-medium">生活刻度 // 2026</span>
          <span>全年在册 {{ entries.length }} 条印记</span>
        </div>
        <button
          @click="showHeatmap = !showHeatmap"
          class="hover:text-[var(--ink-primary)] cursor-pointer hover-underline text-xs"
        >
          {{ showHeatmap ? '收起刻度图 ▲' : '展开刻度图 ▼' }}
        </button>
      </div>

      <div v-if="showHeatmap" class="pt-4 animate-in fade-in duration-200">
        <LifeActivityMap
          :entries="entries"
          :selectedDate="selectedDate"
          @selectDate="handleDateSelect"
        />
      </div>
    </section>

    <!-- 核心主角：纵向生活时间线 (Timeline / Life Stream) -->
    <section id="timeline-section" class="pt-10 sm:pt-14 pb-28">
      
      <!-- 动态标题与视角说明 -->
      <div class="flex items-baseline justify-between pb-6 mb-8 border-b border-[var(--border-subtle)]">
        <div class="space-y-1">
          <div class="flex items-center gap-2">
            <h2 class="text-lg font-serif-editorial font-medium tracking-tight text-[var(--ink-primary)]">
              {{ currentCategoryInfo.title }}
            </h2>
            <span class="text-xs font-mono-archive text-[var(--ink-muted)]">
              （当前视图 {{ entries.length }} 条）
            </span>
          </div>
          <p class="text-xs font-mono-archive text-[var(--ink-muted)]">
            {{ currentCategoryInfo.desc }}
          </p>
        </div>

        <button
          @click="emit('navigate', 'archive')"
          class="text-xs font-mono-archive text-[var(--ink-muted)] hover:text-[var(--ink-primary)] hover-underline cursor-pointer flex items-center gap-1"
        >
          <span>按年归档</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </button>
      </div>

      <!-- 纵向时间线 (带 7 重视角切换) -->
      <LifeStream
        :entries="entries"
        :limit="0"
        :showFilters="true"
        :dateFilter="selectedDate"
        :currentCategory="currentCategory"
        @clearDateFilter="selectedDate = ''"
        @changeCategory="emit('changeCategory', $event)"
        @selectEntry="emit('selectEntry', $event)"
        @openAdd="emit('openAdd')"
      />

    </section>

  </div>
</template>
