-- Website table (optional if using GORM auto migration)
-- PostgreSQL database is already created by POSTGRES_DB, no need for CREATE DATABASE / USE

CREATE TABLE IF NOT EXISTS websites (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    icon VARCHAR(500),
    description TEXT,
    url VARCHAR(500) NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN websites.title IS 'Website title';
COMMENT ON COLUMN websites.icon IS 'Website icon URL or base64';
COMMENT ON COLUMN websites.description IS 'Website description';
COMMENT ON COLUMN websites.url IS 'Website URL';
COMMENT ON COLUMN websites.user_id IS 'User ID';

CREATE INDEX IF NOT EXISTS idx_websites_user_id ON websites(user_id);
CREATE INDEX IF NOT EXISTS idx_websites_created_at ON websites(created_at);

-- Auto update updated_at trigger
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_websites_updated_at ON websites;
CREATE TRIGGER trigger_websites_updated_at
    BEFORE UPDATE ON websites
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
