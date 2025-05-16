-- 偶像团体表
CREATE TABLE idol_groups (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    english_name VARCHAR(100),
    nickname VARCHAR(100),
    company VARCHAR(100),
    founded_date DATE,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 偶像基本信息表
CREATE TABLE idols (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    english_name VARCHAR(100),
    nickname VARCHAR(100),
    birthday DATE,
    height INTEGER,
    weight INTEGER,
    blood_type VARCHAR(10),
    zodiac VARCHAR(20),
    group_id BIGINT REFERENCES idol_groups(id),
    group_name VARCHAR(100),
    position VARCHAR(50),
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 偶像社交媒体账号表
CREATE TABLE idol_socials (
    id BIGSERIAL PRIMARY KEY,
    idol_id BIGINT REFERENCES idols(id) ON DELETE CASCADE,
    platform VARCHAR(50) NOT NULL,
    account_id VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL,
    url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(idol_id, platform, account_id)
);

-- 创建索引
CREATE INDEX idx_idols_group_id ON idols(group_id);
CREATE INDEX idx_idol_socials_idol_id ON idol_socials(idol_id);
CREATE INDEX idx_idol_socials_platform ON idol_socials(platform);

-- 创建更新时间触发器
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_idol_groups_updated_at
    BEFORE UPDATE ON idol_groups
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_idols_updated_at
    BEFORE UPDATE ON idols
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_idol_socials_updated_at
    BEFORE UPDATE ON idol_socials
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column(); 