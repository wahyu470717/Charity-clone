package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config Holds All Configuration For The Application
type Config struct {
	ServerPort   string
	DatabaseUrl  string
	JWTSecret    string
	Environment  string
	CORSOrigins  string
	StoragePath  string
	MaxFileSize  int64
	AllowedTypes []string
	DBConfig     DBConfig
	MinioConfig  MinioConfig
}

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

type MinioConfig struct {
	Endpoint   string
	AccessKey  string
	SecretKey  string
	UseSSL     bool
	BucketName string
	Region     string
}

// getEnvreturn environment variable value or default if not set
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

// getEnvAsInt returns environment variable as int64 or default if not set or invalid
func getEnvAsInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
		log.Printf("Warning: Invalid integer value for %s: %s, using default: %d", key, value, defaultValue)
	}
	return defaultValue
}

// getEnvAsStringSlice returns environment variable as string slice (comma-separated)
func getEnvAsStringSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		items := strings.Split(value, ",")
		for i, item := range items {
			items[i] = strings.TrimSpace(item)
		}
		return items
	}

	return defaultValue
}

func LoadConfig() (*Config, error) {
	//Load .env file if it exist
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dbConfig := DBConfig{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "share_the_meal"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Default allowed file types
	defaultAllowedTypes := []string{"image/jpeg", "image/png", "image/gif"}

	useSSL, _ := strconv.ParseBool(getEnv("MINIO_USE_SSL", "true"))

	minioConfig := MinioConfig{
		Endpoint:   getEnv("MINIO_ENDPOINT", ""),
		AccessKey:  getEnv("MINIO_ACCESS_KEY", ""),
		SecretKey:  getEnv("MINIO_SECRET_KEY", ""),
		UseSSL:     useSSL,
		BucketName: getEnv("MINIO_BUCKET_NAME", ""),
		Region:     getEnv("MINIO_REGION", "us-east-1"),
	}

	// Load main configuration
	config := &Config{
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		JWTSecret:    getEnv("JWT_SECRET", ""),
		Environment:  getEnv("ENVIRONMENT", "development"),
		CORSOrigins:  getEnv("CORS_ORIGINS", "*"),
		StoragePath:  getEnv("STORAGE_PATH", "./uploads"),
		MaxFileSize:  getEnvAsInt64("MAX_FILE_SIZE", 5242880),
		AllowedTypes: getEnvAsStringSlice("ALLOWED_TYPES", defaultAllowedTypes),
		DBConfig:     dbConfig,
		MinioConfig:  minioConfig,
	}

	// Build database URL from individual components
	config.DatabaseUrl = buildDatabaseUrl(dbConfig)

	// Validate required fields
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

// buildDatabaseUrl constructs database URL from DBConfig
func buildDatabaseUrl(db DBConfig) string {
	return "host=" + db.DBHost +
		" port=" + db.DBPort +
		" user=" + db.DBUser +
		" password=" + db.DBPassword +
		" dbname=" + db.DBName +
		" sslmode=" + db.DBSSLMode
}

func validateConfig(config *Config) error {
	if config.JWTSecret == "" {
		log.Println("Warning: JWT_SECRET is not set")
	}

	if config.DBConfig.DBPassword == "" {
		log.Println("Warning: DB_PASSWORD is not set")
	}

		// Validasi MinIO
	if config.MinioConfig.Endpoint == "" {
		log.Println("Warning: MINIO_ENDPOINT is not set")
	}
	if config.MinioConfig.AccessKey == "" {
		log.Println("Warning: MINIO_ACCESS_KEY is not set")
	}
	if config.MinioConfig.SecretKey == "" {
		log.Println("Warning: MINIO_SECRET_KEY is not set")
	}
	if config.MinioConfig.BucketName == "" {
		log.Println("Warning: MINIO_BUCKET_NAME is not set")
	}
	

	// You can add more validation as needed
	return nil
}

var globalConfig *Config

func GetConfig() (*Config, error) {
	if globalConfig == nil {
		var err error
		globalConfig, err = LoadConfig()
		if err != nil {
			return nil, err
		}
	}
	return globalConfig, nil
}

func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}
