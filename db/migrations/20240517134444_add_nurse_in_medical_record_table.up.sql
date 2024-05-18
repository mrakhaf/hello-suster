ALTER TABLE medical_record
ADD userId varchar(255),
ADD FOREIGN KEY (userId) REFERENCES users(id);