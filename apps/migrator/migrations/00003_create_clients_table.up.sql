BEGIN;

CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL REFERENCES client_credentials(id) ON DELETE CASCADE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    profile_picture TEXT
    time_zone VARCHAR(50),
    notification_preferences JSONB DEFAULT '{}'::JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMIT;