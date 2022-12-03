-- name: Getcommit :one
SELECT * FROM commits
WHERE commit_id = $1 LIMIT 1;

-- name: Listcommits :many
SELECT * FROM commits
ORDER BY title;

-- name: Createcommit :one
INSERT INTO commits(
  title, supervisor_id,task_id,comment,category,user_id
) VALUES (
  $1, $2, $3, $4 , $5, $6
)
RETURNING *;

-- name: Deletecommit :exec
DELETE FROM commits
WHERE commit_id = $1;


-- name: Updatecommit :exec
UPDATE commits
set title= $2,
comment = $3
WHERE commit_id = $1;