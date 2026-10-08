<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import HeaderNav from './components/HeaderNav.vue'
import HomeView from './components/HomeView.vue'
import LifeStream from './components/LifeStream.vue'
import ArchiveView from './components/ArchiveView.vue'
import NowView from './components/NowView.vue'
import ProjectsView from './components/ProjectsView.vue'
import AboutView from './components/AboutView.vue'
import FooterArchive from './components/FooterArchive.vue'
import type { LifeEntry, NowStatus, Project } from './types'

const currentView = ref<string>('home')

// 真实生活档案数据 (含默认离线备选)
const entries = ref<LifeEntry[]>([
  {
    id: 1,
    date: '2026.10.09',
    year: '2026',
    month: '10',
    day: '09',
    time: '00:30',
    type: 'thought',
    title: '做一个真正记录生活的网站',
    content: '深夜突然想到：我不想再把自己的网络主页做成一份展示给招聘者或甲方的技术简历了。生活不只有编译、部署和代码架构。下班后迎面吹来的夜风、清晨手冲咖啡的香气、拍下的胶卷、听了整夜的唱片、和朋友玩过的游戏……这些所有具体琐碎的事物，才构成了我真实活着的样子。这里应当是一本安静的个人生活档案馆。',
    meta: '深夜书房随笔',
    tags: '想法,生活志,初心',
    featured: true,
  },
  {
    id: 2,
    date: '2026.10.08',
    year: '2026',
    month: '10',
    day: '08',
    time: '21:40',
    type: 'photo',
    title: '满觉陇的桂花蒸与旁轴底片',
    content: '假期最后几天去了一趟满觉陇。沿山的村落都在做糖桂花，青石板路上落了一层薄薄的金黄，空气甜得发稠。下午四点半的夕阳斜斜穿过樟树叶，光斑在青石砖上摇曳。带了老胶片机拍完了一卷 Kodak 400。阳光晒在皮肤上的温热感，大概就是秋天最好的注脚。',
    images: 'https://images.unsplash.com/photo-1507525428034-b723cf961d3e?auto=format&fit=crop&w=1200&q=80',
    location: '杭州 · 满觉陇',
    meta: 'Contax T2 · 38mm f/2.8 · Kodak Portra 400',
    tags: '摄影,胶片,满觉陇,秋天',
    featured: true,
  },
  {
    id: 3,
    date: '2026.10.08',
    year: '2026',
    month: '10',
    day: '08',
    time: '15:20',
    type: 'coffee',
    title: '浅烘埃塞古吉：92°C 细水慢萃',
    content: '磨了 15g 埃塞俄比亚古吉日晒花魁豆。磨齿带出的干香满是草莓和茉莉花气息。92°C 纯净水，三段式注水萃取。第一段闷蒸 35 秒看咖啡粉像面包一样膨胀，白桃与柑橘酸感明亮。每天看着水柱在滤杯里画同心圆，是让人最沉静的日常仪式。',
    images: 'https://images.unsplash.com/photo-1514432324607-a09d9b4aefdd?auto=format&fit=crop&w=1000&q=80',
    location: '书房窗台前',
    meta: '埃塞古吉花魁 · 15g粉 / 225g水 · 2分18秒',
    tags: '咖啡,手冲,慢生活',
    featured: false,
  },
  {
    id: 4,
    date: '2026.10.07',
    year: '2026',
    month: '10',
    day: '07',
    time: '23:15',
    type: 'music',
    title: '最近开始反复听坂本龙一的《async》',
    content: '深夜关掉大灯，只留一盏暖黄的书桌灯。戴上耳机放《async》。空旷而平静的钢琴琴键、风吹过树梢的采样，整个城市好像都睡熟了。在这样的底噪里看几页书、记几笔日记，心底有一种极其踏实的宁静。',
    images: 'https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?auto=format&fit=crop&w=1000&q=80',
    meta: 'Ryuichi Sakamoto - async (2017) · 推荐曲目：《andata》《solari》',
    tags: '音乐,坂本龙一,深夜',
    featured: true,
  },
  {
    id: 5,
    date: '2026.10.06',
    year: '2026',
    month: '10',
    day: '06',
    time: '22:30',
    type: 'game',
    title: '和朋友玩了一整晚游戏',
    content: '假期倒数第二天，叫上几个好友连麦打了一整晚游戏。从《黑神话：悟空》讨论到 Steam 合作通关，中间点了一大桶冰镇可乐和炸鸡。好久没有这样毫无目的地和朋友大笑、打打闹闹到深夜了。生活需要这样的留白和放空。',
    images: 'https://images.unsplash.com/photo-1550745165-9bc0b252726f?auto=format&fit=crop&w=1000&q=80',
    meta: 'Steam & PS5 连麦联机 · 快乐放空',
    tags: '游戏,朋友,放空',
    featured: false,
  },
  {
    id: 6,
    date: '2026.10.03',
    year: '2026',
    month: '10',
    day: '03',
    time: '20:10',
    type: 'place',
    title: '杨公堤夜骑：风从荷叶上吹过来的温度',
    content: '傍晚换了身衣服去西湖夜骑。从曲院风荷出发，沿着杨公堤一路穿行到南山路。两侧高大的水杉在路灯下投下斑驳长影，秋夜的风穿过枯荷拂在脸上，带着微凉的水汽。把白天的信息和杂念全抛在脑后，出了一身薄汗，整个人仿佛被晚风洗涤过一遍。',
    images: 'https://images.unsplash.com/photo-1502680390469-be75c86b636f?auto=format&fit=crop&w=1200&q=80',
    location: '杭州 · 西湖杨公堤',
    meta: '公路车巡航 · 里程 21.4 km · 均速 22 km/h',
    tags: '骑行,西湖,夜风,行迹',
    featured: true,
  },
  {
    id: 7,
    date: '2026.09.30',
    year: '2026',
    month: '09',
    day: '30',
    time: '18:00',
    type: 'project',
    title: '校园开源协会业务后端 (SUSE-OAA-BACKEND)',
    content: '为四川轻化工大学开源协会重构了核心业务架构。抽离鉴权拦截器与事务层，让新一届协会学弟学妹在举办活动和技术招新时能有一个稳定干净的后端服务。造轮子最开心的瞬间，莫过于亲手写的逻辑真正跑在校园生活里。',
    images: 'https://images.unsplash.com/photo-1517694712202-14dd9538aa97?auto=format&fit=crop&w=1000&q=80',
    meta: 'Go / Gin / GORM / Campus OpenSource Backbone',
    link: 'https://github.com/suse-edu-cn/SUSE-OAA-BACKEND',
    tags: '造物,Go,校园开源',
    featured: false,
  },
  {
    id: 8,
    date: '2026.09.24',
    year: '2026',
    month: '09',
    day: '24',
    time: '16:30',
    type: 'book',
    title: '重读《禅与摩托车维修艺术》：手艺与良质',
    content: '坐在长椅上吹着风读波西格。书中写：‘佛陀或耶稣坐在排气管边，就跟坐在莲花座上一样正常。’ 当你带着真正的良质（Quality）去对待手头的事物——无论是调校一辆自行车、手冲一杯咖啡，还是雕琢一段逻辑——人与工具就合一了。少一点向外张望的功利，多一点对手艺与生活的敬畏。',
    images: 'https://images.unsplash.com/photo-1497633762265-9d179a990aa6?auto=format&fit=crop&w=1000&q=80',
    meta: '罗伯特·M·波西格 著 · 哲学与手艺书摘',
    tags: '阅读,书摘,良质,手艺',
    featured: true,
  },
  {
    id: 9,
    date: '2026.09.15',
    year: '2026',
    month: '09',
    day: '15',
    time: '14:00',
    type: 'purchase',
    title: '陪伴三年的旧物：无刻机械键盘、钢笔与手账',
    content: '擦拭这把用了三年的无刻机械键盘，键帽表面泛出了温润的光泽。旁边是一支百乐钢笔和一个牛皮纸手账。这些天天陪伴我的物件，沉默却忠诚。‘日用即道’，善待身边的每一件物，日子也会变得更踏实有温度。',
    images: 'https://images.unsplash.com/photo-1587829741301-dc798b83add3?auto=format&fit=crop&w=1000&q=80',
    meta: 'Custom Mechanical Keyboard + Pilot 78G + Midori MD',
    tags: '好物,文具,旧物,陪伴',
    featured: false,
  },
])

