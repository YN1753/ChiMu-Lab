package main

import (
	"log"
	"net/http"
	"os"

	"chimu-lab/internal/config"
	"chimu-lab/internal/db"
	"chimu-lab/internal/handlers"
	"chimu-lab/internal/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置并初始化 R2 存储
	cfg := config.LoadConfig()
	if err := storage.InitR2(cfg.R2); err != nil {
		log.Printf("Notice: R2 storage initialization: %v", err)
	} else if storage.IsConfigured() {
		log.Println("Cloudflare R2 storage initialized successfully.")
	} else {
		log.Println("Notice: Cloudflare R2 storage credentials not configured. Text entries & transactions will work normally.")
	}

	// 2. 初始化 SQLite 数据库
	_, err := db.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 3. 生产模式或开发模式
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// 4. CORS 跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(corsConfig))

	// 5. API v1 路由组 (标准结构化数据与持久化服务)
	v1 := r.Group("/api/v1")
	{
		// Life Entries (生活档案)
		v1.GET("/entries", handlers.GetLifeEntriesV1)
		v1.POST("/entries", handlers.CreateLifeEntry)
		v1.GET("/entries/:id", handlers.GetLifeEntryByID)
		v1.PUT("/entries/:id", handlers.UpdateLifeEntry)
		v1.DELETE("/entries/:id", handlers.DeleteLifeEntry)

		// Uploads & Storage (R2 预签名直传与完成)
		v1.POST("/uploads/presign", handlers.PresignUpload)
		v1.POST("/uploads/complete", handlers.CompleteUpload)
		v1.POST("/uploads/cleanup", handlers.CleanupUpload)
		v1.GET("/storage/status", handlers.GetStorageStatus)

		// Transactions (记账独立体系)
		v1.GET("/transactions", handlers.GetTransactions)
		v1.POST("/transactions", handlers.CreateTransaction)
		v1.GET("/transactions/summary", handlers.GetTransactionSummary)
		v1.PUT("/transactions/:id", handlers.UpdateTransaction)
		v1.DELETE("/transactions/:id", handlers.DeleteTransaction)

		// Projects (项目与造物)
		v1.GET("/projects", handlers.GetProjectsV1)
		v1.POST("/projects", handlers.CreateProject)
		v1.GET("/projects/:id", handlers.GetProjectByID)
		v1.PUT("/projects/:id", handlers.UpdateProject)
		v1.DELETE("/projects/:id", handlers.DeleteProject)

		// Stats (生活统计)
		v1.GET("/stats", handlers.GetGlobalStats)
	}

	// 6. 兼容原有旧版路由
	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/entries", handlers.GetLifeEntries)
		api.GET("/now", handlers.GetNowStatus)
		api.GET("/projects", handlers.GetProjects)
		api.GET("/archive", handlers.GetArchiveStats)
		api.GET("/config", handlers.GetConfig)
	}

	// 7. 前端静态文件托管（如果 web/dist 存在）
	if stat, err := os.Stat("web/dist"); err == nil && stat.IsDir() {
		r.Static("/assets", "web/dist/assets")
		r.StaticFile("/favicon.ico", "web/dist/favicon.ico")
		r.StaticFile("/robots.txt", "web/dist/robots.txt")

		// SPA 单页应用回退到 index.html
		r.NoRoute(func(c *gin.Context) {
			c.File("web/dist/index.html")
		})
		log.Println("Serving frontend from web/dist")
	} else {
		r.NoRoute(func(c *gin.Context) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "ChiMu-Lab API service is running. Web frontend dist not found, build it via 'cd web && npm run build'.",
			})
		})
	}

	port := cfg.Port
	log.Printf("ChiMu-Lab Server starting on :%s ...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
