-- name: GetAllM2Modules :many
SELECT *
FROM m2_modules_types;

-- name: GetAllM2ModulesTypes :many
select id, name, module_type_number
from m2_modules_types;

-- name: CreateM2Module :one
INSERT INTO m2_modules_types ("name", module_type_number)
VALUES ($1, $2)
returning *;