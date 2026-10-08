<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Navbar from './components/Navbar.vue'
import Hero from './components/Hero.vue'
import ProjectShowcase from './components/ProjectShowcase.vue'
import LivingFrames from './components/LivingFrames.vue'
import ActivityHub from './components/ActivityHub.vue'
import ComplianceFooter from './components/ComplianceFooter.vue'
import type { Profile, Project, Activity, Stats, SiteConfig, LifeMoment } from './types'

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
  last_updated: '2026-10-08 23:45:00',
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

const moments = ref<LifeMoment[]>([
  {
    id: 1,
    date: '2026.10.08',
    time: '23:40',
    location: '杭州 · 余杭工位',
    weather: '19°C · 秋夜微雨',
    mood: '专注',
    category: 'thought',
    title: '雨夜、通道阻塞与造轮子的意义',
    content: '窗外下着绵密的秋雨。有人曾问我，为什么开源社区那么多现成的调度器和压测工具，还要花时间一行行写 GoLens 和 Go-Load？亲手排查过一次 GMP 的偷取队列、体会过 Channel 阻塞唤醒时的原子语义，那种心里的笃定是调现成库给不了的。代码不是虚张声势的展品，是一块块自己铺上去的青石板。',
    note: '手冲了一杯耶加雪菲。雨水敲打百叶窗的声音很适合写代码。',
    image_url: 'https://images.unsplash.com/photo-1517694712202-14dd9538aa97?auto=format&fit=crop&w=800&q=80',
    camera: 'Leica Q2 · 28mm f/1.7 · ISO 400',
    tags: 'Go,GMP,思考,夜雨',
    likes: 14,
  },
  {
    id: 2,
    date: '2026.10.05',
    time: '17:20',
    location: '杭州 · 满觉陇',
    weather: '23°C · 晴朗薄暮',
    mood: '拾光',
    category: 'film',
    title: '桂花蒸与胶卷里的秋天',
    content: '趁着假期最后两天，骑车去了一趟满觉陇。满山的金桂已经落了一地，空气里全是甜香。随身带了旁轴机拍完了一卷 Kodak 400。阳光透过香樟树叶洒在青石板上，突然觉得不管是写系统架构还是过生活，留白和呼吸感才是最珍贵的底色。',
    note: 'Kodak Portra 400 曝光补偿 +0.3EV，光斑很柔和。',
    image_url: 'https://images.unsplash.com/photo-1507525428034-b723cf961d3e?auto=format&fit=crop&w=800&q=80',
    camera: 'Contax T2 · 38mm f/2.8 · Kodak 400',
    tags: '胶片,满觉陇,秋日,骑行',
    likes: 28,
  },
  {
    id: 3,
    date: '2026.10.01',
    time: '15:00',
    location: '西湖边 · 树荫下长椅',
    weather: '25°C · 微风',
    mood: '松弛',
    category: 'reading',
    title: '《禅与摩托车维修艺术》与工程手艺',
    content: '重读波西格的《禅与摩托车维修艺术》。书中写道：‘佛陀或耶稣坐在排气管边，就跟坐在莲花座上一样正常。’ 当你带着真正的良质（Quality）去调试一段并发竞态，或者去调配一杯手冲水粉比时，工具和人就已经融为一体了。少一点功利的目标，多一点对手艺的敬畏。',
    note: 'P.142 标注：‘如果你对事情感到厌倦，说明你已经失去了与它的活生生的联结。’',
    image_url: 'https://images.unsplash.com/photo-1497633762265-9d179a990aa6?auto=format&fit=crop&w=800&q=80',
    camera: 'Ricoh GR IIIx · 40mm f/2.8',
    tags: '阅读,手艺,良质,哲学',
    likes: 19,
  },
  {
    id: 4,
    date: '2026.09.28',
    time: '22:15',
    location: '杭州 · 迟暮工坊',
    weather: '21°C · 清凉夜风',
    mood: '沉思',
    category: 'coffee',
    title: '曼特宁手冲笔记与 AST 遍历灵感',
    content: '深烘苏门答腊曼特宁，研磨度偏粗，水温 90°C。前段坚果与黑巧的风味很厚重。在等第二段注水浸润的 40 秒里，突然想通了 ArchCanvas 解析 Go struct tag 时的嵌套循环问题：直接在语义树上做局部剪枝，比正则匹配合适得多。生活中的沉淀，往往会在不经意间反哺工程。',
    note: '咖啡粉 16g，水粉比 1:14，两段注水，总萃取用时 2分15秒。',
    image_url: 'https://images.unsplash.com/photo-1501339847302-ac426a4a7cbb?auto=format&fit=crop&w=800&q=80',
    camera: 'Fujifilm X100V · Classic Chrome',
    tags: '手冲,咖啡,ArchCanvas,思考',
    likes: 22,
  },
  {
    id: 5,
    date: '2026.09.20',
    time: '01:30',
    location: '深夜桌面',
    weather: '18°C · 静谧夜色',
    mood: '沉浸',
    category: 'reading',
    title: '坂本龙一的音符与静默的 Goroutine',
    content: '戴上耳机放坂本龙一的《async》。整座城市好像都睡熟了，只有屏幕上微微发光的代码。把 suseoaa 的一个死锁隐患顺藤摸瓜找了出来——是一个 defer unlock 在条件分支里意外漏掉。修掉的那一刻，世界好像都变宽阔了。这种深夜的纯粹，千金不换。',
    note: '推荐曲目：《solari》与《andata》。空旷而平静。',
    image_url: 'https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?auto=format&fit=crop&w=800&q=80',
    camera: 'Sony A7C · 35mm f/1.4 GM',
    tags: '音乐,坂本龙一,深夜编码,suseoaa',
    likes: 31,
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
  site_desc: '迟暮的个人数字工坊与生活档案',
  domain: 'codeactivityhub.top',
  icp_number: '浙ICP备2026081664号',
  icp_link: 'https://beian.miit.gov.cn',
  police_number: '全国公安联网备案审核中',
  police_code: '',
})

onMounted(async () => {
  try {
    const [profRes, projRes, actRes, statRes, cfgRes, momRes] = await Promise.allSettled([
      fetch('/api/profile').then(r => r.ok ? r.json() : null),
      fetch('/api/projects').then(r => r.ok ? r.json() : null),
      fetch('/api/activities').then(r => r.ok ? r.json() : null),
      fetch('/api/stats').then(r => r.ok ? r.json() : null),
      fetch('/api/config').then(r => r.ok ? r.json() : null),
      fetch('/api/moments').then(r => r.ok ? r.json() : null),
    ])

    if (profRes.status === 'fulfilled' && profRes.value) profile.value = profRes.value
    if (projRes.status === 'fulfilled' && projRes.value) projects.value = projRes.value
    if (actRes.status === 'fulfilled' && actRes.value) activities.value = actRes.value
    if (statRes.status === 'fulfilled' && statRes.value) stats.value = statRes.value
    if (cfgRes.status === 'fulfilled' && cfgRes.value) config.value = cfgRes.value
    if (momRes.status === 'fulfilled' && momRes.value) moments.value = momRes.value
  } catch (err) {
    console.log('Using default local seed data', err)
  }
})
</script>

<template>
  <div class="min-h-screen bg-[var(--bg-page)] text-[var(--ink-primary)] flex flex-col justify-between selection:bg-[var(--accent-amber)]/20 selection:text-[var(--accent-amber)] transition-colors duration-400">
    <Navbar />
    <main class="flex-grow">
      <Hero :profile="profile" :stats="stats" />
      <ProjectShowcase :projects="projects" />
      <LivingFrames :moments="moments" />
      <ActivityHub :activities="activities" />
    </main>
    <ComplianceFooter :config="config" />
  </div>
</template>
