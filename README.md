# ChiMu-Lab (迟暮实验室) 🧪

> **Code Activity Hub · 迟暮的个人极客工坊与代码动态中心**
>
> 🌐 线上站点: [codeactivityhub.top](https://codeactivityhub.top)
> 🔗 个人主页 / GitHub: [https://github.com/YN1753](https://github.com/YN1753)

---

## 📖 项目简介

**ChiMu-Lab** 是一个面向全栈开发者的现代化个人实验工坊与作品聚合平台。立足工程美学与极客实践，本站专为技术作品展厅、日常打卡沉淀、代码动态追踪及 ICP / 公安合规备案量身设计。

### ✨ 核心特性

- **👨‍💻 极客个人名片 (Hero Profile)**：展示个人技术履历、技能栈徽章墙（Go, Vue 3, Docker, Linux, SQLite...）、社交链接与状态指示器。
- **🚀 实验室作品展厅 (Showcase)**：多分类筛选（核心平台、实验工坊、实用工具），包含高频面试打卡站（Daily Practice Hub）、分布式存储探测等丰富项目。
- **📊 代码动态与打卡看板 (Code Activity Hub)**：契合域名 `codeactivityhub.top`，展示 GitHub 风格年度活跃度热力图矩阵与近期工程里程碑脉络。
- **🛡️ 备案合规页脚 (Compliance Ready)**：内置工信部 ICP 备案号直达链接（`https://beian.miit.gov.cn`）与全国公安机关互联网站安全管理服务平台标号槽位，符合网安与工信部审核标准。
- **⚡ 现代化科技暗黑美学**：采用 Cyber-clean 深色科技质感、玻璃拟态模糊（Glassmorphism）、微交互悬浮动效与全端响应式适配。

---

## 🛠️ 技术架构

| 模块 | 技术选型 | 说明 |
| :--- | :--- | :--- |
| **后端语言** | **Go 1.27** | 高性能、低内存占用 |
| **Web 框架** | **Gin** (`github.com/gin-gonic/gin`) | 高性能 REST API 路由与中间件 |
| **数据库** | **SQLite** (`github.com/glebarez/sqlite`) | **纯 Go 实现 (CGO-Free)**，跨平台零编译依赖 |
| **ORM** | **GORM** (`gorm.io/gorm`) | 自动迁移模型（Profiles, Projects, Activities, SiteConfigs） |
| **前端框架** | **Vue 3** (Composition API + `<script setup>`) | 现代化响应式前端框架 |
| **构建工具** | **Vite** | 极速秒级 HMR 与打包构建 |
| **样式工程** | **Tailwind CSS v4** (`@tailwindcss/vite`) | 原子化 CSS 极客暗黑主题 |
| **图标库** | **Lucide Icons** (`lucide-vue-next` / `@lucide/vue`) | 统一的高质感极简矢量图标 |

---

## 📂 目录结构

```text
ChiMu-Lab/
├── main.go                     # Go 后端启动入口（集成 SPA 静态资源分发）
├── go.mod / go.sum             # Go 依赖配置
├── internal/
│   ├── models/                 # SQLite 数据库模型 (Profile, Project, Activity, SiteConfig)
│   ├── db/                     # 数据库连接、AutoMigrate 与自动种子填充 (Seed Data)
│   └── handlers/               # REST API 处理器 (/api/health, /api/projects, etc.)
├── web/                        # Vue 3 前端工程
│   ├── index.html              # 页面 HTML 模板与 Meta 标签
│   ├── vite.config.ts          # Vite 配置（Tailwind 插件、路径别名、开发代理）
│   ├── package.json            # 前端依赖配置
│   └── src/
│       ├── App.vue             # 页面主装配与数据驱动
│       ├── main.ts             # 前端入口
│       ├── style.css           # 全局样式与 Tailwind
│       ├── types/              # TypeScript 接口声明
│       └── components/         # 模块组件 (Navbar, Hero, ProjectShowcase, ActivityHub, ComplianceFooter)
└── README.md
```

---

## 🚀 本地快速启动

### 1. 启动后端 (Go)

```bash
# 启动 Go 服务 (默认运行在 :8080，支持 PORT=8090 指定端口)
go run main.go
```

首次运行会自动创建 `data/chimu.db` SQLite 数据库并预置丰富的初始数据。

### 2. 启动前端开发调试 (Vue 3)

```bash
cd web
npm install
npm run dev
```

浏览器打开 `http://localhost:5173` 即可实时预览并热更新。

---

## 📦 生产打包与部署

### 步骤 1：前端编译构建
```bash
cd web
npm run build
```
编译产物将输出至 `web/dist`。

### 步骤 2：后端二进制编译 (Linux CVM 部署)
```bash
# 针对 Linux 服务器进行跨平台编译（由于采用纯 Go SQLite，无任何 CGO 依赖）
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chimu-server main.go
```

### 步骤 3：部署至服务器并通过 Nginx 反向代理
在腾讯云服务器上：
```nginx
server {
    listen 80;
    listen 443 ssl http2;
    server_name codeactivityhub.top www.codeactivityhub.top;

    # SSL 证书配置
    ssl_certificate /etc/nginx/ssl/codeactivityhub.top.crt;
    ssl_certificate_key /etc/nginx/ssl/codeactivityhub.top.key;

    # 反向代理至 Go 服务（或直接指向 web/dist 静态目录）
    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    # 兼容原面试题练习站路径
    location /interview {
        alias /var/www/interview;
        index index.html;
        try_files $uri $uri/ /interview/index.html;
    }
}
```

---

## 📄 License

MIT © [迟暮 (ChiMu)](https://github.com/YN1753)
