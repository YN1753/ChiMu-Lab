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
		&models.LifeMoment{},
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

	// 真实生活切片与日常记录 (Life & Vignettes)
	database.Exec("DELETE FROM life_moments")
	moments := []models.LifeMoment{
		{
			Date:     "2026.10.08",
			Time:     "23:40",
			Location: "杭州 · 余杭工位",
			Weather:  "19°C · 秋夜微雨",
			Mood:     "专注",
			Category: "thought",
			Title:    "雨夜、通道阻塞与造轮子的意义",
			Content:  "窗外下着绵密的秋雨。有人曾问我，为什么开源社区那么多现成的调度器和压测工具，还要花时间一行行写 GoLens 和 Go-Load？亲手排查过一次 GMP 的偷取队列、体会过 Channel 阻塞唤醒时的原子语义，那种心里的笃定是调现成库给不了的。代码不是虚张声势的展品，是一块块自己铺上去的青石板。",
			Note:     "手冲了一杯耶加雪菲。雨水敲打百叶窗的声音很适合写代码。",
			ImageURL: "https://images.unsplash.com/photo-1517694712202-14dd9538aa97?auto=format&fit=crop&w=800&q=80",
			Camera:   "Leica Q2 · 28mm f/1.7 · ISO 400",
			Tags:     "Go,GMP,思考,夜雨",
			Likes:    14,
		},
		{
			Date:     "2026.10.05",
			Time:     "17:20",
			Location: "杭州 · 满觉陇",
			Weather:  "23°C · 晴朗薄暮",
			Mood:     "拾光",
			Category: "film",
			Title:    "桂花蒸与胶卷里的秋天",
			Content:  "趁着国庆假期最后两天，骑车去了一趟满觉陇。满山的金桂已经落了一地，空气里全是甜香。随身带了旁轴机拍完了一卷 Kodak 400。阳光透过香樟树叶洒在青石板上，突然觉得不管是写系统架构还是过生活，留白和呼吸感才是最珍贵的底色。",
			Note:     "Kodak Portra 400 曝光补偿 +0.3EV，光斑很柔和。",
			ImageURL: "https://images.unsplash.com/photo-1507525428034-b723cf961d3e?auto=format&fit=crop&w=800&q=80",
			Camera:   "Contax T2 · 38mm f/2.8 · Kodak Portra 400",
			Tags:     "胶片,满觉陇,秋日,骑行",
			Likes:    28,
		},
		{
			Date:     "2026.10.01",
			Time:     "15:00",
			Location: "西湖边 · 树荫下长椅",
			Weather:  "25°C · 微风",
			Mood:     "松弛",
			Category: "reading",
			Title:    "《禅与摩托车维修艺术》与工程手艺",
			Content:  "重读波西格的《禅与摩托车维修艺术》。书中写道：‘佛陀或耶稣坐在排气管边，就跟坐在莲花座上一样正常。’ 当你带着真正的良质（Quality）去调试一段并发竞态，或者去调配一杯手冲水粉比时，工具和人就已经融为一体了。少一点功利的目标，多一点对手艺的敬畏。",
			Note:     "P.142 标注：‘如果你对事情感到厌倦，说明你已经失去了与它的活生生的联结。’",
			ImageURL: "https://images.unsplash.com/photo-1497633762265-9d179a990aa6?auto=format&fit=crop&w=800&q=80",
			Camera:   "Ricoh GR IIIx · 40mm f/2.8",
			Tags:     "阅读,手艺,良质,哲学",
			Likes:    19,
		},
		{
			Date:     "2026.09.28",
			Time:     "22:15",
			Location: "杭州 · 迟暮工坊",
			Weather:  "21°C · 清凉夜风",
			Mood:     "沉思",
			Category: "coffee",
			Title:    "曼特宁手冲笔记与 AST 遍历灵感",
			Content:  "深烘苏门答腊曼特宁，研磨度偏粗，水温 90°C。前段坚果与黑巧的风味很厚重。在等第二段注水浸润的 40 秒里，突然想通了 ArchCanvas 解析 Go struct tag 时的嵌套循环问题：直接在语义树上做局部剪枝，比正则匹配合适得多。生活中的沉淀，往往会在不经意间反哺工程。",
			Note:     "咖啡粉 16g，水粉比 1:14，两段注水，总萃取用时 2分15秒。",
			ImageURL: "https://images.unsplash.com/photo-1501339847302-ac426a4a7cbb?auto=format&fit=crop&w=800&q=80",
			Camera:   "Fujifilm X100V · Classic Chrome",
			Tags:     "手冲,咖啡,ArchCanvas,思考",
			Likes:    22,
		},
		{
			Date:     "2026.09.20",
			Time:     "01:30",
			Location: "深夜桌面",
			Weather:  "18°C · 静谧夜色",
			Mood:     "沉浸",
			Category: "music",
			Title:    "坂本龙一的音符与静默的 Goroutine",
			Content:  "戴上耳机放坂本龙一的《async》。整座城市好像都睡熟了，只有屏幕上微微发光的代码。把 suseoaa 的一个死锁隐患顺藤摸瓜找了出来——是一个 defer unlock 在条件分支里意外漏掉。修掉的那一刻，世界好像都变宽阔了。这种深夜的纯粹，千金不换。",
			Note:     "推荐曲目：《solari》与《andata》。空旷而平静。",
			ImageURL: "https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?auto=format&fit=crop&w=800&q=80",
			Camera:   "Sony A7C · 35mm f/1.4 GM",
			Tags:     "音乐,坂本龙一,深夜编码,suseoaa",
			Likes:    31,
		},
	}
	for _, m := range moments {
		database.Create(&m)
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

	log.Println("Database initialized and real repos and moments seeded successfully.")
}
