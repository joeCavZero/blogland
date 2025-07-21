-- name: CreateUser :exec
INSERT INTO users (
    email, 
    password, 
    role
) VALUES (
    $1,
    $2,
    $3
);

-- name: GetUserByEmailAndPassword :one
SELECT * FROM users WHERE email = $1 AND password = $2;