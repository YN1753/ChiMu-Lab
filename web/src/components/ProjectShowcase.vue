<script setup lang="ts">
import { ref, computed } from 'vue'
import type { Project } from '../types'
import { FolderGit2, ExternalLink, Github, Sparkles, Box } from 'lucide-vue-next'

const props = defineProps<{
  projects: Project[]
}>()

const activeCategory = ref<string>('all')

const categories = [
  { id: 'all', name: '全部作品' },
  { id: 'core', name: '核心平台' },
  { id: 'lab', name: '实验项目' },
  { id: 'tool', name: '实用工具' },
]

const filteredProjects = computed(() => {
  if (activeCategory.value === 'all') {
    return props.projects
  }
  return props.projects.filter(p => p.category === activeCategory.value)
})

const getStatusBadge = (status: string) => {
  switch (status.toLowerCase()) {
    case 'active':
      return { text: '进行中 / 活跃', class: 'bg-emerald-950 text-emerald-400 border-emerald-800' }
    case 'stable':
      return { text: '稳定生产', class: 'bg-cyan-950 text-cyan-400 border-cyan-800' }
    case 'wip':
      return { text: '开发实验', class: 'bg-amber-950 text-amber-400 border-amber-800' }
    default:
      return { text: '已归档', class: 'bg-slate-800 text-slate-400 border-slate-700' }
  }
}
</script>

<template>
  <section id="projects" class="py-16 border-t border-white/5 relative">
    <div class="max-w-6xl mx-auto px-4 sm:px-6">
      <!-- 标题栏 -->
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-10">
        <div>
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-slate-900 border border-slate-800 text-cyan-400 text-xs font-mono mb-2">
            <Box class="w-3.5 h-3.5" />
            <span>Showcase & Experiments</span>
          </div>
          <h2 class="text-3xl font-bold tracking-tight text-white">
            作品展厅 · 实验室项目
          </h2>
          <p class="text-slate-400 text-sm mt-1 max-w-xl">
            涵盖全栈实战、算法沉淀、自动化管线与高并发工具，每一项实验都经过严谨的工程打磨。
          </p>
        </div>

        <!-- 筛选器标签 -->
        <div class="flex items-center gap-2 p-1.5 rounded-xl bg-slate-900 border border-slate-800 shrink-0">
          <button
            v-for="cat in categories"
            :key="cat.id"
            @click="activeCategory = cat.id"
            :class="[
              'px-3.5 py-1.5 rounded-lg text-xs font-medium transition-all cursor-pointer',
              activeCategory === cat.id
                ? 'bg-cyan-500 text-white shadow-md shadow-cyan-500/20'
                : 'text-slate-400 hover:text-slate-200'
            ]"
          >
            {{ cat.name }}
          </button>
        </div>
      </div>

      <!-- 项目卡片网格 -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div
          v-for="proj in filteredProjects"
          :key="proj.id"
          class="glass-panel rounded-2xl p-6 glow-card border border-white/10 flex flex-col justify-between group relative"
        >
          <!-- 卡片头部 -->
          <div>
            <div class="flex items-start justify-between gap-4 mb-4">
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 rounded-xl bg-slate-800 border border-slate-700 flex items-center justify-center text-cyan-400 group-hover:scale-105 transition-transform">
                  <FolderGit2 class="w-5 h-5" />
                </div>
                <div>
                  <h3 class="text-lg font-bold text-white group-hover:text-cyan-400 transition-colors">
                    {{ proj.title }}
                  </h3>
                  <p class="text-xs text-slate-400 font-medium">
                    {{ proj.subtitle }}
                  </p>
                </div>
              </div>

              <!-- 状态标签 -->
              <span
                :class="[
                  'text-[11px] font-mono px-2 py-0.5 rounded-full border shrink-0',
                  getStatusBadge(proj.status).class
                ]"
              >
                {{ getStatusBadge(proj.status).text }}
              </span>
            </div>

            <!-- 描述 -->
            <p class="text-sm text-slate-300 leading-relaxed mb-6">
              {{ proj.description }}
            </p>
          </div>

          <!-- 卡片底部技术栈与操作链接 -->
          <div>
            <!-- Tags -->
            <div class="flex flex-wrap gap-1.5 mb-5">
              <span
                v-for="tag in proj.tags.split(',')"
                :key="tag"
                class="px-2 py-0.5 rounded bg-slate-900 border border-slate-800 text-[11px] font-mono text-slate-400"
              >
                #{{ tag.trim() }}
              </span>
            </div>

            <!-- Action buttons -->
            <div class="pt-4 border-t border-slate-800/80 flex items-center justify-between">
              <div class="flex items-center gap-3">
                <a
                  v-if="proj.demo_url"
                  :href="proj.demo_url"
                  target="_blank"
                  class="inline-flex items-center gap-1.5 text-xs font-semibold text-cyan-400 hover:text-cyan-300 hover:underline"
                >
                  <span>在线体验 / 访问</span>
                  <ExternalLink class="w-3.5 h-3.5" />
                </a>

                <a
                  v-if="proj.github_url"
                  :href="proj.github_url"
                  target="_blank"
                  class="inline-flex items-center gap-1.5 text-xs font-semibold text-slate-400 hover:text-white"
                >
                  <Github class="w-3.5 h-3.5" />
                  <span>开源仓库</span>
                </a>
              </div>

              <span v-if="proj.featured" class="text-xs font-mono text-amber-400 flex items-center gap-1">
                <Sparkles class="w-3.5 h-3.5" /> 精选工程
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
