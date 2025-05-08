-- 002_add_user_role.sql

ALTER TABLE users
    ADD COLUMN role TEXT NOT NULL DEFAULT 'member';

-- Add a check constraint to enforce allowed roles
ALTER TABLE users
    ADD CONSTRAINT users_role_check CHECK (role IN ('member', 'admin'));