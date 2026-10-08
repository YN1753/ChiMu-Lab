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

// seedData 填充默认种子数据（保证初次运行时内容完整丰富，备案审核合规）
func seedData(database *gorm.DB) {
	var count int64
	database.Model(&models.Profile{}).Count(&count)
	if count == 0 {
		profile := models.Profile{
			Name:     "迟暮",
			Title:    "全栈研发工程师 / Gopher / 独立开发者",
			Bio:      "立足工程美学与极客实践。热衷于 Go 高性能后端、现代前端与自动化技术探索，在数字世界构建优雅而坚固的实验室工具。",
			Avatar:   "https://avatars.githubusercontent.com/u/108920146?v=4",
			Github:   "https://github.com/YN1753",
			Email:    "chimu@codeactivityhub.top",
			Location: "中国 · 杭州",
			Skills:   "Go,Vue 3,TypeScript,Docker,Linux,SQLite,MySQL,Gin,Tailwind CSS,Git,Redis",
		}
		database.Create(&profile)
	}

	database.Model(&models.Project{}).Count(&count)
	if count == 0 {
		projects := []models.Project{
			{
				Title:       "面试练习站 · Daily Practice Hub",
				Subtitle:    "算法每日打卡与高频面试考点演练平台",
				Description: "围绕 ACM 模式、数组贪心算法、高频场景题构建的沉浸式每日刷题平台。支持闭卷练习、代码思路验证与历史归档回溯。",
				Category:    "core",
				Tags:        "Go,Algorithms,ACM,Vue,Tailwind",
				DemoURL:     "/interview",
				GithubURL:   "https://github.com/YN1753",
				Status:      "Active",
				Featured:    true,
				Order:       1,
			},
			{
				Title:       "ChiMu-Lab 迟暮实验室",
				Subtitle:    "极简主义全栈工坊与代码动态聚合站",
				Description: "本站核心工程。基于 Go + Vue 3 + SQLite 构建的轻量级开发实验室，整合项目展厅、工程动态追踪与合规备案中心。",
				Category:    "lab",
				Tags:        "Go,Gin,Vue3,TypeScript,SQLite,Tailwind",
				DemoURL:     "https://codeactivityhub.top",
				GithubURL:   "https://github.com/YN1753/ChiMu-Lab",
				Status:      "Active",
				Featured:    true,
				Order:       2,
			},
			{
				Title:       "GopherSpace 分布式存储探测工具",
				Subtitle:    "高性能并发网络扫描与节点状态同步守护进程",
				Description: "轻量级网络探测器，专为探测节点连通性、实时延迟与健康状态设计，具备极低资源占用与高吞吐并发能力。",
				Category:    "tool",
				Tags:        "Go,Goroutine,Network,Linux,Syscall",
				DemoURL:     "",
				GithubURL:   "https://github.com/YN1753",
				Status:      "Stable",
				Featured:    true,
				Order:       3,
			},
			{
				Title:       "Hermes 自动化工作流引擎",
				Subtitle:    "每日任务调度、内容聚合与静态站点自动化发布管线",
				Description: "定时拉取技术动态与刷题计划，自动化编译生成静态练习页并无缝推送到生产 Web 服务。",
				Category:    "lab",
				Tags:        "Automation,Python,Shell,CI/CD,Linux",
				DemoURL:     "",
				GithubURL:   "https://github.com/YN1753",
				Status:      "Active",
				Featured:    false,
				Order:       4,
			},
		}
		for _, p := range projects {
			database.Create(&p)
		}
	}

	database.Model(&models.Activity{}).Count(&count)
	if count == 0 {
		now := time.Now()
		activities := []models.Activity{
			{
				Date:        now.Format("2006-01-02"),
				Type:        "release",
				Title:       "迟暮实验室 ChiMu-Lab 1.0 正式上线",
				Description: "完成全栈架构搭建，基于 Go + Vue 3 + SQLite 实现单二进制极简部署与现代化响应式 UI。",
				RepoName:    "YN1753/ChiMu-Lab",
				Count:       12,
				Link:        "https://github.com/YN1753/ChiMu-Lab",
			},
			{
				Date:        now.AddDate(0, 0, -2).Format("2006-01-02"),
				Type:        "commit",
				Title:       "重构面试每日题静态排版与移动端自适应",
				Description: "优化 ACM 模式题目展示效果，增加平滑滚动与语法高亮支持。",
				RepoName:    "YN1753/interview-hub",
				Count:       6,
				Link:        "https://github.com/YN1753",
			},
			{
				Date:        now.AddDate(0, 0, -5).Format("2006-01-02"),
				Type:        "study",
				Title:       "贪心与动态规划状态转移专题复习",
				Description: "深度演练相邻约束类贪心原型题目，输出白板题解与复杂度推导笔记。",
				RepoName:    "YN1753/algo-notes",
				Count:       4,
				Link:        "https://github.com/YN1753",
			},
			{
				Date:        now.AddDate(0, 0, -9).Format("2006-01-02"),
				Type:        "milestone",
				Title:       "启用独立主域名 codeactivityhub.top",
				Description: "完成云服务器 DNS 解析、全站 HTTPS/TLS 证书签发与安全加固配置。",
				RepoName:    "YN1753/infra",
				Count:       8,
				Link:        "https://codeactivityhub.top",
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
			SiteDesc:     "Code Activity Hub · 迟暮的个人极客工坊与代码动态中心",
			Domain:       "codeactivityhub.top",
			ICPNumber:    "苏ICP备2024000000号-1", // 用户可在配置或环境变量中覆盖为真实备案号
			ICPLink:      "https://beian.miit.gov.cn",
			PoliceNumber: "公网安备 待审核",
			PoliceCode:   "",
		}
		database.Create(&config)
	}

	log.Println("Database initialized and seeded successfully.")
}
