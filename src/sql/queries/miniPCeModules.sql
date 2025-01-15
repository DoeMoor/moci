-- name: GetAllMiniPCeModulesType :many
SELECT *
FROM mini_pcie_modules_type;

-- name: GetMiniPCeModuleById :one
SELECT mpt.name as name,
       m.*
FROM mini_pcie_modules m
         LEFT JOIN mini_pcie_modules_type mpt ON mpt.id = m.mini_pcie_modules_type_id
WHERE m.id = $1;

-- name: GetAllMiniPCeModules :many
SELECT mpt.name as name,
       m.*
FROM mini_pcie_modules m
         LEFT JOIN mini_pcie_modules_type mpt ON mpt.id = m.mini_pcie_modules_type_id;

-- name: CreateMiniPCeModule :one
INSERT INTO mini_pcie_modules (mini_pcie_modules_type_id, serial_number)
VALUES ($1, $2)
RETURNING *;

