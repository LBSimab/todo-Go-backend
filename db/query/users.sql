-- name: Getuser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: Listusers :many
SELECT * FROM users
ORDER BY full_name;

-- name: Createuser :one
INSERT INTO users(
  full_name, supervisor
) VALUES (
  $1, $2
)
RETURNING *;

-- name: Deleteuser :exec
DELETE FROM users
WHERE id = $1;


-- name: Updateuser :exec
UPDATE users
set full_name = $2,
supervisor = $3
WHERE id = $1;