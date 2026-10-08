package handlers

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"chimu-lab/internal/db"
	"chimu-lab/internal/models"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

// GetProfile 获取个人资料及技能
func GetProfile(c *gin.Context) {
	var profile models.Profile
	if err := db.DB.First(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load profile"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

// GetProjects 获取实验室项目列表
func GetProjects(c *gin.Context) {
	category := c.Query("category")
	var projects []models.Project

	query := db.DB.Order("`order` ASC, id DESC")
	if category != "" && category != "all" {
		query = query.Where("category = ?", category)
	}

	if err := query.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}
	c.JSON(http.StatusOK, projects)
}

// GetActivities 获取代码动态和每日打卡记录
func GetActivities(c *gin.Context) {
	var activities []models.Activity
	if err := db.DB.Order("date DESC, id DESC").Limit(50).Find(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activities"})
		return
	}
	c.JSON(http.StatusOK, activities)
}

// StatsResponse 实验室统计数据与运行时 HUD
type StatsResponse struct {
	TotalProjects   int64   `json:"total_projects"`
	TotalActivities int64   `json:"total_activities"`
	ActiveDays      int     `json:"active_days"`
	UptimeHours     int     `json:"uptime_hours"`
	UptimeSeconds   int64   `json:"uptime_seconds"`
	LastUpdated     string  `json:"last_updated"`
	GoVersion       string  `json:"go_version"`
	Goroutines      int     `json:"goroutines"`
	MemoryAllocMB   float64 `json:"memory_alloc_mb"`
	DatabaseType    string  `json:"database_type"`
	QueryLatencyMs  float64 `json:"query_latency_ms"`
}

// GetStats 获取概览统计数据
func GetStats(c *gin.Context) {
	startQuery := time.Now()

	var projectCount int64
	var activityCount int64

	db.DB.Model(&models.Project{}).Count(&projectCount)
	db.DB.Model(&models.Activity{}).Count(&activityCount)

	queryLatency := float64(time.Since(startQuery).Microseconds()) / 1000.0

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memAllocMB := float64(m.Alloc) / 1024.0 / 1024.0

	uptimeSec := int64(time.Since(startTime).Seconds())
	uptimeHours := int(uptimeSec / 3600)
	if uptimeHours < 1 {
		uptimeHours = 1
	}

	res := StatsResponse{
		TotalProjects:   projectCount,
		TotalActivities: activityCount,
		ActiveDays:      218,
		UptimeHours:     uptimeHours,
		UptimeSeconds:   uptimeSec,
		LastUpdated:     time.Now().Format("2006-01-02 15:04:05"),
		GoVersion:       runtime.Version(),
		Goroutines:      runtime.NumGoroutine(),
		MemoryAllocMB:   float64(int(memAllocMB*100)) / 100.0,
		DatabaseType:    "Pure-Go SQLite (CGO-Free)",
		QueryLatencyMs:  float64(int(queryLatency*100)) / 100.0,
	}
	c.JSON(http.StatusOK, res)
}

// GetConfig 获取网站信息与备案信息
func GetConfig(c *gin.Context) {
	var config models.SiteConfig
	if err := db.DB.First(&config).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch site config"})
		return
	}
	c.JSON(http.StatusOK, config)
}

// HealthCheck 健康探测接口
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":     "operational",
		"service":    "ChiMu-Lab Core API",
		"version":    "2.0.0",
		"runtime":    runtime.Version(),
		"goroutines": runtime.NumGoroutine(),
		"time":       time.Now().Unix(),
		"message":    fmt.Sprintf("ChiMu-Lab Engine running smoothly on %s", runtime.GOARCH),
	})
}
