-- +goose Up
CREATE TABLE grow_log (
	id SERIAL PRIMARY KEY,
	plant_id BIGINT REFERENCES plants(id),
	plant_type TEXT REFERENCES plants(plant_type),
	care_event_id BIGINT REFERENCES events(id),
	care_event_name TEXT REFERENCES events(task_name),
	time_of_care TIMESTAMP,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE grow_log;
