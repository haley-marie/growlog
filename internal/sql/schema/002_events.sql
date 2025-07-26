-- +goose Up
CREATE TABLE events (
	id SERIAL PRIMARY KEY,
	task_name TEXT UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE events;
