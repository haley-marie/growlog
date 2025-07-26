-- name: AddPlant :one
INSERT INTO plants(plant_type, planted_at)
VALUES (
	$1,
	$2
	)
RETURNING *;

-- name: ResetPlants :exec
DELETE FROM plants *;

-- name: GetPlantById :one
SELECT * FROM plants WHERE id = $1;

-- name: GetPlantsByType :many
SELECT * FROM plants
WHERE plant_type = $1
ORDER BY planted_at DESC NULLS LAST;

-- name: RemovePlantById :exec
DELETE FROM plants * WHERE id = $1;

-- name: GetAllPlants :many
SELECT * FROM plants;
