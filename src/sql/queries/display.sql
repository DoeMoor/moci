-- name: GetAllDisplayTypes :many
select display_adapters_type.name as name
from display_adapters_type;

-- name: CreateDisplayType :one
insert into display_adapters_type (name)
values ($1)
returning *;


-- name: GetAllDisplayAdapters :many
select id, display_adapters_type_name, manufacturer_qr_code
from display_adapters;

-- name: GetDisplayAdapterByQR :one
select id, display_adapters_type_name, manufacturer_qr_code
from display_adapters
where manufacturer_qr_code = $1;

-- name: CreateDisplayAdapter :one
INSERT INTO display_adapters (display_adapters_type_name, manufacturer_qr_code)
VALUES ($1, $2)
returning *;