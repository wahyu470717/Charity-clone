CREATE TABLE IF NOT EXISTS donations (
    id SERIAL PRIMARY KEY,
    campaign_id INTEGER REFERENCES campaigns(campaign_id) NOT NULL,
    donor_id INTEGER REFERENCES users(user_id) NOT NULL,
    amount DECIMAL(12, 2) NOT NULL,
    payment_status VARCHAR(20) NOT NULL DEFAULT 'pending',
    payment_method VARCHAR(50),
    transaction_id VARCHAR(255),
    message TEXT,
    anonymous BOOLEAN DEFAULT FALSE,
    created_by VARCHAR(255),
    modified_by VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_donations_campaign_id ON donations(campaign_id);
CREATE INDEX idx_donations_donor_id ON donations(donor_id);
CREATE INDEX idx_donations_payment_status ON donations(payment_status);