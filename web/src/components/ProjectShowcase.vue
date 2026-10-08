<script setup lang="ts">
import type { Project } from '../types'
import { Github, ArrowUpRight, Cpu, Terminal, Network, RefreshCw, Film } from 'lucide-vue-next'

defineProps<{
  projects: Project[]
}>()
</script>

<template>
  <section id="projects" class="py-24 border-t border-[#e8e6df] bg-[#fbfbfa] relative">
    <div class="max-w-6xl mx-auto px-5 sm:px-8 text-left">
      <!-- 分镜标头 (Scene Header) -->
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-14">
        <div>
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-[#edeae1] text-[#716e64] text-xs font-mono mb-3">
            <Film class="w-3.5 h-3.5 text-[#d97706]" />
            <span>Act II · Storyboard & Works · 作品分镜展厅</span>
          </div>
          <h2 class="text-3xl sm:text-4xl font-medium tracking-tight text-[#14151a] font-serif-cinematic">
            工程实践与实验作品
          </h2>
          <p class="text-[#525662] text-sm sm:text-base mt-2 max-w-xl font-normal">
            拒绝扁平浮躁的 Demo。每一个作品如同一部构思严密的分镜，追求确定性的系统机制与真实的生产可用性。
          </p>
        </div>

        <div class="flex items-center gap-2 text-xs font-mono text-[#8c8f9b] bg-white border border-[#e8e6df] px-3.5 py-1.5 rounded-full shadow-2xs">
          <span>4 Principal Projects</span>
        </div>
      </div>

      <!-- 宽银幕分镜卡片布局 -->
      <div class="grid grid-cols-1 md:grid-cols-12 gap-7">
        <!-- 🌟 主分镜大卡片 (12 栏)：今日 · 面试练习站 (Daily Practice Hub) -->
        <div class="md:col-span-12 film-card p-7 sm:p-9 bg-white border border-[#e8e6df]">
          <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-center">
            <!-- 左侧内容与简介 -->
            <div class="lg:col-span-7 space-y-5">
              <div class="flex items-center gap-3">
                <span class="px-2.5 py-0.5 rounded-full bg-[#fef3c7] text-[#92400e] border border-[#fde68a] text-[11px] font-mono font-medium">
                  Sequence 01 · Featured Hub
                </span>
                <span class="text-xs text-[#8c8f9b] font-mono">Algorithm Practice</span>
              </div>
              <h3 class="text-2xl sm:text-3xl font-bold text-[#14151a] font-serif-cinematic tracking-tight">
                今日 · 面试练习站
                <span class="text-sm font-normal text-[#8c8f9b] block sm:inline sm:ml-2 font-sans">
                  (Daily Practice Hub)
                </span>
              </h3>
              <p class="text-[#525662] text-sm leading-relaxed max-w-xl">
                为大厂高频面试考点与 ACM 模式演练量身打造的每日刷题平台。涵盖相邻约束贪心算法、动态规划、双指针原型，支持白板闭卷推演与每日归档回溯。
              </p>

              <!-- 技术标签 -->
              <div class="flex flex-wrap gap-2 pt-1">
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #ACM模式
                </span>
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #相邻约束贪心
                </span>
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #Go/C++双解
                </span>
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #静态自动化生成
                </span>
              </div>

              <!-- 动作按钮 -->
              <div class="pt-3 flex items-center gap-4">
                <a
                  href="/interview"
                  class="px-5 py-2.5 rounded-xl bg-[#14151a] hover:bg-[#252834] text-white font-medium text-xs flex items-center gap-2 shadow-sm transition-all cursor-pointer"
                >
                  <span>立即进入练习站</span>
                  <ArrowUpRight class="w-4 h-4 text-slate-300" />
                </a>
                <a
                  href="https://github.com/YN1753"
                  target="_blank"
                  class="px-4 py-2.5 rounded-xl bg-white hover:bg-[#faf9f5] text-[#525662] hover:text-[#14151a] text-xs font-medium border border-[#e8e6df] flex items-center gap-2 transition-all shadow-2xs"
                >
                  <Github class="w-3.5 h-3.5" />
                  <span>开源仓库</span>
                </a>
              </div>
            </div>

            <!-- 右侧精装印刷质感算法切片卡片 -->
            <div class="lg:col-span-5">
              <div class="rounded-xl p-5 sm:p-6 bg-[#faf9f5] border border-[#e8e6df] shadow-inner font-mono text-xs space-y-3">
                <div class="flex items-center justify-between pb-2 border-b border-[#ece8df]">
                  <span class="text-[#92400e] font-semibold flex items-center gap-1.5">
                    <Terminal class="w-3.5 h-3.5" /> 今日精选题卡 / ACM
                  </span>
                  <span class="text-[10px] text-[#0d766e] bg-[#f0fdf4] px-1.5 py-0.5 rounded border border-[#bbf7d0]">
                    AC 验证通过
                  </span>
                </div>
                <div class="text-[#14151a] font-bold text-sm">
                  P3817 小A的糖果 (洛谷经典原题)
                </div>
                <div class="p-3 rounded-lg bg-white border border-[#e8e6df] text-[#525662] text-[11px] leading-relaxed">
                  💡 <strong class="text-[#14151a]">贪心核心推导</strong>：扫描相邻两盒之和若超 x，<span class="text-[#92400e] font-semibold">只扣减右盒糖果</span>。改左盒仅满足当前对，改右盒不仅满足当前对，还能为后续 (i+1, i+2) 创造最大松弛度。
                </div>
                <div class="flex items-center justify-between text-[10px] text-[#8c8f9b] pt-1">
                  <span>时间复杂度: O(n)</span>
                  <span>空间复杂度: O(1)</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 🌟 分镜卡片 2 (6 栏)：ChiMu-Lab 架构 -->
        <div class="md:col-span-6 film-card p-7 sm:p-8 flex flex-col justify-between bg-white border border-[#e8e6df]">
          <div>
            <div class="flex items-center justify-between mb-4">
              <div class="p-2.5 rounded-xl bg-[#edeae1] text-[#716e64]">
                <Cpu class="w-5 h-5 text-[#d97706]" />
              </div>
              <span class="text-[11px] font-mono px-2 py-0.5 rounded-full bg-[#f0fdf4] text-[#0d766e] border border-[#bbf7d0]">
                Live Production
              </span>
            </div>
            <h3 class="text-xl font-bold text-[#14151a] font-serif-cinematic mb-2">
              ChiMu-Lab 全栈工坊
            </h3>
            <p class="text-sm text-[#525662] leading-relaxed mb-5">
              本站核心架构。采用纯 Go 编写的 SQLite (CGO-Free) 与 Gin 框架，配合 Vue 3 极简前端，构建单二进制静态部署架构，服务常驻内存仅 ~1.76MB。
            </p>

            <!-- 印刷质感架构流 -->
            <div class="p-3.5 rounded-xl bg-[#faf9f5] border border-[#e8e6df] font-mono text-[11px] text-[#525662] space-y-1 mb-4">
              <div class="text-[#8c8f9b] text-[10px]">// Pipeline Architecture</div>
              <div class="text-[#14151a] font-medium">Vue 3 (Vite + Tailwind) 前端</div>
              <div class="text-[#8c8f9b]">   ↓ Nginx 1.24 (Port 443 / TLS 1.3)</div>
              <div class="text-[#0d766e] font-medium">Go 1.27 Gin REST API Core</div>
              <div class="text-[#8c8f9b]">   ↓ GORM AutoMigrate</div>
              <div class="text-[#92400e] font-medium">Pure-Go SQLite (data/chimu.db)</div>
            </div>
          </div>

          <div class="pt-4 border-t border-[#f0ede6] flex items-center justify-between">
            <span class="text-xs font-mono text-[#8c8f9b]">#ZeroCGO #SPA #HighSpeed</span>
            <a href="https://github.com/YN1753/ChiMu-Lab" target="_blank" class="text-xs text-[#14151a] font-semibold hover:underline flex items-center gap-1 font-mono">
              开源代码 <ArrowUpRight class="w-3.5 h-3.5" />
            </a>
          </div>
        </div>

        <!-- 🌟 分镜卡片 3 (6 栏)：GopherSpace 探针 -->
        <div class="md:col-span-6 film-card p-7 sm:p-8 flex flex-col justify-between bg-white border border-[#e8e6df]">
          <div>
            <div class="flex items-center justify-between mb-4">
              <div class="p-2.5 rounded-xl bg-[#edeae1] text-[#716e64]">
                <Network class="w-5 h-5 text-[#d97706]" />
              </div>
              <span class="text-[11px] font-mono px-2 py-0.5 rounded-full bg-[#edeae1] text-[#716e64] border border-[#dedcd5]">
                System Utility
              </span>
            </div>
            <h3 class="text-xl font-bold text-[#14151a] font-serif-cinematic mb-2">
              GopherSpace 分布式网络探针
            </h3>
            <p class="text-sm text-叙述 text-[#525662] leading-relaxed mb-5">
              轻量级并发网络连通性探测器。基于 Goroutine 池与底层原生网络系统调用设计，专用于高并发探测集群节点状态、TCP/HTTP 真实握手时延与保活心跳。
            </p>

            <div class="grid grid-cols-2 gap-3 text-xs font-mono mb-4">
              <div class="p-3 rounded-lg bg-[#faf9f5] border border-[#e8e6df]">
                <div class="text-[#8c8f9b] text-[10px]">CONCURRENCY</div>
                <div class="text-[#14151a] font-bold mt-0.5">10k+ Goroutines</div>
              </div>
              <div class="p-3 rounded-lg bg-[#faf9f5] border border-[#e8e6df]">
                <div class="text-[#8c8f9b] text-[10px]">MEMORY OVERHEAD</div>
                <div class="text-[#0d766e] font-bold mt-0.5">&lt; 15 MB</div>
              </div>
            </div>
          </div>

          <div class="pt-4 border-t border-[#f0ede6] flex items-center justify-between">
            <span class="text-xs font-mono text-[#8c8f9b]">#Syscall #GoroutinePool</span>
            <a href="https://github.com/YN1753" target="_blank" class="text-xs text-[#14151a] font-semibold hover:underline flex items-center gap-1 font-mono">
              工程详情 <ArrowUpRight class="w-3.5 h-3.5" />
            </a>
          </div>
        </div>

        <!-- 🌟 分镜卡片 4 (12 栏横幅)：Hermes 自动化管线 -->
        <div class="md:col-span-12 film-card p-7 sm:p-8 flex flex-col md:flex-row items-center justify-between gap-6 bg-white border border-[#e8e6df]">
          <div class="flex items-center gap-5">
            <div class="p-3 rounded-2xl bg-[#edeae1] text-[#716e64] shrink-0">
              <RefreshCw class="w-6 h-6 text-[#d97706]" />
            </div>
            <div>
              <h4 class="text-lg font-bold text-[#14151a] font-serif-cinematic">
                Hermes 自动化工作流与数据同步管线
              </h4>
              <p class="text-xs sm:text-sm text-[#525662] mt-1 font-normal">
                定时触发打卡聚合、Markdown 静态站点编译与自动化热重载发布，全天候保障实验室站点的持续交付。
              </p>
            </div>
          </div>
          <div class="flex items-center gap-3 shrink-0">
            <span class="text-xs font-mono text-[#0d766e] bg-[#f0fdf4] px-3.5 py-1.5 rounded-full border border-[#bbf7d0]">
              ● Automated Cron Active
            </span>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
