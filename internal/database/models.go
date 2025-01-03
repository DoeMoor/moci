// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

type CanTerminationConf struct {
	ID        uuid.UUID    `json:"id"`
	Can1      sql.NullBool `json:"can_1"`
	Can2      sql.NullBool `json:"can_2"`
	Can3      sql.NullBool `json:"can_3"`
	Can4      sql.NullBool `json:"can_4"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	IsDeleted sql.NullBool `json:"is_deleted"`
}

type Controller struct {
	ID                         uuid.UUID      `json:"id"`
	ControllerTypesID          uuid.NullUUID  `json:"controller_types_id"`
	ProjectsID                 uuid.NullUUID  `json:"projects_id"`
	Description                sql.NullString `json:"description"`
	ControllersPcbHwVersionsID uuid.NullUUID  `json:"controllers_pcb_hw_versions_id"`
	PcbVersionNumber           sql.NullInt32  `json:"pcb_version_number"`
	SerialNumber               sql.NullString `json:"serial_number"`
	ManufacturersID            uuid.NullUUID  `json:"manufacturers_id"`
	AssemblyDate               sql.NullTime   `json:"assembly_date"`
	MacAddress                 sql.NullString `json:"mac_address"`
	SimNumber                  sql.NullString `json:"sim_number"`
	MiniPcieModulesID          uuid.NullUUID  `json:"mini_pcie_modules_id"`
	M2ModulesTypesID           uuid.NullUUID  `json:"m2_modules_types_id"`
	DisplayAdaptersID          uuid.NullUUID  `json:"display_adapters_id"`
	ArticleNumber              sql.NullString `json:"article_number"`
	QrCode                     sql.NullString `json:"qr_code"`
	CanTerminationConfsID      uuid.NullUUID  `json:"can_termination_confs_id"`
	Usb                        sql.NullBool   `json:"usb"`
	Serial                     sql.NullBool   `json:"serial"`
	ManufacturerQrCode         sql.NullString `json:"manufacturer_qr_code"`
	OrderID                    uuid.NullUUID  `json:"order_id"`
	CreatedAt                  sql.NullTime   `json:"created_at"`
	UpdatedAt                  sql.NullTime   `json:"updated_at"`
	IsDeleted                  sql.NullBool   `json:"is_deleted"`
}

type ControllerConf struct {
	ID                    uuid.UUID     `json:"id"`
	ArticleNumber         sql.NullInt32 `json:"article_number"`
	NumberOfModules       sql.NullInt32 `json:"number_of_modules"`
	Serial                sql.NullBool  `json:"serial"`
	Usb                   sql.NullBool  `json:"usb"`
	MiniPcieModulesTypeID uuid.NullUUID `json:"mini_pcie_modules_type_id"`
	M2ModulesTypesID      uuid.NullUUID `json:"m2_modules_types_id"`
	CanTerminationConfID  uuid.NullUUID `json:"can_termination_conf_id"`
	ControllerTypesID     uuid.NullUUID `json:"controller_types_id"`
	CreatedAt             sql.NullTime  `json:"created_at"`
	UpdatedAt             sql.NullTime  `json:"updated_at"`
	IsDeleted             sql.NullBool  `json:"is_deleted"`
}

type ControllerConfIoModuleTypeJunction struct {
	ID              uuid.UUID     `json:"id"`
	ControllerConfs uuid.NullUUID `json:"controller_confs"`
	IoModuleTypes   uuid.NullUUID `json:"io_module_types"`
	SlotNumber      sql.NullInt32 `json:"slot_number"`
	IsDeleted       sql.NullBool  `json:"is_deleted"`
}

type ControllerNote struct {
	ID            uuid.UUID      `json:"id"`
	ControllersID uuid.NullUUID  `json:"controllers_id"`
	Body          sql.NullString `json:"body"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
	IsDeleted     sql.NullBool   `json:"is_deleted"`
}

type ControllerPcbHwVersion struct {
	ID                 uuid.UUID      `json:"id"`
	VersionNumber      sql.NullString `json:"version_number"`
	Revision           sql.NullString `json:"revision"`
	AmountOfConnectors sql.NullInt32  `json:"amount_of_connectors"`
	CreatedAt          sql.NullTime   `json:"created_at"`
	UpdatedAt          sql.NullTime   `json:"updated_at"`
	IsDeleted          sql.NullBool   `json:"is_deleted"`
}

type ControllerPcbHwVersionsRev struct {
	Revision string `json:"revision"`
}

type ControllerSlot struct {
	ID            uuid.UUID     `json:"id"`
	ControllersID uuid.NullUUID `json:"controllers_id"`
	IoModulesID   uuid.NullUUID `json:"io_modules_id"`
	SlotNumber    sql.NullInt32 `json:"slot_number"`
	IsDeleted     sql.NullBool  `json:"is_deleted"`
}

type ControllerType struct {
	ID                 uuid.UUID      `json:"id"`
	Name               sql.NullString `json:"name"`
	IoModuleSlotAmount sql.NullInt32  `json:"io_module_slot_amount"`
	CreatedAt          sql.NullTime   `json:"created_at"`
	UpdatedAt          sql.NullTime   `json:"updated_at"`
	IsDeleted          sql.NullBool   `json:"is_deleted"`
}

type ControllersPinout struct {
	ControllerTypeID                uuid.UUID             `json:"controller_type_id"`
	ControllerPcbHwVersionsNumber   string                `json:"controller_pcb_hw_versions_number"`
	ControllerPcbHwVersionsRevision string                `json:"controller_pcb_hw_versions_revision"`
	SlotPinoutJson                  pqtype.NullRawMessage `json:"slot_pinout_json"`
}

