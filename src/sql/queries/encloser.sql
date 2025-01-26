-- name: GetEncloserId :one
select
    id,
    controller_id,
    serial_number,
    controller_types_id,
    manufacturers_id
from
    enclosure
where
    id = $1;

-- name: GetEncloserBySN :one
select
    id,
    controller_id,
    serial_number,
    controller_types_id,
    manufacturers_id
from
    enclosure
where
    serial_number = $1;

-- name: CreateEncloser :one
insert into
    enclosure (
        controller_id,
        serial_number,
        controller_types_id,
        manufacturers_id
    )
values
    ($1, $2, $3, $4)
on conflict (serial_number) do nothing
returning
    *;
