-- name: GetAllManufacturers :many
SELECT id, name
FROM manufacturers;

-- name: CreateManufacturer :one
INSERT INTO manufacturers (name)
VALUES ($1)
RETURNING *;

-- name: GetManufacturerById :one
SELECT *
FROM manufacturers
WHERE id = $1;