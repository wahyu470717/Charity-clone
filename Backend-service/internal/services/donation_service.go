package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"share-the-meal/internal/dto/request"
	"share-the-meal/internal/dto/response"
	"share-the-meal/internal/models"
	"share-the-meal/internal/repository"
	"share-the-meal/internal/utils"
	"time"
)

type DonationService struct {
	donationRepo     repository.DonationRepositoryInterface
	campaignRepo     repository.CampaignRepositoryInterface
	notificationRepo repository.NotificationsRepositoryInterface
	hub              *utils.Hub
}

func NewDonationService(
	donationRepo repository.DonationRepositoryInterface,
	campaignRepo repository.CampaignRepositoryInterface,
	notificationRepo repository.NotificationsRepositoryInterface,
	hub *utils.Hub,
) *DonationService {
	return &DonationService{
		donationRepo:     donationRepo,
		campaignRepo:     campaignRepo,
		notificationRepo: notificationRepo,
		hub:              hub,
	}
}

func (s *DonationService) CreateDonation(ctx context.Context, req request.DonationRequest, userID int64) (*response.DonationResponse, error) {
	// Create donation
	donation := &models.Donation{
		UserID:      userID,
		CampaignID:  req.CampaignID,
		Amount:      req.Amount,
		IsAnonymous: req.IsAnonymous,
	}

	err := s.donationRepo.CreateDonation(ctx, donation)
	if err != nil {
		return nil, err
	}

	// Update campaign current amount
	campaign, err := s.campaignRepo.GetCampaignByID(ctx, req.CampaignID)
	if err != nil {
		return nil, fmt.Errorf("failed to get campaign: %w", err)
	}

	campaign.Current += req.Amount
	err = s.campaignRepo.UpdateCampaign(ctx, campaign)
	if err != nil {
		return nil, fmt.Errorf("failed to update campaign: %w", err)
	}

	// Send real-time notification
	go func() {
		notification := &models.Notifications{
			UserID:    userID,
			Title:     "Donation Successful",
			Message:   fmt.Sprintf("Your donation of $%.2f has been processed", req.Amount),
			IsRead:    false,
			CreatedBy: "system",
			CreatedAt: time.Now(),
		}
		
		if err := s.notificationRepo.CreateNotification( notification); err != nil {
			log.Printf("Failed to create notification: %v", err)
		}

		// Notify via WebSocket
		msg := map[string]interface{}{
			"type":    "donation",
			"amount":  req.Amount,
			"message": "Thank you for your donation!",
		}
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Failed to marshal websocket message: %v", err)
			return
		}
		s.hub.NotifyUser(userID, jsonMsg)
	}()

	return &response.DonationResponse{
		ID:          donation.ID,
		Amount:      donation.Amount,
		IsAnonymous: donation.IsAnonymous,
		CreatedAt:   donation.CreatedAt,
	}, nil
}

// GetCampaignStats returns statistics for a campaign
func (s *DonationService) GetCampaignStats(ctx context.Context, campaignID int64) (*response.CampaignResponse, error) {
	// Get campaign details
	_, err := s.campaignRepo.GetCampaignByID(ctx, campaignID)
	if err != nil {
		return nil, err
	}

	// Get all donations for the campaign
	donations, err := s.donationRepo.GetCampaignDonations(ctx, campaignID)
	if err != nil {
		return nil, err
	}

	// Calculate statistics
	totalDonations := 0.0
	donorCount := make(map[int64]bool) // Track unique donors

	for _, donation := range donations {
		totalDonations += donation.Amount
		donorCount[donation.UserID] = true
	}

	// Prepare response
	stats := &response.CampaignResponse{
		CampaignID:    campaignID,
		// CampaignTitle: campaign.Title,
		// TargetAmount:  campaign.Target,
		// CurrentAmount: campaign.Current,
		// TotalDonors:   len(donorCount),
		// TotalDonations: len(donations),
		// TotalAmount:   totalDonations,
	}

	return stats, nil
}

// GetUserDonations retrieves donations made by a specific user
func (s *DonationService) GetUserDonations(ctx context.Context, userID int64) ([]models.Donation, error) {
    return s.donationRepo.GetUserDonations(ctx, userID)
}

func (s *DonationService) ListAllDonations(ctx context.Context) ([]*response.DonationResponse, error) {
	donations, err := s.donationRepo.GetAllDonations(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*response.DonationResponse
	for _, d := range donations {
		responses = append(responses, &response.DonationResponse{
			ID:          d.ID,
			// UserID:      d.UserID,
			// CampaignID:  d.CampaignID,
			Amount:      d.Amount,
			IsAnonymous: d.IsAnonymous,
			CreatedAt:   d.CreatedAt,
		})
	}
	return responses, nil
}