package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chimu-lab/internal/config"
	"chimu-lab/internal/db"
	"chimu-lab/internal/handlers"
	"chimu-lab/internal/middleware"
	"chimu-lab/internal/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置并初始化 R2 存储 (支持 .env 自动读取)
	cfg := config.LoadConfig()
	if err := storage.InitR2(cfg.R2); err != nil {
		log.Printf("Notice: R2 storage initialization: %v", err)
	} else if storage.IsConfigured() {
		log.Println("Cloudflare R2 storage initialized successfully.")
	} else {
		log.Println("Notice: Cloudflare R2 storage credentials not configured. Text entries & transactions will work normally.")
	}

	if cfg.AdminAPIKey != "" {
		log.Println("Admin write protection enabled (ADMIN_API_KEY configured).")
	} else {
		log.Println("Notice: ADMIN_API_KEY not set. Write endpoints are open (suitable for local dev). Set ADMIN_API_KEY in .env for production.")
	}

	// 2. 初始化 SQLite 数据库 (已启用 WAL 模式与连接池优化)
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
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Admin-Key"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(corsConfig))

	// 5. API v1 路由组 (标准结构化数据与持久化服务)
	v1 := r.Group("/api/v1")
	{
		// === 公开只读接口 ===
		v1.GET("/entries", handlers.GetLifeEntriesV1)
		v1.GET("/entries/:id", handlers.GetLifeEntryByID)
		v1.GET("/storage/status", handlers.GetStorageStatus)
		v1.GET("/transactions", handlers.GetTransactions)
		v1.GET("/transactions/summary", handlers.GetTransactionSummary)
		v1.GET("/projects", handlers.GetProjectsV1)
		v1.GET("/projects/:id", handlers.GetProjectByID)
		v1.GET("/stats", handlers.GetGlobalStats)

		// === 管理员写权限受保护接口 ===
		write := v1.Group("")
		write.Use(middleware.AdminAuthRequired())
		{
			// Life Entries (生活档案)
			write.POST("/entries", handlers.CreateLifeEntry)
			write.PUT("/entries/:id", handlers.UpdateLifeEntry)
			write.DELETE("/entries/:id", handlers.DeleteLifeEntry)

			// Uploads & Storage (R2 预签名直传与完成)
			write.POST("/uploads/presign", handlers.PresignUpload)
			write.POST("/uploads/complete", handlers.CompleteUpload)
			write.POST("/uploads/cleanup", handlers.CleanupUpload)

			// Transactions (记账独立体系)
			write.POST("/transactions", handlers.CreateTransaction)
			write.PUT("/transactions/:id", handlers.UpdateTransaction)
			write.DELETE("/transactions/:id", handlers.DeleteTransaction)

			// Projects (项目与造物)
			write.POST("/projects", handlers.CreateProject)
			write.PUT("/projects/:id", handlers.UpdateProject)
			write.DELETE("/projects/:id", handlers.DeleteProject)
		}
	}

	// 6. 兼容原有旧版只读路由
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
		r.StaticFile("/favicon.svg", "web/dist/favicon.svg")
		r.StaticFile("/icons.svg", "web/dist/icons.svg")
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
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// 启动 HTTP 服务协程
	go func() {
		log.Printf("ChiMu-Lab Server starting on :%s ...", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// 8. 优雅停机 (Graceful Shutdown)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down ChiMu-Lab server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("ChiMu-Lab server exited cleanly.")
}
