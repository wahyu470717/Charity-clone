package services

import (
	"context"
	"mime/multipart"
	"share-the-meal/internal/dto/request"
	"share-the-meal/internal/dto/response"
	"share-the-meal/internal/models"
	"share-the-meal/internal/repository"
	"share-the-meal/internal/utils"
)

type CampaignService struct {
	campaignRepo repository.CampaignRepositoryInterface
	minioUtil    utils.MinIOUtilInterface
}

func NewCampaignService(campaignRepo repository.CampaignRepositoryInterface, minioUtil utils.MinIOUtilInterface) *CampaignService {
	return &CampaignService{
		campaignRepo: campaignRepo,
		minioUtil:    minioUtil,
	}
}

func (s *CampaignService) CreateCampaign(ctx context.Context, req request.CreateCampaignRequest, file *multipart.FileHeader) (*response.CampaignResponse, error) {
	imageURL, err := s.minioUtil.UploadFile(ctx, file, "campaigns", req.Title)
	if err != nil {
		return nil, err
	}

	campaign := &models.Campaigns{
		Title:       req.Title,
		Description: req.Content,
		Target:      req.Target,
		ImageURL:    imageURL,
		IsActive:    true,
	}

	err = s.campaignRepo.CreateCampaign(ctx, campaign)
	if err != nil {
		return nil, err
	}

	return &response.CampaignResponse{
		CampaignID: campaign.CampaignID,
		Title:      campaign.Title,
		Content:    campaign.Description,
		Target:     campaign.Target,
		Current:    campaign.Current,
		ImageUrl:   campaign.ImageURL,
	}, nil
}

func (s *CampaignService) GetCampaignDetails(ctx context.Context, id int64) (*response.CampaignResponse, error) {
	campaign, err := s.campaignRepo.GetCampaignByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &response.CampaignResponse{
		CampaignID:  campaign.CampaignID,
		Title:       campaign.Title,
		Description: campaign.Description,
		Target:      campaign.Target,
		Current:     campaign.Current,
		ImageUrl:    campaign.ImageURL,
	}, nil
}

// Implement other methods

func (s *CampaignService) ListActiveCampaigns(ctx context.Context) ([]*response.CampaignResponse, error) {
	campaigns, err := s.campaignRepo.ListActiveCampaigns(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*response.CampaignResponse
	for _, c := range campaigns {
		responses = append(responses, &response.CampaignResponse{
			CampaignID: c.CampaignID,
			Title:      c.Title,
			Content:    c.Description,
			Target:     c.Target,
			Current:    c.Current,
			ImageUrl:   c.ImageURL,
		})
	}
	return responses, nil
}

func (s *CampaignService) UpdateCampaign(
	ctx context.Context,
	id int64,
	req request.UpdateCampaignRequest,
	file *multipart.FileHeader,
) (*response.CampaignResponse, error) {
	campaign, err := s.campaignRepo.GetCampaignByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields
	if req.Title != "" {
		campaign.Title = req.Title
	}
	if req.Description != "" {
		campaign.Description = req.Description
	}
	if req.Target > 0 {
		campaign.Target = req.Target
	}

	// Handle new image upload
	if file != nil {
		imageURL, err := s.minioUtil.UploadFile(ctx, file, "campaigns", req.Title)
		if err != nil {
			return nil, err
		}
		campaign.ImageURL = imageURL
	}

	err = s.campaignRepo.UpdateCampaign(ctx, campaign)
	if err != nil {
		return nil, err
	}

	return &response.CampaignResponse{
		CampaignID:  campaign.CampaignID,
		Title:       campaign.Title,
		Description: campaign.Description,
		Target:      campaign.Target,
		ImageUrl:    campaign.ImageURL,
	}, nil
}

func (s *CampaignService) DeleteCampaign(ctx context.Context, id int64) error {
	return s.campaignRepo.DeleteCampaign(ctx, id)
}
