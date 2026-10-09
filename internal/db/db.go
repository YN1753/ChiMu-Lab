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
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	// 自动迁移表结构
	err = database.AutoMigrate(
		&models.LifeEntry{},
		&models.Attachment{},
		&models.Project{},
		&models.Transaction{},
		&models.NowStatus{},
		&models.SiteConfig{},
	)
	if err != nil {
		return nil, err
	}

	DB = database
	seedData(database)
	return database, nil
}

// seedData 填充生活档案与造物项目
func seedData(database *gorm.DB) {
	// 1. 初始化「此刻的我」状态 (Now)
	var nowCount int64
	database.Model(&models.NowStatus{}).Count(&nowCount)
	if nowCount == 0 {
		now := models.NowStatus{
			Building:    "生活自留地 · 真正记录生活的数字档案馆",
			Learning:    "编译器机制、胶片摄影、手冲烘焙萃取",
			Playing:     "黑神话：悟空、塞尔达传说、Steam 联机",
			Listening:   "坂本龙一 - 《async》/《BTTB》",
			Reading:     "罗伯特·波西格 - 《禅与摩托车维修艺术》",
			Thinking:    "如何在快节奏的喧嚣中保持诚恳、安静且有确定性的生活",
			Using:       "MacBook Pro 14, Contax T2, 纯白陶瓷滤杯, Midori MD 手账",
			Location:    "中国 · 杭州 (西湖区 / 余杭)",
			LastUpdated: "2026.10.09",
		}
		database.Create(&now)
	}

	// 2. 初始化项目（作为生活经历的一章：Things I Made）
	var projCount int64
	database.Model(&models.Project{}).Count(&projCount)
	if projCount == 0 {
		projects := []models.Project{
		{
			Title:       "SUSE-OAA-BACKEND",
			Subtitle:    "四川轻化工大学开放原子开源协会 · 业务后端",
			Description: "为高校开源协会搭建的高可用业务后端体系，支撑协会事务流转、招新管理与校园开源数据接口。",
			Story:       "大二时参与协会建设，发现校园组织的招新和事务常散落在群聊里。于是和同伴一起用 Go 搭建了这套业务核心。也是我第一次深入将分层架构与领域逻辑应用到实际生产场景中。",
			Category:    "Campus OpenSource",
			Tags:        "Go,Gin,Campus OpenSource",
			GithubURL:   "https://github.com/suse-edu-cn/SUSE-OAA-BACKEND",
			Status:      "Active",
			Featured:    true,
			Order:       1,
		},
		{
			Title:       "ArchCanvas",
			Subtitle:    "AI 辅助 Go 架构设计画布",
			Description: "让 AI 与开发者一起，从需求语义理解、ER 实体建模到架构设计，快速生成与构建可运行的 Go 工程骨架。",
			Story:       "在画系统架构图时常常觉得图与代码是割裂的。我想做一个直观的画板，把拖拽出来的实体和微服务边界实时映射为干净的 Go struct 代码。也是探索 AI 与工程设计工具融合的小实验。",
			Category:    "AI & Architecture",
			Tags:        "TypeScript,Go,Architecture",
			GithubURL:   "https://github.com/YN1753/ArchCanvas",
			Status:      "Active",
			Featured:    true,
			Order:       2,
		},
		{
			Title:       "GoLens",
			Subtitle:    "基于交互式状态机的 Go 底层机制透视镜",
			Description: "让 GMP 协程调度、三色标记 GC 屏障与 Channel 阻塞机制清晰可见的高交互度运行时可视化工具。",
			Story:       "学 Go 底层时被各种纸上谈兵的图解绕晕了。既然理解不了抽象，就手写一套状态机把它像电影放映机一样在浏览器里跑起来。做完之后，GMP 的工作偷取机制就像齿轮一样印在脑子里了。",
			Category:    "Visualization",
			Tags:        "JavaScript,Go Internals,GMP",
			GithubURL:   "https://github.com/YN1753/GoLens",
			Status:      "Stable",
			Featured:    true,
			Order:       3,
		},
		{
			Title:       "AstraLink-Desktop",
			Subtitle:    "基于 Wails 的图笔记桌面端应用",
			Description: "星链 2.0。探索 Go + 前端混合桌面开发（Wails 架构），支持双向链接、图谱可视化与本地隐私优先的知识管理。",
			Story:       "一直渴望一款数据 100% 留在本地硬盘、不依赖云端账户的个人笔记。用 Wails 把 WebKit 前端和 Go 核心绑在一起，启动轻快，内存只占几十兆。",
			Category:    "Desktop App",
			Tags:        "Go,Wails,Vue,Desktop",
			GithubURL:   "https://github.com/YN1753/AstraLink-Desktop",
			Status:      "Active",
			Featured:    false,
			Order:       4,
		},
		{
			Title:       "Go-Load",
			Subtitle:    "轻量分布式 HTTP 压测探针",
			Description: "自制轻量级高并发压测工具，毫秒级统计 P90 / P99 延迟直方图。",
			Story:       "觉得现成压测工具安装繁琐，就用 Go 原生协程池手打了一个发压器。小巧单文件，丢到服务器上就能测。",
			Category:    "CLI Tool",
			Tags:        "Go,Benchmark,CLI",
			GithubURL:   "https://github.com/YN1753/Go-Load",
			Status:      "Stable",
			Featured:    false,
			Order:       5,
		},
		{
			Title:       "nexus",
			Subtitle:    "轻量 Linux 服务器管理控制台",
			Description: "个人云服务器状态监控与容器服务守护面板。",
			Story:       "管理自己的轻量云服务器时写的工具，随时扫一眼内存、磁盘和守护进程。",
			Category:    "DevOps",
			Tags:        "Vue,Linux,DevOps",
			GithubURL:   "https://github.com/YN1753/nexus",
			Status:      "WIP",
			Featured:    false,
			Order:       6,
		},
		{
			Title:       "DeviceDaily",
			Subtitle:    "支持 Mac 小组件的设备成本统计 App",
			Description: "基于 Swift 原生开发，支持 macOS WidgetKit，记录电子产品日均使用成本。",
			Story:       "想看看手里买了三年的相机和电脑到底均摊到了多少钱一天。写了个简洁的 macOS 桌面 Widget。",
			Category:    "macOS App",
			Tags:        "Swift,macOS,WidgetKit",
			GithubURL:   "https://github.com/YN1753/DeviceDaily",
			Status:      "Active",
			Featured:    false,
			Order:       7,
		},
	}
	for _, p := range projects {
		database.Create(&p)
	}
	}

	// 3. 初始化生活时间线（Life Stream）：自然混合日常、想法、照片、游戏、音乐、地点、造物
	var entryCount int64
	database.Model(&models.LifeEntry{}).Count(&entryCount)
	if entryCount == 0 {
	entries := []models.LifeEntry{
		{
			Date:      "2026.10.09",
			Year:      "2026",
			Month:     "10",
			Day:       "09",
			Time:      "00:30",
			Type:      "thought",
			Title:     "做一个真正记录生活的网站",
			Content:   "深夜突然想到：我不想再把自己的网络主页做成一份展示给招聘者或甲方的技术简历了。生活不只有编译、部署和代码架构。下班后迎面吹来的夜风、清晨手冲咖啡的香气、拍下的胶卷、听了整夜的唱片、和朋友玩过的游戏……这些所有具体琐碎的事物，才构成了我真实活着的样子。这里应当是一本安静的个人生活档案馆。",
			Meta:      "深夜书房随笔",
			Tags:      "想法,生活志,初心",
			Featured:  true,
		},
		{
			Date:      "2026.10.08",
			Year:      "2026",
			Month:     "10",
			Day:       "08",
			Time:      "21:40",
			Type:      "photo",
			Title:     "满觉陇的桂花蒸与旁轴底片",
			Content:   "假期最后几天去了一趟满觉陇。沿山的村落都在做糖桂花，青石板路上落了一层薄薄的金黄，空气甜得发稠。下午四点半的夕阳斜斜穿过樟树叶，光斑在青石砖上摇曳。带了老胶片机拍完了一卷 Kodak 400。阳光晒在皮肤上的温热感，大概就是秋天最好的注脚。",
			Images:    "https://images.unsplash.com/photo-1507525428034-b723cf961d3e?auto=format&fit=crop&w=1200&q=80",
			Location:  "杭州 · 满觉陇",
			Meta:      "Contax T2 · 38mm f/2.8 · Kodak Portra 400",
			Tags:      "摄影,胶片,满觉陇,秋天",
			Featured:  true,
		},
		{
			Date:      "2026.10.08",
			Year:      "2026",
			Month:     "10",
			Day:       "08",
			Time:      "15:20",
			Type:      "coffee",
			Title:     "浅烘埃塞古吉：92°C 细水慢萃",
			Content:   "磨了 15g 埃塞俄比亚古吉日晒花魁豆。磨齿带出的干香满是草莓和茉莉花气息。92°C 纯净水，三段式注水萃取。第一段闷蒸 35 秒看咖啡粉像面包一样膨胀，白桃与柑橘酸感明亮。每天看着水柱在滤杯里画同心圆，是让人最沉静的日常仪式。",
			Images:    "https://images.unsplash.com/photo-1514432324607-a09d9b4aefdd?auto=format&fit=crop&w=1000&q=80",
			Location:  "书房窗台前",
			Meta:      "埃塞古吉花魁 · 15g粉 / 225g水 · 2分18秒",
			Tags:      "咖啡,手冲,慢生活",
			Featured:  false,
		},
		{
			Date:      "2026.10.07",
			Year:      "2026",
			Month:     "10",
			Day:       "07",
			Time:      "23:15",
			Type:      "music",
			Title:     "最近开始反复听坂本龙一的《async》",
			Content:   "深夜关掉大灯，只留一盏暖黄的书桌灯。戴上耳机放《async》。空旷而平静的钢琴琴键、风吹过树梢的采样，整个城市好像都睡熟了。在这样的底噪里看几页书、记几笔日记，心底有一种极其踏实的宁静。",
			Images:    "https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?auto=format&fit=crop&w=1000&q=80",
			Meta:      "Ryuichi Sakamoto - async (2017) · 推荐曲目：《andata》《solari》",
			Tags:      "音乐,坂本龙一,深夜",
			Featured:  true,
		},
		{
			Date:      "2026.10.06",
			Year:      "2026",
			Month:     "10",
			Day:       "06",
			Time:      "22:30",
			Type:      "game",
			Title:     "和朋友玩了一整晚游戏",
			Content:   "假期倒数第二天，叫上几个好友连麦打了一整晚游戏。从《黑神话：悟空》讨论到 Steam 合作通关，中间点了一大桶冰镇可乐和炸鸡。好久没有这样毫无目的地和朋友大笑、打打闹闹到深夜了。生活需要这样的留白和放空。",
			Images:    "https://images.unsplash.com/photo-1550745165-9bc0b252726f?auto=format&fit=crop&w=1000&q=80",
			Meta:      "Steam & PS5 连麦联机 · 快乐放空",
			Tags:      "游戏,朋友,放空",
			Featured:  false,
		},
		{
			Date:      "2026.10.03",
			Year:      "2026",
			Month:     "10",
			Day:       "03",
			Time:      "20:10",
			Type:      "place",
			Title:     "杨公堤夜骑：风从荷叶上吹过来的温度",
			Content:   "傍晚换了身衣服去西湖夜骑。从曲院风荷出发，沿着杨公堤一路穿行到南山路。两侧高大的水杉在路灯下投下斑驳长影，秋夜的风穿过枯荷拂在脸上，带着微凉的水汽。把白天的信息和杂念全抛在脑后，出了一身薄汗，整个人仿佛被晚风洗涤过一遍。",
			Images:    "https://images.unsplash.com/photo-1502680390469-be75c86b636f?auto=format&fit=crop&w=1200&q=80",
			Location:  "杭州 · 西湖杨公堤",
			Meta:      "公路车巡航 · 里程 21.4 km · 均速 22 km/h",
			Tags:      "骑行,西湖,夜风,行迹",
			Featured:  true,
		},
		{
			Date:      "2026.09.30",
			Year:      "2026",
			Month:     "09",
			Day:       "30",
			Time:      "18:00",
			Type:      "project",
			Title:     "校园开源协会业务后端 (SUSE-OAA-BACKEND)",
			Content:   "为四川轻化工大学开源协会重构了核心业务架构。抽离鉴权拦截器与事务层，让新一届协会学弟学妹在举办活动和技术招新时能有一个稳定干净的后端服务。造轮子最开心的瞬间，莫过于亲手写的逻辑真正跑在校园生活里。",
			Images:    "https://images.unsplash.com/photo-1517694712202-14dd9538aa97?auto=format&fit=crop&w=1000&q=80",
			Meta:      "Go / Gin / GORM / Campus OpenSource Backbone",
			Link:      "https://github.com/suse-edu-cn/SUSE-OAA-BACKEND",
			Tags:      "造物,Go,校园开源",
			Featured:  false,
		},
		{
			Date:      "2026.09.24",
			Year:      "2026",
			Month:     "09",
			Day:       "24",
			Time:      "16:30",
			Type:      "book",
			Title:     "重读《禅与摩托车维修艺术》：手艺与良质",
			Content:   "坐在长椅上吹着风读波西格。书中写：‘佛陀或耶稣坐在排气管边，就跟坐在莲花座上一样正常。’ 当你带着真正的良质（Quality）去对待手头的事物——无论是调校一辆自行车、手冲一杯咖啡，还是雕琢一段逻辑——人与工具就合一了。少一点向外张望的功利，多一点对手艺与生活的敬畏。",
			Images:    "https://images.unsplash.com/photo-1497633762265-9d179a990aa6?auto=format&fit=crop&w=1000&q=80",
			Meta:      "罗伯特·M·波西格 著 · 哲学与手艺书摘",
			Tags:      "阅读,书摘,良质,手艺",
			Featured:  true,
		},
		{
			Date:      "2026.09.15",
			Year:      "2026",
			Month:     "09",
			Day:       "15",
			Time:      "14:00",
			Type:      "purchase",
			Title:     "陪伴三年的旧物：无刻机械键盘、钢笔与手账",
			Content:   "擦拭这把用了三年的无刻机械键盘，键帽表面泛出了温润的光泽。旁边是一支百乐钢笔和一个牛皮纸手账。这些天天陪伴我的物件，沉默却忠诚。‘日用即道’，善待身边的每一件物，日子也会变得更踏实有温度。",
			Images:    "https://images.unsplash.com/photo-1587829741301-dc798b83add3?auto=format&fit=crop&w=1000&q=80",
			Meta:      "Custom Mechanical Keyboard + Pilot 78G + Midori MD",
			Tags:      "好物,文具,旧物,陪伴",
			Featured:  false,
		},
		{
			Date:      "2026.09.05",
			Year:      "2026",
			Month:     "09",
			Day:       "05",
			Time:      "21:00",
			Type:      "project",
			Title:     "ArchCanvas：AI 辅助 Go 架构设计画布",
			Content:   "想让 AI 辅助开发不只停留在单文件的补全上，而是从全局需求语义直观生成整套可运行的 Go 架构拓扑。开始构建 ArchCanvas 的 AST 双向同步引擎。",
			Meta:      "TypeScript / Go AST / Architecture Canvas",
			Link:      "https://github.com/YN1753/ArchCanvas",
			Tags:      "造物,AI架构,实验",
			Featured:  false,
		},
		{
			Date:      "2026.08.20",
			Year:      "2026",
			Month:     "08",
			Day:       "20",
			Time:      "19:30",
			Type:      "project",
			Title:     "GoLens：交互式状态机下的运行时透视镜",
			Content:   "把 Go 复杂的 GMP 协程调度、三色标记 GC 写屏障做成交互式状态机动画。亲手敲完一套调度流转，才算真正摸到了并发系统的齿轮。",
			Meta:      "Go Internals Visualization / GMP / GC Barrier",
			Link:      "https://github.com/YN1753/GoLens",
			Tags:      "造物,Go底层,状态机",
			Featured:  false,
		},
		{
			Date:      "2026.07.18",
			Year:      "2026",
			Month:     "07",
			Day:       "18",
			Time:      "05:15",
			Type:      "photo",
			Title:     "东极岛的日出与第一缕海风",
			Content:   "凌晨四点半爬起来走到东福山岛的最东端。海平线先是从深蓝泛出极淡的鹅黄，然后是一抹不可思议的粉紫。浪头撞在黑色的礁石上激起三米高的雪白浪花。坐在礁石上吹着咸湿的海风，天地浩大，人的那些微小烦恼瞬间显得微不足道。",
			Images:    "https://images.unsplash.com/photo-1506744038136-46273834b3fb?auto=format&fit=crop&w=1200&q=80",
			Location:  "舟山 · 东极岛东福山",
			Meta:      "Contax T2 · 38mm f/2.8 · Kodak Ektar 100",
			Tags:      "摄影,旅行,海边,日出",
			Featured:  true,
		},
		{
			Date:      "2026.06.21",
			Year:      "2026",
			Month:     "06",
			Day:       "21",
			Time:      "19:45",
			Type:      "thought",
			Title:     "夏至黄昏的暴雨与空街",
			Content:   "夏至这天傍晚，暴雨毫无预兆地倾盆而下。站在便利店屋檐下等雨停，看着水流顺着沥青路面汇入下水道，空气里满是泥土被雨水浇透后的清新气味。城市突然按下了静音键。很享受这样被意外困住的片刻，不用赶路，只要看雨就好。",
			Meta:      "雨天屋檐随记",
			Tags:      "随笔,夏至,雨天",
			Featured:  false,
		},
		{
			Date:      "2026.05.02",
			Year:      "2026",
			Month:     "05",
			Day:       "02",
			Time:      "14:30",
			Type:      "place",
			Title:     "莫干山竹林徒步：春末初夏的青翠",
			Content:   "五一假期避开人群，钻进莫干山后山的野竹林里徒步了五公里。阳光从密密匝匝的竹叶缝隙里漏下来，脚下的泥土松软带着竹叶香。山涧溪水冰凉刺骨，捧起来洗了把脸，神清气爽。",
			Images:    "https://images.unsplash.com/photo-1448375240586-882707db888b?auto=format&fit=crop&w=1200&q=80",
			Location:  "浙江 · 湖州莫干山",
			Meta:      "山野徒步 · 5.8 km 穿行",
			Tags:      "徒步,山野,自然,行迹",
			Featured:  false,
		},
		{
			Date:      "2026.04.12",
			Year:      "2026",
			Month:     "04",
			Day:       "12",
			Time:      "16:00",
			Type:      "coffee",
			Title:     "哥伦比亚双重厌氧瑰夏：水蜜桃与肉桂尾韵",
			Content:   "收到了新到的哥伦比亚蕙兰双重厌氧瑰夏豆。研磨时的香气像刚切开的多汁水蜜桃。90°C 水温萃取，前段是明艳的水果红茶感，温度降到微温后，肉桂与红糖的尾韵浮现出来。好的豆子确实能让人在舌尖上旅行。",
			Images:    "https://images.unsplash.com/photo-1495474472287-4d71bcdd2085?auto=format&fit=crop&w=1000&q=80",
			Location:  "书房",
			Meta:      "哥伦比亚双重厌氧瑰夏 · V60滤杯",
			Tags:      "咖啡,手冲,瑰夏",
			Featured:  false,
		},
		{
			Date:      "2026.03.20",
			Year:      "2026",
			Month:     "03",
			Day:       "20",
			Time:      "23:40",
			Type:      "music",
			Title:     "深夜重听 Bill Evans《Sunday at the Village Vanguard》",
			Content:   "春分这天夜里，重新翻出 Bill Evans 1961 年在先锋村的现场录音。背景里隐隐约约有酒杯轻碰的声音、低声的交谈，以及保罗·莫蒂安极其克制细腻的刷子鼓点。生活不需要时刻紧绷，爵士乐里那种即兴与松弛，是面对复杂世界的最好解药。",
			Images:    "https://images.unsplash.com/photo-1514525253161-7a46d19cd819?auto=format&fit=crop&w=1000&q=80",
			Meta:      "Bill Evans Trio - Sunday at the Village Vanguard (1961)",
			Tags:      "音乐,爵士乐,深夜,黑胶",
			Featured:  false,
		},
		{
			Date:      "2026.02.10",
			Year:      "2026",
			Month:     "02",
			Day:       "10",
			Time:      "01:20",
			Type:      "game",
			Title:     "春节假期通关《星际拓荒 Outer Wilds》",
			Content:   "在老家安静的深夜里打完了 Outer Wilds 的最后一幕。当所有乐器在宇宙尽头合奏起那首主题曲时，坐在屏幕前久久不能平静。它不仅是一款游戏，更是一首写给好奇心、科学探索与生命终极意义的伟大散文诗。",
			Images:    "https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?auto=format&fit=crop&w=1000&q=80",
			Meta:      "Mobius Digital · 22分钟循环的终点",
			Tags:      "游戏,OuterWilds,科幻,感动",
			Featured:  false,
		},
		{
			Date:      "2026.01.01",
			Year:      "2026",
			Month:     "01",
			Day:       "01",
			Time:      "10:00",
			Type:      "thought",
			Title:     "给二零二六年的自己写一封信",
			Content:   "新年第一天清晨，在手账本第一页写下：新的一年，不求做多宏大的事情，但求认认真真过好具体的一天天。多看几卷胶卷，多冲几杯好咖啡，多去山林和湖边吹风，少一些盲目的自我消耗。真实而安静地生活。",
			Meta:      "新年手账寄语",
			Tags:      "新年,随笔,初心,生活志",
			Featured:  true,
		},
	}
		for _, e := range entries {
			if e.OccurredAt.IsZero() && e.Date != "" {
				tStr := e.Time
				if tStr == "" {
					tStr = "12:00"
				}
				pTime, err := time.Parse("2006.01.02 15:04", e.Date+" "+tStr)
				if err == nil {
					e.OccurredAt = pTime
				}
			}
			database.Create(&e)
		}
	}

	// 4. 回填历史数据中缺失的 OccurredAt 字段
	var zeroEntries []models.LifeEntry
	database.Where("occurred_at = ? OR occurred_at IS NULL", time.Time{}).Find(&zeroEntries)
	for _, ze := range zeroEntries {
		if ze.Date != "" {
			tStr := ze.Time
			if tStr == "" {
				tStr = "12:00"
			}
			parsed, err := time.Parse("2006.01.02 15:04", ze.Date+" "+tStr)
			if err == nil {
				database.Model(&ze).Update("occurred_at", parsed)
			}
		}
	}

	// 5. 初始化基础记账记录 (Transactions)
	var txCount int64
	database.Model(&models.Transaction{}).Count(&txCount)
	if txCount == 0 {
		t1, _ := time.Parse("2006.01.02 15:04", "2026.10.09 14:20")
		t2, _ := time.Parse("2006.01.02 15:04", "2026.10.09 12:15")
		t3, _ := time.Parse("2006.01.02 15:04", "2026.10.08 16:30")
		t4, _ := time.Parse("2006.01.02 15:04", "2026.10.07 10:00")
		sampleTxs := []models.Transaction{
			{
				Amount:        12900, // 129.00
				Type:          "expense",
				Category:      "购物",
				Title:         "MCHOSE A7 无线鼠标",
				Note:          "替换用了四年的旧鼠标",
				PaymentMethod: "微信支付",
				OccurredAt:    t1,
			},
			{
				Amount:        1800, // 18.00
				Type:          "expense",
				Category:      "餐饮",
				Title:         "午饭面条",
				Note:          "公司楼下片儿川",
				PaymentMethod: "支付宝",
				OccurredAt:    t2,
			},
			{
				Amount:        6800, // 68.00
				Type:          "expense",
				Category:      "餐饮",
				Title:         "埃塞俄比亚咖啡生豆",
				Note:          "浅烘花魁 250g",
				PaymentMethod: "微信支付",
				OccurredAt:    t3,
			},
			{
				Amount:        500000, // 5000.00
				Type:          "income",
				Category:      "其他",
				Title:         "稿酬与项目结项",
				Note:          "校园开源平台二期补贴",
				PaymentMethod: "银行转账",
				OccurredAt:    t4,
			},
		}
		for _, tx := range sampleTxs {
			database.Create(&tx)
		}
	}

	// 6. 站点合规配置
	database.Model(&models.SiteConfig{}).Count(&nowCount)
	if nowCount == 0 {
		config := models.SiteConfig{
			SiteName:  "ChiMu · 迟暮的数字生活档案馆",
			SiteDesc:  "A personal archive of things I've done, seen, thought about, and don't want to forget.",
			Domain:    "codeactivityhub.top",
			ICPNumber: "浙ICP备2026081664号",
			ICPLink:   "https://beian.miit.gov.cn",
		}
		database.Create(&config)
	}

	log.Println("Database initialized with life stream entries, craft projects, and transactions.")
}
