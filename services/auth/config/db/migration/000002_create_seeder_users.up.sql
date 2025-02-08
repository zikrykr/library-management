-- Ensure the UUID extension is enabled (Run this once)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Insert users with UUID-based IDs
INSERT INTO users (id, full_name, email, password_hash, created_at, updated_at) VALUES
    (uuid_generate_v4(), 'Admin User', 'admin@example.com', '$2a$12$zM4GpAPX5n7BYJl9h2e.1eIhbDCT.y7/Ky5JKz0jlFTdugyC7E.GS', NOW(), NOW()),
    (uuid_generate_v4(), 'John Doe', 'john@example.com', '$2a$12$zM4GpAPX5n7BYJl9h2e.1eIhbDCT.y7/Ky5JKz0jlFTdugyC7E.GS', NOW(), NOW()),
    (uuid_generate_v4(), 'Jane Doe', 'jane@example.com', '$2a$12$zM4GpAPX5n7BYJl9h2e.1eIhbDCT.y7/Ky5JKz0jlFTdugyC7E.GS', NOW(), NOW()),
    (uuid_generate_v4(), 'Library Staff 1', 'librarian1@example.com', '$2a$12$zM4GpAPX5n7BYJl9h2e.1eIhbDCT.y7/Ky5JKz0jlFTdugyC7E.GS', NOW(), NOW()),
    (uuid_generate_v4(), 'Library Staff 2', 'librarian2@example.com', '$2a$12$zM4GpAPX5n7BYJl9h2e.1eIhbDCT.y7/Ky5JKz0jlFTdugyC7E.GS', NOW(), NOW());