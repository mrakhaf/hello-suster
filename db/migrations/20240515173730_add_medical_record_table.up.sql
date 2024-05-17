CREATE TABLE medical_record (
  id VARCHAR(255) PRIMARY KEY,
  identitynumber BIGINT NOT NULL,
  symptoms TEXT NOT NULL,
  medications TEXT NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (identityNumber) REFERENCES patient(identitynumber)
);
