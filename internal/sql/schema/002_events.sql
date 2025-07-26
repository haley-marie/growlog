-- +goose Up
CREATE TABLE events (
	id SERIAL PRIMARY KEY,
	task_name TEXT UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL,
	last_performed_at TIMESTAMP,
	updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE events;
