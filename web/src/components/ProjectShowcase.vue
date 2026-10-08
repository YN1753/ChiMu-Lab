<script setup lang="ts">
import type { Project } from '../types'
import { Github, ArrowUpRight, Eye, BookOpen, Layers, Terminal, Sparkles } from 'lucide-vue-next'

defineProps<{
  projects: Project[]
}>()
</script>

<template>
  <section id="projects" class="py-24 border-t border-[#e8e6df] bg-[#fbfbfa] relative">
    <div class="max-w-6xl mx-auto px-5 sm:px-8 text-left">
      <!-- 分镜标头 -->
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-14">
        <div>
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-[#edeae1] text-[#716e64] text-xs font-mono mb-3">
            <Layers class="w-3.5 h-3.5 text-[#d97706]" />
            <span>The Workshop · 个人真实仓库与工程陈列</span>
          </div>
          <h2 class="text-3xl sm:text-4xl font-medium tracking-tight text-[#14151a] font-serif-cinematic">
            造过的轮子与真实探索
          </h2>
          <p class="text-[#525662] text-sm sm:text-base mt-2 max-w-xl font-normal">
            全部来自我的 GitHub 真实仓库。不作虚假包装，仅记录自己亲手写下的代码与系统思考。
          </p>
        </div>

        <div class="flex items-center gap-2 text-xs font-mono text-[#8c8f9b] bg-white border border-[#e8e6df] px-3.5 py-1.5 rounded-full shadow-2xs">
          <span>7 Active Repositories</span>
        </div>
      </div>

      <!-- 真实项目陈列 -->
      <div class="grid grid-cols-1 md:grid-cols-12 gap-7">
        <!-- 🥇 核心置顶 #1：SUSE-OAA-BACKEND (大分镜 12 栏) -->
        <div class="md:col-span-12 film-card p-8 sm:p-10 bg-white border border-[#e8e6df]">
          <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-center">
            <div class="lg:col-span-7 space-y-4">
              <div class="flex items-center gap-3">
                <span class="px-2.5 py-0.5 rounded-full bg-[#fef3c7] text-[#92400e] border border-[#fde68a] text-[11px] font-mono font-medium">
                  #1 Core Pinned · suse-edu-cn
                </span>
                <span class="text-xs text-[#8c8f9b] font-mono">Campus OpenSource</span>
              </div>
              <h3 class="text-2xl sm:text-3xl font-bold text-[#14151a] font-serif-cinematic tracking-tight">
                SUSE-OAA-BACKEND
              </h3>
              <p class="text-[#525662] text-sm leading-relaxed max-w-xl">
                为四川轻化工大学开源协会（suse-edu-cn）量身打造的核心业务后台。基于 Go 语言构建，支撑协会事务协作、组织管理、招新流程与校园数据服务接口。
              </p>

              <div class="flex flex-wrap gap-2 pt-1">
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #Go
                </span>
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #Gin
                </span>
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #suse-edu-cn
                </span>
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #CampusOpenSource
                </span>
              </div>

              <div class="pt-3 flex items-center gap-4">
                <a
                  href="https://github.com/suse-edu-cn/SUSE-OAA-BACKEND"
                  target="_blank"
                  class="px-5 py-2.5 rounded-xl bg-[#14151a] hover:bg-[#252834] text-white font-medium text-xs flex items-center gap-2 shadow-sm transition-all"
                >
                  <Github class="w-3.5 h-3.5" />
                  <span>访问 GitHub 仓库</span>
                  <ArrowUpRight class="w-4 h-4 text-slate-300" />
                </a>
              </div>
            </div>

            <!-- 右侧：组织与技术栈速览卡片 -->
            <div class="lg:col-span-5">
              <div class="rounded-xl p-6 bg-[#faf9f5] border border-[#e8e6df] font-mono text-xs space-y-3">
                <div class="flex items-center justify-between pb-2 border-b border-[#ece8df]">
                  <span class="text-[#92400e] font-semibold flex items-center gap-1.5">
                    <Terminal class="w-3.5 h-3.5" /> suse-edu-cn 组织置顶
                  </span>
                  <span class="text-[10px] text-[#0d766e] bg-[#f0fdf4] px-1.5 py-0.5 rounded border border-[#bbf7d0]">
                    Public Repo
                  </span>
                </div>
                <div class="text-[#14151a] font-semibold text-sm">
                  四川轻化工大学 开放原子开源协会
                </div>
                <div class="p-3 rounded-lg bg-white border border-[#e8e6df] text-[#525662] text-[11px] leading-relaxed">
                  负责协会核心微服务架构设计、数据库实体定义、鉴权中间件编写与生产环境接口部署交付。
                </div>
                <div class="flex items-center justify-between text-[10px] text-[#8c8f9b] pt-1">
                  <span>Language: Go</span>
                  <span>Team: suse-edu-cn</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 🥈 核心置顶 #2：ArchCanvas (大分镜 12 栏) -->
        <div class="md:col-span-12 film-card p-8 sm:p-10 bg-white border border-[#e8e6df]">
          <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-center">
            <div class="lg:col-span-7 space-y-4">
              <div class="flex items-center gap-3">
                <span class="px-2.5 py-0.5 rounded-full bg-[#fef3c7] text-[#92400e] border border-[#fde68a] text-[11px] font-mono font-medium">
                  #2 Active Development · YN1753
                </span>
                <span class="text-xs text-[#8c8f9b] font-mono">AI Architecture Tool</span>
              </div>
              <h3 class="text-2xl sm:text-3xl font-bold text-[#14151a] font-serif-cinematic tracking-tight">
                ArchCanvas
              </h3>
              <p class="text-[#525662] text-sm leading-relaxed max-w-xl">
                让 AI 与开发者一起，从需求语义理解、ER 实体关系图到系统架构设计，快速构建与生成可运行的 Go 工程骨架。探索 AI 与现代架构设计工具的深度融合。
              </p>

              <div class="flex flex-wrap gap-2 pt-1">
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #TypeScript
                </span>
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #Go
                </span>
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #Architecture
                </span>
                <span class="px-2.5 py-1 rounded-md bg-[#faf9f5] border border-[#e8e6df] text-xs font-mono text-[#525662]">
                  #AI Canvas
                </span>
              </div>

              <div class="pt-3 flex items-center gap-4">
                <a
                  href="https://github.com/YN1753/ArchCanvas"
                  target="_blank"
                  class="px-5 py-2.5 rounded-xl bg-[#14151a] hover:bg-[#252834] text-white font-medium text-xs flex items-center gap-2 shadow-sm transition-all"
                >
                  <Github class="w-3.5 h-3.5" />
                  <span>访问 GitHub 仓库</span>
                  <ArrowUpRight class="w-4 h-4 text-slate-300" />
                </a>
              </div>
            </div>

            <!-- 右侧：设计理念切片 -->
            <div class="lg:col-span-5">
              <div class="rounded-xl p-6 bg-[#faf9f5] border border-[#e8e6df] font-mono text-xs space-y-3">
                <div class="flex items-center justify-between pb-2 border-b border-[#ece8df]">
                  <span class="text-[#92400e] font-semibold flex items-center gap-1.5">
                    <Sparkles class="w-3.5 h-3.5" /> AI × 架构可视化
                  </span>
                  <span class="text-[10px] text-[#0d766e] bg-[#f0fdf4] px-1.5 py-0.5 rounded border border-[#bbf7d0]">
                    In Active Progress
                  </span>
                </div>
                <div class="text-[#14151a] font-semibold text-sm">
                  从需求文本 ➔ ER 关系 ➔ 可运行 Go 工程
                </div>
                <div class="p-3 rounded-lg bg-white border border-[#e8e6df] text-[#525662] text-[11px] leading-relaxed">
                  不再手写样板代码。将系统边界、领域模型与数据库关系在画布上即时呈现，实现工程逻辑的可视化推演。
                </div>
                <div class="flex items-center justify-between text-[10px] text-[#8c8f9b] pt-1">
                  <span>Stack: TypeScript + Go</span>
                  <span>Author: YN1753</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 🥉 项目 #3：GoLens (6 栏) -->
        <div class="md:col-span-6 film-card p-7 sm:p-8 flex flex-col justify-between bg-white border border-[#e8e6df]">
          <div>
            <div class="flex items-center justify-between mb-4">
              <div class="p-2.5 rounded-xl bg-[#edeae1] text-[#716e64]">
                <Eye class="w-5 h-5 text-[#d97706]" />
              </div>
              <span class="text-[11px] font-mono px-2 py-0.5 rounded-full bg-[#f0fdf4] text-[#0d766e] border border-[#bbf7d0]">
                Go Internals
              </span>
            </div>
            <h3 class="text-xl font-bold text-[#14151a] font-serif-cinematic mb-2">
              GoLens · Go 底层机制透视镜
            </h3>
            <p class="text-sm text-[#525662] leading-relaxed mb-5">
              基于交互式状态机的 Go 底层机制透视镜。让 GMP 协程调度、三色标记 GC 屏障与 Channel 阻塞状态流转清晰可见，将抽象的运行时原理直观化。
            </p>
            <div class="flex flex-wrap gap-1.5 mb-4">
              <span class="px-2 py-0.5 rounded bg-[#faf9f5] border border-[#e8e6df] text-[11px] font-mono text-[#525662]">#GMP</span>
              <span class="px-2 py-0.5 rounded bg-[#faf9f5] border border-[#e8e6df] text-[11px] font-mono text-[#525662]">#GC屏障</span>
              <span class="px-2 py-0.5 rounded bg-[#faf9f5] border border-[#e8e6df] text-[11px] font-mono text-[#525662]">#Channel</span>
            </div>
          </div>
          <div class="pt-4 border-t border-[#f0ede6] flex items-center justify-between">
            <span class="text-xs font-mono text-[#8c8f9b]">#Visualization</span>
            <a href="https://github.com/YN1753/GoLens" target="_blank" class="text-xs text-[#14151a] font-semibold hover:underline flex items-center gap-1 font-mono">
              查看仓库 <ArrowUpRight class="w-3.5 h-3.5" />
            </a>
          </div>
        </div>

        <!-- 📦 项目 #4：AstraLink-Desktop (6 栏) -->
        <div class="md:col-span-6 film-card p-7 sm:p-8 flex flex-col justify-between bg-white border border-[#e8e6df]">
          <div>
            <div class="flex items-center justify-between mb-4">
              <div class="p-2.5 rounded-xl bg-[#edeae1] text-[#716e64]">
                <BookOpen class="w-5 h-5 text-[#d97706]" />
              </div>
              <span class="text-[11px] font-mono px-2 py-0.5 rounded-full bg-[#edeae1] text-[#716e64] border border-[#dedcd5]">
                Wails Desktop
              </span>
            </div>
            <h3 class="text-xl font-bold text-[#14151a] font-serif-cinematic mb-2">
              AstraLink-Desktop · 星链 2.0
            </h3>
            <p class="text-sm text-[#525662] leading-relaxed mb-5">
              基于 Wails 架构的双链图笔记桌面端应用。探索 Go + 前端混合开发，支持双向链接、知识图谱可视化与本地隐私优先的个人知识库体系。
            </p>
            <div class="flex flex-wrap gap-1.5 mb-4">
              <span class="px-2 py-0.5 rounded bg-[#faf9f5] border border-[#e8e6df] text-[11px] font-mono text-[#525662]">#Go</span>
              <span class="px-2 py-0.5 rounded bg-[#faf9f5] border border-[#e8e6df] text-[11px] font-mono text-[#525662]">#Wails</span>
              <span class="px-2 py-0.5 rounded bg-[#faf9f5] border border-[#e8e6df] text-[11px] font-mono text-[#525662]">#Vue</span>
              <span class="px-2 py-0.5 rounded bg-[#faf9f5] border border-[#e8e6df] text-[11px] font-mono text-[#525662]">#Graph</span>
            </div>
          </div>
          <div class="pt-4 border-t border-[#f0ede6] flex items-center justify-between">
            <span class="text-xs font-mono text-[#8c8f9b]">#KnowledgeBase</span>
            <a href="https://github.com/YN1753/AstraLink-Desktop" target="_blank" class="text-xs text-[#14151a] font-semibold hover:underline flex items-center gap-1 font-mono">
              查看仓库 <ArrowUpRight class="w-3.5 h-3.5" />
            </a>
          </div>
        </div>

        <!-- ⚡ 项目 #5：Go-Load (4 栏) -->
        <div class="md:col-span-4 film-card p-6 flex flex-col justify-between bg-white border border-[#e8e6df]">
          <div>
            <div class="flex items-center justify-between mb-3">
              <span class="text-xs font-mono text-[#92400e] font-semibold">Go-Load</span>
              <span class="text-[10px] font-mono px-2 py-0.5 rounded bg-[#f0fdf4] text-[#0d766e]">Stable</span>
            </div>
            <h4 class="text-base font-bold text-[#14151a] font-serif-cinematic mb-1.5">
              轻量高并发压测工具
            </h4>
            <p class="text-xs text-[#525662] leading-relaxed mb-4">
              基于 Go 原生并发模型的高性能 HTTP 压测工具，轻量无依赖，提供精确的时延与吞吐量统计。
            </p>
          </div>
          <div class="pt-3 border-t border-[#f0ede6] flex items-center justify-between">
            <span class="text-[11px] font-mono text-[#8c8f9b]">#Benchmark #CLI</span>
            <a href="https://github.com/YN1753/Go-Load" target="_blank" class="text-xs text-[#14151a] font-semibold flex items-center gap-1">
              Repo <ArrowUpRight class="w-3 h-3" />
            </a>
          </div>
        </div>

        <!-- 🖥️ 项目 #6：nexus (4 栏) -->
        <div class="md:col-span-4 film-card p-6 flex flex-col justify-between bg-white border border-[#e8e6df]">
          <div>
            <div class="flex items-center justify-between mb-3">
              <span class="text-xs font-mono text-[#716e64] font-semibold">nexus</span>
              <span class="text-[10px] font-mono px-2 py-0.5 rounded bg-[#faf9f5] text-[#716e64]">Linux Ops</span>
            </div>
            <h4 class="text-base font-bold text-[#14151a] font-serif-cinematic mb-1.5">
              Linux 服务器管理平台
            </h4>
            <p class="text-xs text-[#525662] leading-relaxed mb-4">
              面向开发者的现代化 Linux 服务器管理控制台，用于便捷监控云服务器硬件、容器与核心守护进程。
            </p>
          </div>
          <div class="pt-3 border-t border-[#f0ede6] flex items-center justify-between">
            <span class="text-[11px] font-mono text-[#8c8f9b]">#Vue #Linux</span>
            <a href="https://github.com/YN1753/nexus" target="_blank" class="text-xs text-[#14151a] font-semibold flex items-center gap-1">
              Repo <ArrowUpRight class="w-3 h-3" />
            </a>
          </div>
        </div>

        <!-- 📱 项目 #7：DeviceDaily (4 栏) -->
        <div class="md:col-span-4 film-card p-6 flex flex-col justify-between bg-white border border-[#e8e6df]">
          <div>
            <div class="flex items-center justify-between mb-3">
              <span class="text-xs font-mono text-[#716e64] font-semibold">DeviceDaily</span>
              <span class="text-[10px] font-mono px-2 py-0.5 rounded bg-[#f0fdf4] text-[#0d766e]">macOS</span>
            </div>
            <h4 class="text-base font-bold text-[#14151a] font-serif-cinematic mb-1.5">
              设备成本统计原生 App
            </h4>
            <p class="text-xs text-[#525662] leading-relaxed mb-4">
              基于 Swift 原生开发，支持 macOS 桌面原生 Widget 小组件，计算并追踪电子设备的日均使用成本。
            </p>
          </div>
          <div class="pt-3 border-t border-[#f0ede6] flex items-center justify-between">
            <span class="text-[11px] font-mono text-[#8c8f9b]">#Swift #WidgetKit</span>
            <a href="https://github.com/YN1753/DeviceDaily" target="_blank" class="text-xs text-[#14151a] font-semibold flex items-center gap-1">
              Repo <ArrowUpRight class="w-3 h-3" />
            </a>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
