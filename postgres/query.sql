-- name: InsertUser :one
INSERT INTO users (
    id, auth_id, email, name, role, image_url, created_at, updated_at, deleted_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9
         )
    RETURNING *;

-- name: GetAllUsers :many
SELECT * from users;
-- name: GetAllCategories :many
SELECT * from categories;