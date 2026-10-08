package db

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"chimu-lab/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化 SQLite 数据库
func InitDB(dbPath string) (*gorm.DB, error) {
	if dbPath == "" {
		dbPath = "data/chimu.db"
	}

	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// 自动迁移表结构
	err = database.AutoMigrate(
		&models.Profile{},
		&models.Project{},
		&models.Activity{},
		&models.SiteConfig{},
	)
	if err != nil {
		return nil, err
	}

	DB = database
	seedData(database)
	return database, nil
}

// seedData 填充迟暮真实的个人项目与仓库数据
func seedData(database *gorm.DB) {
	var count int64
	database.Model(&models.Profile{}).Count(&count)
	if count == 0 {
		profile := models.Profile{
			Name:     "迟暮",
			Title:    "Gopher / 全栈开发者",
			Bio:      "迟暮的个人数字工坊。不迎合外界，只记录自己造过的轮子、踩过的坑与真实的工程探索。",
			Avatar:   "https://avatars.githubusercontent.com/u/108920146?v=4",
			Github:   "https://github.com/YN1753",
			Email:    "chimu@codeactivityhub.top",
			Location: "中国 · 杭州",
			Skills:   "Go,Wails,Vue 3,TypeScript,Docker,Linux,Swift,SQLite,Gin,Tailwind CSS",
		}
		database.Create(&profile)
	}

	// 每次更新时重新同步真实项目
	database.Exec("DELETE FROM projects")
	projects := []models.Project{
		{
			Title:       "SUSE-OAA-BACKEND",
			Subtitle:    "四川轻化工大学开放原子开源协会 · 核心业务后端",
			Description: "为四川轻化工大学开源协会研发的 Go 后端服务体系，支撑协会事务协作、招新管理与数据接口。",
			Category:    "core",
			Tags:        "Go,Gin,suse-edu-cn,Campus OpenSource",
			DemoURL:     "",
			GithubURL:   "https://github.com/suse-edu-cn/SUSE-OAA-BACKEND",
			Status:      "Active",
			Featured:    true,
			Order:       1,
		},
		{
			Title:       "ArchCanvas",
			Subtitle:    "AI 辅助 Go 架构设计画布",
			Description: "让 AI 与开发者一起，从需求语义理解、ER 实体建模到架构设计，快速生成与构建可运行的 Go 工程骨架。",
			Category:    "core",
			Tags:        "TypeScript,Go,Architecture,AI Canvas",
			DemoURL:     "",
			GithubURL:   "https://github.com/YN1753/ArchCanvas",
			Status:      "Active",
			Featured:    true,
			Order:       2,
		},
		{
			Title:       "GoLens",
			Subtitle:    "基于交互式状态机的 Go 底层机制透视镜",
			Description: "让 GMP 协程调度、三色标记 GC 屏障与 Channel 阻塞机制清晰可见的高交互度运行时可视化工具。",
			Category:    "tool",
			Tags:        "JavaScript,Go Internals,GMP,Visualization",
			DemoURL:     "",
			GithubURL:   "https://github.com/YN1753/GoLens",
			Status:      "Stable",
			Featured:    true,
			Order:       3,
		},
		{
			Title:       "AstraLink-Desktop",
			Subtitle:    "基于 Wails 的图笔记桌面端应用",
			Description: "星链 2.0。探索 Go + 前端混合桌面开发（Wails 架构），支持双向链接、图谱可视化与本地隐私优先的知识管理。",
			Category:    "tool",
			Tags:        "Go,Wails,Vue,Desktop,Graph",
			DemoURL:     "",
			GithubURL:   "https://github.com/YN1753/AstraLink-Desktop",
			Status:      "Active",
			Featured:    false,
			Order:       4,
		},
		{
			Title:       "Go-Load",
			Subtitle:    "轻量级高并发 HTTP 压测工具",
			Description: "基于 Go 语言原生并发模型编写的高性能压测工具，轻巧无外部依赖，具备低资源开销与精确的时延吞吐量统计。",
			Category:    "tool",
			Tags:        "Go,Benchmark,High Concurrency,CLI",
			DemoURL:     "",
			GithubURL:   "https://github.com/YN1753/Go-Load",
			Status:      "Stable",
			Featured:    false,
			Order:       5,
		},
		{
			Title:       "nexus",
			Subtitle:    "面向开发者的现代化 Linux 服务器管理平台",
			Description: "轻量化 Linux 运维控制台，用于打理自己的云服务器，监控基础硬件性能、容器与服务状态。",
			Category:    "tool",
			Tags:        "Vue,Linux,DevOps,System",
			DemoURL:     "",
			GithubURL:   "https://github.com/YN1753/nexus",
			Status:      "WIP",
			Featured:    false,
			Order:       6,
		},
		{
			Title:       "DeviceDaily",
			Subtitle:    "支持 Mac 原生小组件的设备成本统计 App",
			Description: "用 Swift 原生开发的实用记账与设备折旧折算工具，支持 macOS 原生 Widget 小组件常驻桌面。",
			Category:    "tool",
			Tags:        "Swift,macOS,WidgetKit,Utility",
			DemoURL:     "",
			GithubURL:   "https://github.com/YN1753/DeviceDaily",
			Status:      "Active",
			Featured:    false,
			Order:       7,
		},
	}
	for _, p := range projects {
		database.Create(&p)
	}

	database.Model(&models.Activity{}).Count(&count)
	if count == 0 {
		now := time.Now()
		activities := []models.Activity{
			{
				Date:        now.Format("2006-01-02"),
				Type:        "commit",
				Title:       "SUSE-OAA-BACKEND 架构优化与接口梳理",
				Description: "重构业务逻辑层与中间件鉴权，提升校园协会服务响应性能。",
				RepoName:    "suse-edu-cn/SUSE-OAA-BACKEND",
				Count:       8,
				Link:        "https://github.com/suse-edu-cn/SUSE-OAA-BACKEND",
			},
			{
				Date:        now.AddDate(0, 0, -1).Format("2006-01-02"),
				Type:        "commit",
				Title:       "ArchCanvas AI 画布核心状态机与实体生成测试",
				Description: "打通从需求文本解析到 Go struct 及 GORM 实体定义代码生成链条。",
				RepoName:    "YN1753/ArchCanvas",
				Count:       12,
				Link:        "https://github.com/YN1753/ArchCanvas",
			},
			{
				Date:        now.AddDate(0, 0, -4).Format("2006-01-02"),
				Type:        "release",
				Title:       "GoLens 运行时可视化透视镜初版调试",
				Description: "完成 GMP 调度状态转移与 Channel 缓冲队列的可视化逻辑。",
				RepoName:    "YN1753/GoLens",
				Count:       5,
				Link:        "https://github.com/YN1753/GoLens",
			},
			{
				Date:        now.AddDate(0, 0, -8).Format("2006-01-02"),
				Type:        "milestone",
				Title:       "AstraLink-Desktop Wails 桌面端跨平台打包",
				Description: "验证 macOS / Windows 双平台打包产物与 SQLite 本地数据存储。",
				RepoName:    "YN1753/AstraLink-Desktop",
				Count:       6,
				Link:        "https://github.com/YN1753/AstraLink-Desktop",
			},
		}
		for _, a := range activities {
			database.Create(&a)
		}
	}

	database.Model(&models.SiteConfig{}).Count(&count)
	if count == 0 {
		config := models.SiteConfig{
			SiteName:     "迟暮实验室 · ChiMu-Lab",
			SiteDesc:     "迟暮的个人数字工坊与工程工作台",
			Domain:       "codeactivityhub.top",
			ICPNumber:    "浙ICP备2026081664号",
			ICPLink:      "https://beian.miit.gov.cn",
			PoliceNumber: "公网安备 待审核",
			PoliceCode:   "",
		}
		database.Create(&config)
	}

	log.Println("Database initialized and real repos seeded successfully.")
}
