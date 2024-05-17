CREATE TABLE medical_record (
    id VARCHAR(255) PRIMARY KEY,
    identitynumber BIGINT NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    phonenumber VARCHAR(255) NOT NULL,
    birthdate VARCHAR(255) NOT NULL,
    gender VARCHAR(255) NOT NULL,
    identityscanimage VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)