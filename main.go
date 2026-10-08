package main

import (
	"log"
	"net/http"
	"os"

	"chimu-lab/internal/db"
	"chimu-lab/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 SQLite 数据库
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/chimu.db"
	}
	_, err := db.InitDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 生产模式或开发模式
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// CORS 跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	// API 路由组
	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
		api.GET("/profile", handlers.GetProfile)
		api.GET("/projects", handlers.GetProjects)
		api.GET("/activities", handlers.GetActivities)
		api.GET("/stats", handlers.GetStats)
		api.GET("/config", handlers.GetConfig)
		api.GET("/moments", handlers.GetMoments)
		api.POST("/moments/:id/like", handlers.LikeMoment)
	}

	// 前端静态文件托管（如果 web/dist 存在）
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("ChiMu-Lab Server starting on :%s ...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
