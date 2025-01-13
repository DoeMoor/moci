CREATE TABLE "controllers" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "controller_types_id" uuid,
  "projects_id" uuid,
  "description" text,
  "controllers_pcb_hw_versions_id" uuid,
  "pcb_version_number" int,
  "serial_number" varchar(20) UNIQUE,
  "manufacturers_id" uuid,
  "assembly_date" timestamp,
  "mac_address" varchar(20),
  "sim_number" varchar(30) UNIQUE,
  "mini_pcie_modules_id" uuid,
  "m2_modules_types_id" uuid,
  "display_adapters_id" uuid,
  "article_number" varchar(10),
  "qr_code" varchar(100),
  "can_termination_confs_id" uuid,
  "usb" bool,
  "serial" bool,
  "manufacturer_qr_code" varchar(50) UNIQUE,
  "order_id" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "controller_types" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar(20) UNIQUE,
  "io_module_slot_amount" int,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "controllers_pinout" (
  "controller_type_id" uuid NOT NULL,
  "controller_pcb_hw_versions_number" varchar(10) DEFAULT '',
  "controller_pcb_hw_versions_revision" varchar(2) DEFAULT '',
  "slot_pinout_json" jsonb,
  PRIMARY KEY ("controller_type_id", "controller_pcb_hw_versions_number", "controller_pcb_hw_versions_revision")
);

CREATE TABLE "controller_notes" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "controllers_id" uuid,
  "body" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "led_daughter_board" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "manufacturer_qr_code" varchar(100) UNIQUE,
  "controllers_id" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "can_termination_confs" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "can_1" bool,
  "can_2" bool,
  "can_3" bool,
  "can_4" bool,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "controller_slot" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "controllers_id" uuid,
  "io_modules_id" uuid,
  "slot_number" int,
  "is_deleted" bool
);

CREATE TABLE "controller_pcb_hw_versions" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "version_number" varchar(10),
  "revision" varchar(2),
  "amount_of_connectors" int,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "controller_pcb_hw_versions_rev" (
  "revision" varchar(2) UNIQUE PRIMARY KEY
);

CREATE TABLE "controller_confs" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "article_number" int UNIQUE,
  "number_of_modules" int,
  "serial" bool,
  "usb" bool,
  "mini_pcie_modules_type_id" uuid,
  "m2_modules_types_id" uuid,
  "can_termination_conf_id" uuid,
  "controller_types_id" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "controller_conf_io_module_type_junction" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "controller_confs" uuid,
  "io_module_types" uuid,
  "slot_number" int,
  "is_deleted" bool
);

CREATE TABLE "io_modules" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "io_module_types_id" uuid,
  "manufacturer_top_qr_code" varchar(50) UNIQUE,
  "manufacturer_bottom_qr_code" varchar(50) UNIQUE,
  "rma_number" varchar(50),
  "io_module_hw_versions_id" uuid,
  "order_id" uuid,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "io_module_hw_versions" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "hw_version" int,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "io_module_types" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "module_type_name" varchar(30),
  "module_type_number" int UNIQUE,
  "pinout_json" jsonb,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "io_module_notes" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "io_modules_id" uuid,
  "body" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "customers" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar(50) UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "customer_notes" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "customers_id" uuid,
  "body" text,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "customer_projects_junction" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "customers_id" uuid,
  "projects_id" uuid,
  "is_deleted" bool
);

CREATE TABLE "orders" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "customers_id" uuid,
  "notes" text,
  "delivery_date" timestamp,
  "invoice_number" varchar(50),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "projects" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar(50) UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "manufacturers" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar(50) UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "m2_modules_types" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar(20) UNIQUE,
  "module_type_number" int,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "display_adapters" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar(20),
  "manufacturer_qr_code" varchar(50) UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "mini_pcie_modules_type" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar(20) UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

