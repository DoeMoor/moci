SELECT 
	io.id,
	iomt.module_type_name,
	iomt.module_type_number,
	io.manufacturer_bottom_qr_code,
	io.manufacturer_top_qr_code,
	io.rma_number,
	io.order_id,
	iohw.hw_version,
	io.order_id,
	iomt.pinout_json
FROM 
	io_modules io
LEFT join io_module_types iomt ON iomt.id = io.io_module_types_id
LEFT join io_module_hw_versions iohw ON iohw.id = io.io_module_hw_versions_id;



SELECT * FROM controller_pcb_hw_versions;

SELECT
  *
FROM
  m2_modules_types;


select * FROM io_modules;

SELECT
  *
FROM
  mini_pcie_modules_type;


SELECT
    mpt.name as rs,
    m.*
FROM
    mini_pcie_modules m
    LEFT JOIN mini_pcie_modules_type mpt ON mpt.id = m.mini_pcie_modules_type_id;


SELECT
  *
FROM
  projects;


SELECT 
	controller_types.name,
	* 
FROM 
	controllers_pinout
LEFT join controller_types ON controllers_pinout.controller_type_id = controller_types.id
;


SELECT
  *
FROM
  controller_types;


select
  *
from
  controllers;


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

---
SELECT 
	m.id,
	io_m_t.module_type_name,
	io_m_t.module_type_number,
	m.manufacturer_top_qr_code,
	m.manufacturer_bottom_qr_code,
	m.rma_number,
	io_module_hw_versions.hw_version,
	orders.invoice_number as order_invoice,
	io_m_t.pinout_json
FROM controller_slot
left join io_modules m ON controller_slot.io_modules_id = m.id
left join io_module_types io_m_t ON m.io_module_types_id = io_m_t.id
LEFT join io_module_hw_versions ON m.io_module_hw_versions_id = io_module_hw_versions.id
LEFT join orders ON m.order_id = orders.id
WHERE controller_slot.controllers_id = (select
  id
from
  controllers limit 1);
---
SELECT 
	m.id,
	io_m_t.module_type_name,
	io_m_t.module_type_number,
	m.manufacturer_top_qr_code,
	m.manufacturer_bottom_qr_code,
	m.rma_number,
	io_module_hw_versions.hw_version,
	orders.invoice_number as order_invoice,
	io_m_t.pinout_json
FROM controller_slot
left join io_modules m ON controller_slot.io_modules_id = m.id
left join io_module_types io_m_t ON m.io_module_types_id = io_m_t.id
LEFT join io_module_hw_versions ON m.io_module_hw_versions_id = io_module_hw_versions.id
LEFT join orders ON m.order_id = orders.id
WHERE controller_slot.controllers_id = (select
  id
from
  controllers limit 1 OFFSET 1);
