-- name: GetAllCanTerminationConfig :many
select id, can_1_terminated, can_2_terminated, can_3_terminated, can_4_terminated
from can_termination_confs;

-- name: GetTerminationConfigIdByCanTerminated :one
select id
from can_termination_confs
where can_1_terminated = $1
  and can_2_terminated = $2
  and can_3_terminated = $3
  and can_4_terminated = $4;

-- name: CreateCanTerminationConfig :one
insert into can_termination_confs (can_1_terminated, can_2_terminated, can_3_terminated,can_4_terminated)
VALUES($1,$2,$3,$4)
on conflict (can_1_terminated, can_2_terminated, can_3_terminated, can_4_terminated) do nothing
returning *;