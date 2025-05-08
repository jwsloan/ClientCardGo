-- 003_create_profile_insights.sql

CREATE TABLE IF NOT EXISTS profile_insights (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    primary_services TEXT NOT NULL,
    client_interactions TEXT NOT NULL,
    pain_points TEXT NOT NULL,
    feature_ideas TEXT,
    submitted_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);