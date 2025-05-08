-- 004_create_invitations.sql

CREATE TABLE IF NOT EXISTS invitations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    token TEXT NOT NULL UNIQUE,
    note TEXT,
    status TEXT NOT NULL CHECK (status IN ('unused', 'used', 'expired')),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    used_at TIMESTAMP WITH TIME ZONE
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_invitations_token ON invitations(token);