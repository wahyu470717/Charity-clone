CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role INTEGER,
    fullname VARCHAR(255),
    profile_picture VARCHAR(255),
    phone_number VARCHAR(20),
    address TEXT,
    is_active BOOL,
    created_by VARCHAR(255),
    modified_by VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);

INSERT INTO users (name, email, password, role) 
VALUES ('Admin', 'admin@sharethemeal.org', '$2a$10$1qAz2wSx3eDc4rFv5tGb5t/yMB3Vjr97gYVZcROVqlfS1QgURXOa.', 1);
