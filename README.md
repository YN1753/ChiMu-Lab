# ChiMu-Lab (迟暮实验室) 🎞️

> **Digital Atelier & Living Archive · 迟暮的私人数字暗房与生活档案**
>
> 🌐 线上站点: [https://codeactivityhub.top](https://codeactivityhub.top)  
> 🔗 GitHub 个人主页: [https://github.com/YN1753](https://github.com/YN1753)  
> 📜 ICP 备案号: [浙ICP备2026081664号](https://beian.miit.gov.cn/)

---

## 📖 项目定位与设计哲学

**ChiMu-Lab** 摒弃了千篇一律的程序员求职简历模板与商业宣传辞令，立足于**“写给自己看”的诚恳态度**，打造为一个充满呼吸感、电影感与实体触感的**个人数字暗房与生活工坊 (Digital Atelier & Living Archive)**。

它融合了双轨叙事：
1. **左手造物（The Craft）**：记录亲手编写的真实工程、底层状态机调度、架构设计画布与踩过的坑。
2. **右手生活（Living Vignettes）**：收集秋日的桂花、夜雨的咖啡、读书与音乐的手记、胶片相纸与吉光片羽。

---

## ✨ 核心亮点与高阶交互

### 1. 📷 3D 拍立得胶卷手记 (`LivingFrames.vue`)
- **真实生活切片流**：包含夜雨中的通道阻塞思考、满觉陇金桂胶片、手冲曼特宁与 AST 遍历反哺、《禅与摩托车维修艺术》阅读手记、坂本龙一的音符与 Goroutine 死锁排查。
- **3D 物理翻转 (Flip Card)**：鼠标滑过时呈现 3D 悬浮透视；点击“翻转查看手记”沿 Y 轴 180° 平滑翻转，背面展现复古明信片纹理、手写便签与杭州邮戳。
- **暖心盖章互动**：点击红心可实时在 SQLite 数据库中完成盖章持久化，伴随机械快门声。

### 2. 🎛️ 电影控制台与 Web Audio 声学生成器 (`Navbar.vue` & `utils/audio.ts`)
- **三套电影 LUT 调色盘**：
  - **暖骨白 (Alabaster)**：温润米白 `#f7f6f2`，纸质漫反射光晕。
  - **琥珀暮 (Sunset Amber)**：柯达 400 暖金黄昏调，落日余晖。
  - **暗房黑 (Midnight Noir)**：暗房深蓝黑与微荧光，深夜沉浸感。
- **零外部音频加载的 Web Audio 原生合成**：
  - **机械快门音**：点击切换主题、翻转卡片或盖章时，合成清脆的相机帘幕快门声。
  - **秋夜微雨自然白噪音**：432Hz 动态粉红噪声带通滤波器，点击即可开启沉浸细雨声。

### 3. 📐 真实造物工坊 · 架构蓝图底片抽屉 (`ProjectShowcase.vue`)
- **#1 首位置顶核心**：[`suse-edu-cn/SUSE-OAA-BACKEND`](https://github.com/suse-edu-cn/SUSE-OAA-BACKEND)（四川轻化工大学开放原子开源协会业务后端服务体系）。
- **#2 架构核心**：[`YN1753/ArchCanvas`](https://github.com/YN1753/ArchCanvas)（AI 辅助 Go 架构设计画布）。
- **#3 底层透视**：[`YN1753/GoLens`](https://github.com/YN1753/GoLens)（基于状态机的 Go GMP / GC 可视化透视镜）。
- **后续真实仓库**：`AstraLink-Desktop` (Wails 图笔记)、`Go-Load` (高并发压测)、`nexus` (基础设施)、`DeviceDaily` (Swift 原生小组件)。
- **蓝图底片抽屉 (Engineering Dossier)**：点击任意项目唤出半透明蓝图网格视窗，展示工程初衷、关键架构切片与一键复制 `git clone` 命令。

### 4. ⏱️ 当下的迟暮 · 实时状态轮盘 (`Hero.vue`)
- 顶部电影打板横标：`SCENE: CHIMU-ATELIER / TAKE: 2026.FALL / HANGZHOU (30.27° N)` 与秒级流动时钟。
- 点击状态胶囊可顺畅切换迟暮此时此刻的状态（手冲耶加雪菲 / 排查 suseoaa 协程 / 西湖边夜骑 / 调优 ArchCanvas AST），带有清脆微音。

---

## 🛠️ 技术栈与架构选型

| 模块 | 技术选型 | 优势与说明 |
| :--- | :--- | :--- |
| **后端语言** | **Go 1.27** | 高性能微架构、毫秒级响应、超低内存开销 (~1.8MB) |
| **Web 框架** | **Gin** | 高性能 REST API，集成 SPA 静态资源路由 |
| **数据库** | **SQLite** (`github.com/glebarez/sqlite`) | **纯 Go 实现 (CGO-Free)**，无外部 C 编译器依赖，轻松 cross-compile |
| **ORM** | **GORM** | 自动迁移模型（Profile, Project, LifeMoment, Activity, SiteConfig） |
| **前端框架** | **Vue 3** (Composition API) | 现代化响应式前端，极速开发体验 |
| **构建工具** | **Vite** | 生产打包耗时仅 ~190ms |
| **样式工程** | **Tailwind CSS v4** | 原子化 CSS，结合电影感自定义主题变量 |
| **声学引擎** | **Web Audio API** | 纯浏览器端算法合成快门与自然雨声，0KB 外部音效加载 |
| **字体呈现** | **Newsreader** + **Plus Jakarta Sans** | 电影海报衬线体与现代几何字体的高级混排 |

---

## 📂 项目结构全景

```text
ChiMu-Lab/
├── main.go                     # Go 后端启动入口（集成 SPA 静态资源分发）
├── go.mod / go.sum             # Go 依赖配置
├── internal/
│   ├── models/                 # SQLite 数据库模型 (Profile, Project, LifeMoment, Activity, SiteConfig)
│   ├── db/                     # 数据库连接、AutoMigrate 与种子数据同步 (seedData)
│   └── handlers/               # REST API 处理器 (/api/projects, /api/moments, /api/stats, etc.)
├── web/                        # Vue 3 前端工程
│   ├── index.html              # 页面 HTML 模板与 Newsreader 字体加载
│   ├── vite.config.ts          # Vite 构建配置
│   ├── package.json            # 前端依赖配置
│   └── src/
│       ├── App.vue             # 页面根组件与全量 API 数据驱动
│       ├── main.ts             # 前端入口
│       ├── style.css           # 电影 LUT 调色盘变量、3D 翻转与蓝图网格 CSS
│       ├── types/              # TypeScript 接口定义 (Project, LifeMoment, etc.)
│       ├── utils/
│       │   └── audio.ts        # Web Audio API 机械快门与秋夜微雨合成器
│       └── components/
│           ├── Navbar.vue           # 顶部电影控制台（LUT 调色盘 + 雨声/快门开关）
│           ├── Hero.vue             # 电影打板头图 + 迟暮当下状态切频轮盘
│           ├── ProjectShowcase.vue  # 造物工坊（suseoaa 置顶 #1 + 蓝图底片抽屉）
│           ├── LivingFrames.vue     # 生活切片（3D 拍立得翻转卡片 + 盖章互动）
│           ├── ActivityHub.vue      # 52周代码热力刻度与语言成分条
│           └── ComplianceFooter.vue # 备案合规页脚（浙ICP备2026081664号直链）
└── README.md
```

---

## 🚀 本地开发与生产部署

### 1. 本地启动
```bash
# 1. 启动 Go 服务 (默认 :8080)
go run main.go

# 2. 启动前端开发调试 (新终端)
cd web && npm run dev
```

### 2. 生产打包与上传部署
```bash
# 1. 前端打包
cd web && npm run build

# 2. 编译 Linux AMD64 静态无依赖二进制
cd ..
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o chimu-server-linux main.go

# 3. 推送至腾讯云服务器并重启服务
tar -czf - chimu-server-linux web/dist | ssh ubuntu@101.35.227.224 "tar -xzf - -C /home/ubuntu/chimu-lab && sudo systemctl restart chimu-lab"
```

---

## 📄 执照与版权

MIT License © 2026 [迟暮 (ChiMu)](https://github.com/YN1753)
