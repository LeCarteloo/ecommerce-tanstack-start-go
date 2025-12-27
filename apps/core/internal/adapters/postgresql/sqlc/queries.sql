-- name: GetUserByID :one
SELECT
    id,
    username,
    email,
    role,
    created_at
FROM
    users
WHERE
    id = $1;
