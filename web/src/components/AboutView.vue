<script setup lang="ts">
import { ArrowUpRight } from 'lucide-vue-next'
</script>

<template>
  <div class="max-w-3xl mx-auto px-5 sm:px-8 py-20 sm:py-28 text-left">
    
    <!-- 标头 -->
    <header class="pb-10 border-b border-[var(--border-subtle)] mb-14 space-y-3">
      <div class="flex items-center gap-3">
        <h1 class="font-serif-editorial text-4xl sm:text-5xl font-normal text-[var(--ink-primary)]">
          关于
        </h1>
        <span class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
          // ABOUT
        </span>
      </div>
      <p class="font-serif-editorial text-sm sm:text-base text-[var(--ink-secondary)]">
        一段关于我、这处空间与生活档案馆初衷的独白。
      </p>
    </header>

    <!-- 叙事正文：极其自然的人性化散文，绝非求职简历技能条 -->
    <div class="space-y-12 font-serif-editorial text-lg sm:text-xl text-[var(--ink-secondary)] leading-relaxed">
      
      <!-- 我是谁 -->
      <section class="space-y-4">
        <h2 class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
          // 我是谁 · WHO I AM
        </h2>
        <p class="text-[var(--ink-primary)]">
          我是 <strong class="font-medium">迟暮</strong>。现生活在杭州。
        </p>
        <p>
          我喜欢具体的、有体温的事物。清晨手冲咖啡时漫开的花果香气、骑着公路车穿过西湖杨公堤时迎面扑来的湿润晚风、老旁轴相机快门闭合时清脆的一声咔哒，以及在深夜书房里用 Go 雕琢出一套干净的并发状态机。
        </p>
        <p>
          我不喜欢虚饰的概念与空洞的头衔。生活终究是由一顿顿具体的饭、一趟趟具体的骑行、一次次真实的思索拼接而成的。
        </p>
      </section>

      <!-- 为什么做这个网站 -->
      <section class="space-y-4 pt-4 border-t border-[var(--border-subtle)]">
        <h2 class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
          // 为什么做这个网站 · WHY THIS ARCHIVE EXISTS
        </h2>
        <p class="text-[var(--ink-primary)] italic">
          “生活大于项目，真实大于表演。”
        </p>
        <p>
          在当下的互联网上，很多人的个人网站变成了单向度的“职业履历展厅”：满屏的技能徽章、KPI 指标、格式化的项目卡片与给招聘者看的说辞。
        </p>
        <p>
          但我渐渐发现，那些真正构成我生命质感的东西——读完一本书时的触动、和朋友在深夜连麦打游戏的畅快、秋天满觉陇落满青石板的金桂底片——在那些标准的简历模板里根本无处安放。
        </p>
        <p>
          所以我决定彻底推翻传统的作品集结构，亲手做这本<strong>数字生活档案馆</strong>。时间是这里唯一的刻度。无论是一张照片、一段随笔、一个亲手造的轮子，还是一张循环播放的唱片，它们都是我生命中真真切切发生过的事情。
        </p>
      </section>

      <!-- 喜欢与珍视的事物 -->
      <section class="space-y-4 pt-4 border-t border-[var(--border-subtle)]">
        <h2 class="font-mono-archive text-xs tracking-wider text-[var(--ink-muted)]">
          // 珍视的事物 · THINGS I VALUE
        </h2>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-6 text-base font-normal pt-2">
          <div class="space-y-2">
            <span class="font-mono-archive text-xs text-[var(--ink-muted)] block">// 精神给养</span>
            <p>坂本龙一、爵士乐、波西格《禅与摩托车维修艺术》、独立电影、安静的书房时光。</p>
          </div>
          <div class="space-y-2">
            <span class="font-mono-archive text-xs text-[var(--ink-muted)] block">// 日常手艺</span>
            <p>手冲咖啡 (V60 三段式)、西湖夜骑、Kodak Portra 胶卷、亲手编写的高性能小工具。</p>
          </div>
        </div>
      </section>

      <!-- 联系方式 -->
      <footer class="pt-8 border-t border-[var(--border-subtle)] font-mono-archive text-xs text-[var(--ink-muted)] flex flex-wrap items-center justify-between gap-4">
        <span>来信交流：chimu@codeactivityhub.top</span>
        <a
          href="https://github.com/YN1753"
          target="_blank"
          rel="noopener noreferrer"
          class="text-[var(--ink-primary)] hover-underline inline-flex items-center gap-1"
        >
          <span>GitHub @YN1753</span>
          <ArrowUpRight class="w-3 h-3" />
        </a>
      </footer>

    </div>

  </div>
</template>
