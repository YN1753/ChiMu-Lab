package config

import (
	"bufio"
	"os"
	"strings"
)

type R2Config struct {
	AccountID       string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	Endpoint        string
	PublicDomain    string
}

type Config struct {
	DBPath      string
	Port        string
	AdminAPIKey string
	R2          R2Config
}

var AppConfig Config

// loadDotEnv 轻量读取当前工作目录下的 .env 文件
func loadDotEnv() {
	file, err := os.Open(".env")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		val = strings.Trim(val, `"'`)
		// 仅在环境变量未显式设置时填充
		if os.Getenv(key) == "" {
			_ = os.Setenv(key, val)
		}
	}
}

func LoadConfig() Config {
	loadDotEnv()

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/chimu.db"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	adminKey := strings.TrimSpace(os.Getenv("ADMIN_API_KEY"))

	accountID := strings.TrimSpace(os.Getenv("R2_ACCOUNT_ID"))
	endpoint := strings.TrimSpace(os.Getenv("R2_ENDPOINT"))
	if endpoint == "" && accountID != "" {
		endpoint = "https://" + accountID + ".r2.cloudflarestorage.com"
	}

	r2Cfg := R2Config{
		AccountID:       accountID,
		AccessKeyID:     strings.TrimSpace(os.Getenv("R2_ACCESS_KEY_ID")),
		SecretAccessKey: strings.TrimSpace(os.Getenv("R2_SECRET_ACCESS_KEY")),
		BucketName:      strings.TrimSpace(os.Getenv("R2_BUCKET_NAME")),
		Endpoint:        endpoint,
		PublicDomain:    strings.TrimRight(strings.TrimSpace(os.Getenv("R2_PUBLIC_DOMAIN")), "/"),
	}

	AppConfig = Config{
		DBPath:      dbPath,
		Port:        port,
		AdminAPIKey: adminKey,
		R2:          r2Cfg,
	}

	return AppConfig
}

func (r *R2Config) IsConfigured() bool {
	return r.AccessKeyID != "" && r.SecretAccessKey != "" && r.BucketName != "" && r.Endpoint != ""
}
