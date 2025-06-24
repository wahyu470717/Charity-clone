CREATE TABLE IF NOT EXISTS campaign (
    id SERIAL PRIMARY KEY,
    campaign_id INTEGER REFERENCES campaigns(campaign_id) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    image_url VARCHAR(255),
    created_by VARCHAR(255),
    modified_by VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_campaign_id ON campaign(id);