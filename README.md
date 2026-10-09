# 散页 · Loose Leaf

> **Digital Atelier & Living Archive · 迟暮的数字生活档案馆与造物工坊**
>
> 线上站点: [https://codeactivityhub.top](https://codeactivityhub.top)  
> 主理人 GitHub: [https://github.com/YN1753](https://github.com/YN1753)  
> 备案号: [浙ICP备2026081664号](https://beian.miit.gov.cn/)

---

## 项目定位与设计哲学

**《散页》 (Loose Leaf)** 摒弃了千篇一律的程序员技能简历模板与商业宣传辞令，立足于**“写给自己看”的诚恳态度**，打造为一个充满纸张微噪点质感、出版物美感与时间沉淀质感的**个人数字生活档案馆 (Digital Atelier & Living Archive)**。主理人为**迟暮**。

核心理念：**生活大于项目，真实大于表演**。
- **时间按顺序流淌的真实印记**：收集秋日的桂花、夜雨的咖啡、读书与音乐的手记、胶片相纸与吉光片羽。
- **造物作为人生经历的一章**：记录亲手编写的真实工程、底层状态机调度、架构设计画布与踩过的坑。
- **柴米油盐融入生活流**：独立记账模块，把握日常真实开销与生活步调。

---

## ✨ 核心亮点与交互设计

### 1. 📜 纵向脊柱时间线 (`LifeStream.vue`)
- **多态内容排版 (Polymorphic Layouts)**：
  - **想法 / 随笔 (Thought)**：免标题，舒适大字号引言体与轻盈随手记。
  - **记账微卡 (Transaction)**：单行账本微卡片，收支自然融于时间流。
  - **胶卷相纸 (Photo)**：全幅大图与胶卷参数注脚。
  - **造物手艺 (Project)**：工程初衷、关键技术切片与源码直链。
  - **日常生活 (Daily)**：咖啡慢萃、黑胶听音、西湖夜骑与好友联机。
- **七重視角灵活切换**：全部记录、日常、想法、照片、记账、项目、收藏一键筛选。

### 2. 📅 2026 生活刻度热力图 (`LifeActivityMap.vue`)
- **全年 52 周刻度网格**：每一格都是真实度过的一天，支持悬浮浏览当日生活印记。
- **一键点击日期穿透**：点击任意有记录的刻度格，无缝联动定位并聚焦当天时间线。

### 3. 💳 独立记账体系 (`TransactionsView.vue`)
- **月度核心指标**：本月总支出、总收入与结余实时核算。
- **按日汇聚流水**：自动按自然日汇聚账目列表，支持内联编辑与分类管理。
- **金额采用整型“分”存储**：避免浮点数精度截断误差。

### 4. 🔒 管理员暗房钥匙 (`AdminAuthModal.vue` & `middleware/auth.go`)
- **写接口安全防护**：通过 `ADMIN_API_KEY` 保护所有新建、编辑、删除与存储上传接口。
- **极简本地暗房钥匙**：前端提供密码弹窗，输入一次自动存入本地浏览器，长期静默免密畅享记录。

### 5. ☁️ 对象存储直传 (`ImageUploader.vue` & `storage/r2.go`)
- **Cloudflare R2 预签名 PUT 直传**：浏览器直接将媒体文件上传至对象存储，不消耗云主机中转带宽。
- **零配置优雅降级**：未配置 R2 时纯文字记录与记账 100% 正常运行，弹窗清晰呈现存储状态。

### 6. 🌧️ Web Audio 自然声学生成器 (`utils/audio.ts`)
- **零外部音频加载**：432Hz 动态粉红噪声带通滤波器，算法实时合成沉浸式秋夜微雨白噪音。
- **按钮微触感**：高频晶莹微音效，营造机械键盘般的实体操控触感。

---

## 🛠️ 技术栈与架构选型

| 模块 | 技术选型 | 说明 |
| :--- | :--- | :--- |
| **后端语言** | **Go 1.24+** | 极速响应、内存占用低 (~1.8MB) |
| **Web 框架** | **Gin** | 高性能 REST API，内置优雅停机与 SPA 静态文件托管 |
| **数据库** | **SQLite** (`github.com/glebarez/sqlite`) | 纯 Go 实现 (CGO-Free)，开启 **WAL 模式**与连接池调优 |
| **ORM** | **GORM** | 自动迁移模型（LifeEntry, Attachment, Project, Transaction, NowStatus, SiteConfig） |
| **前端框架** | **Vue 3** (Composition API) | 响应式组件化开发 |
| **构建工具** | **Vite** | 极速热更新，生产打包构建耗时 ~250ms |
| **样式工程** | **Tailwind CSS v4** | 现代原子化 CSS，出版物纸质漫反射色彩变量 |
| **声学引擎** | **Web Audio API** | 纯浏览器端算法实时合成雨声，0KB 外部音效文件 |

---

## 📂 项目结构全景

```text
ChiMu-Lab/
├── main.go                     # Go 后端启动入口（集成鉴权、SPA 托管与优雅停机）
├── Dockerfile                  # 多阶段轻量生产容器构建
├── docker-compose.yml          # Docker 一键编排配置
├── go.mod / go.sum             # Go 依赖配置
├── .env.example                # 环境变量配置模板
├── internal/
│   ├── config/                 # 配置中心（支持 .env 自动加载与 R2/Admin 配置）
│   ├── db/                     # SQLite 数据库连接（WAL 模式）、迁移与种子数据
│   ├── middleware/             # 中间件（AdminAuthRequired 管理员鉴权）
│   ├── models/                 # 数据模型 (LifeEntry, Attachment, Project, Transaction, etc.)
│   ├── storage/                # Cloudflare R2 S3 预签名直传客户端
│   └── handlers/               # RESTful API 处理器与弹性时间解析器
├── web/                        # Vue 3 前端工程
│   ├── index.html              # 页面 HTML 模板与字体加载
│   ├── vite.config.ts          # Vite 构建配置
│   ├── package.json            # 前端依赖配置
│   └── src/
│       ├── App.vue             # 页面根组件与统一弹窗分发
│       ├── main.ts             # 前端入口
│       ├── style.css           # 纸质色调、Newsreader 衬线字体与等宽排版变量
│       ├── types/              # TypeScript 核心类型定义
│       ├── utils/
│       │   ├── api.ts          # 统一 API 请求封装（自动附加 Bearer 鉴权）
│       │   └── audio.ts        # Web Audio API 微音效与秋夜微雨合成器
│       └── components/
│           ├── HeaderNav.vue        # 顶部极简导航（单轨栏目 + 设置浮层 + 雨声）
│           ├── HomeView.vue         # 首页视图（自白开篇 + 52周生活刻度 + 时间流）
│           ├── LifeStream.vue       # 纵向脊柱时间线（多态排版与7重视角筛选）
│           ├── LifeActivityMap.vue  # 2026 生活刻度热力图（52周动态刻度与日期穿透）
│           ├── TransactionsView.vue # 独立记账视图（月度收支指标与按日流水）
│           ├── ProjectsView.vue     # 造物视图（作为经历一章的项目故事与源码）
│           ├── NowView.vue          # 此刻视图（Now Page 个人状态切片）
│           ├── ArchiveView.vue      # 归档视图（按年回溯与月份清单）
│           ├── AboutView.vue        # 关于视图（生活自白与档案馆初衷）
│           ├── NewEntryModal.vue    # 新增记录/记账弹窗
│           ├── EntryDetailModal.vue # 记录详情与编辑/删除弹窗
│           ├── AdminAuthModal.vue   # 暗房管理员钥匙配置弹窗
│           ├── StorageModal.vue     # 对象存储连接状态弹窗
│           ├── StatsModal.vue       # 生活刻度统计指标弹窗
│           ├── ImageUploader.vue    # R2 预签名媒体直传与多图上传器
│           ├── FloatingActionButton.vue # 右下角极简新增浮动按钮
│           └── FooterArchive.vue    # 备案合规页脚（浙ICP备2026081664号直链）
└── README.md
```

---

## 🚀 本地开发与生产部署

### 1. 本地启动开发

```bash
# 1. 复制配置文件 (可选)
cp .env.example .env

# 2. 启动 Go 后端服务 (默认 :8080)
go run main.go

# 3. 启动前端开发调试 (打开新终端)
cd web && npm run dev
```

### 2. 运行单元测试

```bash
go test -v ./...
```

### 3. Docker 容器化部署 (推荐)

```bash
# 构建并后台启动
docker compose up -d --build

# 查看运行日志
docker compose logs -f
```

### 4. 传统静态编译与主机部署

```bash
# 1. 前端生产打包
cd web && npm run build
cd ..

# 2. 编译 Linux AMD64 静态二进制
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o chimu-server-linux main.go

# 3. 推送至远程服务器并平滑重启服务
tar -czf - chimu-server-linux web/dist | ssh user@<your-server-ip> "tar -xzf - -C /home/ubuntu/chimu-lab && sudo systemctl restart chimu-lab"
```

---

## 📄 执照与版权

MIT License © 2026 [迟暮 (ChiMu)](https://github.com/YN1753)
