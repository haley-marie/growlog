-- name: CreateLogEntry :one
INSERT INTO grow_log(plant_id, plant_type, care_event_id, care_event_name, time_of_care)
SELECT
		p.id,
		p.plant_type,
		e.id,
		e.task_name,
		$3
FROM plants p, events e
WHERE p.id = $1 AND e.id = $2
RETURNING *;

-- name: GetLogEntryById :one
SELECT * FROM grow_log WHERE id = $1;

-- name: ListAllLogs :many
SELECT * FROM grow_log;

-- name: GetLogsForPlant :many
SELECT * FROM grow_log WHERE plant_id = $1;

-- name: GetLogsForPlantType :many
SELECT * FROM grow_log WHERE plant_type = $1;

-- name: GetLogsForCareEvent :many
SELECT * FROM grow_log WHERE care_event_id = $1;

-- name: ResetLogs :exec
DELETE FROM grow_log *;

-- name: DeleteLogById :exec
DELETE FROM grow_log * WHERE id = $1;

