-- 密码表（若使用 GORM 自动迁移，此脚本可选）
-- PostgreSQL 下数据库已由 POSTGRES_DB 创建，无需 CREATE DATABASE / USE

CREATE TABLE IF NOT EXISTS passwords (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    user_id BIGINT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN passwords.title IS '密码标题';
COMMENT ON COLUMN passwords.description IS '密码描述';
COMMENT ON COLUMN passwords.user_id IS '用户ID';
COMMENT ON COLUMN passwords.password IS '加密后的密码';

CREATE INDEX IF NOT EXISTS idx_passwords_user_id ON passwords(user_id);
CREATE INDEX IF NOT EXISTS idx_passwords_created_at ON passwords(created_at);

-- 自动更新 updated_at 的触发器
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trigger_passwords_updated_at ON passwords;
CREATE TRIGGER trigger_passwords_updated_at
    BEFORE UPDATE ON passwords
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
