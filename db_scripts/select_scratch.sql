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



SELECT *
FROM controller_pcb_hw_versions;

SELECT *
FROM m2_modules_types;

select name, module_type_number as moduleTypeNumber
from m2_modules_types;

select *
from manufacturers;

select *
FROM io_modules;

SELECT *
FROM mini_pcie_modules_type;


SELECT mpt.name as name,
       m.*
FROM mini_pcie_modules m
         LEFT JOIN mini_pcie_modules_type mpt ON mpt.id = m.mini_pcie_modules_type_id;


SELECT *
FROM projects;



SELECT id, controller_types.name
FROM controller_types;


select *
from controllers;

select display_adapters_type_name as type_name, manufacturer_qr_code as qr_code
from display_adapters;

SELECT c.id,
       ct.name                       AS controller_type_name,
       ct.io_module_slot_amount      AS io_module_socket_amount,
       p.name                        AS project_name,
       c.description,
       concat(
               pcbv.version_number,
               pcbrev.revision
       )                             AS pcb_hw_version,
       c.pcb_version_number,
       c.serial_number               as controller_serial_number,
       m.name                        AS manufacturer_name,
       c.assembly_date,
       c.mac_address,
       c.sim_number,
       mpt.name                      AS mini_pcie_modules_name,
       mpt.module_type_number        as mini_pcie_module_type_number,
       mpcie.serial_number           AS mini_pcie_serial_number,
       m2t.name                      AS m2_module_name,
       m2t.module_type_number        AS m2_module_type_number,
       ldb.manufacturer_qr_code      as LED_board,
       da.display_adapters_type_name as display_type,
       da.manufacturer_qr_code       as display_manufacturer_qr_code,
       c.article_number,
       c.qr_code                     as controller_info_qr_code,
       ctc.can_1_terminated,
       ctc.can_2_terminated,
       ctc.can_3_terminated,
       ctc.can_4_terminated,
       c.usb,
       c.serial,
       c.manufacturer_qr_code        as controller_manufacturer_qr_code,
       c.order_id,
       c.created_at,
       c.updated_at,
       c.is_deleted,
       ct.slot_pinout_json
FROM controllers c
         LEFT JOIN controller_types ct ON c.controller_types_id = ct.id
         LEFT JOIN projects p ON c.projects_id = p.id
         LEFT JOIN controller_pcb_hw_versions pcbv ON c.controllers_pcb_hw_versions_id = pcbv.id
         LEFT JOIN manufacturers m ON c.manufacturers_id = m.id
         LEFT JOIN mini_pcie_modules mpcie ON c.mini_pcie_modules_id = mpcie.id
         LEFT JOIN m2_modules_types m2t ON c.m2_modules_types_id = m2t.id
         left join display_adapters da on c.display_adapters_id = da.id
         LEFT JOIN can_termination_confs ctc ON c.can_termination_confs_id = ctc.id
         LEFT JOIN controller_pcb_hw_versions_rev pcbrev ON pcbv.revision = pcbrev.revision
         LEFT JOIN mini_pcie_modules_type mpt ON mpcie.mini_pcie_modules_type_id = mpt.id
         left join led_daughter_board ldb on c.id = ldb.controllers_id;


---
select *
from led_daughter_board;


---
SELECT m.id,
       io_m_t.module_type_name,
       io_m_t.module_type_number,
       m.manufacturer_top_qr_code,
       m.manufacturer_bottom_qr_code,
       m.rma_number,
       io_module_hw_versions.hw_version,
       io_m_t.pinout_json
FROM controller_slot
         left join io_modules m ON controller_slot.io_modules_id = m.id
         left join io_module_types io_m_t ON m.io_module_types_id = io_m_t.id
         LEFT join io_module_hw_versions ON m.io_module_hw_versions_id = io_module_hw_versions.id
WHERE controller_slot.controllers_id = (select id
                                        from controllers
                                        limit 1);
---
SELECT m.id,
       io_m_t.module_type_name,
       io_m_t.module_type_number,
       m.manufacturer_top_qr_code,
       m.manufacturer_bottom_qr_code,
       m.rma_number,
       io_module_hw_versions.hw_version,
       io_m_t.pinout_json
FROM controller_slot
         left join io_modules m ON controller_slot.io_modules_id = m.id
         left join io_module_types io_m_t ON m.io_module_types_id = io_m_t.id
         LEFT join io_module_hw_versions ON m.io_module_hw_versions_id = io_module_hw_versions.id
WHERE controller_slot.controllers_id = (select id
                                        from controllers
                                        limit 1 OFFSET 1);
