CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,           -- Auto-incrementing primary key
    firstname VARCHAR(255) NOT NULL,  -- Firstname field, cannot be null
    lastname VARCHAR(255) NOT NULL,  -- Lastname field, cannot be null
    email VARCHAR(255) UNIQUE NOT NULL, -- Unique email field, cannot be null
    password TEXT NOT NULL,     -- Password hash, cannot be null
    refreshToken VARCHAR(128) DEFAULT NULL, -- Refresh token, null by default
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp of creation
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP -- Timestamp of last update
);

