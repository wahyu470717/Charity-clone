package handlers

import (
	"net/http"
	"share-the-meal/internal/dto/response"
	"share-the-meal/internal/repository"
	"share-the-meal/internal/services"
	"share-the-meal/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type CampaignHandler struct {
	campaignService *services.CampaignService
	logger          *zap.Logger
}

// func NewCampaignHandler(db *pgxpool.Pool, logger *zap.Logger) *CampaignHandler {
// 	campaignRepo := repository.NewCampaignRepository(db, "public")
// 	minioUtil := utils.GetMinIOUtil()
// 	campaignService := services.NewCampaignService(campaignRepo, minioUtil)
// 	return &CampaignHandler{
// 		campaignService: campaignService,
// 		logger:          logger,
// 	}
// }

// Ganti inisialisasi
func NewCampaignHandler(db *pgxpool.Pool, logger *zap.Logger) *CampaignHandler {
    campaignRepo := repository.NewCampaignRepository(db, "public")
    minioUtil := utils.GetMinIOUtil()
    campaignService := services.NewCampaignService(campaignRepo, minioUtil)
    return &CampaignHandler{
        campaignService: campaignService,
        logger:          logger,
    }
}

func (h *CampaignHandler) ListActiveCampaigns(c *gin.Context) {
	campaigns, err := h.campaignService.ListActiveCampaigns(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to list campaigns", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to list campaigns"))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(campaigns))
}

func (h *CampaignHandler) GetCampaignDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid campaign ID"))
		return
	}

	campaign, err := h.campaignService.GetCampaignDetails(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get campaign details", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to get campaign details"))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(campaign))
}