package models

import (
	"time"
)

// Profile 个人资料
type Profile struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Bio       string    `json:"bio"`
	Avatar    string    `json:"avatar"`
	Github    string    `json:"github"`
	Email     string    `json:"email"`
	Location  string    `json:"location"`
	Skills    string    `json:"skills"` // 逗号分隔或 JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Project 实验室/作品项目
type Project struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	Category    string    `json:"category"` // lab, open-source, study, tool
	Tags        string    `json:"tags"`     // 逗号分隔的标签: Go, Vue3, Docker
	DemoURL     string    `json:"demo_url"`
	GithubURL   string    `json:"github_url"`
	Status      string    `json:"status"` // Active, Stable, WIP, Archived
	Featured    bool      `json:"featured"`
	Order       int       `json:"order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Activity 代码动态与日常打卡
type Activity struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Date        string    `json:"date"` // YYYY-MM-DD
	Type        string    `json:"type"` // commit, release, study, milestone
	Title       string    `json:"title"`
	Description string    `json:"description"`
	RepoName    string    `json:"repo_name"`
	Count       int       `json:"count"` // 提交/活跃次数
	Link        string    `json:"link"`
	CreatedAt   time.Time `json:"created_at"`
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
