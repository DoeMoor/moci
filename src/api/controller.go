package api

import (
	"encoding/json"
	// "fmt"
	"log"
	"time"

	// "github.com/doemoor/moci/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type GetControllerJsonResponse struct {
	ID                            string          `json:"id"`
	ControllerTypeName            string          `json:"controllerTypeName"`
	IoModuleSocketAmount          int             `json:"ioModuleSocketAmount"`
	ProjectName                   string          `json:"projectName"`
	Description                   string          `json:"description"`
	PcbHwVersion                  string          `json:"pcbHwVersion"`
	PcbVersionNumber              int             `json:"pcbVersionNumber"`
	ControllerSerialNumber        string          `json:"controllerSerialNumber"`
	ManufacturerName              string          `json:"manufacturerName"`
	AssemblyDate                  time.Time       `json:"assemblyDate"`
	MacAddress                    string          `json:"macAddress"`
	SimNumber                     string          `json:"simNumber"`
	EncloserSerialNumber          string          `json:"encloserSerialNumber"`
	EncloserManufacturerName      string          `json:"encloserManufacturerName"`
	MiniPcieModulesName           string          `json:"miniPcieModulesName"`
	MiniPcieModuleTypeNumber      int             `json:"miniPcieModuleTypeNumber"`
	MiniPcieSerialNumber          string          `json:"miniPcieSerialNumber"`
	M2ModuleName                  string          `json:"m2ModuleName"`
	M2ModuleTypeNumber            int             `json:"m2ModuleTypeNumber"`
	LedBoard                      string          `json:"ledBoard"`
	DisplayType                   string          `json:"displayType"`
	DisplayManufacturerQrCode     string          `json:"displayManufacturerQrCode"`
	ArticleNumber                 string          `json:"articleNumber"`
	ControllerInfoQrCode          string          `json:"controllerInfoQrCode"`
	Can1Terminated                bool            `json:"can1Terminated"`
	Can2Terminated                bool            `json:"can2Terminated"`
	Can3Terminated                bool            `json:"can3Terminated"`
	Can4Terminated                bool            `json:"can4Terminated"`
	Usb                           bool            `json:"usb"`
	Serial                        bool            `json:"serial"`
	ControllerManufacturerQrCodes string          `json:"controllerManufacturerQrCodes"`
	OrderID                       uuid.UUID       `json:"orderId"`
	CreatedAt                     time.Time       `json:"createdAt"`
	UpdatedAt                     time.Time       `json:"updatedAt"`
	IsDeleted                     bool            `json:"isDeleted"`
	SlotPinoutJson                json.RawMessage `json:"slotPinoutJson"`
}

func (cnf *ApiConfig) GetAllControllers(c *fiber.Ctx) error {
	allController, err := cnf.DbQueries.GetAllControllers(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		log.Println(err)
		return err
	}

	var allControllerJson []GetControllerJsonResponse

	for _, controller := range allController {
		allControllerJson = append(allControllerJson, GetControllerJsonResponse{
			ID:                            controller.ID.String(),
			ControllerTypeName:            controller.ControllerTypeName.String,
			IoModuleSocketAmount:          int(controller.IoModuleSocketAmount.Int32),
			ProjectName:                   controller.ProjectName.String,
			Description:                   controller.Description.String,
			PcbHwVersion:                  controller.PcbHwVersion.(string),
			PcbVersionNumber:              int(controller.PcbVersionNumber.Int32),
			ControllerSerialNumber:        controller.ControllerSerialNumber.String,
			ManufacturerName:              controller.ManufacturerName.String,
			AssemblyDate:                  controller.AssemblyDate.Time,
			MacAddress:                    controller.MacAddress.String,
			SimNumber:                     controller.SimNumber.String,
			EncloserSerialNumber:          controller.EnclosureSerialNumber.String,
			EncloserManufacturerName:      controller.EnclosureManufacturer.String,
			MiniPcieModulesName:           controller.MiniPcieModulesName.String,
			MiniPcieModuleTypeNumber:      int(controller.MiniPcieModuleTypeNumber.Int32),
			MiniPcieSerialNumber:          controller.MiniPcieSerialNumber.String,
			M2ModuleName:                  controller.M2ModuleName.String,
			M2ModuleTypeNumber:            int(controller.M2ModuleTypeNumber.Int32),
			LedBoard:                      controller.LedBoard.String,
			DisplayType:                   controller.DisplayType.String,
			DisplayManufacturerQrCode:     controller.DisplayManufacturerQrCode.String,
			ArticleNumber:                 controller.ArticleNumber.String,
			ControllerInfoQrCode:          controller.ControllerInfoQrCode.String,
			Can1Terminated:                controller.Can1Terminated.Bool,
			Can2Terminated:                controller.Can2Terminated.Bool,
			Can3Terminated:                controller.Can3Terminated.Bool,
			Can4Terminated:                controller.Can4Terminated.Bool,
			Usb:                           controller.Usb.Bool,
			Serial:                        controller.Serial.Bool,
			ControllerManufacturerQrCodes: controller.ControllerManufacturerQrCode.String,
			OrderID:                       controller.OrderID.UUID,
			CreatedAt:                     controller.CreatedAt.Time,
			UpdatedAt:                     controller.UpdatedAt.Time,
			SlotPinoutJson:                controller.SlotPinoutJson.RawMessage,
			IsDeleted:                     controller.IsDeleted.Bool,
		})
	}

	return c.JSON(allControllerJson)
}

