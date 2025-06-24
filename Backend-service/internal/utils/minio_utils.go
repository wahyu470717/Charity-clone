package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"regexp"
	"share-the-meal/internal/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOUtil struct {
	client     *minio.Client
	bucketName string
}

type MinIOUtilInterface interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, folderPath, title string) (string, error)
	GetFileURL(filename string) string
	// HandleFileUpload(c interface{}, fieldName string) (*multipart.FileHeader, error) // Gin Context diganti interface{} agar mockable
	DeleteFile(ctx context.Context, fileName string) error
}

var MinIOUtilInstance *MinIOUtil

func InitMinIOUtil(cfg *config.MinioConfig) error {
	// Parse the endpoint to remove any protocol prefix
	endpoint := cfg.Endpoint

	// Remove protocol prefix if present (http:// or https://)
	if strings.HasPrefix(endpoint, "http://") {
		endpoint = strings.TrimPrefix(endpoint, "http://")
	} else if strings.HasPrefix(endpoint, "https://") {
		endpoint = strings.TrimPrefix(endpoint, "https://")
	}

	// If there's any path in the URL, remove it
	if parsedURL, err := url.Parse("http://" + endpoint); err == nil {
		endpoint = parsedURL.Host
	}

	// Initialize minio client object
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})

	if err != nil {
		return fmt.Errorf("failed to create MinIO client: %v", err)
	}

	// Create the bucket if it doesn't exist
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, cfg.BucketName)
	if err != nil {
		return fmt.Errorf("failed to check bucket existence: %v", err)
	}

	if !exists {
		err = minioClient.MakeBucket(ctx, cfg.BucketName, minio.MakeBucketOptions{
			Region: cfg.Region,
		})

		if err != nil {
			return fmt.Errorf("failed to create bucket: %v", err)
		}
	}

	MinIOUtilInstance = &MinIOUtil{
		client:     minioClient,
		bucketName: cfg.BucketName,
	}

	return nil
}

func GetMinIOUtil() MinIOUtilInterface {
	return MinIOUtilInstance
}

func (m *MinIOUtil) UploadFile(ctx context.Context, file *multipart.FileHeader, folderPath, title string) (string, error) {
	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	sanitizedTitle := sanitizeForFilename(title)
	ext := strings.ToLower(filepath.Ext(file.Filename))
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	filename := fmt.Sprintf("%s_%d%s", sanitizedTitle, timestamp, ext)

	if !strings.HasSuffix(folderPath, "/") {
		folderPath += "/"
	}

	objectName := folderPath + filename

	// Upload file ke MinIO
	_, err = m.client.PutObject(
		ctx,
		m.bucketName,
		objectName,
		src,
		file.Size,
		minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")},
	)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	return filename, nil
}

func sanitizeForFilename(input string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	sanitized := reg.ReplaceAllString(input, "_")
	sanitized = strings.Trim(sanitized, "_")
	return sanitized
}

func (m *MinIOUtil) GetFileURL(filename string) string {
	if filename == "" {
		return ""
	}
	return fmt.Sprintf("http://%s/%s/%s", m.client.EndpointURL().Host, m.bucketName, filename)
}

// Gin middleware to handle file upload
func HandleFileUpload(c *gin.Context, fieldName string) (*multipart.FileHeader, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return nil, fmt.Errorf("failed to get file from form: %v", err)
	}

	// Validate file size (e.g., 5MB limit)
	if file.Size > 5<<20 {
		return nil, fmt.Errorf("file size exceeds 5MB limit")
	}

	// Validate file extension
	ext := filepath.Ext(file.Filename)
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}
	if !allowedExts[ext] {
		return nil, fmt.Errorf("invalid file extension. Only jpg, jpeg, png, webp are allowed")
	}

	return file, nil
}

func (m *MinIOUtil) DeleteFile(ctx context.Context, fileName string) error {
	if fileName == "" {
		return nil
	}
	err := m.client.RemoveObject(ctx, m.bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file from MinIO: %w", err)
	}
	return nil
}
