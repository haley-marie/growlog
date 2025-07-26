-- name: CreateCareEvent :one
INSERT INTO events(task_name)
VALUES (
	$1
	)
RETURNING *;

-- name: ListCareEvents :many
SELECT * FROM events;

-- name: GetEventById :one
SELECT * FROM events WHERE id = $1;
