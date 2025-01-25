-- name: GetLedBoardByQR :one
select id, manufacturer_qr_code, controllers_id
from led_daughter_board
where manufacturer_qr_code = $1;

-- name: CreateLedBoard :one
insert into led_daughter_board (manufacturer_qr_code, controllers_id)
values ($1,$2)
on conflict (manufacturer_qr_code) do nothing
returning id;