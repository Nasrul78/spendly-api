-- name: CreateCategory :one
INSERT INTO categories (user_id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetCategoryByID :one
SELECT * FROM categories
WHERE id = $1 AND user_id = $2;

-- name: GetCategoriesByUserID :many
SELECT * FROM categories
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateCategory :one
UPDATE categories
SET name = $1, updated_at = NOW()
WHERE id = $2 AND user_id = $3
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1 AND user_id = $2;
