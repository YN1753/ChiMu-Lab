<script setup lang="ts">
import type { Project } from '../types'
import { ArrowUpRight, Github } from 'lucide-vue-next'

defineProps<{
  projects: Project[]
}>()
</script>

<template>
  <div class="max-w-4xl mx-auto px-5 sm:px-8 py-20 sm:py-28 text-left">
    
    <!-- 标头 -->
    <header class="pb-12 border-b border-[var(--border-subtle)] mb-14 space-y-4">
      <div class="flex items-center gap-3">
        <h1 class="font-serif-editorial text-4xl sm:text-5xl font-normal text-[var(--ink-primary)]">
          造物
        </h1>
        <span class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
          // THINGS I MADE
        </span>
      </div>
      <p class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
        手艺、工具与好奇心驱动的尝试。
      </p>
      <p class="text-base sm:text-lg text-[var(--ink-secondary)] font-serif-editorial max-w-2xl leading-relaxed pt-2">
        编程只是我体验生活、表达创造力的一门手艺。就像手冲一杯咖啡、或者打磨一件木器一样，在逻辑的泥土里亲手塑造出确定性的东西。这里记录着我做过的工具、踩过的坑，以及那些在大学和深夜工位里诞生的真实经历。
      </p>
    </header>

    <!-- 项目故事清单 (Editorial Stories, 非模板化 Card 墙) -->
    <div class="divide-y divide-[var(--border-subtle)] space-y-16">
      
      <article
        v-for="(proj, idx) in projects"
        :key="proj.id"
        class="pt-14 first:pt-0 space-y-6"
      >
        <!-- 序号、状态与归属 -->
        <div class="flex items-center justify-between text-xs font-mono-archive text-[var(--ink-muted)]">
          <div class="flex items-center gap-2">
            <span class="text-[var(--ink-primary)] font-semibold">0{{ idx + 1 }}.</span>
            <span class="tracking-wider">{{ proj.category }}</span>
          </div>
          <span class="border border-[var(--border-subtle)] px-2 py-0.5 rounded-xs text-[10px] uppercase">
            {{ proj.status }}
          </span>
        </div>

        <!-- 标题与副标 -->
        <div>
          <h2 class="text-2xl sm:text-3xl font-serif-editorial font-medium text-[var(--ink-primary)]">
            {{ proj.title }}
          </h2>
          <p class="text-sm font-mono-archive text-[var(--ink-muted)] mt-1">
            {{ proj.subtitle }}
          </p>
        </div>

        <!-- 人生经历叙事 (Story: 为什么开始、怎么做、遇到了什么) -->
        <div class="space-y-3 font-serif-editorial text-base sm:text-lg text-[var(--ink-secondary)] leading-relaxed max-w-2xl">
          <p class="text-[var(--ink-primary)] font-medium">
            {{ proj.description }}
          </p>
          <p v-if="proj.story" class="text-[var(--ink-secondary)] italic text-sm sm:text-base border-l-2 border-[var(--border-subtle)] pl-4 py-1">
            “{{ proj.story }}”
          </p>
        </div>

        <!-- 标签与外部链接 -->
        <div class="pt-2 flex flex-wrap items-center justify-between gap-4 text-xs font-mono-archive">
          <div class="flex flex-wrap gap-2 text-[var(--ink-muted)]">
            <span v-for="tag in proj.tags.split(',')" :key="tag">
              #{{ tag.trim() }}
            </span>
          </div>

          <a
            :href="proj.github_url"
            target="_blank"
            rel="noopener noreferrer"
            class="inline-flex items-center gap-1.5 text-[var(--ink-primary)] hover-underline pb-0.5"
          >
            <Github class="w-3.5 h-3.5" />
            <span>查看 GitHub 源码</span>
            <ArrowUpRight class="w-3 h-3" />
          </a>
        </div>

      </article>

    </div>

  </div>
</template>
