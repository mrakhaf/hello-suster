CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    nip BIGINT NOT NULL UNIQUE,
    password VARCHAR(255),
    identityscanimage VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX IF NOT EXISTS index_name
	ON user USING(lower(name));

CREATE INDEX index_nip_string ON user CAST(nip AS VARCHAR);    