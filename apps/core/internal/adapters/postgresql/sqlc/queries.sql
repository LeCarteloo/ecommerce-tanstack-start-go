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

-- name: RegisterUser :one
INSERT INTO users
    (username, email, password, role)
VALUES
    ($1, $2, $3, 'user')
RETURNING
    id,
    username,
    email,
    role,
    created_at;
