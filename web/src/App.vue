<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Navbar from './components/Navbar.vue'
import Hero from './components/Hero.vue'
import ProjectShowcase from './components/ProjectShowcase.vue'
import EngineeringNotes from './components/EngineeringNotes.vue'
import ActivityHub from './components/ActivityHub.vue'
import ComplianceFooter from './components/ComplianceFooter.vue'
import type { Profile, Project, Activity, Stats, SiteConfig } from './types'

const profile = ref<Profile | null>({
  id: 1,
  name: '迟暮',
  title: '全栈研发工程师 / Gopher / 独立开发者',
  bio: '立足工程美学与极客实践。热衷于 Go 高性能后端、现代前端架构与自动化技术探索，在数字世界构建优雅而坚固的实验室工具。',
  avatar: 'https://avatars.githubusercontent.com/u/108920146?v=4',
  github: 'https://github.com/YN1753',
  email: 'chimu@codeactivityhub.top',
  location: '中国 · 杭州',
  skills: 'Go (Core),Goroutine,Vue 3,TypeScript,Docker,Linux,SQLite,Gin,Tailwind CSS,Git',
})

const stats = ref<Stats | null>({
  total_projects: 4,
  total_activities: 32,
  active_days: 218,
  uptime_hours: 48,
  uptime_seconds: 172800,
  last_updated: '2026-10-08 22:00:00',
  go_version: 'go1.27',
  goroutines: 6,
  memory_alloc_mb: 3.8,
  database_type: 'Pure-Go SQLite (CGO-Free)',
  query_latency_ms: 0.22,
})

const projects = ref<Project[]>([
  {
    id: 1,
    title: '今日 · 面试练习站 (Daily Practice Hub)',
    subtitle: '算法每日打卡与高频面试考点演练平台',
    description: '围绕 ACM 模式、数组贪心算法、高频场景题构建的沉浸式每日刷题平台。支持闭卷练习、代码思路验证与历史归档回溯。',
    category: 'core',
    tags: 'Go,Algorithms,ACM,Vue,Tailwind',
    demo_url: '/interview',
    github_url: 'https://github.com/YN1753',
    status: 'Active',
    featured: true,
    order: 1,
  },
  {
    id: 2,
    title: 'ChiMu-Lab 迟暮实验室',
    subtitle: '极简主义全栈工坊与代码动态聚合站',
    description: '本站核心工程。基于 Go + Vue 3 + SQLite 构建的轻量级开发实验室，整合项目展厅、工程动态追踪与合规备案中心。',
    category: 'lab',
    tags: 'Go,Gin,Vue3,TypeScript,SQLite,Tailwind',
    demo_url: 'https://codeactivityhub.top',
    github_url: 'https://github.com/YN1753/ChiMu-Lab',
    status: 'Active',
    featured: true,
    order: 2,
  },
  {
    id: 3,
    title: 'GopherSpace 分布式网络探针',
    subtitle: '高性能并发网络扫描与节点状态同步守护进程',
    description: '轻量级网络探测器，专为探测节点连通性、实时延迟与健康状态设计，具备极低资源占用与高吞吐并发能力。',
    category: 'tool',
    tags: 'Go,Goroutine,Network,Linux,Syscall',
    demo_url: '',
    github_url: 'https://github.com/YN1753',
    status: 'Stable',
    featured: true,
    order: 3,
  },
  {
    id: 4,
    title: 'Hermes 自动化工作流管线',
    subtitle: '每日任务调度、内容聚合与静态站点自动化发布管线',
    description: '定时拉取技术动态与刷题计划，自动化编译生成静态练习页并无缝推送到生产 Web 服务。',
    category: 'lab',
    tags: 'Automation,Python,Shell,CI/CD,Linux',
    demo_url: '',
    github_url: 'https://github.com/YN1753',
    status: 'Active',
    featured: false,
    order: 4,
  },
])

const activities = ref<Activity[]>([
  {
    id: 1,
    date: '2026-10-08',
    type: 'release',
    title: 'ChiMu-Lab 2.0 视觉重构：引入 Bento Grid 布局与 System HUD',
    description: '全新升级极客科技美学，引入运行时性能监控面板、算法推导笔记与 GitHub 年度活跃矩阵。',
    repo_name: 'YN1753/ChiMu-Lab',
    count: 15,
    link: 'https://github.com/YN1753/ChiMu-Lab',
  },
  {
    id: 2,
    date: '2026-10-08',
    type: 'milestone',
    title: '完成工信部备案号注入与 TLS 1.3 证书握手测试',
    description: '绑定浙ICP备2026081664号，打通 443 端口安全组与 Nginx 反向代理链路。',
    repo_name: 'YN1753/infra',
    count: 8,
    link: 'https://codeactivityhub.top',
  },
  {
    id: 3,
    date: '2026-10-06',
    type: 'commit',
    title: '重构面试每日题移动端响应式布局与 ACM 样例排版',
    description: '优化 ACM 模式题目展示效果，增加平滑滚动与语法高亮支持。',
    repo_name: 'YN1753/interview-hub',
    count: 6,
    link: 'https://github.com/YN1753',
  },
  {
    id: 4,
    date: '2026-10-03',
    type: 'study',
    title: 'P3817 贪心相邻约束题解证明与白板思路复盘',
    description: '深度演练相邻约束类贪心原型题目，输出白板题解与复杂度推导笔记。',
    repo_name: 'YN1753/algo-notes',
    count: 4,
    link: 'https://github.com/YN1753',
  },
])

const config = ref<SiteConfig | null>({
  id: 1,
  site_name: '迟暮实验室 · ChiMu-Lab',
  site_desc: 'Code Activity Hub · 迟暮的个人极客工坊与代码动态中心',
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
      <EngineeringNotes />
      <ActivityHub :activities="activities" />
    </main>
    <ComplianceFooter :config="config" />
  </div>
</template>
