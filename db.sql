CREATE TABLE users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL
);
CREATE TABLE refresh_tokens (
    token_id UUID PRIMARY KEY,         
    token_hash VARCHAR(255) NOT NULL,                   
    ip_address VARCHAR(45) UNIQUE NOT NULL,
);
