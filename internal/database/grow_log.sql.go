// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: grow_log.sql

package database

import (
	"context"
	"database/sql"
)

const createLogEntry = `-- name: CreateLogEntry :one
INSERT INTO grow_log(plant_id, plant_type, care_event_id, care_event_name, time_of_care)
SELECT
		p.id,
		p.plant_type,
		e.id,
		e.task_name,
		$3
FROM plants p, events e
WHERE p.id = $1 AND e.id = $2
RETURNING id, plant_id, plant_type, care_event_id, care_event_name, time_of_care, created_at, updated_at
`

type CreateLogEntryParams struct {
	ID         int32
	ID_2       int32
	TimeOfCare sql.NullTime
}

func (q *Queries) CreateLogEntry(ctx context.Context, arg CreateLogEntryParams) (GrowLog, error) {
	row := q.db.QueryRowContext(ctx, createLogEntry, arg.ID, arg.ID_2, arg.TimeOfCare)
	var i GrowLog
	err := row.Scan(
		&i.ID,
		&i.PlantID,
		&i.PlantType,
		&i.CareEventID,
		&i.CareEventName,
		&i.TimeOfCare,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteLogById = `-- name: DeleteLogById :exec
DELETE FROM grow_log * WHERE id = $1
`

func (q *Queries) DeleteLogById(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteLogById, id)
	return err
}

const getLogEntryById = `-- name: GetLogEntryById :one
SELECT id, plant_id, plant_type, care_event_id, care_event_name, time_of_care, created_at, updated_at FROM grow_log WHERE id = $1
`

func (q *Queries) GetLogEntryById(ctx context.Context, id int32) (GrowLog, error) {
	row := q.db.QueryRowContext(ctx, getLogEntryById, id)
	var i GrowLog
	err := row.Scan(
		&i.ID,
		&i.PlantID,
		&i.PlantType,
		&i.CareEventID,
		&i.CareEventName,
		&i.TimeOfCare,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLogsForCareEvent = `-- name: GetLogsForCareEvent :many
SELECT id, plant_id, plant_type, care_event_id, care_event_name, time_of_care, created_at, updated_at FROM grow_log WHERE care_event_id = $1
`

func (q *Queries) GetLogsForCareEvent(ctx context.Context, careEventID sql.NullInt64) ([]GrowLog, error) {
	rows, err := q.db.QueryContext(ctx, getLogsForCareEvent, careEventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GrowLog
	for rows.Next() {
		var i GrowLog
		if err := rows.Scan(
			&i.ID,
			&i.PlantID,
			&i.PlantType,
			&i.CareEventID,
			&i.CareEventName,
			&i.TimeOfCare,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogsForPlant = `-- name: GetLogsForPlant :many
SELECT id, plant_id, plant_type, care_event_id, care_event_name, time_of_care, created_at, updated_at FROM grow_log WHERE plant_id = $1
`

func (q *Queries) GetLogsForPlant(ctx context.Context, plantID sql.NullInt64) ([]GrowLog, error) {
	rows, err := q.db.QueryContext(ctx, getLogsForPlant, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GrowLog
	for rows.Next() {
		var i GrowLog
		if err := rows.Scan(
			&i.ID,
			&i.PlantID,
			&i.PlantType,
			&i.CareEventID,
			&i.CareEventName,
			&i.TimeOfCare,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLogsForPlantType = `-- name: GetLogsForPlantType :many
SELECT id, plant_id, plant_type, care_event_id, care_event_name, time_of_care, created_at, updated_at FROM grow_log WHERE plant_type = $1
`

func (q *Queries) GetLogsForPlantType(ctx context.Context, plantType sql.NullString) ([]GrowLog, error) {
	rows, err := q.db.QueryContext(ctx, getLogsForPlantType, plantType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GrowLog
	for rows.Next() {
		var i GrowLog
		if err := rows.Scan(
			&i.ID,
			&i.PlantID,
			&i.PlantType,
			&i.CareEventID,
			&i.CareEventName,
			&i.TimeOfCare,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllLogs = `-- name: ListAllLogs :many
SELECT id, plant_id, plant_type, care_event_id, care_event_name, time_of_care, created_at, updated_at FROM grow_log
`

func (q *Queries) ListAllLogs(ctx context.Context) ([]GrowLog, error) {
	rows, err := q.db.QueryContext(ctx, listAllLogs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GrowLog
	for rows.Next() {
		var i GrowLog
		if err := rows.Scan(
			&i.ID,
			&i.PlantID,
			&i.PlantType,
			&i.CareEventID,
			&i.CareEventName,
			&i.TimeOfCare,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const resetLogs = `-- name: ResetLogs :exec
DELETE FROM grow_log *
`

func (q *Queries) ResetLogs(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, resetLogs)
	return err
}
