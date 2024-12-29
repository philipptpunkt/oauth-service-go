BEGIN;

CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL REFERENCES client_credentials(id) ON DELETE CASCADE,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    organisation VARCHAR(255),
    job_title VARCHAR(100),
    profile_picture TEXT, -- URL to the uploaded image
    time_zone VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMIT;
