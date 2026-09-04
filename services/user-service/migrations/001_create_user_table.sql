CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,

    email TEXT UNIQUE NOT NULL,

    password_hash TEXT NOT NULL,

    first_name TEXT,

    last_name TEXT,

    created_at TIMESTAMP DEFAULT NOW(),

    updated_at TIMESTAMP DEFAULT NOW()
);