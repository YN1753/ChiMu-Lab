package handlers

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"chimu-lab/internal/db"
	"chimu-lab/internal/models"
	"chimu-lab/internal/storage"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

// 统一标准响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Code:    status,
		Message: message,
		Data:    nil,
	})
}

// ParseFlexTime 弹性解析前端传入的各种时间格式 (含 HTML5 datetime-local 标准格式 YYYY-MM-DDTHH:mm)
func ParseFlexTime(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Now(), nil
	}
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02T15:04",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"2006.01.02 15:04",
		"2006.01.02",
	}
	for _, l := range layouts {
		if t, err := time.ParseInLocation(l, s, time.Local); err == nil {
			return t, nil
		}
	}
	return time.Now(), fmt.Errorf("unable to parse time string: %s", s)
}

// formatDatesFromTime 根据 OccurredAt 自动格式化旧版兼容字段
func formatDatesFromTime(e *models.LifeEntry) {
	if e.OccurredAt.IsZero() {
		e.OccurredAt = time.Now()
	}
	e.Year = fmt.Sprintf("%04d", e.OccurredAt.Year())
	e.Month = fmt.Sprintf("%02d", e.OccurredAt.Month())
	e.Day = fmt.Sprintf("%02d", e.OccurredAt.Day())
	e.Time = fmt.Sprintf("%02d:%02d", e.OccurredAt.Hour(), e.OccurredAt.Minute())
	e.Date = fmt.Sprintf("%s.%s.%s", e.Year, e.Month, e.Day)
}

// populateAttachmentURLs 填充附件的公开或代理访问 URL
func populateAttachmentURLs(atts []models.Attachment) []models.Attachment {
	for i := range atts {
		atts[i].URL = storage.GetPublicURL(atts[i].ObjectKey)
	}
	return atts
}

// ==========================================
// 1. Life Entry CRUD (生活档案接口)
// ==========================================

// CreateEntryRequest 新增生活记录请求载荷
type CreateEntryRequest struct {
	Type          string    `json:"type" binding:"required"` // life, thought, photo, project, collection
	Title         string    `json:"title"`
	Content       string    `json:"content" binding:"required"`
	OccurredAt    string    `json:"occurred_at"` // ISO8601 或 "2006-01-02T15:04" 或 "2006-01-02 15:04" 或 "2006.01.02"
	Location      string    `json:"location"`
	ProjectID     *uint     `json:"project_id"`
	AttachmentIDs []uint    `json:"attachment_ids"`
}

// CreateLifeEntry 创建一条新的生活档案记录
func CreateLifeEntry(c *gin.Context) {
	var req CreateEntryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "内容必填，且必须指定有效类型")
		return
	}

	occurred, _ := ParseFlexTime(req.OccurredAt)

	entry := models.LifeEntry{
		Type:       req.Type,
		Title:      strings.TrimSpace(req.Title),
		Content:    strings.TrimSpace(req.Content),
		OccurredAt: occurred,
		Location:   strings.TrimSpace(req.Location),
		ProjectID:  req.ProjectID,
	}
	formatDatesFromTime(&entry)

	// 保存主体
	if err := db.DB.Create(&entry).Error; err != nil {
		Error(c, http.StatusInternalServerError, "保存记录失败: "+err.Error())
		return
	}

	// 关联已上传的附件
	if len(req.AttachmentIDs) > 0 {
		db.DB.Model(&models.Attachment{}).
			Where("id IN ?", req.AttachmentIDs).
			Update("entry_id", entry.ID)

		// 同步到旧版 images 字段以保证老组件兼容
		var atts []models.Attachment
		db.DB.Where("id IN ?", req.AttachmentIDs).Find(&atts)
		if len(atts) > 0 {
			var urls []string
			for _, a := range atts {
				urls = append(urls, storage.GetPublicURL(a.ObjectKey))
			}
			entry.Images = strings.Join(urls, ",")
			db.DB.Model(&entry).Update("images", entry.Images)
		}
	}

	// 重新预加载关联
	db.DB.Preload("Attachments").Preload("Project").First(&entry, entry.ID)
	entry.Attachments = populateAttachmentURLs(entry.Attachments)

	Success(c, entry)
}