type Customer struct {
	ID        uuid.UUID      `json:"id"`
	Name      sql.NullString `json:"name"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	IsDeleted sql.NullBool   `json:"is_deleted"`
}

type CustomerNote struct {
	ID          uuid.UUID      `json:"id"`
	CustomersID uuid.NullUUID  `json:"customers_id"`
	Body        sql.NullString `json:"body"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	IsDeleted   sql.NullBool   `json:"is_deleted"`
}

type CustomerProjectsJunction struct {
	ID          uuid.UUID     `json:"id"`
	CustomersID uuid.NullUUID `json:"customers_id"`
	ProjectsID  uuid.NullUUID `json:"projects_id"`
	IsDeleted   sql.NullBool  `json:"is_deleted"`
}

type DisplayAdapter struct {
	ID                 uuid.UUID      `json:"id"`
	Name               sql.NullString `json:"name"`
	ManufacturerQrCode sql.NullString `json:"manufacturer_qr_code"`
	CreatedAt          sql.NullTime   `json:"created_at"`
	UpdatedAt          sql.NullTime   `json:"updated_at"`
	IsDeleted          sql.NullBool   `json:"is_deleted"`
}

type IoModule struct {
	ID                       uuid.UUID      `json:"id"`
	IoModuleTypesID          uuid.NullUUID  `json:"io_module_types_id"`
	ManufacturerTopQrCode    sql.NullString `json:"manufacturer_top_qr_code"`
	ManufacturerBottomQrCode sql.NullString `json:"manufacturer_bottom_qr_code"`
	RmaNumber                sql.NullString `json:"rma_number"`
	IoModuleHwVersionsID     uuid.NullUUID  `json:"io_module_hw_versions_id"`
	OrderID                  uuid.NullUUID  `json:"order_id"`
	CreatedAt                sql.NullTime   `json:"created_at"`
	UpdatedAt                sql.NullTime   `json:"updated_at"`
	IsDeleted                sql.NullBool   `json:"is_deleted"`
}

type IoModuleHwVersion struct {
	ID        uuid.UUID     `json:"id"`
	HwVersion sql.NullInt32 `json:"hw_version"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	IsDeleted sql.NullBool  `json:"is_deleted"`
}

type IoModuleNote struct {
	ID          uuid.UUID      `json:"id"`
	IoModulesID uuid.NullUUID  `json:"io_modules_id"`
	Body        sql.NullString `json:"body"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	IsDeleted   sql.NullBool   `json:"is_deleted"`
}

type IoModuleType struct {
	ID               uuid.UUID             `json:"id"`
	ModuleTypeName   sql.NullString        `json:"module_type_name"`
	ModuleTypeNumber sql.NullInt32         `json:"module_type_number"`
	PinoutJson       pqtype.NullRawMessage `json:"pinout_json"`
	CreatedAt        sql.NullTime          `json:"created_at"`
	UpdatedAt        sql.NullTime          `json:"updated_at"`
	IsDeleted        sql.NullBool          `json:"is_deleted"`
}

type LedDaughterBoard struct {
	ID                 uuid.UUID      `json:"id"`
	ManufacturerQrCode sql.NullString `json:"manufacturer_qr_code"`
	ControllersID      uuid.NullUUID  `json:"controllers_id"`
	CreatedAt          sql.NullTime   `json:"created_at"`
	UpdatedAt          sql.NullTime   `json:"updated_at"`
	IsDeleted          sql.NullBool   `json:"is_deleted"`
}

type M2ModulesType struct {
	ID               uuid.UUID      `json:"id"`
	Name             sql.NullString `json:"name"`
	ModuleTypeNumber sql.NullInt32  `json:"module_type_number"`
	CreatedAt        sql.NullTime   `json:"created_at"`
	UpdatedAt        sql.NullTime   `json:"updated_at"`
	IsDeleted        sql.NullBool   `json:"is_deleted"`
}

type Manufacturer struct {
	ID        uuid.UUID      `json:"id"`
	Name      sql.NullString `json:"name"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	IsDeleted sql.NullBool   `json:"is_deleted"`
}

type MiniPcieModule struct {
	ID                    uuid.UUID      `json:"id"`
	MiniPcieModulesTypeID uuid.NullUUID  `json:"mini_pcie_modules_type_id"`
	SerialNumber          sql.NullString `json:"serial_number"`
	CreatedAt             sql.NullTime   `json:"created_at"`
	UpdatedAt             sql.NullTime   `json:"updated_at"`
	IsDeleted             sql.NullBool   `json:"is_deleted"`
}

type MiniPcieModulesType struct {
	ID        uuid.UUID      `json:"id"`
	Name      sql.NullString `json:"name"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	IsDeleted sql.NullBool   `json:"is_deleted"`
}

type Order struct {
	ID            uuid.UUID      `json:"id"`
	CustomersID   uuid.NullUUID  `json:"customers_id"`
	Notes         sql.NullString `json:"notes"`
	DeliveryDate  sql.NullTime   `json:"delivery_date"`
	InvoiceNumber sql.NullString `json:"invoice_number"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
	IsDeleted     sql.NullBool   `json:"is_deleted"`
}

type Project struct {
	ID        uuid.UUID      `json:"id"`
	Name      sql.NullString `json:"name"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	IsDeleted sql.NullBool   `json:"is_deleted"`
}
