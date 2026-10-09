package storage

import (
	"context"
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"chimu-lab/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

var (
	s3Client       *s3.Client
	presignClient  *s3.PresignClient
	r2Bucket       string
	r2PublicDomain string
	isConfigured   bool

	safeFilenameRegex = regexp.MustCompile(`[^a-zA-Z0-9._-]`)
)

// InitR2 初始化 Cloudflare R2 / S3 客户端
func InitR2(cfg config.R2Config) error {
	if !cfg.IsConfigured() {
		isConfigured = false
		return nil
	}

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               cfg.Endpoint,
			HostnameImmutable: true,
			SigningRegion:     "auto",
		}, nil
	})

	awsCfg, err := s3config.LoadDefaultConfig(context.TODO(),
		s3config.WithEndpointResolverWithOptions(customResolver),
		s3config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, "")),
		s3config.WithRegion("auto"),
	)
	if err != nil {
		return fmt.Errorf("failed to load AWS config for R2: %w", err)
	}

	s3Client = s3.NewFromConfig(awsCfg)
	presignClient = s3.NewPresignClient(s3Client)
	r2Bucket = cfg.BucketName
	r2PublicDomain = cfg.PublicDomain
	isConfigured = true

	return nil
}

// IsConfigured 检查 R2 是否已配置就绪
func IsConfigured() bool {
	return isConfigured
}

// GetStorageStatus 获取存储服务当前公开状态 (不暴露密钥)
func GetStorageStatus() map[string]interface{} {
	cfg := config.AppConfig.R2
	return map[string]interface{}{
		"provider":      "Cloudflare R2",
		"configured":    isConfigured,
		"bucket":        cfg.BucketName,
		"public_domain": cfg.PublicDomain,
		"endpoint":      cfg.Endpoint,
	}
}

// sanitizeFilename 净化文件名，防止路径注入
func sanitizeFilename(raw string) string {
	base := filepath.Base(raw)
	ext := filepath.Ext(base)
	nameWithoutExt := strings.TrimSuffix(base, ext)

	safeName := safeFilenameRegex.ReplaceAllString(nameWithoutExt, "_")
	if len(safeName) > 50 {
		safeName = safeName[:50]
	}
	if safeName == "" {
		safeName = "file"
	}

	safeExt := strings.ToLower(safeFilenameRegex.ReplaceAllString(ext, ""))
	return safeName + safeExt
}

// GeneratePresignedPutURL 生成用于客户端直接上传到 R2 的预签名 PUT URL
func GeneratePresignedPutURL(filename, contentType string, size int64) (uploadID, objectKey, presignedURL string, expiresIn int, err error) {
	if !isConfigured {
		return "", "", "", 0, errors.New("R2 存储未配置，暂无法生成预签名上传凭证")
	}

	now := time.Now()
	cleanName := sanitizeFilename(filename)
	uid := uuid.New().String()[:8]

	// 设计规范：entries/{year}/{month}/{day}/{uuid}-{safe_filename}
	objectKey = fmt.Sprintf("entries/%04d/%02d/%02d/%s-%s",
		now.Year(), now.Month(), now.Day(), uid, cleanName)

	expiresIn = 900 // 15 分钟
	expiresDuration := time.Duration(expiresIn) * time.Second

	putInput := &s3.PutObjectInput{
		Bucket:        aws.String(r2Bucket),
		Key:           aws.String(objectKey),
		ContentType:   aws.String(contentType),
		ContentLength: aws.Int64(size),
	}

	presignedReq, err := presignClient.PresignPutObject(context.TODO(), putInput, s3.WithPresignExpires(expiresDuration))
	if err != nil {
		return "", "", "", 0, fmt.Errorf("failed to presign PUT request: %w", err)
	}

	uploadID = uuid.New().String()
	return uploadID, objectKey, presignedReq.URL, expiresIn, nil
}

// VerifyObjectExists 确认对象已实际上传到 R2，并获取真实大小与类型
func VerifyObjectExists(objectKey string) (size int64, contentType string, err error) {
	if !isConfigured {
		return 0, "", errors.New("R2 存储未配置")
	}

	headInput := &s3.HeadObjectInput{
		Bucket: aws.String(r2Bucket),
		Key:    aws.String(objectKey),
	}

	res, err := s3Client.HeadObject(context.TODO(), headInput)
	if err != nil {
		return 0, "", fmt.Errorf("object not found in R2: %w", err)
	}

	if res.ContentLength != nil {
		size = *res.ContentLength
	}
	if res.ContentType != nil {
		contentType = *res.ContentType
	}

	return size, contentType, nil
}

// DeleteObject 从 R2 删除指定对象
func DeleteObject(objectKey string) error {
	if !isConfigured || objectKey == "" {
		return nil
	}

	delInput := &s3.DeleteObjectInput{
		Bucket: aws.String(r2Bucket),
		Key:    aws.String(objectKey),
	}

	_, err := s3Client.DeleteObject(context.TODO(), delInput)
	return err
}

// GetPublicURL 计算公开访问 URL
func GetPublicURL(objectKey string) string {
	if objectKey == "" {
		return ""
	}
	if strings.HasPrefix(objectKey, "http://") || strings.HasPrefix(objectKey, "https://") {
		return objectKey
	}
	if r2PublicDomain != "" {
		return fmt.Sprintf("%s/%s", strings.TrimRight(r2PublicDomain, "/"), path.Clean(objectKey))
	}
	if config.AppConfig.R2.Endpoint != "" && r2Bucket != "" {
		return fmt.Sprintf("%s/%s/%s", strings.TrimRight(config.AppConfig.R2.Endpoint, "/"), r2Bucket, path.Clean(objectKey))
	}
	return objectKey
}
