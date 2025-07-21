-- name: CreateSessionToken :exec
INSERT INTO session_tokens (
    user_id,
    token
) VALUES (
    $1,
    $2
);

-- name: GetSessionTokenByIDAndToken :one
SELECT * FROM session_tokens WHERE session_tokens.user_id = $1 AND session_tokens.token = $2;

-- name: DeleteOldSessionTokens :exec
DELETE FROM session_tokens WHERE created < NOW() - $1::BIGINT * INTERVAL '1 seconds';