// 「此刻的我」当前状态
const nowStatus = ref<NowStatus | null>({
  building: '生活自留地 · 真正记录生活的数字档案馆',
  learning: '编译器底层机制、旁轴胶片冲洗、手冲浅烘水温曲线',
  playing: '黑神话：悟空、塞尔达传说、偶尔和朋友 Steam 连麦放空',
  listening: '坂本龙一 - 《async》/《BTTB》',
  reading: '罗伯特·M·波西格 - 《禅与摩托车维修艺术》',
  thinking: '如何在喧嚣的信息时代，认认真真过好具体的生活，保持安宁与自洽。',
  using: 'MacBook Pro 14, Contax T2, 纯白陶瓷滤杯, Midori MD 方格手账',
  location: '中国 · 杭州 (西湖区 / 余杭)',
  last_updated: '2026.10.09',
})

// 造物项目列表 (Things I made)
const projects = ref<Project[]>([
  {
    id: 1,
    title: 'SUSE-OAA-BACKEND',
    subtitle: '四川轻化工大学开放原子开源协会 · 业务后端',
    description: '为高校开源协会搭建的高可用业务后端体系，支撑协会事务流转、招新管理与校园开源数据接口。',
    story: '大二时参与协会建设，发现校园组织的招新和事务常散落在群聊里。于是和同伴一起用 Go 搭建了这套业务核心。也是我第一次深入将分层架构与领域逻辑应用到实际生产场景中。',
    category: 'Campus OpenSource',
    tags: 'Go,Gin,Campus OpenSource',
    github_url: 'https://github.com/suse-edu-cn/SUSE-OAA-BACKEND',
    status: 'Active',
    featured: true,
    order: 1,
  },
  {
    id: 2,
    title: 'ArchCanvas',
    subtitle: 'AI 辅助 Go 架构设计画布',
    description: '让 AI 与开发者一起，从需求语义理解、ER 实体建模到架构设计，快速生成与构建可运行的 Go 工程骨架。',
    story: '在画系统架构图时常常觉得图与代码是割裂的。我想做一个直观的画板，把拖拽出来的实体和微服务边界实时映射为干净的 Go struct 代码。也是探索 AI 与工程设计工具融合的小实验。',
    category: 'AI & Architecture',
    tags: 'TypeScript,Go,Architecture',
    github_url: 'https://github.com/YN1753/ArchCanvas',
    status: 'Active',
    featured: true,
    order: 2,
  },
  {
    id: 3,
    title: 'GoLens',
    subtitle: '基于交互式状态机的 Go 底层机制透视镜',
    description: '让 GMP 协程调度、三色标记 GC 屏障与 Channel 阻塞机制清晰可见的高交互度运行时可视化工具。',
    story: '学 Go 底层时被各种纸上谈兵的图解绕晕了。既然理解不了抽象，就手写一套状态机把它像电影放映机一样在浏览器里跑起来。做完之后，GMP 的工作偷取机制就像齿轮一样印在脑子里了。',
    category: 'Visualization',
    tags: 'JavaScript,Go Internals,GMP',
    github_url: 'https://github.com/YN1753/GoLens',
    status: 'Stable',
    featured: true,
    order: 3,
  },
  {
    id: 4,
    title: 'AstraLink-Desktop',
    subtitle: '基于 Wails 的图笔记桌面端应用',
    description: '星链 2.0。探索 Go + 前端混合桌面开发（Wails 架构），支持双向链接、图谱可视化与本地隐私优先的知识管理。',
    story: '一直渴望一款数据 100% 留在本地硬盘、不依赖云端账户的个人笔记。用 Wails 把 WebKit 前端和 Go 核心绑在一起，启动轻快，内存只占几十兆。',
    category: 'Desktop App',
    tags: 'Go,Wails,Vue,Desktop',
    github_url: 'https://github.com/YN1753/AstraLink-Desktop',
    status: 'Active',
    featured: false,
    order: 4,
  },
  {
    id: 5,
    title: 'Go-Load',
    subtitle: '轻量分布式 HTTP 压测探针',
    description: '自制轻量级高并发压测工具，毫秒级统计 P90 / P99 延迟直方图。',
    story: '觉得现成压测工具安装繁琐，就用 Go 原生协程池手打了一个发压器。小巧单文件，丢到服务器上就能测。',
    category: 'CLI Tool',
    tags: 'Go,Benchmark,CLI',
    github_url: 'https://github.com/YN1753/Go-Load',
    status: 'Stable',
    featured: false,
    order: 5,
  },
  {
    id: 6,
    title: 'nexus',
    subtitle: '轻量 Linux 服务器管理控制台',
    description: '个人云服务器状态监控与容器服务守护面板。',
    story: '管理自己的轻量云服务器时写的工具，随时扫一眼内存、磁盘和守护进程。',
    category: 'DevOps',
    tags: 'Vue,Linux,DevOps',
    github_url: 'https://github.com/YN1753/nexus',
    status: 'WIP',
    featured: false,
    order: 6,
  },
  {
    id: 7,
    title: 'DeviceDaily',
    subtitle: '支持 Mac 小组件的设备成本统计 App',
    description: '基于 Swift 原生开发，支持 macOS WidgetKit，记录电子产品日均使用成本。',
    story: '想看看手里买了三年的相机和电脑到底均摊到了多少钱一天。写了个简洁的 macOS 桌面 Widget。',
    category: 'macOS App',
    tags: 'Swift,macOS,WidgetKit',
    github_url: 'https://github.com/YN1753/DeviceDaily',
    status: 'Active',
    featured: false,
    order: 7,
  },
])

