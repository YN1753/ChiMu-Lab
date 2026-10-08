package models

import (
	"time"
)

// LifeEntry 统一生活时间线档案条目 (Life Stream)
// 记录生活的所有方方面面：日常、照片、碎片化记录、想法、项目、游戏、音乐、书、旅行/地点、买过的东西...
type LifeEntry struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Date           string    `json:"date"` // e.g. "2026.10.09"
	Year           string    `json:"year"` // e.g. "2026"
	Month          string    `json:"month"` // e.g. "10"
	Day            string    `json:"day"` // e.g. "09"
	Time           string    `json:"time"` // e.g. "23:45"
	Type           string    `json:"type"` // thought, photo, moment, place, music, book, game, project, purchase, gear
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Images         string    `json:"images"` // 图片 URL 列表（逗号分隔或单个大图）
	Location       string    `json:"location"`
	Meta           string    `json:"meta"` // 专辑名 / 书籍作者 / 游戏平台 / 相机镜头 / 购买品类
	Tags           string    `json:"tags"`
	Link           string    `json:"link"`
	RelatedProject string    `json:"related_project"`
	Featured       bool      `json:"featured"`
	CreatedAt      time.Time `json:"created_at"`
}

// NowStatus 「此刻的我」当前状态
type NowStatus struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Building    string    `json:"building"`
	Learning    string    `json:"learning"`
	Playing     string    `json:"playing"`
	Listening   string    `json:"listening"`
	Reading     string    `json:"reading"`
	Thinking    string    `json:"thinking"`
	Using       string    `json:"using"`
	Location    string    `json:"location"`
	LastUpdated string    `json:"last_updated"`
}

// Project 造物与手艺 (Things I Made - 人生经历的一章)
type Project struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	Story       string    `json:"story"` // 为什么开始、怎么做、遇到了什么
	Category    string    `json:"category"`
	Tags        string    `json:"tags"`
	DemoURL     string    `json:"demo_url"`
	GithubURL   string    `json:"github_url"`
	Status      string    `json:"status"`
	Featured    bool      `json:"featured"`
	Order       int       `json:"order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SiteConfig 站点与备案合规配置
type SiteConfig struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	SiteName     string    `json:"site_name"`
	SiteDesc     string    `json:"site_desc"`
	Domain       string    `json:"domain"`
	ICPNumber    string    `json:"icp_number"`    // 工信部 ICP 备案号
	ICPLink      string    `json:"icp_link"`      // https://beian.miit.gov.cn
	PoliceNumber string    `json:"police_number"` // 公安联网备案号
	PoliceCode   string    `json:"police_code"`   // 公安备案编号
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
