package routes

import (
	"net/http"
	"share-the-meal/internal/handlers"
	"share-the-meal/internal/middleware"

	"share-the-meal/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine, db *pgxpool.Pool, logger *zap.Logger, hub *utils.Hub) {
	r.Use(gin.Recovery())

	authHandler := handlers.NewAuthHandler(db, logger)
	campaignHandler := handlers.NewCampaignHandler(db, logger)
	donationHandler := handlers.NewDonationHandler(db, logger)
	userHandler := handlers.NewUserHandler(db, logger)
	cmsHandler := handlers.NewCMSHandler(db, logger)
	companyHandler := handlers.NewCompanyHandler(logger)
	notificationHandler := handlers.NewNotificationHandler(db, logger)

	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		role, _ := c.Get("role")
		utils.ServeWs(hub, c.Writer, c.Request, userID.(int64), role.(string))
	})

	r.LoadHTMLGlob("templates/*")
	swaggerHandler := handlers.NewSwaggerHandler(logger)

	r.GET("/swagger.yaml", swaggerHandler.ServeSwaggerYAML)
	r.GET("/docs", swaggerHandler.ServeSwaggerUI)

r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := r.Group("/api/v1")
	{
		authRoutes := apiV1.Group("/auth-management")
		{
			authRoutes.POST("/sign-in", authHandler.SignInUser)
			authRoutes.POST("/register", authHandler.RegisterUser)
			authRoutes.POST("/forgot-password", authHandler.ForgetPassword)
			authRoutes.PUT("/change-password", authHandler.ChangePassword)
		}

		// Public routes
		publicRoutes := apiV1.Group("/public")
		{
			publicRoutes.GET("/campaigns", campaignHandler.ListActiveCampaigns)
			publicRoutes.GET("/campaigns/:id", campaignHandler.GetCampaignDetails)
			publicRoutes.GET("/company-profile", companyHandler.GetCompanyProfile)
		}

		// Authenticated routes
		auth := apiV1.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			// User routes
			userRoutes := auth.Group("/users")
			{
				userRoutes.GET("/profile", userHandler.GetUserProfile)
				userRoutes.PUT("/profile", userHandler.UpdateUserProfile)
				userRoutes.GET("/notifications", notificationHandler.GetUserNotifications)
			}

			// Donation routes
			donationRoutes := auth.Group("/donations")
			{
				donationRoutes.POST("", donationHandler.CreateDonation)
				donationRoutes.GET("", donationHandler.GetUserDonations)
			}

			// Real-time notifications
			auth.GET("/notifications/ws", func(c *gin.Context) {
				userID, _ := c.Get("userID")
				role, _ := c.Get("role")
				utils.ServeWs(hub, c.Writer, c.Request, userID.(int64), role.(string))
			})
		}

		// CMS routes (Admin only)
		cms := apiV1.Group("/cms")
		cms.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("superadmin"))
		{
			cms.POST("/campaigns", cmsHandler.CreateCampaign)
			cms.PUT("/campaigns/:id", cmsHandler.UpdateCampaign)
			cms.DELETE("/campaigns/:id", cmsHandler.DeleteCampaign)
			cms.GET("/campaigns/:id/stats", cmsHandler.GetCampaignStats)
			cms.GET("/donations", cmsHandler.ListAllDonations)
			cms.PUT("/company-profile", companyHandler.UpdateCompanyProfile)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found",
		})
	})
}
