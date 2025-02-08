BEGIN;

-- drop role column from users table
ALTER TABLE users DROP COLUMN role;

-- drop role_enum type
DROP TYPE role_enum;

COMMIT;