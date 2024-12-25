BEGIN;

CREATE TABLE client_refresh_tokens (
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL REFERENCES client_credentials(id) ON DELETE CASCADE,
    token TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMIT;