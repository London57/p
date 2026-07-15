-- name: GetAllP :many
SELECT * FROM p;

-- name: GetOneP :one
SELECT * FROM p
WHERE id = $1;

-- name: InsertOneP :one
INSERT INTO p (name, num)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateOneP :one
UPDATE p SET
    name = COALESCE(sqlc.narg(name), name),
    num = COALESCE(sqlc.narg(num), num)
WHERE id = sqlc.arg(id)
RETURNING *;