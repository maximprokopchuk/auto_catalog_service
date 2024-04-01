-- name: GetCarModel :one
SELECT * FROM car_model
WHERE id = $1 LIMIT 1;

-- name: ListCarModels :many
SELECT * FROM car_model
ORDER BY name;

-- name: CreateCarModel :one
INSERT INTO car_model (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: DeleteCarModel :exec
DELETE FROM car_model
WHERE id = $1;

-- name: CreateComponent :one
INSERT INTO component (
  name, car_model_id, parent_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTopLevelComponentsByCarModel :many
SELECT * FROM component
WHERE car_model_id = $1 LIMIT 1;


-- name: GetChildComponentsByComponent :many
SELECT * FROM component
WHERE parent_id = $1 LIMIT 1;

-- name: DeleteComponent :exec
DELETE FROM component
WHERE id = $1;

-- name: UpdateComponent :one
UPDATE component
SET name = $1
RETURNING *;
