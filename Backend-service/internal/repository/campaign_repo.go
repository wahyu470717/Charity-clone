package repository

import (
	"context"
	"fmt"
	"share-the-meal/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CampaignRepositoryInterface interface {
	CreateCampaign(ctx context.Context, campaign *models.Campaigns) error
	GetCampaignByID(ctx context.Context, id int64) (*models.Campaigns, error)
	UpdateCampaign(ctx context.Context, campaign *models.Campaigns) error
	DeleteCampaign(ctx context.Context, id int64) error
	ListActiveCampaigns(ctx context.Context) ([]models.Campaigns, error)
}

type CampaignRepository struct {
	db     *pgxpool.Pool
	schema string
}

func NewCampaignRepository(db *pgxpool.Pool, schema string) *CampaignRepository {
	return &CampaignRepository{
		db:     db,
		schema: schema,
	}
}

func (r *CampaignRepository) CreateCampaign(ctx context.Context, campaign *models.Campaigns) error {
	query := `
		INSERT INTO campaigns (title, description, target_amount, current_amount, image_url, is_active, created_at, modified_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	return r.db.QueryRow(ctx, query,
		campaign.Title,
		campaign.Description,
		campaign.Target,
		campaign.Current,
		campaign.ImageURL,
		campaign.IsActive,
		time.Now(),
		time.Now(),
	).Scan(&campaign.CampaignID)
}

func (r *CampaignRepository) GetCampaignByID(ctx context.Context, id int64) (*models.Campaigns, error) {
	query := `
		SELECT id, title, description,  target_amount, current_amount, image_url, is_active, created_at, modified_at
		FROM campaigns
		WHERE id = $1
	`

	var campaign models.Campaigns
	err := r.db.QueryRow(ctx, query, id).Scan(
		&campaign.CampaignID,
		&campaign.Title,
		&campaign.Description,
		&campaign.Target,
		&campaign.Current,
		&campaign.ImageURL,
		&campaign.IsActive,
		&campaign.CreatedAt,
		&campaign.ModifiedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("campaign not found")
		}
		return nil, err
	}

	return &campaign, nil
}

func (r *CampaignRepository) UpdateCampaign(ctx context.Context, campaign *models.Campaigns) error {
	query := `
        UPDATE campaigns 
        SET title = $1, 
            description = $2, 
            target_amount = $3, 
            current_amount = $4, 
            image_url = $5, 
            is_active = $6,
            modified_at = $7
        WHERE id = $8
    `

	_, err := r.db.Exec(ctx, query,
		campaign.Title,
		campaign.Description,
		campaign.Target,
		campaign.Current,
		campaign.ImageURL,
		campaign.IsActive,
		time.Now(),
		campaign.CampaignID,
	)
	return err
}

func (r *CampaignRepository) DeleteCampaign(ctx context.Context, id int64) error {
	query := `DELETE FROM campaigns WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *CampaignRepository) ListActiveCampaigns(ctx context.Context) ([]models.Campaigns, error) {
	query := `
        SELECT id, title, description,  target_amount, current_amount, image_url, is_active, created_at, modified_at
        FROM campaigns
        WHERE is_active = true
    `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []models.Campaigns
	for rows.Next() {
		var c models.Campaigns
		err := rows.Scan(
			&c.CampaignID,
			&c.Title,
			&c.Description,
			&c.Target,
			&c.Current,
			&c.ImageURL,
			&c.IsActive,
			&c.CreatedAt,
			&c.ModifiedAt,
		)
		if err != nil {
			return nil, err
		}
		campaigns = append(campaigns, c)
	}

	return campaigns, nil
}

// Implement other methods: UpdateCampaign, DeleteCampaign, ListActiveCampaigns
