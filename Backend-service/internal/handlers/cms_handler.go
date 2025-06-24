package handlers

import (
	"mime/multipart"
	"net/http"
	"share-the-meal/internal/dto/request"
	"share-the-meal/internal/dto/response"
	"share-the-meal/internal/repository"
	"share-the-meal/internal/utils"

	// "share-the-meal/internal/middleware"
	"share-the-meal/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type CMSHandler struct {
	campaignService *services.CampaignService
	donationService *services.DonationService
	logger          *zap.Logger
}

func NewCMSHandler(db *pgxpool.Pool, logger *zap.Logger) *CMSHandler {
	campaignRepo := repository.NewCampaignRepository(db, "public")
	donationRepo := repository.NewDonationRepository(db, "public")
	minioUtil := utils.GetMinIOUtil()

	return &CMSHandler{
		campaignService: services.NewCampaignService(campaignRepo, minioUtil),
		donationService: services.NewDonationService(donationRepo, campaignRepo, nil, nil),
		logger:          logger,
	}
}

// CreateCampaign godoc
// @Summary Create a new campaign
// @Description Create a new campaign (Superadmin only)
// @Tags CMS
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Campaign title"
// @Param description formData string true "Campaign description"
// @Param target formData number true "Target amount"
// @Param image formData file true "Campaign image"
// @Security BearerAuth
// @Success 201 {object} response.APIResponse{data=response.CampaignResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /cms/campaigns [post]
func (h *CMSHandler) CreateCampaign(c *gin.Context) {
	var req request.CreateCampaignRequest
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request"))
		return
	}

	file, err := utils.HandleFileUpload(c, "image")
	if err != nil {
		h.logger.Error("File upload error", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Image is required"))
		return
	}

	campaign, err := h.campaignService.CreateCampaign(c.Request.Context(), req, file)
	if err != nil {
		h.logger.Error("Failed to create campaign", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to create campaign"))
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse(campaign))
}

// GetCampaignStats godoc
// @Summary Get campaign statistics
// @Description Get statistics for a campaign (Superadmin only)
// @Tags CMS
// @Produce json
// @Param id path int true "Campaign ID"
// @Security BearerAuth
// @Success 200 {object} response.APIResponse{data=response.CampaignResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 403 {object} response.APIResponse "Forbidden"
// @Failure 404 {object} response.APIResponse
// @Router /cms/campaigns/{id}/stats [get]
func (h *CMSHandler) GetCampaignStats(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid campaign ID"))
		return
	}

	stats, err := h.donationService.GetCampaignStats(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get stats", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to get stats"))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(stats))
}

// UpdateCampaign godoc
// @Summary Update a campaign
// @Description Update an existing campaign (Superadmin only)
// @Tags CMS
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Campaign ID"
// @Param title formData string false "Campaign title"
// @Param description formData string false "Campaign description"
// @Param target formData number false "Target amount"
// @Param image formData file false "Campaign image"
// @Security BearerAuth
// @Success 200 {object} response.APIResponse{data=response.CampaignResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /cms/campaigns/{id} [put]
func (h *CMSHandler) UpdateCampaign(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid campaign ID"))
		return
	}

	var req request.UpdateCampaignRequest
	if err := c.ShouldBind(&req); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request"))
		return
	}

	var file *multipart.FileHeader
	if formFile, err := c.FormFile("image"); err == nil {
		file = formFile
	}

	campaign, err := h.campaignService.UpdateCampaign(c.Request.Context(), id, req, file)
	if err != nil {
		h.logger.Error("Failed to update campaign", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to update campaign"))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(campaign))
}

// DeleteCampaign godoc
// @Summary Delete a campaign
// @Description Delete a campaign (Superadmin only)
// @Tags CMS
// @Produce json
// @Param id path int true "Campaign ID"
// @Security BearerAuth
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /cms/campaigns/{id} [delete]
func (h *CMSHandler) DeleteCampaign(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid campaign ID"))
		return
	}

	err = h.campaignService.DeleteCampaign(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("Failed to delete campaign", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to delete campaign"))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse("Campaign deleted successfully"))
}

// ListAllDonations godoc
// @Summary List all donations
// @Description Get all donations (Superadmin only)
// @Tags CMS
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.APIResponse{data=[]response.DonationResponse}
// @Failure 500 {object} response.APIResponse
// @Router /cms/donations [get]
func (h *CMSHandler) ListAllDonations(c *gin.Context) {
	donations, err := h.donationService.ListAllDonations(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get donations", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to get donations"))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(donations))
}
