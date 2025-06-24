package handlers

import (
	"net/http"
	// "strconv"
	"share-the-meal/internal/dto/request"
	"share-the-meal/internal/dto/response"
	"share-the-meal/internal/repository"
	"share-the-meal/internal/services"
	"share-the-meal/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type DonationHandler struct {
	donationService *services.DonationService
	logger          *zap.Logger
}

func NewDonationHandler(db *pgxpool.Pool, logger *zap.Logger) *DonationHandler {
	donationRepo := repository.NewDonationRepository(db, "public")
	campaignRepo := repository.NewCampaignRepository(db, "public")
	notificationRepo := repository.NewNotificationsRepository(db, "public")

	// Inisialisasi hub (bisa dari tempat lain atau buat baru jika perlu)
	hub := utils.NewHub()
	go hub.Run()

	return &DonationHandler{
		donationService: services.NewDonationService(
			donationRepo,
			campaignRepo,
			notificationRepo,
			hub,
		),
		logger: logger,
	}
}

func (h *DonationHandler) CreateDonation(c *gin.Context) {
	var req request.DonationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request"))
		return
	}

	// Ambil userID dari context (setelah autentikasi)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse("Unauthorized"))
		return
	}

	donation, err := h.donationService.CreateDonation(c.Request.Context(), req, userID.(int64))
	if err != nil {
		h.logger.Error("Failed to create donation", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to create donation"))
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse(donation))
}

func (h *DonationHandler) GetUserDonations(c *gin.Context) {
	// Ambil userID dari context (setelah autentikasi)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse("Unauthorized"))
		return
	}

	donations, err := h.donationService.GetUserDonations(c.Request.Context(), userID.(int64))
	if err != nil {
		h.logger.Error("Failed to get donations", zap.Error(err))
		c.JSON(http.StatusInternalServerError, response.ErrorResponse("Failed to get donations"))
		return
	}

	// Konversi ke response DTO
	var res []response.DonationResponse
	for _, d := range donations {
		res = append(res, response.DonationResponse{
			ID:          d.ID,
			Amount:      d.Amount,
			IsAnonymous: d.IsAnonymous,
			CreatedAt:   d.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, response.SuccessResponse(res))
}
