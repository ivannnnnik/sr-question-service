CREATE TABLE IF NOT EXISTS question (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(128) NOT NULL UNIQUE,
    category VARCHAR(128) NOT NULL,
    difficulty INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);