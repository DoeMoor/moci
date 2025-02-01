-- name: GetAllIoModules :many
SELECT io.id,
       iomt.module_type_name,
       iomt.module_type_number,
       io.manufacturer_bottom_qr_code,
       io.manufacturer_top_qr_code,
       io.rma_number,
       io.order_id,
       iohw.hw_version,
       orders.invoice_number,
       customers."name"  as "customer name",
       conts.controllers_id,
       conts.slot_number as "controller slot number",
       iomt.pinout_json
FROM io_modules io
         LEFT join io_module_types iomt ON iomt.id = io.io_module_types_id
         LEFT join io_module_hw_versions iohw ON iohw.id = io.io_module_hw_versions_id
         LEFT join orders on orders.id = io.order_id
         LEFT join customers ON customers.id = orders.customers_id
         LEFT join controller_slot conts ON conts.io_modules_id = io.id;

-- name: GetIoModuleById :one
SELECT io.id,
       iomt.module_type_name,
       iomt.module_type_number,
       io.manufacturer_bottom_qr_code,
       io.manufacturer_top_qr_code,
       io.rma_number,
       io.order_id,
       iohw.hw_version,
       orders.invoice_number,
       customers."name"  as "customer name",
       conts.controllers_id,
       conts.slot_number as "controller slot number",
       iomt.pinout_json
FROM io_modules io
         LEFT join io_module_types iomt ON iomt.id = io.io_module_types_id
         LEFT join io_module_hw_versions iohw ON iohw.id = io.io_module_hw_versions_id
         LEFT join orders on orders.id = io.order_id
         LEFT join customers ON customers.id = orders.customers_id
         LEFT join controller_slot conts ON conts.io_modules_id = io.id
WHERE io.id = $1;

-- name: CreateIoModule :one
INSERT INTO io_modules (io_module_types_id,
                        manufacturer_top_qr_code,
                        manufacturer_bottom_qr_code,
                        rma_number,
                        io_module_hw_versions_id,
                        order_id)
VALUES ($1, $2, $3, $4, $5, $6)
returning
    *;

-- name: CreateIoModuleHWVersion :one
INSERT INTO io_module_hw_versions (hw_version)
VALUES ($1)
returning
    *;

-- name: GetAllIoModuleTypes :many
select id, module_type_name, module_type_number
from io_module_types;

-- name: CreateIoModuleType :one
INSERT INTO io_module_types (module_type_name, module_type_number, pinout_json)
VALUES ($1, $2, $3)
returning
    *;