const navigateTo = (view: string) => {
  currentView.value = view
  window.location.hash = view
}

// 监听 hash 路由
onMounted(async () => {
  const hash = window.location.hash.replace('#', '')
  if (['home', 'life', 'archive', 'now', 'projects', 'about'].includes(hash)) {
    currentView.value = hash
  }

  window.addEventListener('hashchange', () => {
    const h = window.location.hash.replace('#', '')
    if (['home', 'life', 'archive', 'now', 'projects', 'about'].includes(h)) {
      currentView.value = h
    }
  })

  // 从后端实时拉取生活档案与状态
  try {
    const [entriesRes, nowRes, projRes] = await Promise.allSettled([
      fetch('/api/entries').then(r => r.ok ? r.json() : null),
      fetch('/api/now').then(r => r.ok ? r.json() : null),
      fetch('/api/projects').then(r => r.ok ? r.json() : null),
    ])

    if (entriesRes.status === 'fulfilled' && entriesRes.value && entriesRes.value.length > 0) {
      entries.value = entriesRes.value
    }
    if (nowRes.status === 'fulfilled' && nowRes.value) {
      nowStatus.value = nowRes.value
    }
    if (projRes.status === 'fulfilled' && projRes.value && projRes.value.length > 0) {
      projects.value = projRes.value
    }
  } catch {
    // 自动回退本地优雅预置数据
  }
})

