-- name: CreateUserWatch :one
INSERT INTO user_watches (
  user_id,
  watch_reason
) VALUES (
  $1, $2
) RETURNING *;


-- name: GetUserWatch :one
SELECT * FROM user_watches
WHERE user_id = $1 LIMIT 1;


-- name: GetUserWatchForUpdate :one
SELECT * FROM user_watches
WHERE user_id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: ListUsersOnWatch :many
SELECT * FROM user_watches
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateUserWatch :one
UPDATE user_watches
SET watch_reason = $2
WHERE user_id = $1
RETURNING *;


-- name: DeleteUserFromWatch :exec
DELETE FROM user_watches
WHERE user_id = $1;
