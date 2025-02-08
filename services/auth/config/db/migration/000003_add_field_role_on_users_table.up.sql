BEGIN;

-- create role_enum type
CREATE TYPE role_enum AS ENUM ('user', 'admin');

-- add role column to users table
ALTER TABLE users ADD COLUMN role role_enum DEFAULT 'user';

COMMIT;