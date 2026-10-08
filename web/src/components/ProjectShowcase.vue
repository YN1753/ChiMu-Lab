<script setup lang="ts">
import { ref } from 'vue'
import type { Project } from '../types'
import { Github, ArrowUpRight, Eye, Layers, Terminal, X, Copy, Check, ExternalLink, ShieldCheck } from 'lucide-vue-next'
import { audio } from '../utils/audio'

const props = defineProps<{
  projects: Project[]
}>()

// 选中的档案底片抽屉
const activeModalProject = ref<Project | null>(null)
const copied = ref(false)

const openProjectModal = (proj: Project) => {
  activeModalProject.value = proj
  copied.value = false
  audio.playShutter()
}

const closeModal = () => {
  activeModalProject.value = null
  audio.playTink()
}

const copyCloneCommand = (url: string) => {
  navigator.clipboard.writeText(`git clone ${url}.git`)
  copied.value = true
  audio.playTink()
  setTimeout(() => {
    copied.value = false
  }, 2000)
}

// 对应每个项目的详细架构与初衷档案
const getProjectDossier = (title: string) => {
  switch (title) {
    case 'SUSE-OAA-BACKEND':
      return {
        org: '四川轻化工大学开放原子开源协会 (suse-edu-cn)',
        origin: '为高校开源协会搭建独立自主、高可用、可长期维护的业务后端体系，支撑招新流转、技术协作与校园开源活动。',
        highlights: [
          '分层架构：Controller -> Service -> Repository 清晰解耦',
          '中间件闭环：基于 JWT 的上下文鉴权与动态权限校验',
          '高内聚领域：活动事件、成员档案、面试流转三位一体',
        ],
        codeHint: 'go run main.go --config=configs/config.yaml',
      }
    case 'ArchCanvas':
      return {
        org: '迟暮个人实验仓库 (YN1753)',
        origin: '探索 AI 时代软件架构设计的表达方式。摆脱传统孤立的代码编写，从需求语义直接投射到 ER 实体与 Go struct 架构全景。',
        highlights: [
          '双向同步：画布节点拖拽与 Go AST 代码 AST 双向保持一致',
          '自动建模：从对话需求直接输出标准化 GORM / Ent 架构骨架',
          '架构可视化：将抽象的微服务边界转化为物理拓扑图谱',
        ],
        codeHint: 'npm run dev # 启动 AI 架构画布编辑器',
      }
    case 'GoLens':
      return {
        org: '迟暮个人实验仓库 (YN1753)',
        origin: 'Go 语言底层机制晦涩抽象。用状态机和交互式动画把 GMP 调度、三色标记 GC 和 Channel 缓冲在浏览器中直观展现。',
        highlights: [
          'GMP 真实模拟：偷取队列 (Work Stealing) 与 sysmon 抢占可视化',
          '三色标记法：灰色工作集遍历与写屏障 (Write Barrier) 动态推演',
          '通道交互：无缓冲与有缓冲通道的锁竞争与唤醒演示',
        ],
        codeHint: 'open index.html # 基于原生状态机的无依赖透视镜',
      }
    case 'AstraLink-Desktop':
      return {
        org: '迟暮个人实验仓库 (YN1753)',
        origin: '探索 Go 语言在现代桌面应用上的表现。基于 Wails 2.0 打造双向链接图笔记，数据 100% 留存在本地 SQLite。',
        highlights: [
          'Wails 架构：原生 WebKit 视图绑定 Go 核心高性能后端',
          '本地隐私优先：无任何第三方上云侵入，SQLite 纯本地加密',
          '图谱联想：知识卡片双向链接的力导向图实时重算',
        ],
        codeHint: 'wails build # 打包 macOS / Windows 原生跨平台二进制',
      }
    case 'Go-Load':
      return {
        org: '迟暮个人实验仓库 (YN1753)',
        origin: '自制轻量级分布式压测探针。不依赖臃肿的 Java JMeter，用纯粹的 Go 协程池测出高并发下的服务吞吐量。',
        highlights: [
          '极低开销：单个工作节点以极低 CPU 占用维持数万并发连接',
          '实时直方图：P90 / P99 毫秒级延迟窗口滑动采样',
          '分布式发压：Master-Worker 模式协调多节点流量注入',
        ],
        codeHint: 'go run cmd/load/main.go -c 1000 -n 50000 http://target',
      }
    default:
      return {
        org: '迟暮个人实验仓库 (YN1753)',
        origin: '记录个人日常工程探索与基础设施工具链，坚持手打每个核心模块。',
        highlights: [
          '实战导向：以解决个人或实际协作中的真实痛点为出发点',
          '代码干净：注重并发安全性、内存开销与日志追踪',
        ],
        codeHint: 'git clone https://github.com/YN1753/' + title + '.git',
      }
  }
}
</script>

