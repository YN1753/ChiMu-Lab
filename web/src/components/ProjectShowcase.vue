<script setup lang="ts">
import type { Project } from '../types'
import { Github, Box, ArrowUpRight, Cpu, Terminal, Network, RefreshCw, Layers } from 'lucide-vue-next'

defineProps<{
  projects: Project[]
}>()
</script>

<template>
  <section id="projects" class="py-20 border-t border-white/[0.06] relative">
    <div class="max-w-6xl mx-auto px-4 sm:px-6">
      <!-- 模块标题栏 -->
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-12 text-left">
        <div>
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-white/[0.04] border border-white/[0.08] text-cyan-400 text-xs font-mono mb-3">
            <Box class="w-3.5 h-3.5" />
            <span>Curated Showcase · 实验室展厅</span>
          </div>
          <h2 class="text-3xl sm:text-4xl font-extrabold tracking-tight text-white">
            工程实践与实验作品
          </h2>
          <p class="text-slate-400 text-sm sm:text-base mt-2 max-w-xl">
            不堆砌浮躁的 Demo。每一个工程均注重底层机制、确定性控制与真实生产可用性。
          </p>
        </div>

        <!-- 状态指示 -->
        <div class="flex items-center gap-2 px-3.5 py-1.5 rounded-xl bg-white/[0.03] border border-white/[0.07] text-xs font-mono text-slate-400 shrink-0">
          <Layers class="w-3.5 h-3.5 text-cyan-400" />
          <span>4 Production Projects</span>
        </div>
      </div>

      <!-- Bento Grid 便当盒网格布局 -->
      <div class="grid grid-cols-1 md:grid-cols-12 gap-6 text-left">
        <!-- 🌟 大 Bento 卡片 (占据 12 栏)：今日 · 面试练习站 (Daily Practice Hub) -->
        <div class="md:col-span-12 bento-card p-6 sm:p-8 bg-gradient-to-br from-[#0c0f1a] to-[#0a0d16] border border-white/[0.08]">
          <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-center">
            <div class="lg:col-span-7 space-y-4">
              <div class="flex items-center gap-3">
                <span class="px-2.5 py-0.5 rounded-md bg-cyan-950 text-cyan-400 border border-cyan-800 text-[11px] font-mono">
                  Featured Core Platform
                </span>
                <span class="text-xs text-slate-500 font-mono">Daily Coding Practice</span>
              </div>
              <h3 class="text-2xl sm:text-3xl font-extrabold text-white tracking-tight">
                今日 · 面试练习站
                <span class="text-sm font-normal text-slate-400 block sm:inline sm:ml-2">
                  (Daily Practice Hub)
                </span>
              </h3>
              <p class="text-slate-300 text-sm leading-relaxed max-w-xl">
                为大厂高频面试考点与 ACM 模式演练量身打造的每日刷题站。涵盖数组贪心约束、动态规划、双指针原型，支持白板闭卷推演与每日归档复习。
              </p>

              <!-- 技术标签 -->
              <div class="flex flex-wrap gap-2 pt-2">
                <span class="px-2.5 py-1 rounded-md bg-white/[0.04] border border-white/[0.06] text-xs font-mono text-slate-300">
                  #ACM模式
                </span>
                <span class="px-2.5 py-1 rounded-md bg-white/[0.04] border border-white/[0.06] text-xs font-mono text-slate-300">
                  #贪心算法
                </span>
                <span class="px-2.5 py-1 rounded-md bg-white/[0.04] border border-white/[0.06] text-xs font-mono text-slate-300">
                  #Go/C++双解
                </span>
                <span class="px-2.5 py-1 rounded-md bg-white/[0.04] border border-white/[0.06] text-xs font-mono text-slate-300">
                  #自动化发布
                </span>
              </div>

              <!-- 动作按钮 -->
              <div class="pt-4 flex items-center gap-4">
                <a
                  href="/interview"
                  class="px-5 py-2.5 rounded-xl bg-cyan-500 hover:bg-cyan-400 text-slate-950 font-bold text-xs flex items-center gap-2 shadow-lg shadow-cyan-500/20 transition-all cursor-pointer"
                >
                  <span>立即进入练习站</span>
                  <ArrowUpRight class="w-4 h-4" />
                </a>
                <a
                  href="https://github.com/YN1753"
                  target="_blank"
                  class="px-4 py-2.5 rounded-xl bg-white/[0.04] hover:bg-white/[0.08] text-slate-300 text-xs font-medium border border-white/[0.08] flex items-center gap-2 transition-all"
                >
                  <Github class="w-3.5 h-3.5" />
                  <span>开源仓库</span>
                </a>
              </div>
            </div>

            <!-- 右侧仿真题目卡片展示（展现真才实学，摆脱空壳感） -->
            <div class="lg:col-span-5">
              <div class="rounded-xl p-5 bg-[#07090e] border border-white/[0.08] shadow-inner font-mono text-xs space-y-3">
                <div class="flex items-center justify-between pb-2 border-b border-white/[0.06]">
                  <span class="text-cyan-400 font-semibold flex items-center gap-1.5">
                    <Terminal class="w-3.5 h-3.5" /> 今日精选题卡
                  </span>
                  <span class="text-[10px] text-emerald-400 bg-emerald-950/60 px-1.5 py-0.5 rounded border border-emerald-900">
                    AC Verified
                  </span>
                </div>
                <div class="text-slate-200 font-bold text-sm">
                  P3817 小A的糖果 (ACM模式)
                </div>
                <div class="p-2.5 rounded bg-white/[0.02] border border-white/[0.04] text-slate-400 text-[11px] leading-relaxed">
                  💡 <span class="text-amber-300 font-semibold">解题思路</span>：从左到右扫描，相邻两盒之和若超 x，<strong class="text-slate-200">仅扣减右盒糖果</strong>。改左盒会破坏前面约束，改右盒兼顾后序，贪心最优。
                </div>
                <div class="flex items-center justify-between text-[10px] text-slate-500 pt-1">
                  <span>时间复杂度: O(n)</span>
                  <span>空间复杂度: O(1)</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 🌟 Bento 卡片 2 (占据 6 栏)：ChiMu-Lab 架构 -->
        <div class="md:col-span-6 bento-card p-6 sm:p-7 flex flex-col justify-between">
          <div>
            <div class="flex items-center justify-between mb-4">
              <div class="p-2.5 rounded-xl bg-sky-950/80 border border-sky-800/80 text-sky-400">
                <Cpu class="w-5 h-5" />
              </div>
              <span class="text-[11px] font-mono px-2 py-0.5 rounded-full bg-emerald-950 text-emerald-400 border border-emerald-800">
                Active Production
              </span>
            </div>
            <h3 class="text-xl font-bold text-white mb-2">
              ChiMu-Lab 全栈工坊
            </h3>
            <p class="text-sm text-slate-300 leading-relaxed mb-5">
              本站核心架构。采用纯 Go 驱动的 SQLite (CGO-Free) 与 Gin 框架，配合 Vue 3 前端，实现全静态零动态依赖二进制部署，单进程常驻内存仅 ~3.8MB。
            </p>

            <!-- 架构流水线图示 -->
            <div class="p-3 rounded-lg bg-black/40 border border-white/[0.05] font-mono text-[11px] text-slate-400 space-y-1 mb-4">
              <div class="text-slate-500 text-[10px]">// Pipeline Architecture</div>
              <div class="text-cyan-400">Vue 3 (Vite + Tailwind v4)</div>
              <div class="text-slate-600">   ↓ HTTPS Reverse Proxy (Port 443)</div>
              <div class="text-emerald-400">Go 1.27 Core Server (Gin REST API)</div>
              <div class="text-slate-600">   ↓ GORM AutoMigrate</div>
              <div class="text-amber-400">Pure-Go SQLite (data/chimu.db)</div>
            </div>
          </div>

          <div class="pt-3 border-t border-white/[0.06] flex items-center justify-between">
            <span class="text-xs font-mono text-slate-400">#ZeroCGO #SPA #HighSpeed</span>
            <a href="https://github.com/YN1753/ChiMu-Lab" target="_blank" class="text-xs text-cyan-400 hover:underline flex items-center gap-1 font-mono">
              源码仓库 <ArrowUpRight class="w-3.5 h-3.5" />
            </a>
          </div>
        </div>

        <!-- 🌟 Bento 卡片 3 (占据 6 栏)：GopherSpace 探针 -->
        <div class="md:col-span-6 bento-card p-6 sm:p-7 flex flex-col justify-between">
          <div>
            <div class="flex items-center justify-between mb-4">
              <div class="p-2.5 rounded-xl bg-purple-950/80 border border-purple-800/80 text-purple-400">
                <Network class="w-5 h-5" />
              </div>
              <span class="text-[11px] font-mono px-2 py-0.5 rounded-full bg-cyan-950 text-cyan-400 border border-cyan-800">
                System Utility
              </span>
            </div>
            <h3 class="text-xl font-bold text-white mb-2">
              GopherSpace 分布式网络探针
            </h3>
            <p class="text-sm text-slate-300 leading-relaxed mb-5">
              轻量级并发网络连通性探测器。基于 Goroutine 池与原生系统调用设计，专用于高并发探测集群节点状态、TCP/HTTP 真实握手时延与保活心跳。
            </p>

            <div class="grid grid-cols-2 gap-2 text-xs font-mono mb-4">
              <div class="p-2.5 rounded-lg bg-white/[0.02] border border-white/[0.05]">
                <div class="text-slate-500 text-[10px]">CONCURRENCY</div>
                <div class="text-white font-bold mt-0.5">10k+ Goroutines</div>
              </div>
              <div class="p-2.5 rounded-lg bg-white/[0.02] border border-white/[0.05]">
                <div class="text-slate-500 text-[10px]">MEMORY OVERHEAD</div>
                <div class="text-emerald-400 font-bold mt-0.5">&lt; 15 MB</div>
              </div>
            </div>
          </div>

          <div class="pt-3 border-t border-white/[0.06] flex items-center justify-between">
            <span class="text-xs font-mono text-slate-400">#Syscall #GoroutinePool</span>
            <a href="https://github.com/YN1753" target="_blank" class="text-xs text-cyan-400 hover:underline flex items-center gap-1 font-mono">
              工程详情 <ArrowUpRight class="w-3.5 h-3.5" />
            </a>
          </div>
        </div>

        <!-- 🌟 Bento 卡片 4 (占据 12 栏横幅)：Hermes 自动化管线 -->
        <div class="md:col-span-12 bento-card p-6 sm:p-7 flex flex-col md:flex-row items-center justify-between gap-6">
          <div class="flex items-center gap-4">
            <div class="p-3 rounded-2xl bg-amber-950/80 border border-amber-800/80 text-amber-400 shrink-0">
              <RefreshCw class="w-6 h-6" />
            </div>
            <div>
              <h4 class="text-lg font-bold text-white">
                Hermes 自动化工作流与数据同步管线
              </h4>
              <p class="text-xs sm:text-sm text-slate-400 mt-1">
                定时触发打卡聚合、Markdown 静态站点编译与自动化热重载发布，全天候保障实验室站点的持续交付。
              </p>
            </div>
          </div>
          <div class="flex items-center gap-3 shrink-0">
            <span class="text-xs font-mono text-emerald-400 bg-emerald-950/60 px-3 py-1.5 rounded-lg border border-emerald-900">
              ● Automated Cron Active
            </span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