func (cnf *ApiConfig) GetControllerById(c *fiber.Ctx) error {

	if c.Params("id") == "" {
		c.Response().SetStatusCode(404)
		log.Println(c.OriginalURL(), " id is empty")
		return c.SendString("id is empty")
	}

	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		c.Response().SetStatusCode(404)
		log.Println("error parsing uuid: ", c.OriginalURL(), "\n", "  error: ", err)
		return c.SendString("wrong uuid")
	}

	controllerById, err := cnf.DbQueries.GetControllerById(c.Context(), uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.Response().SetStatusCode(404)
			return c.SendString("controller not found")
		}
		c.Response().SetStatusCode(500)
		log.Println("Error getting controller by id from database: ", err)
		return err
	}

	controllerJson := GetControllerJsonResponse{
		ID:                            controllerById.ID.String(),
		ControllerTypeName:            controllerById.ControllerTypeName.String,
		IoModuleSocketAmount:          int(controllerById.IoModuleSocketAmount.Int32),
		ProjectName:                   controllerById.ProjectName.String,
		Description:                   controllerById.Description.String,
		PcbHwVersion:                  controllerById.PcbHwVersion.(string),
		PcbVersionNumber:              int(controllerById.PcbVersionNumber.Int32),
		ControllerSerialNumber:        controllerById.ControllerSerialNumber.String,
		ManufacturerName:              controllerById.ManufacturerName.String,
		AssemblyDate:                  controllerById.AssemblyDate.Time,
		MacAddress:                    controllerById.MacAddress.String,
		SimNumber:                     controllerById.SimNumber.String,
		MiniPcieModulesName:           controllerById.MiniPcieModulesName.String,
		MiniPcieModuleTypeNumber:      int(controllerById.MiniPcieModuleTypeNumber.Int32),
		MiniPcieSerialNumber:          controllerById.MiniPcieSerialNumber.String,
		M2ModuleName:                  controllerById.M2ModuleName.String,
		M2ModuleTypeNumber:            int(controllerById.M2ModuleTypeNumber.Int32),
		LedBoard:                      controllerById.LedBoard.String,
		DisplayType:                   controllerById.DisplayType.String,
		DisplayManufacturerQrCode:     controllerById.DisplayManufacturerQrCode.String,
		ArticleNumber:                 controllerById.ArticleNumber.String,
		ControllerInfoQrCode:          controllerById.ControllerInfoQrCode.String,
		Can1Terminated:                controllerById.Can1Terminated.Bool,
		Can2Terminated:                controllerById.Can2Terminated.Bool,
		Can3Terminated:                controllerById.Can3Terminated.Bool,
		Can4Terminated:                controllerById.Can4Terminated.Bool,
		Usb:                           controllerById.Usb.Bool,
		Serial:                        controllerById.Serial.Bool,
		ControllerManufacturerQrCodes: controllerById.ControllerManufacturerQrCode.String,
		OrderID:                       controllerById.OrderID.UUID,
		CreatedAt:                     controllerById.CreatedAt.Time,
		UpdatedAt:                     controllerById.UpdatedAt.Time,
		IsDeleted:                     controllerById.IsDeleted.Bool,
		SlotPinoutJson:                controllerById.SlotPinoutJson.RawMessage,
	}

	return c.JSON(controllerJson)
}

func (cnf *ApiConfig) GetAllControllerTypes(c *fiber.Ctx) error {

	allControllerTypes, err := cnf.DbQueries.GetALLControllerTypes(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		log.Println(err)
		return err
	}
	type controllerTypesForJson struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}

	var result []controllerTypesForJson

	for _, controllerType := range allControllerTypes {
		result = append(result, controllerTypesForJson{
			ID:   controllerType.ID,
			Name: controllerType.Name.String,
		})
	}

	return c.JSON(result)
}

