package repository

import (
	"context"
	"share-the-meal/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DonationRepositoryInterface interface {
	CreateDonation(ctx context.Context, donation *models.Donation) error
	GetDonationByID(ctx context.Context, id int64) (*models.Donation, error)
	GetUserDonations(ctx context.Context, userID int64) ([]models.Donation, error)
	GetCampaignDonations(ctx context.Context, campaignID int64) ([]models.Donation, error)
	GetAllDonations(ctx context.Context) ([]models.Donation, error)
}

type DonationRepository struct {
	db     *pgxpool.Pool
	schema string
}

func NewDonationRepository(db *pgxpool.Pool, schema string) *DonationRepository {
	return &DonationRepository{
		db:     db,
		schema: schema,
	}
}

func (r *DonationRepository) CreateDonation(ctx context.Context, donation *models.Donation) error {
	query := `
		INSERT INTO donations (user_id, campaign_id, amount, is_anonymous, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	return r.db.QueryRow(ctx, query,
		donation.UserID,
		donation.CampaignID,
		donation.Amount,
		donation.IsAnonymous,
		time.Now(),
	).Scan(&donation.ID)
}

func (r *DonationRepository) GetDonationByID(ctx context.Context, id int64) (*models.Donation, error) {
	query := `
		SELECT id, user_id, campaign_id, amount, is_anonymous, created_at
		FROM donations
		WHERE id = $1
	`

	var donation models.Donation
	err := r.db.QueryRow(ctx, query, id).Scan(
		&donation.ID,
		&donation.UserID,
		&donation.CampaignID,
		&donation.Amount,
		&donation.IsAnonymous,
		&donation.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &donation, nil
}

func (r *DonationRepository) GetUserDonations(ctx context.Context, userID int64) ([]models.Donation, error) {
	query := `
		SELECT id, user_id, campaign_id, amount, is_anonymous, created_at
		FROM donations
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var donations []models.Donation
	for rows.Next() {
		var d models.Donation
		if err := rows.Scan(
			&d.ID,
			&d.UserID,
			&d.CampaignID,
			&d.Amount,
			&d.IsAnonymous,
			&d.CreatedAt,
		); err != nil {
			return nil, err
		}
		donations = append(donations, d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return donations, nil
}

func (r *DonationRepository) GetCampaignDonations(ctx context.Context, campaignID int64) ([]models.Donation, error) {
	query := `
		SELECT id, user_id, campaign_id, amount, is_anonymous, created_at
		FROM donations
		WHERE campaign_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var donations []models.Donation
	for rows.Next() {
		var d models.Donation
		if err := rows.Scan(
			&d.ID,
			&d.UserID,
			&d.CampaignID,
			&d.Amount,
			&d.IsAnonymous,
			&d.CreatedAt,
		); err != nil {
			return nil, err
		}
		donations = append(donations, d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return donations, nil
}

func (r *DonationRepository) GetAllDonations(ctx context.Context) ([]models.Donation, error) {
	query := `
		SELECT id, user_id, campaign_id, amount, is_anonymous, created_at
		FROM donations
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var donations []models.Donation
	for rows.Next() {
		var d models.Donation
		if err := rows.Scan(
			&d.ID,
			&d.UserID,
			&d.CampaignID,
			&d.Amount,
			&d.IsAnonymous,
			&d.CreatedAt,
		); err != nil {
			return nil, err
		}
		donations = append(donations, d)
	}

	return donations, nil
}
