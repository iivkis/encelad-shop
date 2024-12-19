CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY, 
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    role VARCHAR(128) NOT NULL DEFAULT 'user',
    hashed_password VARCHAR(60) NOT NULL
);