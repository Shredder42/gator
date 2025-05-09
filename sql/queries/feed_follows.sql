-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *
)
SELECT iff.id
, iff.created_at
, iff.updated_at
, iff.user_id
, iff.feed_id
, u.name AS user_name
, f.name AS feed_name
FROM inserted_feed_follows iff
JOIN users u ON iff.user_id = u.id
JOIN feeds f ON iff.feed_id = f.id
;

-- name: GetFeedFollowsForUser :many
SELECT ff.id
, ff.created_at
, ff.updated_at
, ff.user_id
, ff.feed_id
, u.name AS user_name
, f.name AS feed_name
FROM feed_follows ff
JOIN users u ON ff.user_id = u.id
JOIN feeds f ON ff.feed_id = f.id
WHERE ff.user_id = $1
;