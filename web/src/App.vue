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
  title: '全栈研发工程师 / Gopher / 独立开发者',
  bio: '立足工程美学与极客实践。热衷于 Go 高性能后端、现代前端与自动化技术探索，在数字世界构建优雅而坚固的实验室工具。',
  avatar: 'https://avatars.githubusercontent.com/u/108920146?v=4',
  github: 'https://github.com/YN1753',
  email: 'chimu@codeactivityhub.top',
  location: '中国 · 杭州',
  skills: 'Go,Vue 3,TypeScript,Docker,Linux,SQLite,MySQL,Gin,Tailwind CSS,Git,Redis',
})

const stats = ref<Stats | null>({
  total_projects: 4,
  total_activities: 28,
  active_days: 180,
  uptime_hours: 24,
  last_updated: '2026-10-08 21:00',
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
    title: 'GopherSpace 分布式存储探测工具',
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
    title: 'Hermes 自动化工作流引擎',
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
    title: '迟暮实验室 ChiMu-Lab 1.0 正式上线',
    description: '完成全栈架构搭建，基于 Go + Vue 3 + SQLite 实现单二进制极简部署与现代化响应式 UI。',
    repo_name: 'YN1753/ChiMu-Lab',
    count: 12,
    link: 'https://github.com/YN1753/ChiMu-Lab',
  },
  {
    id: 2,
    date: '2026-10-06',
    type: 'commit',
    title: '重构面试每日题静态排版与移动端自适应',
    description: '优化 ACM 模式题目展示效果，增加平滑滚动与语法高亮支持。',
    repo_name: 'YN1753/interview-hub',
    count: 6,
    link: 'https://github.com/YN1753',
  },
  {
    id: 3,
    date: '2026-10-03',
    type: 'study',
    title: '贪心与动态规划状态转移专题复习',
    description: '深度演练相邻约束类贪心原型题目，输出白板题解与复杂度推导笔记。',
    repo_name: 'YN1753/algo-notes',
    count: 4,
    link: 'https://github.com/YN1753',
  },
  {
    id: 4,
    date: '2026-09-29',
    type: 'milestone',
    title: '启用独立主域名 codeactivityhub.top',
    description: '完成云服务器 DNS 解析、全站 HTTPS/TLS 证书签发与安全加固配置。',
    repo_name: 'YN1753/infra',
    count: 8,
    link: 'https://codeactivityhub.top',
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

// 从后端接口动态拉取数据
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
  <div class="min-h-screen bg-[#0b0f19] text-slate-100 flex flex-col justify-between selection:bg-cyan-500/25 selection:text-cyan-300">
    <Navbar />
    <main class="flex-grow">
      <Hero :profile="profile" :stats="stats" />
      <ProjectShowcase :projects="projects" />
      <ActivityHub :activities="activities" />
    </main>
    <ComplianceFooter :config="config" />
  </div>
</template>
