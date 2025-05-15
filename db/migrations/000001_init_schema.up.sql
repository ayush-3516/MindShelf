-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create links table
CREATE TABLE links (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    url TEXT NOT NULL,
    description TEXT,
    domain VARCHAR(255),
    favicon_url TEXT,
    is_favorite BOOLEAN DEFAULT FALSE,
    archived BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT url_user_unique UNIQUE (url, user_id)
);

-- Create index on user_id for faster link retrieval
CREATE INDEX idx_links_user_id ON links(user_id);
CREATE INDEX idx_links_domain ON links(domain);
CREATE INDEX idx_links_is_favorite ON links(is_favorite);
CREATE INDEX idx_links_archived ON links(archived);

-- Create tags table
CREATE TABLE tags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT name_user_unique UNIQUE (name, user_id)
);

-- Create index on user_id for faster tag retrieval
CREATE INDEX idx_tags_user_id ON tags(user_id);

-- Create link_tags join table
CREATE TABLE link_tags (
    link_id UUID NOT NULL REFERENCES links(id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (link_id, tag_id)
);

-- Create index for faster tag-based searches
CREATE INDEX idx_link_tags_link_id ON link_tags(link_id);
CREATE INDEX idx_link_tags_tag_id ON link_tags(tag_id);

-- Create function to update timestamps
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create triggers for updated_at
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

CREATE TRIGGER update_links_updated_at
BEFORE UPDATE ON links
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();

