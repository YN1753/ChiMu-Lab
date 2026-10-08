<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Navbar from './components/Navbar.vue'
import Hero from './components/Hero.vue'
import ProjectShowcase from './components/ProjectShowcase.vue'
import ActivityHub from './components/ActivityHub.vue'
import ComplianceFooter from './components/ComplianceFooter.vue'
import type { Profile, Project, Activity, Stats, SiteConfig } from './types'

const profile = ref<Profile | null>({
  id: 1,
  name: '迟暮',
  title: 'Gopher / 全栈开发者',
  bio: '迟暮的个人数字工坊。不迎合外界，只记录自己造过的轮子、踩过的坑与真实的工程探索。',
  avatar: 'https://avatars.githubusercontent.com/u/108920146?v=4',
  github: 'https://github.com/YN1753',
  email: 'chimu@codeactivityhub.top',
  location: '中国 · 杭州',
  skills: 'Go,Wails,Vue 3,TypeScript,Docker,Linux,Swift,SQLite,Gin,Tailwind CSS',
})

const stats = ref<Stats | null>({
  total_projects: 7,
  total_activities: 24,
  active_days: 218,
  uptime_hours: 48,
  uptime_seconds: 172800,
  last_updated: '2026-10-08 23:00:00',
  go_version: 'go1.27',
  goroutines: 4,
  memory_alloc_mb: 1.76,
  database_type: 'Pure-Go SQLite (CGO-Free)',
  query_latency_ms: 0.24,
})

const projects = ref<Project[]>([
  {
    id: 1,
    title: 'SUSE-OAA-BACKEND',
    subtitle: '四川轻化工大学开放原子开源协会 · 核心业务后端',
    description: '为四川轻化工大学开源协会研发的 Go 后端服务体系，支撑协会事务协作、组织管理、招新流程与校园数据服务接口。',
    category: 'core',
    tags: 'Go,Gin,suse-edu-cn,Campus OpenSource',
    demo_url: '',
    github_url: 'https://github.com/suse-edu-cn/SUSE-OAA-BACKEND',
    status: 'Active',
    featured: true,
    order: 1,
  },
  {
    id: 2,
    title: 'ArchCanvas',
    subtitle: 'AI 辅助 Go 架构设计画布',
    description: '让 AI 与开发者一起，从需求语义理解、ER 实体关系图到系统架构设计，快速构建与生成可运行的 Go 工程骨架。探索 AI 与现代架构设计工具的深度融合。',
    category: 'core',
    tags: 'TypeScript,Go,Architecture,AI Canvas',
    demo_url: '',
    github_url: 'https://github.com/YN1753/ArchCanvas',
    status: 'Active',
    featured: true,
    order: 2,
  },
  {
    id: 3,
    title: 'GoLens',
    subtitle: '基于交互式状态机的 Go 底层机制透视镜',
    description: '让 GMP 协程调度、三色标记 GC 屏障与 Channel 阻塞状态流转清晰可见的高交互度运行时可视化工具。',
    category: 'tool',
    tags: 'JavaScript,Go Internals,GMP,Visualization',
    demo_url: '',
    github_url: 'https://github.com/YN1753/GoLens',
    status: 'Stable',
    featured: true,
    order: 3,
  },
  {
    id: 4,
    title: 'AstraLink-Desktop',
    subtitle: '基于 Wails 的图笔记桌面端应用',
    description: '星链 2.0。探索 Go + 前端混合开发（Wails 架构），支持双向链接、知识图谱可视化与本地隐私优先的个人知识库体系。',
    category: 'tool',
    tags: 'Go,Wails,Vue,Desktop,Graph',
    demo_url: '',
    github_url: 'https://github.com/YN1753/AstraLink-Desktop',
    status: 'Active',
    featured: false,
    order: 4,
  },
  {
    id: 5,
    title: 'Go-Load',
    subtitle: '轻量级高并发 HTTP 压测工具',
    description: '基于 Go 语言原生并发模型编写的高性能压测工具，轻巧无依赖，具备极低资源开销与精确的时延吞吐量统计。',
    category: 'tool',
    tags: 'Go,Benchmark,High Concurrency,CLI',
    demo_url: '',
    github_url: 'https://github.com/YN1753/Go-Load',
    status: 'Stable',
    featured: false,
    order: 5,
  },
  {
    id: 6,
    title: 'nexus',
    subtitle: '面向开发者的现代化 Linux 服务器管理平台',
    description: '面向开发者的现代化 Linux 服务器管理控制台，用于便捷监控云服务器硬件、容器与核心守护进程。',
    category: 'tool',
    tags: 'Vue,Linux,DevOps,System',
    demo_url: '',
    github_url: 'https://github.com/YN1753/nexus',
    status: 'WIP',
    featured: false,
    order: 6,
  },
  {
    id: 7,
    title: 'DeviceDaily',
    subtitle: '支持 Mac 原生小组件的设备成本统计 App',
    description: '基于 Swift 原生开发，支持 macOS 桌面原生 Widget 小组件，计算并追踪电子设备的日均使用成本。',
    category: 'tool',
    tags: 'Swift,macOS,WidgetKit,Utility',
    demo_url: '',
    github_url: 'https://github.com/YN1753/DeviceDaily',
    status: 'Active',
    featured: false,
    order: 7,
  },
])

