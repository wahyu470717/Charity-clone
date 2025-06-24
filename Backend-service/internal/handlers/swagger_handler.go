package handlers

import (
	"net/http"
	// "os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SwaggerHandler struct {
	logger *zap.Logger
}

func NewSwaggerHandler(logger *zap.Logger) *SwaggerHandler {
	return &SwaggerHandler{logger: logger}
}

func (h *SwaggerHandler) ServeSwaggerYAML(c *gin.Context) {
	// Dapatkan path root project
	_, currentFile, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(filepath.Dir(filepath.Dir(currentFile)))
	swaggerPath := filepath.Join(rootDir, "docs", "swagger.yaml")

	h.logger.Info("Serving Swagger YAML", zap.String("path", swaggerPath))
	c.File(swaggerPath)
}

func (h *SwaggerHandler) ServeSwaggerUI(c *gin.Context) {
	c.HTML(http.StatusOK, "swagger_index.html", gin.H{
		"Title":       "Share The Meal API",
		"SwaggerURL":  "/swagger.yaml",
		"DeepLinking": true,
	})
}
