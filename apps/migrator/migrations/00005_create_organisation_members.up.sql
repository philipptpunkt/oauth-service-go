BEGIN;

CREATE TABLE organisation_members (
    id SERIAL PRIMARY KEY,
    organisation_id INT NOT NULL REFERENCES organisations(id) ON DELETE CASCADE,
    client_id INT NOT NULL REFERENCES client_credentials(id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL,
    job_title VARCHAR(100),
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (organisation_id, client_id)
);

COMMIT;