// GetLifeEntriesV1 列表查询生活档案 (支持分页与类型筛选)
func GetLifeEntriesV1(c *gin.Context) {
	entryType := c.Query("type")
	year := c.Query("year")
	category := c.Query("category") // 兼容 5 大核心视角 (daily, thought, project, collection, all)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "50"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 50
	}

	query := db.DB.Model(&models.LifeEntry{}).
		Preload("Attachments").
		Preload("Project").
		Order("occurred_at DESC, id DESC")

	if entryType != "" && entryType != "all" {
		query = query.Where("type = ?", entryType)
	}
	if year != "" && year != "all" {
		query = query.Where("year = ?", year)
	}

	// 核心视角动态映射
	if category != "" && category != "all" {
		switch category {
		case "daily":
			query = query.Where("type IN ('daily', 'life', 'coffee', 'place', 'moment')")
		case "thought":
			query = query.Where("type IN ('thought', 'idea', 'note', 'essay')")
		case "project":
			query = query.Where("type IN ('project', 'code', 'craft')")
		case "collection":
			query = query.Where("type IN ('collection', 'music', 'book', 'game', 'purchase', 'gear', 'movie')")
		case "photo":
			query = query.Where("type = 'photo' OR images != ''")
		case "transaction":
			query = query.Where("type = 'transaction'")
		}
	}

	var total int64
	query.Count(&total)

	var entries []models.LifeEntry
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&entries).Error; err != nil {
		Error(c, http.StatusInternalServerError, "获取生活档案失败")
		return
	}

	for i := range entries {
		entries[i].Attachments = populateAttachmentURLs(entries[i].Attachments)
		// 如果 images 为空但有附件，反向补齐 images 供首屏渲染
		if entries[i].Images == "" && len(entries[i].Attachments) > 0 {
			entries[i].Images = entries[i].Attachments[0].URL
		}
	}

	Success(c, gin.H{
		"total":     total,
		"page":      page,
		"page_size": pageSize,
		"entries":   entries,
	})
}

// GetLifeEntryByID 获取单条生活档案详情
func GetLifeEntryByID(c *gin.Context) {
	id := c.Param("id")
	var entry models.LifeEntry
	if err := db.DB.Preload("Attachments").Preload("Project").First(&entry, id).Error; err != nil {
		Error(c, http.StatusNotFound, "未找到该条生活记录")
		return
	}

	entry.Attachments = populateAttachmentURLs(entry.Attachments)
	if entry.Images == "" && len(entry.Attachments) > 0 {
		entry.Images = entry.Attachments[0].URL
	}

	Success(c, entry)
}

// UpdateEntryRequest 更新生活记录请求载荷（字段皆可选）
type UpdateEntryRequest struct {
	Type          string `json:"type"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	OccurredAt    string `json:"occurred_at"`
	Location      string `json:"location"`
	ProjectID     *uint  `json:"project_id"`
	AttachmentIDs []uint `json:"attachment_ids"`
}

// UpdateLifeEntry 更新生活记录
func UpdateLifeEntry(c *gin.Context) {
	id := c.Param("id")
	var entry models.LifeEntry
	if err := db.DB.Preload("Attachments").First(&entry, id).Error; err != nil {
		Error(c, http.StatusNotFound, "未找到该条生活记录")
		return
	}

	var req UpdateEntryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "请求参数无效")
		return
	}

	if req.Type != "" {
		entry.Type = req.Type
	}
	if req.Title != "" {
		entry.Title = strings.TrimSpace(req.Title)
	}
	if req.Content != "" {
		entry.Content = strings.TrimSpace(req.Content)
	}
	entry.Location = strings.TrimSpace(req.Location)
	entry.ProjectID = req.ProjectID

	if req.OccurredAt != "" {
		if t, err := ParseFlexTime(req.OccurredAt); err == nil {
			entry.OccurredAt = t
		}
		formatDatesFromTime(&entry)
	}

	if err := db.DB.Save(&entry).Error; err != nil {
		Error(c, http.StatusInternalServerError, "更新记录失败")
		return
	}

	// 更新关联附件
	if req.AttachmentIDs != nil {
		// 解除未包含的附件
		db.DB.Model(&models.Attachment{}).
			Where("entry_id = ? AND id NOT IN ?", entry.ID, req.AttachmentIDs).
			Update("entry_id", nil)

		// 绑定新选中的附件
		if len(req.AttachmentIDs) > 0 {
			db.DB.Model(&models.Attachment{}).
				Where("id IN ?", req.AttachmentIDs).
				Update("entry_id", entry.ID)
		}
	}

	db.DB.Preload("Attachments").Preload("Project").First(&entry, entry.ID)
	entry.Attachments = populateAttachmentURLs(entry.Attachments)

	Success(c, entry)
}

// DeleteLifeEntry 删除生活记录及关联的附件与 R2 对象
func DeleteLifeEntry(c *gin.Context) {
	id := c.Param("id")
	var entry models.LifeEntry
	if err := db.DB.Preload("Attachments").First(&entry, id).Error; err != nil {
		Error(c, http.StatusNotFound, "记录不存在或已被删除")
		return
	}

	// 级联清理 R2 存储中的对象文件
	for _, att := range entry.Attachments {
		if att.ObjectKey != "" {
			_ = storage.DeleteObject(att.ObjectKey)
		}
	}

	// 删除数据库记录 (GORM 会自动外键级联删除 attachments)
	if err := db.DB.Delete(&entry).Error; err != nil {
		Error(c, http.StatusInternalServerError, "删除失败: "+err.Error())
		return
	}

	Success(c, gin.H{"id": id, "deleted": true})
}

// ==========================================
// 2. Uploads & Storage (R2 媒体上传接口)
// ==========================================

type PresignRequest struct {
	Filename    string `json:"filename" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
	Size        int64  `json:"size" binding:"required"`
}

// PresignUpload 生成 Presigned PUT URL
func PresignUpload(c *gin.Context) {
	var req PresignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "缺少必要的文件信息 (filename, content_type, size)")
		return
	}

	// 大小限制：单文件最大 20MB
	maxSize := int64(20 * 1024 * 1024)
	if req.Size > maxSize {
		Error(c, http.StatusBadRequest, "文件过大，单张图片不能超过 20MB")
		return
	}

	// MIME 类型基本检查
	ct := strings.ToLower(req.ContentType)
	if !strings.HasPrefix(ct, "image/") && !strings.HasPrefix(ct, "video/") && ct != "application/pdf" {
		Error(c, http.StatusBadRequest, "目前仅支持上传图片或媒体文件")
		return
	}

	uploadID, objectKey, presignedURL, expiresIn, err := storage.GeneratePresignedPutURL(req.Filename, req.ContentType, req.Size)
	if err != nil {
		Error(c, http.StatusServiceUnavailable, err.Error())
		return
	}

	Success(c, gin.H{
		"upload_id":   uploadID,
		"object_key":  objectKey,
		"upload_url":  presignedURL,
		"expires_in":  expiresIn,
	})
}

type CompleteUploadRequest struct {
	UploadID         string `json:"upload_id"`
	ObjectKey        string `json:"object_key" binding:"required"`
	OriginalFilename string `json:"original_filename"`
	MimeType         string `json:"mime_type"`
	Size             int64  `json:"size"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
}

// CompleteUpload 确认上传完成并创建 Attachment 实体
func CompleteUpload(c *gin.Context) {
	var req CompleteUploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "缺少 object_key")
		return
	}

	// 如果 R2 已配置，验证对象确实存在于存储桶中
	if storage.IsConfigured() {
		realSize, realContentType, err := storage.VerifyObjectExists(req.ObjectKey)
		if err != nil {
			Error(c, http.StatusBadRequest, "存储端未检测到上传的文件，请重试")
			return
		}
		if realSize > 0 {
			req.Size = realSize
		}
		if realContentType != "" {
			req.MimeType = realContentType
		}
	}

	att := models.Attachment{
		ObjectKey:        req.ObjectKey,
		OriginalFilename: req.OriginalFilename,
		MimeType:         req.MimeType,
		Size:             req.Size,
		Width:            req.Width,
		Height:           req.Height,
	}

	if err := db.DB.Create(&att).Error; err != nil {
		Error(c, http.StatusInternalServerError, "创建附件记录失败")
		return
	}

	att.URL = storage.GetPublicURL(att.ObjectKey)
	Success(c, att)
}

type CleanupUploadRequest struct {
	ObjectKeys []string `json:"object_keys" binding:"required"`
}

// CleanupUpload 清理未保存引用的孤立上传对象
func CleanupUpload(c *gin.Context) {
	var req CleanupUploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "缺少 object_keys")
		return
	}

	cleaned := 0
	for _, key := range req.ObjectKeys {
		if key != "" {
			var att models.Attachment
			// 仅允许清理在数据库中且未关联生活条目的孤立附件，防止越权传入他人 ObjectKey 随意删除
			if err := db.DB.Where("object_key = ? AND entry_id IS NULL", key).First(&att).Error; err == nil {
				_ = storage.DeleteObject(key)
				db.DB.Delete(&att)
				cleaned++
			}
		}
	}

	Success(c, gin.H{"cleaned_count": cleaned})
}

// GetStorageStatus 获取 R2 配置状态
func GetStorageStatus(c *gin.Context) {
	Success(c, storage.GetStorageStatus())
}

// ==========================================
// 3. Transactions CRUD (记账接口)
// ==========================================

type TransactionRequest struct {
	Amount        int64  `json:"amount" binding:"required"` // 分
	Type          string `json:"type" binding:"required"`   // income, expense
	Category      string `json:"category" binding:"required"`
	Title         string `json:"title"`
	Note          string `json:"note"`
	PaymentMethod string `json:"payment_method"`
	OccurredAt    string `json:"occurred_at"`
}

// CreateTransaction 记一笔
func CreateTransaction(c *gin.Context) {
	var req TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "请填写金额、类型(支出/收入)与分类")
		return
	}

	occurred, _ := ParseFlexTime(req.OccurredAt)

	tx := models.Transaction{
		Amount:        req.Amount,
		Type:          req.Type,
		Category:      strings.TrimSpace(req.Category),
		Title:         strings.TrimSpace(req.Title),
		Note:          strings.TrimSpace(req.Note),
		PaymentMethod: strings.TrimSpace(req.PaymentMethod),
		OccurredAt:    occurred,
	}

	if err := db.DB.Create(&tx).Error; err != nil {
		Error(c, http.StatusInternalServerError, "记账失败: "+err.Error())
		return
	}

	Success(c, tx)
}

// GetTransactions 列表查询账目 (支持按年、按月)
func GetTransactions(c *gin.Context) {
	yearStr := c.Query("year")
	monthStr := c.Query("month")

	query := db.DB.Order("occurred_at DESC, id DESC")

	if yearStr != "" {
		y, err := strconv.Atoi(yearStr)
		if err == nil {
			if monthStr != "" {
				m, err := strconv.Atoi(monthStr)
				if err == nil {
					start := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local)
					end := start.AddDate(0, 1, 0)
					query = query.Where("occurred_at >= ? AND occurred_at < ?", start, end)
				}
			} else {
				start := time.Date(y, 1, 1, 0, 0, 0, 0, time.Local)
				end := start.AddDate(1, 0, 0)
				query = query.Where("occurred_at >= ? AND occurred_at < ?", start, end)
			}
		}
	}

	var txs []models.Transaction
	if err := query.Find(&txs).Error; err != nil {
		Error(c, http.StatusInternalServerError, "获取记账列表失败")
		return
	}

	Success(c, txs)
}

// GetTransactionSummary 月度记账统计
func GetTransactionSummary(c *gin.Context) {
	yearStr := c.DefaultQuery("year", fmt.Sprintf("%04d", time.Now().Year()))
	monthStr := c.DefaultQuery("month", fmt.Sprintf("%02d", time.Now().Month()))

	y, _ := strconv.Atoi(yearStr)
	m, _ := strconv.Atoi(monthStr)

	start := time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, 0)

	var txs []models.Transaction
	db.DB.Where("occurred_at >= ? AND occurred_at < ?", start, end).Find(&txs)

	var totalExpense int64
	var totalIncome int64
	categoryBreakdown := make(map[string]int64)

	for _, tx := range txs {
		if tx.Type == "expense" {
			totalExpense += tx.Amount
			categoryBreakdown[tx.Category] += tx.Amount
		} else if tx.Type == "income" {
			totalIncome += tx.Amount
		}
	}

	Success(c, gin.H{
		"year":               y,
		"month":              m,
		"total_expense":      totalExpense,
		"total_income":       totalIncome,
		"balance":            totalIncome - totalExpense,
		"category_breakdown": categoryBreakdown,
		"count":              len(txs),
	})
}

// UpdateTransaction 更新账目
func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var tx models.Transaction
	if err := db.DB.First(&tx, id).Error; err != nil {
		Error(c, http.StatusNotFound, "账目记录不存在")
		return
	}

	var req TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "请求参数无效")
		return
	}

	tx.Amount = req.Amount
	tx.Type = req.Type
	tx.Category = strings.TrimSpace(req.Category)
	tx.Title = strings.TrimSpace(req.Title)
	tx.Note = strings.TrimSpace(req.Note)
	tx.PaymentMethod = strings.TrimSpace(req.PaymentMethod)

	if req.OccurredAt != "" {
		if t, err := ParseFlexTime(req.OccurredAt); err == nil {
			tx.OccurredAt = t
		}
	}

	if err := db.DB.Save(&tx).Error; err != nil {
		Error(c, http.StatusInternalServerError, "更新账目失败")
		return
	}

	Success(c, tx)
}

// DeleteTransaction 删除账目
func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Transaction{}, id).Error; err != nil {
		Error(c, http.StatusInternalServerError, "删除失败")
		return
	}
	Success(c, gin.H{"id": id, "deleted": true})
}

// ==========================================
// 4. Projects CRUD (项目与造物接口)
// ==========================================

type ProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	Tags        string `json:"tags"`
	GithubURL   string `json:"github_url"`
	Status      string `json:"status"`
}

func CreateProject(c *gin.Context) {
	var req ProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "项目名称必填")
		return
	}

	p := models.Project{
		Name:        strings.TrimSpace(req.Name),
		Title:       strings.TrimSpace(req.Name),
		Description: strings.TrimSpace(req.Description),
		Content:     strings.TrimSpace(req.Content),
		Story:       strings.TrimSpace(req.Content),
		Category:    strings.TrimSpace(req.Category),
		Tags:        strings.TrimSpace(req.Tags),
		GithubURL:   strings.TrimSpace(req.GithubURL),
		Status:      strings.TrimSpace(req.Status),
	}
	if p.Status == "" {
		p.Status = "Active"
	}

	if err := db.DB.Create(&p).Error; err != nil {
		Error(c, http.StatusInternalServerError, "创建项目失败")
		return
	}

	Success(c, p)
}

func GetProjectsV1(c *gin.Context) {
	var projects []models.Project
	db.DB.Order("`order` ASC, id DESC").Find(&projects)
	Success(c, projects)
}

func GetProjectByID(c *gin.Context) {
	id := c.Param("id")
	var p models.Project
	if err := db.DB.First(&p, id).Error; err != nil {
		Error(c, http.StatusNotFound, "项目不存在")
		return
	}

	// 同时拉取关联该项目的 Life Entries
	var entries []models.LifeEntry
	db.DB.Where("project_id = ?", p.ID).Preload("Attachments").Order("occurred_at DESC").Find(&entries)
	for i := range entries {
		entries[i].Attachments = populateAttachmentURLs(entries[i].Attachments)
	}

	Success(c, gin.H{
		"project": p,
		"entries": entries,
	})
}

func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var p models.Project
	if err := db.DB.First(&p, id).Error; err != nil {
		Error(c, http.StatusNotFound, "项目不存在")
		return
	}

	var req ProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "请求参数无效")
		return
	}

	p.Name = strings.TrimSpace(req.Name)
	p.Title = strings.TrimSpace(req.Name)
	p.Description = strings.TrimSpace(req.Description)
	p.Content = strings.TrimSpace(req.Content)
	p.Story = strings.TrimSpace(req.Content)
	p.Category = strings.TrimSpace(req.Category)
	p.Tags = strings.TrimSpace(req.Tags)
	p.GithubURL = strings.TrimSpace(req.GithubURL)
	p.Status = strings.TrimSpace(req.Status)

	db.DB.Save(&p)
	Success(c, p)
}

func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	db.DB.Delete(&models.Project{}, id)
	Success(c, gin.H{"id": id, "deleted": true})
}

// ==========================================
// 5. Global Stats (生活档案统计概览)
// ==========================================

func GetGlobalStats(c *gin.Context) {
	var totalEntries int64
	db.DB.Model(&models.LifeEntry{}).Count(&totalEntries)

	var totalPhotos int64
	db.DB.Model(&models.Attachment{}).Count(&totalPhotos)
	if totalPhotos == 0 {
		db.DB.Model(&models.LifeEntry{}).Where("type = 'photo' OR images != ''").Count(&totalPhotos)
	}

	var totalProjects int64
	db.DB.Model(&models.Project{}).Count(&totalProjects)

	// 当月支出
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, 0)

	var monthExpense int64
	type Result struct {
		Total int64
	}
	var res Result
	db.DB.Model(&models.Transaction{}).
		Select("COALESCE(SUM(amount), 0) as total").
		Where("type = 'expense' AND occurred_at >= ? AND occurred_at < ?", start, end).
		Scan(&res)
	monthExpense = res.Total

	Success(c, gin.H{
		"total_entries": totalEntries,
		"total_photos":  totalPhotos,
		"total_projects": totalProjects,
		"month_expense": monthExpense,
		"current_year":  now.Year(),
		"current_month": int(now.Month()),
	})
}

// ==========================================
// 6. Legacy Compatible Routes (平滑兼容原有旧路由)
// ==========================================

func GetLifeEntries(c *gin.Context) {
	entryType := c.Query("type")
	year := c.Query("year")

	var entries []models.LifeEntry
	query := db.DB.Preload("Attachments").Preload("Project").Order("occurred_at DESC, id DESC")

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

	for i := range entries {
		entries[i].Attachments = populateAttachmentURLs(entries[i].Attachments)
		if entries[i].Images == "" && len(entries[i].Attachments) > 0 {
			entries[i].Images = entries[i].Attachments[0].URL
		}
	}

	c.JSON(http.StatusOK, entries)
}

func GetNowStatus(c *gin.Context) {
	var now models.NowStatus
	if err := db.DB.First(&now).Error; err != nil {
		c.JSON(http.StatusOK, models.NowStatus{
			Building:    "迟暮生活档案馆 · 真正记录生活的数字空间",
			Learning:    "编译器底层机制、旁轴胶片冲洗、手冲浅烘水温曲线",
			Playing:     "黑神话：悟空、塞尔达传说、Steam 连麦放空",
			Listening:   "坂本龙一 - 《async》/《BTTB》",
			Reading:     "罗伯特·M·波西格 - 《禅与摩托车维修艺术》",
			Thinking:    "如何在喧嚣的信息时代，认认真真过好具体的生活，保持安宁与自洽。",
			Using:       "MacBook Pro 14, Contax T2, 纯白陶瓷滤杯, Midori MD 方格手账",
			Location:    "中国 · 杭州 (西湖区 / 余杭)",
			LastUpdated: "2026.10.09",
		})
		return
	}
	c.JSON(http.StatusOK, now)
}

func GetProjects(c *gin.Context) {
	var projects []models.Project
	query := db.DB.Order("`order` ASC, id DESC")
	if err := query.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}
	c.JSON(http.StatusOK, projects)
}

type YearArchiveStat struct {
	Year       string         `json:"year"`
	TotalCount int64          `json:"total_count"`
	TypeCounts map[string]int `json:"type_counts"`
	Months     []string       `json:"months"`
}

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

func GetConfig(c *gin.Context) {
	var config models.SiteConfig
	if err := db.DB.First(&config).Error; err != nil {
		c.JSON(http.StatusOK, models.SiteConfig{
			SiteName:  "迟暮的生活档案馆",
			Domain:    "codeactivityhub.top",
			ICPNumber: "浙ICP备2026081664号",
			ICPLink:   "https://beian.miit.gov.cn",
		})
		return
	}
	c.JSON(http.StatusOK, config)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":     "operational",
		"service":    "ChiMu Life Archive Core",
		"runtime":    runtime.Version(),
		"storage_r2": storage.IsConfigured(),
		"time":       time.Now().Unix(),
		"message":    fmt.Sprintf("Archive engine running quietly on %s", runtime.GOARCH),
	})
}
