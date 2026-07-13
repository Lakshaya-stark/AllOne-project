CREATE TABLE users (

    id UUID PRIMARY KEY,

    username VARCHAR(50) UNIQUE NOT NULL,

    email VARCHAR(255) UNIQUE NOT NULL,

    password_hash TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL,

    updated_at TIMESTAMP NOT NULL

);