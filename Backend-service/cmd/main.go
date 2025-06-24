package main

import (
	"fmt"
	"log"
	"share-the-meal/internal/config"
	"share-the-meal/internal/middleware"
	"share-the-meal/internal/routes"
	"share-the-meal/internal/storage"
	"share-the-meal/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	_ "share-the-meal/docs"
)

// @title Share The Meal API
// @version 1.0
// @description API for Share The Meal donation platform
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey bearerToken
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token
func main() {
	// Initialize logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize JWT utility
	utils.InitJWTUtil(cfg.JWTSecret)

	// Connect to database
	pool, err := storage.ConnectDB(&cfg.DBConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	err = utils.InitMinIOUtil(&cfg.MinioConfig)
	if err != nil {
		log.Fatalf("Failed to initialize MinIO: %v", err)
	}

	// Set Gin mode
	if !cfg.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router
	router := gin.Default()

	router.Static("/docs", "./docs")

	// Add middleware
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.ZapLogger(logger))

	hub := utils.NewHub()
	go hub.Run()

	// Setup routes
	routes.SetupRoutes(router, pool, logger, hub)

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	fmt.Printf("🚀 Server starting on port %s\n", cfg.ServerPort)

	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
