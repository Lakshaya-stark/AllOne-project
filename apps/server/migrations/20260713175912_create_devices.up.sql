CREATE TABLE devices (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100),
    platform VARCHAR(50),
    device_type VARCHAR(50),
    public_key TEXT,
    last_seen TIMESTAMP,
    created_at TIMESTAMP NOT NULL
);