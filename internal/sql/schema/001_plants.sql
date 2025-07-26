-- +goose Up
CREATE TABLE plants (
	id SERIAL PRIMARY KEY,
	plant_type TEXT UNIQUE NOT NULL,
	planted_at TIMESTAMP,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE plants;