const activities = ref<Activity[]>([
  {
    id: 1,
    date: '2026-10-08',
    type: 'commit',
    title: 'SUSE-OAA-BACKEND 核心业务接口与鉴权中间件重构',
    description: '优化校园协会事务协作微服务接口，重构逻辑层响应结构与错误捕获。',
    repo_name: 'suse-edu-cn/SUSE-OAA-BACKEND',
    count: 8,
    link: 'https://github.com/suse-edu-cn/SUSE-OAA-BACKEND',
  },
  {
    id: 2,
    date: '2026-10-08',
    type: 'commit',
    title: 'ArchCanvas AI 需求解析与 Go 实体代码生成打通',
    description: '实现从 ER 图实体到 Go struct 与 GORM 模型代码的自动化生成管线。',
    repo_name: 'YN1753/ArchCanvas',
    count: 12,
    link: 'https://github.com/YN1753/ArchCanvas',
  },
  {
    id: 3,
    date: '2026-09-20',
    type: 'release',
    title: 'GoLens 状态机透视镜：GMP 协程调度流转可视化',
    description: '完成调度器状态转移交互逻辑，使 Goroutine 在 P/M 上的流转过程直观呈现。',
    repo_name: 'YN1753/GoLens',
    count: 5,
    link: 'https://github.com/YN1753/GoLens',
  },
  {
    id: 4,
    date: '2026-05-18',
    type: 'milestone',
    title: 'AstraLink-Desktop Wails 桌面端跨平台打包验证',
    description: '完成星链 2.0 本地知识库双链图谱渲染与桌面端原生打包。',
    repo_name: 'YN1753/AstraLink-Desktop',
    count: 6,
    link: 'https://github.com/YN1753/AstraLink-Desktop',
  },
])

const config = ref<SiteConfig | null>({
  id: 1,
  site_name: '迟暮实验室 · ChiMu-Lab',
  site_desc: '迟暮的个人数字工坊与工程工作台',
  domain: 'codeactivityhub.top',
  icp_number: '浙ICP备2026081664号',
  icp_link: 'https://beian.miit.gov.cn',
  police_number: '全国公安联网备案审核中',
  police_code: '',
})

onMounted(async () => {
  try {
    const [profRes, projRes, actRes, statRes, cfgRes] = await Promise.allSettled([
      fetch('/api/profile').then(r => r.ok ? r.json() : null),
      fetch('/api/projects').then(r => r.ok ? r.json() : null),
      fetch('/api/activities').then(r => r.ok ? r.json() : null),
      fetch('/api/stats').then(r => r.ok ? r.json() : null),
      fetch('/api/config').then(r => r.ok ? r.json() : null),
    ])

    if (profRes.status === 'fulfilled' && profRes.value) profile.value = profRes.value
    if (projRes.status === 'fulfilled' && projRes.value) projects.value = projRes.value
    if (actRes.status === 'fulfilled' && actRes.value) activities.value = actRes.value
    if (statRes.status === 'fulfilled' && statRes.value) stats.value = statRes.value
    if (cfgRes.status === 'fulfilled' && cfgRes.value) config.value = cfgRes.value
  } catch (err) {
    console.log('Using default local seed data', err)
  }
})
</script>

<template>
  <div class="min-h-screen bg-[#f7f6f2] text-[#14151a] flex flex-col justify-between selection:bg-[#d97706]/15 selection:text-[#92400e]">
    <Navbar />
    <main class="flex-grow">
      <Hero :profile="profile" :stats="stats" />
      <ProjectShowcase :projects="projects" />
      <ActivityHub :activities="activities" />
    </main>
    <ComplianceFooter :config="config" />
  </div>
</template>
