SELECT
	TABLE_NAME
FROM
	INFORMATION_SCHEMA.TABLES
WHERE
	TABLE_SCHEMA = 'public';

DROP TABLE IF EXISTS "controller_types",
"controllers",
"projects",
"controller_pcb_hw_versions",
"manufacturers",
"mini_pcie_modules",
"mini_pcie_modules_type",
"m2_modules_types",
"display_adapters",
"can_termination_confs",
"orders",
"controller_connector_types",
"controller_connector_number",
"controller_connector_types_pcb_hw_version_revisions_junction",
"controller_pcb_hw_version_revisions",
"controller_notes",
"controller_slot",
"io_modules",
"controller_confs",
"controller_conf_io_module_type_junction",
"io_module_types",
"io_module_hw_versions",
"io_module_notes",
"customers",
"customer_notes",
"customer_projects_junction",
"controller_pcb_hw_versions_rev",
"controller_connector_installed",
"controllers_pinout",
"led_daughter_board" CASCADE;