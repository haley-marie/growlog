-- +goose Up
INSERT INTO events VALUES 
	(1, 'Water'),
	(2, 'Fertilize'),
	(3, 'Weed'),
	(4, 'Prune'),
	(5, 'Harvest');

SELECT setval('events_id_seq', (SELECT MAX(id) FROM events)); 

-- +goose Down
DELETE FROM events WHERE task_name IN ('Water', 'Fertilize', 'Weed', 'Prune', 'Harvest');

SELECT setval('events_id_seq', 1);
