CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL
);
CREATE TABLE refresh_tokens (
    token_id UUID PRIMARY KEY,         
    token_hash VARCHAR(255) NOT NULL,                   
    ip_address VARCHAR(45) UNIQUE NOT NULL
);

INSERT INTO users (id, email) VALUES 
('5e27b8c0-8b4f-45f6-9c47-bc9bbdb5a9ea', 'john.doe@example.com'),
('2a7d8e8a-345e-40d7-a0df-8ab3f3eb73b6', 'jane.smith@example.com'),
('c9bf9e57-1685-4c89-bafb-f04d596f723d', 'michael.brown@example.com');