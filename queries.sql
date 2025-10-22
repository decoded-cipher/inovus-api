-- users
-- name: get-all-users
-- Get all users with pagination support
SELECT * FROM users ORDER BY id LIMIT $1 OFFSET $2;

-- name: get-users-count
-- Get total count of users
SELECT COUNT(*) FROM users;

-- name: get-user
-- Get a single user by id, username or email
SELECT * FROM users WHERE
    CASE
        WHEN $1 > 0 THEN id = $1
        WHEN $2 != '' THEN username = $2
        WHEN $3 != '' THEN email = $3
    END;

-- name: create-user
-- Create a new user
INSERT INTO users (username, email, password_hash) 
VALUES ($1, $2, $3) 
RETURNING *;

-- name: update-user
-- Update an existing user
UPDATE users 
SET username = $1, email = $2, password_hash = $3, updated_at = CURRENT_TIMESTAMP 
WHERE id = $4 
RETURNING *;

-- name: delete-user
-- Delete a user by id
DELETE FROM users WHERE id = $1;
