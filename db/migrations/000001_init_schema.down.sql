-- Drop triggers
DROP TRIGGER IF EXISTS update_links_updated_at ON links;
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop link_tags table
DROP TABLE IF EXISTS link_tags;

-- Drop tags table
DROP TABLE IF EXISTS tags;

-- Drop links table
DROP TABLE IF EXISTS links;

-- Drop users table
DROP TABLE IF EXISTS users;

-- Drop UUID extension
DROP EXTENSION IF EXISTS "uuid-ossp";

