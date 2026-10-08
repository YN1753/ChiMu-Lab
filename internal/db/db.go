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

	// 真实生活方方面面记录 (Life Chronicles & Facets)
	database.Exec("DELETE FROM life_moments")
	moments := []models.LifeMoment{
		{
			Date:     "2026.10.08",
			Time:     "21:30",
			Location: "杭州 · 西湖杨公堤",
			Weather:  "20°C · 暮秋凉风",
			Mood:     "放空",
			Category: "cycling",
			Title:    "杨公堤夜骑：风从荷叶上吹过来的温度",
			Content:  "下班后没有直接回书房，换了身衣服去西湖夜骑。从曲院风荷出发，沿着杨公堤一直骑到南山路。路两侧是高大的水杉，秋天的晚风穿过荷叶扑在脸上，带着微凉的水汽。把城市的喇叭声和屏幕的蓝光全抛在身后，出了一身薄汗，整个人好像被这阵风重新洗刷了一遍。",
			Quote:    "“骑车的时候，世界退居到两侧，心跳是唯一的节拍器。”",
			Note:     "路线：曲院风荷 -> 杨公堤 -> 虎跑路 -> 钱塘江绿道，总里程 21.4 km。速度 22km/h。",
			ImageURL: "https://images.unsplash.com/photo-1502680390469-be75c86b636f?auto=format&fit=crop&w=800&q=80",
			MetaInfo: "公路车巡航 · 21.4km · 爬升 85m",
			Tags:     "夜骑,西湖,生活行迹,松弛",
			Likes:    34,
		},
		{
			Date:     "2026.10.05",
			Time:     "16:40",
			Location: "杭州 · 满觉陇青石板路",
			Weather:  "23°C · 金桂初放",
			Mood:     "拾光",
			Category: "photo",
			Title:    "满觉陇的桂花蒸与旁轴底片",
			Content:  "赶在假期的尾巴，带了老胶片机去满觉陇。沿山的村落都在做糖桂花，青石板路上落了一层薄薄的金黄，空气甜得发稠。下午四点半的夕阳斜斜穿过樟树叶，光斑在墙面上摇曳。按下快门的那一瞬，时间好像被装进了小小的暗盒里。",
			Quote:    "“光线是时间的影印件，而胶卷留下了温度。”",
			Note:     "Contax T2 · 38mm f/2.8 Carl Zeiss，Kodak Portra 400，光圈 f/4，曝光补偿 +0.3EV。",
			ImageURL: "https://images.unsplash.com/photo-1507525428034-b723cf961d3e?auto=format&fit=crop&w=800&q=80",
			MetaInfo: "Contax T2 · Kodak Portra 400 · 38mm",
			Tags:     "胶片,满觉陇,桂花,摄影",
			Likes:    42,
		},
		{
			Date:     "2026.10.03",
			Time:     "10:15",
			Location: "书房窗前木桌",
			Weather:  "22°C · 晨光微熹",
			Mood:     "惬意",
			Category: "coffee",
			Title:    "浅烘埃塞古吉与 92°C 细水慢萃",
			Content:  "清晨的书房很安静。磨了 15g 埃塞俄比亚古吉产区的花魁日晒豆，磨齿带出的干香满是草莓和茉莉花气息。92°C 纯净水，三段式注水萃取。第一段闷蒸 35 秒看咖啡粉像面包一样膨胀，看着水柱在滤杯里画同心圆，是每天最让人沉静下来的仪式。",
			Quote:    "“把水注入咖啡粉的过程，就像把耐心倾注给生活。”",
			Note:     "水粉比 1:15，粉量 15g，注水总量 225g，总萃取耗时 2分18秒。白桃与柑橘酸感明亮。",
			ImageURL: "https://images.unsplash.com/photo-1514432324607-a09d9b4aefdd?auto=format&fit=crop&w=800&q=80",
			MetaInfo: "埃塞俄比亚 古吉花魁 · 92°C · V60 滤杯",
			Tags:     "手冲,咖啡,日常仪式,慢生活",
			Likes:    29,
		},
		{
			Date:     "2026.09.30",
			Time:     "14:20",
			Location: "西湖边 · 树荫下长椅",
			Weather:  "24°C · 微风拂面",
			Mood:     "沉思",
			Category: "reading",
			Title:    "重读《禅与摩托车维修艺术》：手艺与良质",
			Content:  "坐在长椅上吹着湖风读波西格。书中写：‘佛陀或耶稣坐在排气管边，就跟坐在莲花座上一样正常。’ 当你带着真正的良质（Quality）去面对一件具体的事物——无论是调校一辆自行车的刹车皮、手冲一杯咖啡，抑或是雕琢一行 Go 代码——工具和人就合二为一了。少一点向外张望的功利，多一点对手艺本身的敬畏。",
			Quote:    "“如果你对事情感到厌倦，说明你已经失去了与它的活生生的联结。”",
			Note:     "罗伯特·M·波西格 著，重庆出版社。随手在 P.168 折了角，记下了关于专注的感悟。",
			ImageURL: "https://images.unsplash.com/photo-1497633762265-9d179a990aa6?auto=format&fit=crop&w=800&q=80",
			MetaInfo: "《禅与摩托车维修艺术》· 案头书摘",
			Tags:     "读书,手艺,良质,哲学",
			Likes:    36,
		},
		{
			Date:     "2026.09.25",
			Time:     "23:00",
			Location: "深夜工位 · 暖光台灯下",
			Weather:  "19°C · 秋夜微凉",
			Mood:     "专注",
			Category: "music",
			Title:    "黑胶唱片与坂本龙一的音符",
			Content:  "深夜十一点，把房间大灯关掉，只留一盏暖黄的台灯。戴上耳机放坂本龙一的《async》。空旷而平静的钢琴音、风吹过树梢的采样，整个世界好像都睡着了。在这样的底噪里写两行字，看看书，心里非常踏实。音乐是精神的庇护所。",
			Quote:    "“生命是脆弱的，但音乐能把那一瞬的真实凝固成永恒。”",
			Note:     "推荐循环曲目：《andata》与《solari》。静谧、克制、富有呼吸感。",
			ImageURL: "https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?auto=format&fit=crop&w=800&q=80",
			MetaInfo: "坂本龙一 · 《async》· 24bit/96kHz 无损",
			Tags:     "音乐,坂本龙一,深夜,治愈",
			Likes:    45,
		},
		{
			Date:     "2026.09.18",
			Time:     "22:30",
			Location: "杭州 · 窗台前",
			Weather:  "21°C · 夜雨敲窗",
			Mood:     "自洽",
			Category: "thought",
			Title:    "写给自己：在喧嚣时代认认真真生活",
			Content:  "窗外下着绵绵的秋雨。很多人把生活过成了一场给别人看的展览，急着证明自己掌握了什么技术、走到了什么高度。可日子终究是自己过的：饭要一口一口吃，觉要踏踏实实睡。代码也好，爱好也罢，都只是人生长河里的浪花。认认真真生活，诚恳对待每一顿饭、每一趟骑行、每一个念头，足矣。",
			Quote:    "“不向外界讨要意义，生活的意义就在生活的每一个具体细节里。”",
			Note:     "深夜随笔。泡了一杯温热的陈皮老白茶，听雨声入睡。",
			ImageURL: "https://images.unsplash.com/photo-1517694712202-14dd9538aa97?auto=format&fit=crop&w=800&q=80",
			MetaInfo: "雨夜随想 · 杭州生活手记",
			Tags:     "随想,自洽,内心平宁,夜雨",
			Likes:    58,
		},
		{
			Date:     "2026.09.10",
			Time:     "15:00",
			Location: "书房桌面",
			Weather:  "26°C · 晴朗",
			Mood:     "爱物",
			Category: "gear",
			Title:    "陪伴三年的旧物：机械键盘、钢笔与手账",
			Content:  "清理桌面时擦拭这把用了三年的无刻机械键盘。键帽表面已经泛出了温润的光泽，青轴的手感依旧干脆。旁边是一支用了很久的百乐钢笔和一个牛皮纸手账。这些天天陪伴我的物件，沉默却忠诚。人与器物之间的相处，时间久了也会生出情谊。",
			Quote:    "“日用即道。善待陪伴你的每一件工具。”",
			Note:     "桌面好物：定制无刻键盘、百乐 78G 钢笔、Midori MD 方格本。简约耐看。",
			ImageURL: "https://images.unsplash.com/photo-1587829741301-dc798b83add3?auto=format&fit=crop&w=800&q=80",
			MetaInfo: "EDC 好物 · 桌面日常 · 物与心",
			Tags:     "好物,桌面,文具,陪伴",
			Likes:    27,
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