func (cnf *ApiConfig) GetAllControllerPcbHwVersions(c *fiber.Ctx) error {

	allControllerPcbHwVersions, err := cnf.DbQueries.GetAllControllersPcbHwVersions(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		log.Println(err)
		return err
	}
	type controllerPcbHwVersionsForJson struct {
		ID          string `json:"id"`
		Version     string `json:"version"`
		Revision    string `json:"revision"`
		FullVersion string `json:"fullVersion"`
	}

	var result []controllerPcbHwVersionsForJson

	for _, version := range allControllerPcbHwVersions {
		if version.Revision.String == "" {
			result = append(result, controllerPcbHwVersionsForJson{
				ID:          version.ID.String(),
				Version:     version.VersionNumber.String,
				Revision:    version.Revision.String,
				FullVersion: version.VersionNumber.String,
			})
		}

		if version.Revision.String != "" {
			result = append(result, controllerPcbHwVersionsForJson{
				ID:          version.ID.String(),
				Version:     version.VersionNumber.String,
				Revision:    version.Revision.String,
				FullVersion: version.VersionNumber.String + "-" + version.Revision.String,
			})
		}
	}

	return c.JSON(result)
}

func (cnf *ApiConfig) GetControllerTypesPinout(c *fiber.Ctx) error {

	if c.Params("id") == "" {
		c.Response().SetStatusCode(404)
		log.Println(c.OriginalURL(), " name is empty")
		return c.SendString("name is empty")
	}

	typeID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		c.Response().SetStatusCode(404)
		log.Println("error parsing uuid: ", c.OriginalURL(), "\n", "  error: ", err)
		return c.SendString("wrong uuid")
	}

	controllerTypesPinout, err := cnf.DbQueries.GetControllerTypesPinout(
		c.Context(),
		typeID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.Response().SetStatusCode(404)
			return c.SendString("controller type not found")
		}
		c.Response().SetStatusCode(500)
		log.Println(err)
		return err
	}
	return c.JSON(controllerTypesPinout.RawMessage)
}

func (cnf *ApiConfig) CreateController(c *fiber.Ctx) error {

	if c.Get("Content-Type") != "application/json" {
		c.Response().SetStatusCode(400)
		return c.SendString("Content-Type must be application/json")
	}

	type Controller struct {
		TypesID            string `json:"typesId"`
		ProjectsID         string `json:"projectsId"`
		Description        string `json:"description"`
		PcbHwVersionsID    string `json:"pcbHwVersionsId"`
		SerialNumber       string `json:"serialNumber"`
		ManufacturersID    string `json:"manufacturersId"`
		MacAddress         string `json:"macAddress"`
		SimNumber          string `json:"simNumber"`
		M2ModulesTypesID   string `json:"m2ModulesTypesId"`
		ArticleNumber      string `json:"articleNumber"`
		InfoQrCode             string `json:"infoQrCode"`
		USB                bool   `json:"usb"`
		Serial             bool   `json:"serial"`
		ManufacturerQrCode string `json:"manufacturerQrCode"`
		OrderID            string `json:"orderId"`
		AssemblyDate       string `json:"assemblyDate"`
	}

	type CanTermination struct {
		Can1Terminated bool `json:"can1Terminated"`
		Can2Terminated bool `json:"can2Terminated"`
		Can3Terminated bool `json:"can3Terminated"`
		Can4Terminated bool `json:"can4Terminated"`
	}

	type Display struct {
		TypeID             string `json:"typeId"`
		ManufacturerQrCode string `json:"manufacturerQrCode"`
	}

	type Encloser struct {
		ManufacturerID string `json:"manufacturerId"`
		SerialNumber   string `json:"serialNumber"`
	}

	type MiniPcie struct {
		ModulesID    string `json:"modulesId"`
		SerialNumber string `json:"serialNumber"`
	}

	type newController struct {
		CustomerID     string         `json:"customerId"`
		Controller     Controller     `json:"controller"`
		CanTermination CanTermination `json:"canTermination"`
		Display        Display        `json:"display"`
		Encloser       Encloser       `json:"encloser"`
		MiniPcie       MiniPcie       `json:"miniPcie"`
		LedBoardQrCode string         `json:"ledBoardQrCode"`
	}
	
	var controller newController

	if err := c.BodyParser(&controller); err != nil {
		log.Println("CreateController - BodyParser error: ", err)
		c.Response().SetStatusCode(400)
		return c.SendString("json parse error")
	}

	return c.JSON(controller)
}
