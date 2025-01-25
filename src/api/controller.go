package api

import (
	"database/sql"
	"encoding/json"

	// "fmt"
	"log"
	"time"

	// "github.com/doemoor/moci/internal/database"
	"github.com/doemoor/moci/internal/database"
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
type AddNewController struct {
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
	InfoQrCode         string `json:"infoQrCode"`
	USB                bool   `json:"usb"`
	Serial             bool   `json:"serial"`
	ManufacturerQrCode string `json:"manufacturerQrCode"`
	OrderID            string `json:"orderId"`
	AssemblyDate       string `json:"assemblyDate"`
}

type Display struct {
	Id                 uuid.UUID `json:"id"`
	TypeID             string    `json:"typeId"`
	ManufacturerQrCode string    `json:"manufacturerQrCode"`
}

type Encloser struct {
	Id             uuid.UUID `json:"id"`
	ManufacturerID string    `json:"manufacturerId"`
	SerialNumber   string    `json:"serialNumber"`
}

type MiniPcie struct {
	Id           uuid.UUID `json:"id"`
	ModulesID    string    `json:"modulesId"`
	SerialNumber string    `json:"serialNumber"`
}

type CanTermination struct {
	Id             uuid.UUID `json:"id"`
	Can1Terminated bool      `json:"can1Terminated"`
	Can2Terminated bool      `json:"can2Terminated"`
	Can3Terminated bool      `json:"can3Terminated"`
	Can4Terminated bool      `json:"can4Terminated"`
}

func (cnf *ApiConfig) GetAllControllers(c *fiber.Ctx) error {
	allController, err := cnf.DbQ.GetAllControllers(c.Context())
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

	controllerById, err := cnf.DbQ.GetControllerById(c.Context(), uuid)
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

	allControllerTypes, err := cnf.DbQ.GetALLControllerTypes(c.Context())
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

	allControllerPcbHwVersions, err := cnf.DbQ.GetAllControllersPcbHwVersions(c.Context())
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

	controllerTypesPinout, err := cnf.DbQ.GetControllerTypesPinout(
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

	type newControllerJson struct {
		CustomerID     string           `json:"customerId"`
		Controller     AddNewController `json:"controller"`
		CanTermination CanTermination   `json:"canTermination"`
		Display        Display          `json:"display"`
		Encloser       Encloser         `json:"encloser"`
		MiniPcie       MiniPcie         `json:"miniPcie"`
		LedBoardQrCode string           `json:"ledBoardQrCode"`
	}

	var newController newControllerJson

	if err := c.BodyParser(&newController); err != nil {
		log.Println("CreateController - BodyParser error: ", err)
		c.Response().SetStatusCode(400)
		return c.SendString("json parse error")
	}

	//begin assembly of new controller object

	// validate controller can termination
	canParam, canErr := validateCanTerminationConf(newController.CanTermination)
	if canErr != nil {
		log.Println("CreateController - CanTermination error: ", canErr)
		c.Response().SetStatusCode(400)
		return c.SendString("Can termination parse error")
	}
	// try to get termination config id
	canTerminationID, err := cnf.DbQ.GetTerminationConfigIdByCanTerminated(c.Context(), canParam)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			log.Println("CreateController - GetTerminationConfigIdByCanTerminated error: ", err)
			c.Response().SetStatusCode(400)
			return c.SendString("Can termination parse error")
		}
	}
	newController.CanTermination.Id.Scan(canTerminationID)

	// check if display is present in db
	var displaySN sql.NullString
	if err := displaySN.Scan(newController.Display.ManufacturerQrCode); err != nil {
		log.Println("CreateController - Display parse error: ", err)
		c.Response().SetStatusCode(400)
		return c.SendString("Display parse error")
	}
	display, displayErr := cnf.DbQ.GetDisplayAdapterByQR(c.Context(), displaySN)
	if displayErr != nil {
		if displayErr.Error() != "sql: no rows in result set" {
			log.Println("CreateController - GetDisplayAdapterByQR error: ", displayErr)
			c.Response().SetStatusCode(400)
			return c.SendString("Display parse error")
		}
	}
	if display.ID.String() != "00000000-0000-0000-0000-000000000000" {
		log.Println("CreateController - Display already exists: ")
		return c.SendString("Display already exists")
	}

	// check if encloser is present in db
	var encloserSN sql.NullString
	if err := encloserSN.Scan(newController.Encloser.SerialNumber); err != nil {
		log.Println("CreateController - Encloser parse error: ", err)
		c.Response().SetStatusCode(400)
		return c.SendString("Encloser parse error")
	}
	encloser, encloserErr := cnf.DbQ.GetEncloserBySN(c.Context(), encloserSN)
	if encloserErr != nil {
		if encloserErr.Error() != "sql: no rows in result set" {
			log.Println("CreateController - GetEncloserBySN error: ", encloserErr)
			c.Response().SetStatusCode(400)
			return c.SendString("Encloser parse error")
		}
	}
	if encloser.ID.String() != "00000000-0000-0000-0000-000000000000" {
		log.Println("CreateController - Encloser already exists: ", encloser)
		return c.SendString("Encloser already exists")
	}

	// check if miniPcie is present in db
	var miniPcieSN sql.NullString
	if err := miniPcieSN.Scan(newController.MiniPcie.SerialNumber); err != nil {
		log.Println("CreateController - MiniPcie parse error: ", err)
		c.Response().SetStatusCode(400)
		return c.SendString("MiniPcie parse error")
	}
	miniPcie, miniPcieErr := cnf.DbQ.GetMiniPCIeModuleBySN(c.Context(), miniPcieSN)
	if miniPcieErr != nil {
		if miniPcieErr.Error() != "sql: no rows in result set" {
			log.Println("CreateController - GetMiniPcieBySN error: ", miniPcieErr)
			c.Response().SetStatusCode(400)
			return c.SendString("MiniPcie parse error")
		}
	}
	if miniPcie.ID.String() != "00000000-0000-0000-0000-000000000000" {
		log.Println("CreateController - MiniPcie already exists: ", miniPcie)
		return c.SendString("MiniPcie already exists")
	}

	// check if ledBoard is present in db
	var ledBoardSN sql.NullString
	if err := ledBoardSN.Scan(newController.LedBoardQrCode); err != nil {
		log.Println("CreateController - LedBoard parse error: ", err)
		c.Response().SetStatusCode(400)
		return c.SendString("LedBoard parse error")
	}
	ledBoard, ledBoardErr := cnf.DbQ.GetLedBoardByQR(c.Context(), ledBoardSN)
	if ledBoardErr != nil {
		if ledBoardErr.Error() != "sql: no rows in result set" {
			log.Println("CreateController - GetLedBoardByQR error: ", ledBoardErr)
			c.Response().SetStatusCode(400)
			return c.SendString("LedBoard parse error")
		}
	}
	if ledBoard.ID.String() != "00000000-0000-0000-0000-000000000000" {
		log.Println("CreateController - LedBoard already exists: ", ledBoard)
		return c.SendString("LedBoard already exists")
	}

	// end assembly

	return c.JSON(newController)
}

func validateCanTerminationConf(ct CanTermination) (database.GetTerminationConfigIdByCanTerminatedParams, error) {
	var canParam database.GetTerminationConfigIdByCanTerminatedParams
	canErr := canParam.Can1Terminated.Scan(ct.Can1Terminated)
	canErr = canParam.Can2Terminated.Scan(ct.Can2Terminated)
	canErr = canParam.Can3Terminated.Scan(ct.Can3Terminated)
	canErr = canParam.Can4Terminated.Scan(ct.Can4Terminated)
	return canParam, canErr
}
