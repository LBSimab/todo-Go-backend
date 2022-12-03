-- name: Gettask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

-- name: Listtasks :many
SELECT * FROM tasks
ORDER BY name;

-- name: Createtask :one
INSERT INTO tasks(
  name, supervisor,category
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: Deletetask :exec
DELETE FROM tasks
WHERE id = $1;


-- name: Updatetask :exec
UPDATE tasks
set name = $2,
supervisor = $3
WHERE id = $1;