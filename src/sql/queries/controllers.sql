-- name: GetAllControllers :many
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
       enc.serial_number             as enclosure_serial_number,
       encman.name                   as enclosure_manufacturer,
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
         left join enclosure enc on c.id = enc.controller_id
         left join manufacturers encman on enc.manufacturers_id = encman.id
         LEFT JOIN mini_pcie_modules mpcie ON c.mini_pcie_modules_id = mpcie.id
         LEFT JOIN m2_modules_types m2t ON c.m2_modules_types_id = m2t.id
         left join display_adapters da on c.display_adapters_id = da.id
         LEFT JOIN can_termination_confs ctc ON c.can_termination_confs_id = ctc.id
         LEFT JOIN controller_pcb_hw_versions_rev pcbrev ON pcbv.revision = pcbrev.revision
         LEFT JOIN mini_pcie_modules_type mpt ON mpcie.mini_pcie_modules_type_id = mpt.id
         left join led_daughter_board ldb on c.id = ldb.controllers_id;

-- name: GetControllerById :one
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
       enc.serial_number             as enclosure_serial_number,
       encman.name                   as enclosure_manufacturer,
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
         left join enclosure enc on c.id = enc.controller_id
         left join manufacturers encman on enc.manufacturers_id = encman.id
         LEFT JOIN mini_pcie_modules mpcie ON c.mini_pcie_modules_id = mpcie.id
         LEFT JOIN m2_modules_types m2t ON c.m2_modules_types_id = m2t.id
         left join display_adapters da on c.display_adapters_id = da.id
         LEFT JOIN can_termination_confs ctc ON c.can_termination_confs_id = ctc.id
         LEFT JOIN controller_pcb_hw_versions_rev pcbrev ON pcbv.revision = pcbrev.revision
         LEFT JOIN mini_pcie_modules_type mpt ON mpcie.mini_pcie_modules_type_id = mpt.id
         left join led_daughter_board ldb on c.id = ldb.controllers_id
WHERE c.id = $1;

-- name: GetALLControllerTypes :many
SELECT id, controller_types.name
FROM controller_types;

-- name: GetAllControllersPcbHwVersions :many
select id, version_number, revision
from controller_pcb_hw_versions;

-- name: GetControllerTypesPinout :one
select slot_pinout_json
from controller_types
where id = $1;