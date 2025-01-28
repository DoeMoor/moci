package pkg

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type DBControllerJsonToFrontend struct {
	ID                        string          `json:"id"`
	TypeName                  string          `json:"typeName"`
	IoModuleSocketAmount      int             `json:"ioModuleSocketAmount"`
	ProjectName               string          `json:"projectName"`
	Description               string          `json:"description"`
	PcbHwVersion              string          `json:"pcbHwVersion"`
	PcbVersionNumber          int             `json:"pcbVersionNumber"`
	SerialNumber              string          `json:"serialNumber"`
	ManufacturerName          string          `json:"manufacturerName"`
	AssemblyDate              time.Time       `json:"assemblyDate"`
	MacAddress                string          `json:"macAddress"`
	SimNumber                 string          `json:"simNumber"`
	EncloserSerialNumber      string          `json:"encloserSerialNumber"`
	EncloserManufacturerName  string          `json:"encloserManufacturerName"`
	MiniPcieModulesName       string          `json:"miniPcieModulesName"`
	MiniPcieModuleTypeNumber  int             `json:"miniPcieModuleTypeNumber"`
	MiniPcieSerialNumber      string          `json:"miniPcieSerialNumber"`
	M2ModuleName              string          `json:"m2ModuleName"`
	M2ModuleTypeNumber        int             `json:"m2ModuleTypeNumber"`
	LedBoardQrCode            string          `json:"ledBoardQrCode"`
	DisplayType               string          `json:"displayType"`
	DisplayManufacturerQrCode string          `json:"displayManufacturerQrCode"`
	ArticleNumber             string          `json:"articleNumber"`
	InfoQrCode                string          `json:"infoQrCode"`
	Can1Terminated            bool            `json:"can1Terminated"`
	Can2Terminated            bool            `json:"can2Terminated"`
	Can3Terminated            bool            `json:"can3Terminated"`
	Can4Terminated            bool            `json:"can4Terminated"`
	Usb                       bool            `json:"usb"`
	Serial                    bool            `json:"serial"`
	ManufacturerQrCodes       string          `json:"manufacturerQrCodes"`
	OrderID                   uuid.UUID       `json:"orderId"`
	CreatedAt                 time.Time       `json:"createdAt"`
	UpdatedAt                 time.Time       `json:"updatedAt"`
	IsDeleted                 bool            `json:"isDeleted"`
	SlotPinoutJson            json.RawMessage `json:"slotPinoutJson"`
}
type AddNewController struct {
	TypesID            uuid.UUID `json:"typesId"`
	ProjectsID         uuid.UUID `json:"projectsId"`
	Description        string    `json:"description"`
	PcbHwVersionsID    uuid.UUID `json:"pcbHwVersionsId"`
	SerialNumber       string    `json:"serialNumber"`
	ManufacturersID    uuid.UUID `json:"manufacturersId"`
	MacAddress         string    `json:"macAddress"`
	SimNumber          string    `json:"simNumber"`
	M2ModulesTypesID   uuid.UUID `json:"m2ModulesTypesId"`
	ArticleNumber      string    `json:"articleNumber"`
	InfoQrCode         string    `json:"infoQrCode"`
	USB                bool      `json:"usb"`
	Serial             bool      `json:"serial"`
	ManufacturerQrCode string    `json:"manufacturerQrCode"`
	OrderID            string    `json:"orderId"`
	AssemblyDate       string    `json:"assemblyDate"`
}

type Display struct {
	Id                 uuid.UUID `json:"id"`
	TypeID             string    `json:"typeId"`
	ManufacturerQrCode string    `json:"manufacturerQrCode"`
}

type Encloser struct {
	Id             uuid.UUID `json:"id"`
	ControllerID   uuid.UUID `json:"controllerId"`
	TypeID         string    `json:"typeId"`
	ManufacturerID uuid.UUID `json:"manufacturerId"`
	SerialNumber   string    `json:"serialNumber"`
}

type MiniPcie struct {
	Id           uuid.UUID `json:"id"`
	TypeID       string    `json:"typeId"`
	SerialNumber string    `json:"serialNumber"`
}

type CanTermination struct {
	Id             uuid.UUID `json:"id"`
	Can1Terminated bool      `json:"can1Terminated"`
	Can2Terminated bool      `json:"can2Terminated"`
	Can3Terminated bool      `json:"can3Terminated"`
	Can4Terminated bool      `json:"can4Terminated"`
}

type NewControllerJson struct {
	Id             uuid.UUID        `json:"id"`
	CustomerID     string           `json:"customerId"`
	Controller     AddNewController `json:"controller"`
	CanTermination CanTermination   `json:"canTermination"`
	Display        Display          `json:"display"`
	Encloser       Encloser         `json:"encloser"`
	MiniPcie       MiniPcie         `json:"miniPcie"`
	LedBoardQrCode string           `json:"ledBoardQrCode"`
	LedBoardID     uuid.UUID        `json:"ledBoardId"`
}

type NewControllerJsonFromFrontend struct {
	Id                        uuid.UUID `json:"id"`
	CustomerID                uuid.UUID `json:"customerId"`
	TypesID                   uuid.UUID `json:"typesId"`
	ProjectsID                uuid.UUID `json:"projectsId"`
	Description               *string   `json:"description"`
	PcbHwVersionsID           uuid.UUID `json:"pcbHwVersionsId"`
	SerialNumber              string    `json:"serialNumber"`
	ManufacturersID           uuid.UUID `json:"manufacturersId"`
	MacAddress                string    `json:"macAddress"`
	SimNumber                 *string   `json:"simNumber"`
	M2ModulesTypesID          uuid.UUID `json:"m2ModulesTypesId"`
	ArticleNumber             *string   `json:"articleNumber"`
	InfoQrCode                *string   `json:"infoQrCode"`
	Usb                       bool      `json:"usb"`
	Serial                    bool      `json:"serial"`
	ManufacturerQrCode        *string   `json:"manufacturerQrCode"`
	OrderID                   uuid.UUID `json:"orderId"`
	AssemblyDate              string    `json:"assemblyDate"`
	CanTerminationID          uuid.UUID `json:"canTerminationId"`
	Can1Terminated            bool      `json:"can1Terminated"`
	Can2Terminated            bool      `json:"can2Terminated"`
	Can3Terminated            bool      `json:"can3Terminated"`
	Can4Terminated            bool      `json:"can4Terminated"`
	DisplayId                 uuid.UUID `json:"displayId"`
	DisplayTypeName           string    `json:"displayTypeName"`
	DisplayManufacturerQrCode string    `json:"displayManufacturerQrCode"`
	EncloserID                uuid.UUID `json:"encloserId"`
	EncloserManufacturerID    uuid.UUID `json:"encloserManufacturerId"`
	EncloserSerialNumber      string    `json:"encloserSerialNumber"`
	MiniPcieID                uuid.UUID `json:"miniPcieId"`
	MiniPcieTypeID            uuid.UUID `json:"miniPcieTypeId"`
	MiniPcieSerialNumber      string    `json:"miniPcieSerialNumber"`
	LedBoardID                uuid.UUID `json:"ledBoardId"`
	LedBoardQrCode            string    `json:"ledBoardQrCode"`
}
