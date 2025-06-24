CREATE TABLE IF NOT EXISTS roles (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL UNIQUE,
    role_description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_by VARCHAR(255),
    modified_by VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert default roles
INSERT INTO roles (role_id, role_name, role_description, created_by, modified_by) VALUES
(1, 'superadmin', 'Super Administrator with full system access', 'system', 'system'),
(2, 'donor', 'Regular donor who can make donations', 'system', 'system'),
(3, 'recipient', 'Recipients who can create campaigns', 'system', 'system');