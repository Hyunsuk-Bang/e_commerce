-- name: GetItem :one
SELECT * FROM item
WHERE id = $1 LIMIT 1;

-- name: ListItems :many
SELECT * FROM item
ORDER BY name;
