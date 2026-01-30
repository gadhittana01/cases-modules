package utils

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type BaseConfig struct {
	MigrationURL string
	DBName       string
}

type Config struct {
	DBConnString         string
	Port                 string
	MigrationURL         string
	DBName               string
	JWTSecret            string
	StripeKey            string
	StripeSecret         string
	StripeWebhookSecret  string
	FrontendURL          string
	SupabaseURL          string
	SupabaseKey          string
	SupabaseStorageBucket string
	StorageEndpoint      string
	StorageRegion        string
	StorageAccessKey     string
	StorageSecretKey     string
	StorageBucket        string
	PusherAppID          string
	PusherKey            string
	PusherSecret         string
	PusherCluster        string
}

func LoadEnv(paths ...string) {
	if len(paths) > 0 {
		godotenv.Load(paths...)
	} else {
		godotenv.Load()
	}
}

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func ValidateConfig(config *Config) error {
	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}
	return nil
}

func CheckAndSetConfig(configPath, configName string) *Config {
	LoadEnv(configPath + "/" + configName + ".env")

	return &Config{
		DBConnString:         GetEnv("DB_CONN_STRING", "postgresql://postgres.mhheblvgktovcrdcsjdo:yrYqVWR2Q1ASnPBc@aws-1-ap-southeast-1.pooler.supabase.com:5432/postgres"),
		Port:                 GetEnv("PORT", "8000"),
		MigrationURL:         GetEnv("MIGRATION_URL", "file://db/migration"),
		DBName:               GetEnv("DB_NAME", "postgres"),
		JWTSecret:            GetEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		StripeKey:            GetEnv("STRIPE_PUBLISHABLE_KEY", ""),
		StripeSecret:         GetEnv("STRIPE_SECRET_KEY", ""),
		StripeWebhookSecret:  GetEnv("STRIPE_WEBHOOK_SECRET", ""),
		FrontendURL:          GetEnv("FRONTEND_URL", "http://localhost:3000"),
		SupabaseURL:          GetEnv("SUPABASE_URL", ""),
		SupabaseKey:          GetEnv("SUPABASE_KEY", ""),
		SupabaseStorageBucket: GetEnv("SUPABASE_STORAGE_BUCKET", "cases-files"),
		StorageEndpoint:      GetEnv("STORAGE_ENDPOINT", ""),
		StorageRegion:        GetEnv("STORAGE_REGION", "ap-southeast-1"),
		StorageAccessKey:     GetEnv("STORAGE_ACCESS_KEY", ""),
		StorageSecretKey:     GetEnv("STORAGE_SECRET_KEY", ""),
		StorageBucket:        GetEnv("STORAGE_BUCKET", "my-bucket"),
		PusherAppID:          GetEnv("PUSHER_APP_ID", ""),
		PusherKey:            GetEnv("PUSHER_KEY", ""),
		PusherSecret:         GetEnv("PUSHER_SECRET", ""),
		PusherCluster:        GetEnv("PUSHER_CLUSTER", "ap1"),
	}
}
