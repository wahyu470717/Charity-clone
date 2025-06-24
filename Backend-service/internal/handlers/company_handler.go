package handlers

import (
	"net/http"
	"share-the-meal/internal/dto/request"
	"share-the-meal/internal/dto/response"
	"share-the-meal/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CompanyHandler struct {
	logger *zap.Logger
}

func NewCompanyHandler(logger *zap.Logger) *CompanyHandler {
	return &CompanyHandler{logger: logger}
}

// GetCompanyProfile godoc
// @Summary Get company profile
// @Description Get the company profile information
// @Tags Company
// @Produce json
// @Success 200 {object} response.APIResponse{data=utils.CompanyProfile}
// @Router /public/company-profile [get]
func (h *CompanyHandler) GetCompanyProfile(c *gin.Context) {
	profile := utils.GetCompanyProfile()
	c.JSON(http.StatusOK, response.SuccessResponse(profile))
}

// UpdateCompanyProfile godoc
// @Summary Update company profile
// @Description Update the company profile (Superadmin only)
// @Tags Company
// @Accept json
// @Produce json
// @Param profile body request.UpdateCompanyProfileRequest true "Company profile"
// @Security BearerAuth
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /cms/company-profile [put]
func (h *CompanyHandler) UpdateCompanyProfile(c *gin.Context) {
	var req request.UpdateCompanyProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Invalid request"))
		return
	}

	utils.UpdateCompanyProfile(utils.CompanyProfile{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Phone:       req.Phone,
		Email:       req.Email,
		Website:     req.Website,
	})

	c.JSON(http.StatusOK, response.SuccessResponse("Company profile updated"))
}
