package models

import (
	"time"
)

// LifeEntry 统一生活时间线档案条目 (Life Stream)
// 支持 5 种核心生活形态：life(生活记录), thought(想法), photo(照片), project(项目), collection(收藏)
type LifeEntry struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Type        string       `gorm:"index" json:"type"` // life, thought, photo, project, collection (兼容原有子类型)
	Title       string       `json:"title"`
	Content     string       `gorm:"type:text" json:"content"`
	OccurredAt  time.Time    `gorm:"index" json:"occurred_at"`
	Location    string       `json:"location"`
	ProjectID   *uint        `gorm:"index" json:"project_id"`
	Project     *Project     `gorm:"foreignKey:ProjectID" json:"project,omitempty"`
	Attachments []Attachment `gorm:"foreignKey:EntryID;constraint:OnDelete:CASCADE" json:"attachments"`

	// 兼容原有旧字段，保证旧数据展示与格式平滑迁移
	Date           string `json:"date,omitempty"` // e.g. "2026.10.09"
	Year           string `json:"year,omitempty"` // e.g. "2026"
	Month          string `json:"month,omitempty"`
	Day            string `json:"day,omitempty"`
	Time           string `json:"time,omitempty"`
	Images         string `json:"images,omitempty"`
	Meta           string `json:"meta,omitempty"`
	Tags           string `json:"tags,omitempty"`
	Link           string `json:"link,omitempty"`
	RelatedProject string `json:"related_project,omitempty"`
	Featured       bool   `json:"featured,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Attachment 附件与多媒体对象关联表
type Attachment struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	EntryID          *uint     `gorm:"index" json:"entry_id"`
	ObjectKey        string    `gorm:"uniqueIndex" json:"object_key"`
	OriginalFilename string    `json:"original_filename"`
	MimeType         string    `json:"mime_type"`
	Size             int64     `json:"size"`
	Width            int       `json:"width,omitempty"`
	Height           int       `json:"height,omitempty"`
	URL              string    `gorm:"-" json:"url"` // 动态计算得出的公开或代理可访问 URL
	CreatedAt        time.Time `json:"created_at"`
}

// Project 长期进行的事情、创造与造物
type Project struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"` // 兼容 title
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle,omitempty"`
	Description string    `json:"description"`
	Content     string    `gorm:"type:text" json:"content"`
	Story       string    `gorm:"type:text" json:"story,omitempty"`
	Category    string    `json:"category,omitempty"`
	Tags        string    `json:"tags,omitempty"`
	DemoURL     string    `json:"demo_url,omitempty"`
	GithubURL   string    `json:"github_url,omitempty"`
	Status      string    `json:"status,omitempty"`
	Featured    bool      `json:"featured,omitempty"`
	Order       int       `json:"order,omitempty"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Transaction 独立记账记录 (单位：分，避免浮点数精度误差，如 1250 表示 12.50 元)
type Transaction struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Amount        int64     `gorm:"not null" json:"amount"` // 分
	Type          string    `gorm:"index;not null" json:"type"` // "income" 或 "expense"
	Category      string    `gorm:"index;not null" json:"category"` // 餐饮, 交通, 购物, 娱乐, 学习, 住房, 医疗, 其他
	Title         string    `json:"title"`
	Note          string    `json:"note"`
	PaymentMethod string    `json:"payment_method"`
	OccurredAt    time.Time `gorm:"index;not null" json:"occurred_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
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
