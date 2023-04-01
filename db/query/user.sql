-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    full_name,
    phone_number,
    gender,
    type,
    avatar_url,
    hashed_password
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;
