-- Remove foreign key from urls table
ALTER TABLE urls DROP COLUMN IF EXISTS user_id;

-- Drop users table
DROP TABLE IF EXISTS users;
