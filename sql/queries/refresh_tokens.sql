-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, user_id, expires_at)
VALUES ($1, $2, $3)
RETURNING token, created_at, updated_at, user_id, expires_at, revoked_at;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW(), updated_at = NOW()
WHERE token = $1;

-- name: GetUserFromRefreshToken :one
SELECT u.id, u.created_at, u.updated_at, u.email
FROM refresh_tokens rt
JOIN users u ON u.id = rt.user_id
WHERE rt.token = $1
  AND rt.revoked_at IS NULL
  AND rt.expires_at > NOW();
