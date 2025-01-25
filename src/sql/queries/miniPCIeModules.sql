-- name: GetAllMiniPCIeModulesType :many
SELECT id, name, module_type_number
FROM mini_pcie_modules_type;

-- name: GetMiniPCIeModuleById :one
SELECT mpt.name as name,
       m.*
FROM mini_pcie_modules m
         LEFT JOIN mini_pcie_modules_type mpt ON mpt.id = m.mini_pcie_modules_type_id
WHERE m.id = $1;

-- name: GetMiniPCIeModuleBySN :one
select id, mini_pcie_modules_type_id, serial_number
from mini_pcie_modules
where serial_number = $1;

-- name: GetAllMiniPCIeModules :many
SELECT mpt.name as name,
       m.*
FROM mini_pcie_modules m
         LEFT JOIN mini_pcie_modules_type mpt ON mpt.id = m.mini_pcie_modules_type_id;

-- name: CreateMiniPCIeModule :one
INSERT INTO mini_pcie_modules (mini_pcie_modules_type_id, serial_number)
VALUES ($1, $2)
RETURNING *;

