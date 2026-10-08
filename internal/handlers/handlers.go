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

// GetLifeEntries 获取统一生活时间线档案
func GetLifeEntries(c *gin.Context) {
	entryType := c.Query("type")
	year := c.Query("year")

	var entries []models.LifeEntry
	query := db.DB.Order("date DESC, id DESC")

	if entryType != "" && entryType != "all" {
		query = query.Where("type = ?", entryType)
	}
	if year != "" && year != "all" {
		query = query.Where("year = ?", year)
	}

	if err := query.Find(&entries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch life entries"})
		return
	}
	c.JSON(http.StatusOK, entries)
}

// GetNowStatus 获取「此刻的我」状态
func GetNowStatus(c *gin.Context) {
	var now models.NowStatus
	if err := db.DB.First(&now).Error; err != nil {
		c.JSON(http.StatusOK, models.NowStatus{
			Building:    "ChiMu Life Archive & Campus Map",
			Learning:    "Compiler internals, Film photography, Pour-over methods",
			Playing:     "Black Myth: Wukong, Zelda",
			Listening:   "Ryuichi Sakamoto - async",
			Reading:     "Zen and the Art of Motorcycle Maintenance",
			Thinking:    "How to live with quiet certainty and honest observation",
			Using:       "MacBook Pro 14, Leica Q2, Midori MD notebook",
			Location:    "Hangzhou, China (30.27° N, 120.15° E)",
			LastUpdated: "2026.10.09",
		})
		return
	}
	c.JSON(http.StatusOK, now)
}

// GetProjects 获取我做过的项目 (Things I made)
func GetProjects(c *gin.Context) {
	var projects []models.Project
	query := db.DB.Order("`order` ASC, id DESC")

	if err := query.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}
	c.JSON(http.StatusOK, projects)
}

// YearArchiveStat 年度生活统计元数据
type YearArchiveStat struct {
	Year        string         `json:"year"`
	TotalCount  int64          `json:"total_count"`
	TypeCounts  map[string]int `json:"type_counts"`
	Months      []string       `json:"months"`
}

// GetArchiveStats 获取年度归档概览
func GetArchiveStats(c *gin.Context) {
	var entries []models.LifeEntry
	db.DB.Find(&entries)

	yearMap := make(map[string]*YearArchiveStat)

	for _, e := range entries {
		y := e.Year
		if y == "" {
			y = "2026"
		}
		stat, exists := yearMap[y]
		if !exists {
			stat = &YearArchiveStat{
				Year:       y,
				TotalCount: 0,
				TypeCounts: make(map[string]int),
				Months:     []string{},
			}
			yearMap[y] = stat
		}
		stat.TotalCount++
		stat.TypeCounts[e.Type]++
	}

	var res []*YearArchiveStat
	for _, v := range yearMap {
		res = append(res, v)
	}

	c.JSON(http.StatusOK, res)
}

// GetConfig 获取网站与备案信息
func GetConfig(c *gin.Context) {
	var config models.SiteConfig
	if err := db.DB.First(&config).Error; err != nil {
		c.JSON(http.StatusOK, models.SiteConfig{
			SiteName:  "ChiMu Archive",
			Domain:    "codeactivityhub.top",
			ICPNumber: "浙ICP备2026081664号",
			ICPLink:   "https://beian.miit.gov.cn",
		})
		return
	}
	c.JSON(http.StatusOK, config)
}

// HealthCheck 健康状态
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "operational",
		"service": "ChiMu Life Archive Core",
		"runtime": runtime.Version(),
		"time":    time.Now().Unix(),
		"message": fmt.Sprintf("Archive engine running quietly on %s", runtime.GOARCH),
	})
}
