-- name: getAllIoModules :many
SELECT 
	io.id,
	iomt.module_type_name,
	iomt.module_type_number,
	io.manufacturer_bottom_qr_code,
	io.manufacturer_top_qr_code,
	io.rma_number,
	io.order_id,
	iohw.hw_version,
	orders.invoice_number,
	customers."name",
	conts.controllers_id,
	conts.slot_number as "controller slot number",
	iomt.pinout_json
FROM 
	io_modules io
LEFT join io_module_types iomt ON iomt.id = io.io_module_types_id
LEFT join io_module_hw_versions iohw ON iohw.id = io.io_module_hw_versions_id
LEFT join orders on orders.id = io.order_id
LEFT join customers ON customers.id = orders.customers_id
LEFT join controller_slot conts ON conts.io_modules_id = io.id;

-- name: getIoModuleById :one
SELECT 
	io.id,
	iomt.module_type_name,
	iomt.module_type_number,
	io.manufacturer_bottom_qr_code,
	io.manufacturer_top_qr_code,
	io.rma_number,
	io.order_id,
	iohw.hw_version,
	orders.invoice_number,
	customers."name",
	conts.controllers_id,
	conts.slot_number as "controller slot number",
	iomt.pinout_json
FROM 
	io_modules io
LEFT join io_module_types iomt ON iomt.id = io.io_module_types_id
LEFT join io_module_hw_versions iohw ON iohw.id = io.io_module_hw_versions_id
LEFT join orders on orders.id = io.order_id
LEFT join customers ON customers.id = orders.customers_id
LEFT join controller_slot conts ON conts.io_modules_id = io.id
WHERE io.id = $1;