watch(currentView, (newV) => {
  window.location.hash = newV
})
</script>

<template>
  <div class="min-h-screen bg-[var(--bg-archive)] text-[var(--ink-primary)] flex flex-col justify-between selection:bg-[var(--accent-warm)]/15 selection:text-[var(--accent-warm)] transition-colors duration-300">
    
    <!-- 极简克制顶部导航 -->
    <HeaderNav :currentView="currentView" @navigate="navigateTo" />

    <!-- 主视图区 -->
    <main class="flex-grow">
      
      <!-- 1. HOME 视图：克制开篇封面 + RECENTLY 最近时间流 -->
      <HomeView
        v-if="currentView === 'home'"
        :entries="entries"
        @navigate="navigateTo"
      />

      <!-- 2. LIFE 视图：完整的生活时间线档案 -->
      <div v-else-if="currentView === 'life'" class="max-w-5xl mx-auto px-5 sm:px-8 py-20 sm:py-28 text-left">
        <header class="pb-10 border-b border-[var(--border-subtle)] mb-12 space-y-3">
          <h1 class="font-serif-editorial text-4xl sm:text-5xl font-normal text-[var(--ink-primary)]">
            LIFE STREAM
          </h1>
          <p class="font-mono-archive text-xs uppercase tracking-widest text-[var(--ink-muted)]">
            A chronological stream of moments, memories, thoughts, and things made.
          </p>
        </header>
        <LifeStream :entries="entries" :showFilters="true" />
      </div>

      <!-- 3. ARCHIVE 视图：按年/按月索引 -->
      <ArchiveView
        v-else-if="currentView === 'archive'"
        :entries="entries"
      />

      <!-- 4. NOW 视图：「此刻的我」 -->
      <NowView
        v-else-if="currentView === 'now'"
        :now="nowStatus"
      />

      <!-- 5. PROJECTS 视图：作为人生经历的一章 (Things I Made) -->
      <ProjectsView
        v-else-if="currentView === 'projects'"
        :projects="projects"
      />

      <!-- 6. ABOUT 视图：自然真诚的个人自白与档案馆初衷 -->
      <AboutView
        v-else-if="currentView === 'about'"
      />

    </main>

    <!-- 极简合规页脚（保留工信部合规备案直链） -->
    <FooterArchive />

  </div>
</template>
