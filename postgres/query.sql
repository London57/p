-- name: GetAllP :many
select * from p;

-- name: GetOneP :one
SELECT * FROM p
WHERE id = $1;

-- name: UpdateOneP :one
UPDATE p set
    name = COALESCE(sqlc.narg(name), name),
    num = COALESCE(sqlc.narg(num), num)
WHERE id = sqlc.arg(id)
RETURNING *;