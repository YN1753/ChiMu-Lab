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
    time: '21:30',
    location: '杭州 · 西湖杨公堤',
    weather: '20°C · 暮秋凉风',
    mood: '放空',
    category: 'cycling',
    title: '杨公堤夜骑：风从荷叶上吹过来的温度',
    content: '下班后没有直接回书房，换了身衣服去西湖夜骑。从曲院风荷出发，沿着杨公堤一直骑到南山路。路两侧是高大的水杉，秋天的晚风穿过荷叶扑在脸上，带着微凉的水汽。把城市的喇叭声和屏幕的蓝光全抛在身后，出了一身薄汗，整个人好像被这阵风重新洗刷了一遍。',
    quote: '“骑车的时候，世界退居到两侧，心跳是唯一的节拍器。”',
    note: '路线：曲院风荷 -> 杨公堤 -> 虎跑路 -> 钱塘江绿道，总里程 21.4 km。速度 22km/h。',
    image_url: 'https://images.unsplash.com/photo-1502680390469-be75c86b636f?auto=format&fit=crop&w=800&q=80',
    meta_info: '公路车巡航 · 21.4km · 爬升 85m',
    tags: '夜骑,西湖,生活行迹,松弛',
    likes: 34,
  },
  {
    id: 2,
    date: '2026.10.05',
    time: '16:40',
    location: '杭州 · 满觉陇青石板路',
    weather: '23°C · 金桂初放',
    mood: '拾光',
    category: 'photo',
    title: '满觉陇的桂花蒸与旁轴底片',
    content: '赶在假期的尾巴，带了老胶片机去满觉陇。沿山的村落都在做糖桂花，青石板路上落了一层薄薄的金黄，空气甜得发稠。下午四点半的夕阳斜斜穿过樟树叶，光斑在墙面上摇曳。按下快门的那一瞬，时间好像被装进了小小的暗盒里。',
    quote: '“光线是时间的影印件，而胶卷留下了温度。”',
    note: 'Contax T2 · 38mm f/2.8 Carl Zeiss，Kodak Portra 400，光圈 f/4，曝光补偿 +0.3EV。',
    image_url: 'https://images.unsplash.com/photo-1507525428034-b723cf961d3e?auto=format&fit=crop&w=800&q=80',
    meta_info: 'Contax T2 · Kodak Portra 400 · 38mm',
    tags: '胶片,满觉陇,桂花,摄影',
    likes: 42,
  },
  {
    id: 3,
    date: '2026.10.03',
    time: '10:15',
    location: '书房窗前木桌',
    weather: '22°C · 晨光微熹',
    mood: '惬意',
    category: 'coffee',
    title: '浅烘埃塞古吉与 92°C 细水慢萃',
    content: '清晨的书房很安静。磨了 15g 埃塞俄比亚古吉产区的花魁日晒豆，磨齿带出的干香满是草莓和茉莉花气息。92°C 纯净水，三段式注水萃取。第一段闷蒸 35 秒看咖啡粉像面包一样膨胀，看着水柱在滤杯里画同心圆，是每天最让人沉静下来的仪式。',
    quote: '“把水注入咖啡粉的过程，就像把耐心倾注给生活。”',
    note: '水粉比 1:15，粉量 15g，注水总量 225g，总萃取耗时 2分18秒。白桃与柑橘酸感明亮。',
    image_url: 'https://images.unsplash.com/photo-1514432324607-a09d9b4aefdd?auto=format&fit=crop&w=800&q=80',
    meta_info: '埃塞俄比亚 古吉花魁 · 92°C · V60 滤杯',
    tags: '手冲,咖啡,日常仪式,慢生活',
    likes: 29,
  },
  {
    id: 4,
    date: '2026.09.30',
    time: '14:20',
    location: '西湖边 · 树荫下长椅',
    weather: '24°C · 微风拂面',
    mood: '沉思',
    category: 'reading',
    title: '重读《禅与摩托车维修艺术》：手艺与良质',
    content: '坐在长椅上吹着湖风读波西格。书中写：‘佛陀或耶稣坐在排气管边，就跟坐在莲花座上一样正常。’ 当你带着真正的良质（Quality）去面对一件具体的事物——无论是调校一辆自行车的刹车皮、手冲一杯咖啡，抑或是雕琢一行 Go 代码——工具和人就合二为一了。少一点向外张望的功利，多一点对手艺本身的敬畏。',
    quote: '“如果你对事情感到厌倦，说明你已经失去了与它的活生生的联结。”',
    note: '罗伯特·M·波西格 著，重庆出版社。随手在 P.168 折了角，记下了关于专注的感悟。',
    image_url: 'https://images.unsplash.com/photo-1497633762265-9d179a990aa6?auto=format&fit=crop&w=800&q=80',
    meta_info: '《禅与摩托车维修艺术》· 案头书摘',
    tags: '读书,手艺,良质,哲学',
    likes: 36,
  },
  {
    id: 5,
    date: '2026.09.25',
    time: '23:00',
    location: '深夜工位 · 暖光台灯下',
    weather: '19°C · 秋夜微凉',
    mood: '专注',
    category: 'music',
    title: '黑胶唱片与坂本龙一的音符',
    content: '深夜十一点，把房间大灯关掉，只留一盏暖黄的台灯。戴上耳机放坂本龙一的《async》。空旷而平静的钢琴音、风吹过树梢的采样，整个世界好像都睡着了。在这样的底噪里写两行字，看看书，心里非常踏实。音乐是精神的庇护所。',
    quote: '“生命是脆弱的，但音乐能把那一瞬的真实凝固成永恒。”',
    note: '推荐循环曲目：《andata》与《solari》。静谧、克制、富有呼吸感。',
    image_url: 'https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?auto=format&fit=crop&w=800&q=80',
    meta_info: '坂本龙一 · 《async》· 24bit/96kHz 无损',
    tags: '音乐,坂本龙一,深夜,治愈',
    likes: 45,
  },
  {
    id: 6,
    date: '2026.09.18',
    time: '22:30',
    location: '杭州 · 窗台前',
    weather: '21°C · 夜雨敲窗',
    mood: '自洽',
    category: 'thought',
    title: '写给自己：在喧嚣时代认认真真生活',
    content: '窗外下着绵绵的秋雨。很多人把生活过成了一场给别人看的展览，急着证明自己掌握了什么技术、走到了什么高度。可日子终究是自己过的：饭要一口一口吃，觉要踏踏实实睡。代码也好，爱好也罢，都只是人生长河里的浪花。认认真真生活，诚恳对待每一顿饭、每一趟骑行、每一个念头，足矣。',
    quote: '“不向外界讨要意义，生活的意义就在生活的每一个具体细节里。”',
    note: '深夜随笔。泡了一杯温热的陈皮老白茶，听雨声入睡。',
    image_url: 'https://images.unsplash.com/photo-1517694712202-14dd9538aa97?auto=format&fit=crop&w=800&q=80',
    meta_info: '雨夜随想 · 杭州生活手记',
    tags: '随想,自洽,内心平宁,夜雨',
    likes: 58,
  },
  {
    id: 7,
    date: '2026.09.10',
    time: '15:00',
    location: '书房桌面',
    weather: '26°C · 晴朗',
    mood: '爱物',
    category: 'gear',
    title: '陪伴三年的旧物：机械键盘、钢笔与手账',
    content: '清理桌面时擦拭这把用了三年的无刻机械键盘。键帽表面已经泛出了温润的光泽，青轴的手感依旧干脆。旁边是一支用了很久的百乐钢笔和一个牛皮纸手账。这些天天陪伴我的物件，沉默却忠诚。人与器物之间的相处，时间久了也会生出情谊。',
    quote: '“日用即道。善待陪伴你的每一件工具。”',
    note: '桌面好物：定制无刻键盘、百乐 78G 钢笔、Midori MD 方格本。简约耐看。',
    image_url: 'https://images.unsplash.com/photo-1587829741301-dc798b83add3?auto=format&fit=crop&w=800&q=80',
    meta_info: 'EDC 好物 · 桌面日常 · 物与心',
    tags: '好物,桌面,文具,陪伴',
    likes: 27,
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
      <LivingFrames :moments="moments" />
      <ProjectShowcase :projects="projects" />
      <ActivityHub :activities="activities" />
    </main>
    <ComplianceFooter :config="config" />
  </div>
</template>