CREATE TABLE "mini_pcie_modules" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "mini_pcie_modules_type_id" uuid,
  "serial_number" varchar(20) UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "is_deleted" bool
);

ALTER TABLE "controllers" ADD FOREIGN KEY ("controller_types_id") REFERENCES "controller_types" ("id");

ALTER TABLE "controllers" ADD FOREIGN KEY ("projects_id") REFERENCES "projects" ("id");

ALTER TABLE "controllers" ADD FOREIGN KEY ("controllers_pcb_hw_versions_id") REFERENCES "controller_pcb_hw_versions" ("id");

ALTER TABLE "controllers" ADD FOREIGN KEY ("manufacturers_id") REFERENCES "manufacturers" ("id");

ALTER TABLE "controllers" ADD FOREIGN KEY ("mini_pcie_modules_id") REFERENCES "mini_pcie_modules" ("id");

ALTER TABLE "controllers" ADD FOREIGN KEY ("m2_modules_types_id") REFERENCES "m2_modules_types" ("id");

ALTER TABLE "controllers" ADD FOREIGN KEY ("display_adapters_id") REFERENCES "display_adapters" ("id");

ALTER TABLE "controllers" ADD FOREIGN KEY ("can_termination_confs_id") REFERENCES "can_termination_confs" ("id");

ALTER TABLE "controllers" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "controller_notes" ADD FOREIGN KEY ("controllers_id") REFERENCES "controllers" ("id");

ALTER TABLE "led_daughter_board" ADD FOREIGN KEY ("controllers_id") REFERENCES "controllers" ("id");

ALTER TABLE "controller_slot" ADD FOREIGN KEY ("controllers_id") REFERENCES "controllers" ("id");

ALTER TABLE "controller_slot" ADD FOREIGN KEY ("io_modules_id") REFERENCES "io_modules" ("id");

ALTER TABLE "controller_pcb_hw_versions" ADD FOREIGN KEY ("revision") REFERENCES "controller_pcb_hw_versions_rev" ("revision");

ALTER TABLE "controller_confs" ADD FOREIGN KEY ("mini_pcie_modules_type_id") REFERENCES "mini_pcie_modules_type" ("id");

ALTER TABLE "controller_confs" ADD FOREIGN KEY ("m2_modules_types_id") REFERENCES "m2_modules_types" ("id");

ALTER TABLE "controller_confs" ADD FOREIGN KEY ("can_termination_conf_id") REFERENCES "can_termination_confs" ("id");

ALTER TABLE "controller_confs" ADD FOREIGN KEY ("controller_types_id") REFERENCES "controller_types" ("id");

ALTER TABLE "controller_conf_io_module_type_junction" ADD FOREIGN KEY ("controller_confs") REFERENCES "controller_confs" ("id");

ALTER TABLE "controller_conf_io_module_type_junction" ADD FOREIGN KEY ("io_module_types") REFERENCES "io_module_types" ("id");

ALTER TABLE "io_modules" ADD FOREIGN KEY ("io_module_types_id") REFERENCES "io_module_types" ("id");

ALTER TABLE "io_modules" ADD FOREIGN KEY ("io_module_hw_versions_id") REFERENCES "io_module_hw_versions" ("id");

ALTER TABLE "io_modules" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "io_module_notes" ADD FOREIGN KEY ("io_modules_id") REFERENCES "io_modules" ("id");

ALTER TABLE "customer_notes" ADD FOREIGN KEY ("customers_id") REFERENCES "customers" ("id");

ALTER TABLE "customer_projects_junction" ADD FOREIGN KEY ("customers_id") REFERENCES "customers" ("id");

ALTER TABLE "customer_projects_junction" ADD FOREIGN KEY ("projects_id") REFERENCES "projects" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("customers_id") REFERENCES "customers" ("id");

ALTER TABLE "mini_pcie_modules" ADD FOREIGN KEY ("mini_pcie_modules_type_id") REFERENCES "mini_pcie_modules_type" ("id");