<template>
  <section id="projects" class="py-24 border-t border-[var(--border-color)] bg-[var(--bg-surface-subtle)]/50 relative transition-colors duration-400">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-left">
      
      <!-- 展台分镜标头 -->
      <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-16">
        <div>
          <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-[var(--bg-surface)] text-[var(--ink-secondary)] text-xs font-mono mb-3 border border-[var(--border-color)] shadow-2xs">
            <Layers class="w-3.5 h-3.5 text-[var(--accent-amber)]" />
            <span>THE CRAFT · 真实造物工坊档案</span>
          </div>
          <h2 class="text-3xl sm:text-4xl font-medium tracking-tight text-[var(--ink-primary)] font-serif-cinematic">
            工程造物与核心仓库
          </h2>
          <p class="text-sm text-[var(--ink-secondary)] mt-2 font-mono">
            来自 GitHub 真实公开仓库 · 点击任意项目查看「蓝图底片档案」
          </p>
        </div>

        <div class="flex items-center gap-3 text-xs font-mono text-[var(--ink-muted)]">
          <span class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-full bg-[var(--bg-surface)] border border-[var(--border-color)] text-[var(--accent-teal)]">
            <ShieldCheck class="w-3.5 h-3.5" />
            <span>SUSE-OAA 首位置顶</span>
          </span>
          <span class="px-2 py-1 bg-[var(--bg-surface)] rounded border border-[var(--border-color)]">
            共 {{ projects.length }} 件展品
          </span>
        </div>
      </div>

      <!-- 项目卡片矩阵 -->
      <div class="space-y-8">
        
        <!-- 重点置顶 #1：SUSE-OAA-BACKEND 宽银幕电影特刊展卡 -->
        <div
          v-if="projects.length > 0 && projects[0].title === 'SUSE-OAA-BACKEND'"
          class="film-card p-6 sm:p-10 relative overflow-hidden group border-2 border-[var(--accent-amber)]/20 hover:border-[var(--accent-amber)]/50 transition-all"
        >
          <!-- 胶卷打孔边缘装饰 -->
          <div class="flex items-center justify-between pb-6 border-b border-[var(--border-color)] mb-8">
            <div class="flex items-center gap-3">
              <span class="px-2.5 py-1 rounded-md bg-[var(--accent-amber)] text-white text-[11px] font-mono font-semibold">
                ACT I · 核心置顶 #1
              </span>
              <span class="text-xs font-mono text-[var(--ink-muted)]">
                suse-edu-cn / 四川轻化工大学开放原子开源协会
              </span>
            </div>
            <div class="text-[11px] font-mono text-[var(--accent-teal)] flex items-center gap-1.5">
              <span class="w-2 h-2 rounded-full bg-[var(--accent-teal)] animate-ping"></span>
              <span>生产环境运行中</span>
            </div>
          </div>

          <div class="grid grid-cols-1 lg:grid-cols-12 gap-8 items-center">
            <div class="lg:col-span-7 space-y-4">
              <h3 class="text-2xl sm:text-3xl font-serif-cinematic font-semibold text-[var(--ink-primary)]">
                {{ projects[0].title }}
              </h3>
              <p class="text-sm font-medium text-[var(--accent-amber)] font-mono">
                {{ projects[0].subtitle }}
              </p>
              <p class="text-sm text-[var(--ink-secondary)] leading-relaxed pt-1">
                {{ projects[0].description }}
              </p>

              <!-- 标签与技术栈 -->
              <div class="flex flex-wrap gap-2 pt-3">
                <span
                  v-for="tag in projects[0].tags.split(',')"
                  :key="tag"
                  class="px-2.5 py-1 rounded bg-[var(--bg-surface-subtle)] text-[11px] font-mono text-[var(--ink-secondary)] border border-[var(--border-color)]"
                >
                  {{ tag.trim() }}
                </span>
              </div>
            </div>

            <!-- 右侧操作与蓝图检视按钮 -->
            <div class="lg:col-span-5 flex flex-col justify-center items-start lg:items-end gap-3.5 bg-[var(--bg-surface-subtle)]/70 p-6 rounded-2xl border border-[var(--border-color)]">
              <div class="text-xs font-mono text-[var(--ink-muted)] text-left lg:text-right w-full">
                <div>仓库归属：四川轻化工大学开源协会</div>
                <div class="text-[var(--accent-amber)] mt-1 font-semibold">Campus OpenSource Backbone</div>
              </div>

              <div class="flex flex-wrap gap-3 w-full pt-2">
                <button
                  @click="openProjectModal(projects[0])"
                  class="flex-1 py-2.5 px-4 rounded-xl bg-[var(--ink-primary)] hover:opacity-90 text-[var(--bg-page)] text-xs font-mono font-medium transition-all flex items-center justify-center gap-2 cursor-pointer shadow-sm"
                >
                  <Eye class="w-3.5 h-3.5" />
                  <span>检视架构底片</span>
                </button>

                <a
                  :href="projects[0].github_url"
                  target="_blank"
                  rel="noopener noreferrer"
                  @click="audio.playShutter()"
                  class="py-2.5 px-4 rounded-xl bg-[var(--bg-surface)] hover:bg-[var(--bg-surface-subtle)] border border-[var(--border-color)] text-[var(--ink-primary)] text-xs font-mono font-medium transition-all flex items-center gap-1.5"
                >
                  <Github class="w-3.5 h-3.5" />
                  <span>GitHub</span>
                  <ArrowUpRight class="w-3 h-3 text-[var(--ink-muted)]" />
                </a>
              </div>
            </div>
          </div>
        </div>

        <!-- 其余真实项目卡片网格 (从 #2 ArchCanvas 开始) -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="(project, idx) in projects.slice(1)"
            :key="project.id"
            class="film-card p-6 flex flex-col justify-between group hover:border-[var(--accent-amber)]/40"
          >
            <div>
              <!-- 顶部编号与分镜序号 -->
              <div class="flex items-center justify-between pb-3 border-b border-[var(--border-color)] mb-4 text-[11px] font-mono text-[var(--ink-muted)]">
                <span class="text-[var(--ink-secondary)] font-semibold">
                  FRAME 0{{ idx + 2 }} / 0{{ projects.length }}
                </span>
                <span class="px-2 py-0.5 rounded bg-[var(--bg-surface-subtle)] border border-[var(--border-color)] text-[var(--ink-secondary)]">
                  {{ project.status }}
                </span>
              </div>

              <h4 class="text-xl font-serif-cinematic font-semibold text-[var(--ink-primary)] group-hover:text-[var(--accent-amber)] transition-colors">
                {{ project.title }}
              </h4>
              <p class="text-xs font-medium text-[var(--accent-amber)] font-mono mt-1">
                {{ project.subtitle }}
              </p>
              <p class="text-xs text-[var(--ink-secondary)] leading-relaxed mt-3 line-clamp-3">
                {{ project.description }}
              </p>
            </div>

            <div class="pt-6 mt-6 border-t border-[var(--border-color)]">
              <!-- 标签 -->
              <div class="flex flex-wrap gap-1.5 mb-4">
                <span
                  v-for="tag in project.tags.split(',').slice(0, 3)"
                  :key="tag"
                  class="px-2 py-0.5 rounded bg-[var(--bg-surface-subtle)] text-[10px] font-mono text-[var(--ink-muted)] border border-[var(--border-color)]"
                >
                  {{ tag.trim() }}
                </span>
              </div>

              <!-- 交互按钮 -->
              <div class="flex items-center justify-between pt-1">
                <button
                  @click="openProjectModal(project)"
                  class="text-xs font-mono text-[var(--ink-primary)] hover:text-[var(--accent-amber)] font-medium inline-flex items-center gap-1 cursor-pointer transition-colors"
                >
                  <Eye class="w-3.5 h-3.5" />
                  <span>蓝图底片</span>
                </button>

                <a
                  :href="project.github_url"
                  target="_blank"
                  rel="noopener noreferrer"
                  @click="audio.playShutter()"
                  class="p-2 rounded-lg bg-[var(--bg-surface-subtle)] hover:bg-[var(--ink-primary)] hover:text-white text-[var(--ink-secondary)] transition-all"
                  title="在 GitHub 查看"
                >
                  <ArrowUpRight class="w-3.5 h-3.5" />
                </a>
              </div>
            </div>

          </div>
        </div>

      </div>

    </div>

    <!-- 交互式蓝图底片抽屉 / Dossier Modal -->
    <div
      v-if="activeModalProject"
      class="fixed inset-0 z-50 flex items-center justify-center p-4 sm:p-6 bg-black/60 backdrop-blur-md animate-fade-in"
      @click.self="closeModal"
    >
      <div class="relative w-full max-w-2xl bg-[var(--bg-surface)] border border-[var(--border-color)] rounded-2xl shadow-2xl p-6 sm:p-8 overflow-hidden text-left">
        
        <!-- 蓝图微网格纹理背景 -->
        <div class="absolute inset-0 blueprint-grid opacity-30 pointer-events-none"></div>

        <!-- 顶部关闭与标头 -->
        <div class="relative flex items-center justify-between pb-4 border-b border-[var(--border-color)] mb-6">
          <div class="flex items-center gap-2 text-xs font-mono text-[var(--ink-muted)]">
            <Terminal class="w-4 h-4 text-[var(--accent-amber)]" />
            <span>ENGINEERING DOSSIER · {{ activeModalProject.title }}</span>
          </div>
          <button
            @click="closeModal"
            class="p-1.5 rounded-lg text-[var(--ink-muted)] hover:text-[var(--ink-primary)] hover:bg-[var(--bg-surface-subtle)] transition-colors"
          >
            <X class="w-5 h-5" />
          </button>
        </div>

        <div class="relative space-y-5">
          <div>
            <h3 class="text-2xl font-serif-cinematic font-semibold text-[var(--ink-primary)]">
              {{ activeModalProject.title }}
            </h3>
            <p class="text-xs font-mono text-[var(--accent-amber)] mt-1">
              {{ getProjectDossier(activeModalProject.title).org }}
            </p>
          </div>

          <!-- 初衷与背景 -->
          <div class="p-4 rounded-xl bg-[var(--bg-surface-subtle)] border border-[var(--border-color)]">
            <h5 class="text-xs font-mono font-semibold text-[var(--ink-primary)] uppercase tracking-wider mb-1.5">
              Why I Built This · 造物初衷
            </h5>
            <p class="text-xs text-[var(--ink-secondary)] leading-relaxed">
              {{ getProjectDossier(activeModalProject.title).origin }}
            </p>
          </div>

          <!-- 核心架构亮点 -->
          <div>
            <h5 class="text-xs font-mono font-semibold text-[var(--ink-primary)] uppercase tracking-wider mb-2">
              Architectural Slices · 架构切片
            </h5>
            <ul class="space-y-2 text-xs text-[var(--ink-secondary)] font-mono">
              <li
                v-for="(item, i) in getProjectDossier(activeModalProject.title).highlights"
                :key="i"
                class="flex items-start gap-2 bg-[var(--bg-surface)] p-2 rounded-lg border border-[var(--border-color)]"
              >
                <span class="text-[var(--accent-teal)] font-bold">#0{{ i + 1 }}</span>
                <span>{{ item }}</span>
              </li>
            </ul>
          </div>

          <!-- 快速运行命令与克隆 -->
          <div class="p-3 rounded-xl bg-[#14151a] text-slate-200 font-mono text-xs flex items-center justify-between">
            <code class="text-teal-300 truncate pr-2">
              git clone {{ activeModalProject.github_url }}.git
            </code>
            <button
              @click="copyCloneCommand(activeModalProject.github_url)"
              class="px-2.5 py-1 rounded bg-white/10 hover:bg-white/20 text-white text-[11px] flex items-center gap-1 transition-all cursor-pointer shrink-0"
            >
              <Check v-if="copied" class="w-3.5 h-3.5 text-teal-400" />
              <Copy v-else class="w-3.5 h-3.5" />
              <span>{{ copied ? '已复制' : '复制命令' }}</span>
            </button>
          </div>

          <!-- 底部直达 GitHub -->
          <div class="pt-4 border-t border-[var(--border-color)] flex items-center justify-between">
            <span class="text-xs font-mono text-[var(--ink-muted)]">
              状态：{{ activeModalProject.status }}
            </span>
            <a
              :href="activeModalProject.github_url"
              target="_blank"
              rel="noopener noreferrer"
              @click="audio.playShutter()"
              class="px-4 py-2 rounded-xl bg-[var(--ink-primary)] hover:opacity-90 text-[var(--bg-page)] text-xs font-mono font-medium transition-all flex items-center gap-1.5"
            >
              <span>直达官方仓库</span>
              <ExternalLink class="w-3.5 h-3.5" />
            </a>
          </div>

        </div>

      </div>
    </div>

  </section>
</template>
