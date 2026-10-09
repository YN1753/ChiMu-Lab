<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import HeaderNav from './components/HeaderNav.vue'
import HomeView from './components/HomeView.vue'
import ArchiveView from './components/ArchiveView.vue'
import NowView from './components/NowView.vue'
import ProjectsView from './components/ProjectsView.vue'
import AboutView from './components/AboutView.vue'
import FooterArchive from './components/FooterArchive.vue'
import TransactionsView from './components/TransactionsView.vue'
import NewEntryModal from './components/NewEntryModal.vue'
import EntryDetailModal from './components/EntryDetailModal.vue'
import StorageModal from './components/StorageModal.vue'
import StatsModal from './components/StatsModal.vue'
import AdminAuthModal from './components/AdminAuthModal.vue'
import FloatingActionButton from './components/FloatingActionButton.vue'
import type {
  LifeEntry,
  NowStatus,
  Project,
  ArchiveCategory,
  StorageStatus,
  GlobalStats,
  Transaction,
} from './types'
import { entryMatchesCategory } from './types'
import {
  fetchEntries,
  fetchEntryById,
  fetchStorageStatus,
  fetchGlobalStats,
  fetchProjects,
  fetchTransactions,
} from './utils/api'

const currentView = ref<string>('home')

// 真实生活档案数据 (涵盖 2026 全年真实生活印记，含默认离线备选)
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
  {
    id: 10,
    date: '2026.09.05',
    year: '2026',
    month: '09',
    day: '05',
    time: '21:00',
    type: 'project',
    title: 'ArchCanvas：AI 辅助 Go 架构设计画布',
    content: '想让 AI 辅助开发不只停留在单文件的补全上，而是从全局需求语义直观生成整套可运行的 Go 架构拓扑。开始构建 ArchCanvas 的 AST 双向同步引擎。',
    meta: 'TypeScript / Go AST / Architecture Canvas',
    link: 'https://github.com/YN1753/ArchCanvas',
    tags: '造物,AI架构,实验',
    featured: false,
  },
  {
    id: 11,
    date: '2026.08.20',
    year: '2026',
    month: '08',
    day: '20',
    time: '19:30',
    type: 'project',
    title: 'GoLens：交互式状态机下的运行时透视镜',
    content: '把 Go 复杂的 GMP 协程调度、三色标记 GC 写屏障做成交互式状态机动画。亲手敲完一套调度流转，才算真正摸到了并发系统的齿轮。',
    meta: 'Go Internals Visualization / GMP / GC Barrier',
    link: 'https://github.com/YN1753/GoLens',
    tags: '造物,Go底层,状态机',
    featured: false,
  },
  {
    id: 12,
    date: '2026.07.18',
    year: '2026',
    month: '07',
    day: '18',
    time: '05:15',
    type: 'photo',
    title: '东极岛的日出与第一缕海风',
    content: '凌晨四点半爬起来走到东福山岛的最东端。海平线先是从深蓝泛出极淡的鹅黄，然后是一抹不可思议的粉紫。浪头撞在黑色的礁石上激起三米高的雪白浪花。坐在礁石上吹着咸湿的海风，天地浩大，人的那些微小烦恼瞬间显得微不足道。',
    images: 'https://images.unsplash.com/photo-1506744038136-46273834b3fb?auto=format&fit=crop&w=1200&q=80',
    location: '舟山 · 东极岛东福山',
    meta: 'Contax T2 · 38mm f/2.8 · Kodak Ektar 100',
    tags: '摄影,旅行,海边,日出',
    featured: true,
  },
  {
    id: 13,
    date: '2026.06.21',
    year: '2026',
    month: '06',
    day: '21',
    time: '19:45',
    type: 'thought',
    title: '夏至黄昏的暴雨与空街',
    content: '夏至这天傍晚，暴雨毫无预兆地倾盆而下。站在便利店屋檐下等雨停，看着水流顺着沥青路面汇入下水道，空气里满是泥土被雨水浇透后的清新气味。城市突然按下了静音键。很享受这样被意外困住的片刻，不用赶路，只要看雨就好。',
    meta: '雨天屋檐随记',
    tags: '随笔,夏至,雨天',
    featured: false,
  },
  {
    id: 14,
    date: '2026.05.02',
    year: '2026',
    month: '05',
    day: '02',
    time: '14:30',
    type: 'place',
    title: '莫干山竹林徒步：春末初夏的青翠',
    content: '五一假期避开人群，钻进莫干山后山的野竹林里徒步了五公里。阳光从密密匝匝的竹叶缝隙里漏下来，脚下的泥土松软带着竹叶香。山涧溪水冰凉刺骨，捧起来洗了把脸，神清气爽。',
    images: 'https://images.unsplash.com/photo-1448375240586-882707db888b?auto=format&fit=crop&w=1200&q=80',
    location: '浙江 · 湖州莫干山',
    meta: '山野徒步 · 5.8 km 穿行',
    tags: '徒步,山野,自然,行迹',
    featured: false,
  },
  {
    id: 15,
    date: '2026.04.12',
    year: '2026',
    month: '04',
    day: '12',
    time: '16:00',
    type: 'coffee',
    title: '哥伦比亚双重厌氧瑰夏：水蜜桃与肉桂尾韵',
    content: '收到了新到的哥伦比亚蕙兰双重厌氧瑰夏豆。研磨时的香气像刚切开的多汁水蜜桃。90°C 水温萃取，前段是明艳的水果红茶感，温度降到微温后，肉桂与红糖的尾韵浮现出来。好的豆子确实能让人在舌尖上旅行。',
    images: 'https://images.unsplash.com/photo-1495474472287-4d71bcdd2085?auto=format&fit=crop&w=1000&q=80',
    location: '书房',
    meta: '哥伦比亚双重厌氧瑰夏 · V60滤杯',
    tags: '咖啡,手冲,瑰夏',
    featured: false,
  },
  {
    id: 16,
    date: '2026.03.20',
    year: '2026',
    month: '03',
    day: '20',
    time: '23:40',
    type: 'music',
    title: '深夜重听 Bill Evans《Sunday at the Village Vanguard》',
    content: '春分这天夜里，重新翻出 Bill Evans 1961 年在先锋村的现场录音。背景里隐隐约约有酒杯轻碰的声音、低声的交谈，以及保罗·莫蒂安极其克制细腻的刷子鼓点。生活不需要时刻紧绷，爵士乐里那种即兴与松弛，是面对复杂世界的最好解药。',
    images: 'https://images.unsplash.com/photo-1514525253161-7a46d19cd819?auto=format&fit=crop&w=1000&q=80',
    meta: 'Bill Evans Trio - Sunday at the Village Vanguard (1961)',
    tags: '音乐,爵士乐,深夜,黑胶',
    featured: false,
  },
  {
    id: 17,
    date: '2026.02.10',
    year: '2026',
    month: '02',
    day: '10',
    time: '01:20',
    type: 'game',
    title: '春节假期通关《星际拓荒 Outer Wilds》',
    content: '在老家安静的深夜里打完了 Outer Wilds 的最后一幕。当所有乐器在宇宙尽头合奏起那首主题曲时，坐在屏幕前久久不能平静。它不仅是一款游戏，更是一首写给好奇心、科学探索与生命终极意义的伟大散文诗。',
    images: 'https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?auto=format&fit=crop&w=1000&q=80',
    meta: 'Mobius Digital · 22分钟循环的终点',
    tags: '游戏,OuterWilds,科幻,感动',
    featured: false,
  },
  {
    id: 18,
    date: '2026.01.01',
    year: '2026',
    month: '01',
    day: '01',
    time: '10:00',
    type: 'thought',
    title: '给二零二六年的自己写一封信',
    content: '新年第一天清晨，在手账本第一页写下：新的一年，不求做多宏大的事情，但求认认真真过好具体的一天天。多看几卷胶卷，多冲几杯好咖啡，多去山林和湖边吹风，少一些盲目的自我消耗。真实而安静地生活。',
    meta: '新年手账寄语',
    tags: '新年,随笔,初心,生活志',
    featured: true,
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

// 造物项目列表 (Things I made - 强调经历与手艺，非商业堆叠)
const projects = ref<Project[]>([
  {
    id: 1,
    title: 'SUSE-OAA-BACKEND',
    subtitle: '四川轻化工大学开放原子开源协会 · 业务后端',
    description: '为高校开源协会搭建的高可用业务后端体系，支撑协会事务流转、招新管理与校园开源数据接口。',
    story: '大二时参与协会建设，发现校园组织的招新和事务常散落在群聊里。于是和同伴一起用 Go 搭建了这套业务核心。也是我第一次深入将分层架构与领域逻辑应用到实际生产场景中。',
    category: '校园开源',
    tags: 'Go,Gin,校园开源',
    github_url: 'https://github.com/suse-edu-cn/SUSE-OAA-BACKEND',
    status: '运行中',
    featured: true,
    order: 1,
  },
  {
    id: 2,
    title: 'ArchCanvas',
    subtitle: 'AI 辅助 Go 架构设计画布',
    description: '让 AI 与开发者一起，从需求语义理解、ER 实体建模到架构设计，快速生成与构建可运行的 Go 工程骨架。',
    story: '在画系统架构图时常常觉得图与代码是割裂的。我想做一个直观的画板，把拖拽出来的实体和微服务边界实时映射为干净的 Go struct 代码。也是探索 AI 与工程设计工具融合的小实验。',
    category: '架构实验',
    tags: 'TypeScript,Go,Architecture',
    github_url: 'https://github.com/YN1753/ArchCanvas',
    status: '迭代中',
    featured: true,
    order: 2,
  },
  {
    id: 3,
    title: 'GoLens',
    subtitle: '基于交互式状态机的 Go 底层机制透视镜',
    description: '让 GMP 协程调度、三色标记 GC 屏障与 Channel 阻塞机制清晰可见的高交互度运行时可视化工具。',
    story: '学 Go 底层时被各种纸上谈兵的图解绕晕了。既然理解不了抽象，就手写一套状态机把它像电影放映机一样在浏览器里跑起来。做完之后，GMP 的工作偷取机制就像齿轮一样印在脑子里了。',
    category: '底层透视',
    tags: 'JavaScript,Go Internals,GMP',
    github_url: 'https://github.com/YN1753/GoLens',
    status: '已稳定',
    featured: true,
    order: 3,
  },
  {
    id: 4,
    title: 'AstraLink-Desktop',
    subtitle: '基于 Wails 的图笔记桌面端应用',
    description: '星链 2.0。探索 Go + 前端混合桌面开发（Wails 架构），支持双向链接、图谱可视化与本地隐私优先的知识管理。',
    story: '一直渴望一款数据 100% 留在本地硬盘、不依赖云端账户的个人笔记。用 Wails 把 WebKit 前端和 Go 核心绑在一起，启动轻快，内存只占几十兆。',
    category: '桌面工具',
    tags: 'Go,Wails,Vue,Desktop',
    github_url: 'https://github.com/YN1753/AstraLink-Desktop',
    status: '自用中',
    featured: false,
    order: 4,
  },
  {
    id: 5,
    title: 'Go-Load',
    subtitle: '轻量分布式 HTTP 压测探针',
    description: '自制轻量级高并发压测工具，毫秒级统计 P90 / P99 延迟直方图。',
    story: '觉得现成压测工具安装繁琐，就用 Go 原生协程池手打了一个发压器。小巧单文件，丢到服务器上就能测。',
    category: '命令行工具',
    tags: 'Go,Benchmark,CLI',
    github_url: 'https://github.com/YN1753/Go-Load',
    status: '已归档',
    featured: false,
    order: 5,
  },
  {
    id: 6,
    title: 'nexus',
    subtitle: '轻量 Linux 服务器管理控制台',
    description: '个人云服务器状态监控与容器服务守护面板。',
    story: '管理自己的轻量云服务器时写的工具，随时扫一眼内存、磁盘和守护进程。',
    category: '运维监控',
    tags: 'Vue,Linux,DevOps',
    github_url: 'https://github.com/YN1753/nexus',
    status: '维护中',
    featured: false,
    order: 6,
  },
  {
    id: 7,
    title: 'DeviceDaily',
    subtitle: '支持 Mac 小组件的设备成本统计 App',
    description: '基于 Swift 原生开发，支持 macOS WidgetKit，记录电子产品日均使用成本。',
    story: '想看看手里买了三年的相机和电脑到底均摊到了多少钱一天。写了个简洁的 macOS 桌面 Widget。',
    category: 'macOS 原生',
    tags: 'Swift,macOS,WidgetKit',
    github_url: 'https://github.com/YN1753/DeviceDaily',
    status: '日常自用',
    featured: false,
    order: 7,
  },
])

const currentCategory = ref<ArchiveCategory>('all')

const selectedEntry = ref<LifeEntry | null>(null)
const isNewEntryOpen = ref(false)
const isStorageOpen = ref(false)
const isStatsOpen = ref(false)
const isAdminAuthOpen = ref(false)
const storageStatus = ref<StorageStatus | null>(null)
const globalStats = ref<GlobalStats | null>(null)

// 记账列表 (与生活流交融)
const transactionsList = ref<Transaction[]>([
  {
    id: 1,
    amount: 12900,
    type: 'expense',
    category: '购物',
    title: 'MCHOSE A7 无线鼠标',
    note: '替换用了四年的旧鼠标',
    payment_method: '微信支付',
    occurred_at: '2026-10-09T14:20:00Z',
  },
  {
    id: 2,
    amount: 1800,
    type: 'expense',
    category: '餐饮',
    title: '午饭面条',
    note: '公司楼下片儿川',
    payment_method: '支付宝',
    occurred_at: '2026-10-09T12:15:00Z',
  },
  {
    id: 3,
    amount: 6800,
    type: 'expense',
    category: '餐饮',
    title: '埃塞俄比亚咖啡生豆',
    note: '浅烘花魁 250g',
    payment_method: '微信支付',
    occurred_at: '2026-10-08T16:30:00Z',
  },
  {
    id: 4,
    amount: 500000,
    type: 'income',
    category: '其他',
    title: '稿酬与项目结项',
    note: '校园开源平台二期补贴',
    payment_method: '银行转账',
    occurred_at: '2026-10-07T10:00:00Z',
  },
])

const storageConfigured = computed(() => Boolean(storageStatus.value?.configured))

const txToEntry = (tx: Transaction): LifeEntry => {
  const d = new Date(tx.occurred_at)
  const pad = (n: number) => String(n).padStart(2, '0')
  const year = !isNaN(d.getTime()) ? String(d.getFullYear()) : '2026'
  const month = !isNaN(d.getTime()) ? pad(d.getMonth() + 1) : '10'
  const day = !isNaN(d.getTime()) ? pad(d.getDate()) : '09'
  const time = !isNaN(d.getTime()) ? `${pad(d.getHours())}:${pad(d.getMinutes())}` : ''

  return {
    id: 100000 + tx.id,
    type: 'transaction',
    title: tx.title || tx.category || '记账',
    content: tx.note || '',
    occurred_at: tx.occurred_at,
    date: `${year}.${month}.${day}`,
    year,
    month,
    day,
    time,
    amount: tx.amount,
    tx_type: tx.type,
    category: tx.category,
    payment_method: tx.payment_method,
    is_transaction: true,
  }
}

// 融合了生活记录与记账的完整时间流
const allTimelineEntries = computed(() => {
  const txEntries = transactionsList.value.map(txToEntry)
  const combined = [...entries.value, ...txEntries]
  combined.sort((a, b) => {
    const timeA = a.occurred_at
      ? new Date(a.occurred_at).getTime()
      : a.date
      ? new Date(a.date.replace(/\./g, '-')).getTime()
      : 0
    const timeB = b.occurred_at
      ? new Date(b.occurred_at).getTime()
      : b.date
      ? new Date(b.date.replace(/\./g, '-')).getTime()
      : 0
    return timeB - timeA
  })
  return combined
})

const filteredEntriesByCategory = computed(() => {
  if (currentCategory.value === 'all') {
    return allTimelineEntries.value
  }
  return allTimelineEntries.value.filter((e) => entryMatchesCategory(e, currentCategory.value))
})

const handleCategoryChange = (cat: ArchiveCategory) => {
  currentCategory.value = cat
  if (currentView.value !== 'home' && currentView.value !== 'life') {
    currentView.value = 'home'
    window.location.hash = 'home'
  }
}

const navigateTo = (view: string) => {
  currentView.value = view
  window.location.hash = view
}

const openEntryDetail = (entry: LifeEntry) => {
  if (entry.is_transaction) {
    navigateTo('transactions')
    return
  }
  selectedEntry.value = entry
  window.location.hash = `entry/${entry.id}`
}

const closeEntryDetail = () => {
  selectedEntry.value = null
  if (window.location.hash.startsWith('#entry/')) {
    window.location.hash = currentView.value
  }
}

const openStatsModal = async () => {
  isStatsOpen.value = true
  try {
    globalStats.value = await fetchGlobalStats()
  } catch {
    // ignore
  }
}

const openStorageModal = async () => {
  isStorageOpen.value = true
  try {
    storageStatus.value = await fetchStorageStatus()
  } catch {
    // ignore
  }
}

const handleEntryCreated = (newEntry: LifeEntry) => {
  entries.value = [newEntry, ...entries.value.filter((e) => e.id !== newEntry.id)]
  refreshData()
  if (currentView.value !== 'home' && currentView.value !== 'life') {
    navigateTo('home')
  }
}

const handleEntryUpdated = (updated: LifeEntry) => {
  const idx = entries.value.findIndex((e) => e.id === updated.id)
  if (idx !== -1) {
    entries.value[idx] = updated
  }
  selectedEntry.value = updated
  refreshData()
}

const handleEntryDeleted = (id: number) => {
  entries.value = entries.value.filter((e) => e.id !== id)
  if (selectedEntry.value?.id === id) {
    selectedEntry.value = null
  }
  if (window.location.hash.startsWith('#entry/')) {
    window.location.hash = currentView.value
  }
  refreshData()
}

const handleTxCreated = (_tx: Transaction) => {
  refreshData()
}

const handleTxUpdated = (_tx: Transaction) => {
  refreshData()
}

const handleTxDeleted = (_id: number) => {
  refreshData()
}

const handleProjectCreated = (newProj: Project) => {
  projects.value = [newProj, ...projects.value]
  refreshData()
}

const refreshData = async () => {
  try {
    const [entriesData, statsData, statusData, txData] = await Promise.allSettled([
      fetchEntries({ page: 1, page_size: 50 }),
      fetchGlobalStats(),
      fetchStorageStatus(),
      fetchTransactions(),
    ])
    if (entriesData.status === 'fulfilled' && entriesData.value && entriesData.value.length > 0) {
      entries.value = entriesData.value
    }
    if (statsData.status === 'fulfilled' && statsData.value) {
      globalStats.value = statsData.value
    }
    if (statusData.status === 'fulfilled' && statusData.value) {
      storageStatus.value = statusData.value
    }
    if (txData.status === 'fulfilled' && txData.value && txData.value.length > 0) {
      transactionsList.value = txData.value
    }
  } catch {
    // ignore
  }
}

const handleHashChange = async () => {
  const hash = window.location.hash.replace('#', '')
  if (hash.startsWith('entry/')) {
    const idStr = hash.replace('entry/', '')
    const id = parseInt(idStr, 10)
    if (!isNaN(id)) {
      const found = entries.value.find((e) => e.id === id)
      if (found) {
        selectedEntry.value = found
      } else {
        try {
          const fetched = await fetchEntryById(id)
          if (fetched) selectedEntry.value = fetched
        } catch {
          // ignore
        }
      }
    }
    return
  }

  if (['home', 'life', 'transactions', 'archive', 'now', 'projects', 'about'].includes(hash)) {
    currentView.value = hash
    selectedEntry.value = null
  }
}

// 监听 hash 路由与拉取数据
onMounted(async () => {
  await handleHashChange()
  window.addEventListener('hashchange', handleHashChange)

  // 初始拉取
  await refreshData()

  try {
    const [nowRes, projRes] = await Promise.allSettled([
      fetch('/api/now').then((r) => (r.ok ? r.json() : null)),
      fetchProjects(),
    ])

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
  if (!window.location.hash.startsWith('#entry/')) {
    window.location.hash = newV
  }
})
</script>

<template>
  <div class="min-h-screen bg-[var(--bg-archive)] text-[var(--ink-primary)] flex flex-col justify-between selection:bg-[var(--accent-warm)]/15 selection:text-[var(--accent-warm)] transition-colors duration-300">
    
    <!-- 极简克制顶部导航（集成右上角「＋」与「我的生活⌄」） -->
    <HeaderNav
      :currentView="currentView"
      :currentCategory="currentCategory"
      @navigate="navigateTo"
      @changeCategory="handleCategoryChange"
      @openAdd="isNewEntryOpen = true"
      @openStats="openStatsModal"
      @openStorage="openStorageModal"
      @openAuth="isAdminAuthOpen = true"
    />

    <!-- 主视图区 -->
    <main class="flex-grow">
      
      <!-- 1. HOME 视图：克制开篇封面 + 2026生活刻度热力图 + 当下剪影 + 近况生活流 -->
      <HomeView
        v-if="currentView === 'home'"
        :entries="filteredEntriesByCategory"
        :allEntries="allTimelineEntries"
        :now="nowStatus"
        :currentCategory="currentCategory"
        @navigate="navigateTo"
        @changeCategory="handleCategoryChange"
        @selectEntry="openEntryDetail"
        @openAdd="isNewEntryOpen = true"
      />

      <!-- 2. TRANSACTIONS 视图：独立记账模块 -->
      <TransactionsView
        v-else-if="currentView === 'transactions'"
        @openAdd="isNewEntryOpen = true"
        @transactionDeleted="handleTxDeleted"
        @transactionUpdated="handleTxUpdated"
      />

      <!-- 4. ARCHIVE 视图：按年/按月索引 -->
      <ArchiveView
        v-else-if="currentView === 'archive'"
        :entries="entries"
      />

      <!-- 5. NOW 视图：「此刻的我」 -->
      <NowView
        v-else-if="currentView === 'now'"
        :now="nowStatus"
      />

      <!-- 6. PROJECTS 视图：作为人生经历的一章 (造物 / Things I Made) -->
      <ProjectsView
        v-else-if="currentView === 'projects'"
        :projects="projects"
      />

      <!-- 7. ABOUT 视图：自然真诚的个人自白与档案馆初衷 -->
      <AboutView
        v-else-if="currentView === 'about'"
      />

    </main>

    <!-- 极简合规页脚（保留工信部合规备案直链） -->
    <FooterArchive />

    <!-- 全局快捷新增浮动操作按钮 (桌面端 + 移动端右下角) -->
    <FloatingActionButton @click="isNewEntryOpen = true" />

    <!-- 1. 新增生活记录/记账 弹窗 -->
    <NewEntryModal
      :isOpen="isNewEntryOpen"
      :storageConfigured="storageConfigured"
      :projects="projects"
      @close="isNewEntryOpen = false"
      @entryCreated="handleEntryCreated"
      @transactionCreated="handleTxCreated"
      @projectCreated="handleProjectCreated"
    />

    <!-- 2. 生活记录详情 / 编辑 / 删除 弹窗 -->
    <EntryDetailModal
      :entry="selectedEntry"
      :storageConfigured="storageConfigured"
      :projects="projects"
      @close="closeEntryDetail"
      @updated="handleEntryUpdated"
      @deleted="handleEntryDeleted"
    />

    <!-- 3. 对象存储状态 弹窗 -->
    <StorageModal
      :isOpen="isStorageOpen"
      :status="storageStatus"
      @close="isStorageOpen = false"
    />

    <!-- 4. 统计指标 弹窗 -->
    <StatsModal
      :isOpen="isStatsOpen"
      :stats="globalStats"
      @close="isStatsOpen = false"
    />

    <!-- 5. 暗房管理员钥匙 弹窗 -->
    <AdminAuthModal
      :isOpen="isAdminAuthOpen"
      @close="isAdminAuthOpen = false"
    />

  </div>
</template>
