-- name: CreateMeal :one
INSERT INTO meals (
    username,
    name,
    description,
    calories,
    type
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetMeal :one
SELECT * FROM meals
WHERE id = $1 LIMIT 1;

-- name: ListMeals :many
SELECT * FROM meals
WHERE 
    username = sqlc.arg(username)
    AND (
        (created_at >= sqlc.narg('from_date')::date OR sqlc.narg('from_date') IS NULL) AND
        (created_at <= sqlc.narg('to_date')::date OR sqlc.narg('to_date') IS NULL)
    )
ORDER BY created_at
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');