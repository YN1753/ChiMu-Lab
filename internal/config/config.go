package config

import (
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
	DBPath string
	Port   string
	R2     R2Config
}

var AppConfig Config

func LoadConfig() Config {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/chimu.db"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

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
		DBPath: dbPath,
		Port:   port,
		R2:     r2Cfg,
	}

	return AppConfig
}

func (r *R2Config) IsConfigured() bool {
	return r.AccessKeyID != "" && r.SecretAccessKey != "" && r.BucketName != "" && r.Endpoint != ""
}
