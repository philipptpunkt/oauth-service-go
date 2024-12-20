BEGIN;

CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    client_credentials_id INT UNIQUE NOT NULL REFERENCES client_credentials(id) ON DELETE CASCADE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    company_name VARCHAR(255) NOT NULL,
    client_id UUID DEFAULT gen_random_uuid() UNIQUE,
    client_secret UUID DEFAULT gen_random_uuid() UNIQUE,
    role VARCHAR(50) DEFAULT 'admin',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

COMMIT;