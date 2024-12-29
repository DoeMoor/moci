-- name: GetAllControllers :many
SELECT
  c.id,
  ct.name AS controller_type_name,
  ct.io_module_slot_amount AS io_module_socket_amount,
  p.name AS project_name,
  c.description,
  concat(
    pcbv.version_number,
    pcbrev.revision
  ) AS pcb_hw_version,
  c.pcb_version_number,
  c.serial_number,
  m.name AS manufacturer_name,
  c.assembly_date,
  c.mac_address,
  c.sim_number,
  mpt.name AS mini_pcie_modules_name,
  mpcie.serial_number AS mini_pcie_serial_number,
  m2t.name AS m2_module_name,
  m2t.module_type_number AS m2_module_type_number,
  c.display_adapters_id,
  c.article_number,
  c.qr_code,
  ctc.can_1,
  ctc.can_2,
  ctc.can_3,
  ctc.can_4,
  c.usb,
  c.serial,
  c.manufacturer_qr_code,
  c.order_id,
  c.created_at,
  c.updated_at,
  c.is_deleted,
  cp.slot_pinout_json AS slots_pinout_json
FROM
  controllers c
  LEFT JOIN controller_types ct ON c.controller_types_id = ct.id
  LEFT JOIN projects p ON c.projects_id = p.id
  LEFT JOIN controller_pcb_hw_versions pcbv ON c.controllers_pcb_hw_versions_id = pcbv.id
  LEFT JOIN manufacturers m ON c.manufacturers_id = m.id
  LEFT JOIN mini_pcie_modules mpcie ON c.mini_pcie_modules_id = mpcie.id
  LEFT JOIN m2_modules_types m2t ON c.m2_modules_types_id = m2t.id
  LEFT JOIN can_termination_confs ctc ON c.can_termination_confs_id = ctc.id
  LEFT JOIN controller_pcb_hw_versions_rev pcbrev ON pcbv.revision = pcbrev.revision
  LEFT JOIN mini_pcie_modules_type mpt ON mpcie.mini_pcie_modules_type_id = mpt.id
  LEFT JOIN controllers_pinout cp 
    ON ct.id = cp.controller_type_id 
    AND pcbv.version_number = cp.controller_pcb_hw_versions_number 
    AND COALESCE(pcbrev.revision, '') = cp.controller_pcb_hw_versions_revision;

-- name: GetControllerById :one
SELECT
  c.id,
  ct.name AS controller_type_name,
  ct.io_module_slot_amount AS io_module_socket_amount,
  p.name AS project_name,
  c.description,
  concat(
    pcbv.version_number,
    pcbrev.revision
  ) AS pcb_hw_version,
  c.pcb_version_number,
  c.serial_number,
  m.name AS manufacturer_name,
  c.assembly_date,
  c.mac_address,
  c.sim_number,
  mpt.name AS mini_pcie_modules_name,
  mpcie.serial_number AS mini_pcie_serial_number,
  m2t.name AS m2_module_name,
  m2t.module_type_number AS m2_module_type_number,
  c.display_adapters_id,
  c.article_number,
  c.qr_code,
  ctc.can_1,
  ctc.can_2,
  ctc.can_3,
  ctc.can_4,
  c.usb,
  c.serial,
  c.manufacturer_qr_code,
  c.order_id,
  c.created_at,
  c.updated_at,
  c.is_deleted,
  cp.slot_pinout_json AS slots_pinout_json
FROM
  controllers c
  LEFT JOIN controller_types ct ON c.controller_types_id = ct.id
  LEFT JOIN projects p ON c.projects_id = p.id
  LEFT JOIN controller_pcb_hw_versions pcbv ON c.controllers_pcb_hw_versions_id = pcbv.id
  LEFT JOIN manufacturers m ON c.manufacturers_id = m.id
  LEFT JOIN mini_pcie_modules mpcie ON c.mini_pcie_modules_id = mpcie.id
  LEFT JOIN m2_modules_types m2t ON c.m2_modules_types_id = m2t.id
  LEFT JOIN can_termination_confs ctc ON c.can_termination_confs_id = ctc.id
  LEFT JOIN controller_pcb_hw_versions_rev pcbrev ON pcbv.revision = pcbrev.revision
  LEFT JOIN mini_pcie_modules_type mpt ON mpcie.mini_pcie_modules_type_id = mpt.id
  LEFT JOIN controllers_pinout cp 
    ON ct.id = cp.controller_type_id 
    AND pcbv.version_number = cp.controller_pcb_hw_versions_number 
    AND COALESCE(pcbrev.revision, '') = cp.controller_pcb_hw_versions_revision
WHERE
  c.id = $1;