-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, api_key ) VALUES ($1, $2, $3, $4, gen_random_uuid()) RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1 LIMIT 1;