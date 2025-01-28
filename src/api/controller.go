package api

import (
	"database/sql"

	// "fmt"
	"log"

	"github.com/doemoor/moci/internal/database"
	"github.com/doemoor/moci/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (cnf *ApiCfg) GetAllControllers(c *fiber.Ctx) error {

	allController, err := cnf.DbQ.GetAllControllers(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		log.Println(err)
		return err
	}

	allControllerJson := mapControllersToJson(allController)

	return c.JSON(allControllerJson)
}

func (cnf *ApiCfg) GetControllerById(c *fiber.Ctx) error {

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

	controllerJson, shouldReturn, err := controllerByID(cnf, c, uuid)
	if shouldReturn {
		return err
	}

	return c.JSON(controllerJson)
}

func (cnf *ApiCfg) GetAllControllerTypes(c *fiber.Ctx) error {

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

func (cnf *ApiCfg) GetAllControllerPcbHwVersions(c *fiber.Ctx) error {

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

func (cnf *ApiCfg) GetControllerTypesPinout(c *fiber.Ctx) error {

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

func (cnf *ApiCfg) CreateController(c *fiber.Ctx) error {

	if c.Get("Content-Type") != "application/json" {
		c.Response().SetStatusCode(400)
		return c.SendString("Content-Type must be application/json")
	}

	var newController pkg.NewControllerJsonFromFrontend

	if err := c.BodyParser(&newController); err != nil {
		log.Println("CreateController - BodyParser error: ", err)
		c.Response().SetStatusCode(400)
		return c.SendString("json parse error")
	}

	//begin assembly of new controller object

	// validate controller can termination
	canParam, canErr := validateCanTerminationConfParam(newController)
	if canErr != nil {
		log.Println("CreateController - pkg.CanTermination error: ", canErr)
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
	// no config found create new
	if canTerminationID.String() != "00000000-0000-0000-0000-000000000000" {
		newController.CanTerminationID = canTerminationID
	} else {
		newCanTerminationId, err := cnf.DbQ.CreateCanTerminationConfig(c.Context(), database.CreateCanTerminationConfigParams{
			Can1Terminated: canParam.Can1Terminated,
			Can2Terminated: canParam.Can2Terminated,
			Can3Terminated: canParam.Can3Terminated,
			Can4Terminated: canParam.Can4Terminated,
		})
		if err != nil {
			log.Println("CreateController - CreateCanTerminationConfig error: ", err)
			c.Response().SetStatusCode(400)
			return c.SendString("Can termination parse error")
		}
		newController.CanTerminationID = newCanTerminationId.ID
	}

	// check if display is present in db
	var displaySN sql.NullString
	if err := displaySN.Scan(newController.DisplayManufacturerQrCode); err != nil {
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
		c.Response().SetStatusCode(400)
		return c.SendString("Display already exists")
	}

	// check if encloser is present in db
	var encloserSN sql.NullString
	if err := encloserSN.Scan(newController.EncloserSerialNumber); err != nil {
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
		c.Response().SetStatusCode(400)
		return c.SendString("Encloser already exists")
	}

	// check if miniPcie is present in db
	var miniPcieSN sql.NullString
	if err := miniPcieSN.Scan(newController.MiniPcieSerialNumber); err != nil {
		log.Println("CreateController - pkg.MiniPcie parse error: ", err)
		c.Response().SetStatusCode(400)
		return c.SendString("pkg.MiniPcie parse error")
	}
	miniPcie, miniPcieErr := cnf.DbQ.GetMiniPCIeModuleBySN(c.Context(), miniPcieSN)
	if miniPcieErr != nil {
		if miniPcieErr.Error() != "sql: no rows in result set" {
			log.Println("CreateController - GetMiniPcieBySN error: ", miniPcieErr)
			c.Response().SetStatusCode(400)
			return c.SendString("pkg.MiniPcie parse error")
		}
	}
	if miniPcie.ID.String() != "00000000-0000-0000-0000-000000000000" {
		log.Println("CreateController - pkg.MiniPcie already exists: ", miniPcie)
		c.Response().SetStatusCode(400)
		return c.SendString("pkg.MiniPcie already exists")
	}

	// check if ledBoard is present in db
	var ledBoardQR sql.NullString
	if err := ledBoardQR.Scan(newController.LedBoardQrCode); err != nil {
		log.Println("CreateController - LedBoard parse error: ", err)
		c.Response().SetStatusCode(400)
		return c.SendString("LedBoard parse error")
	}
	ledBoard, ledBoardErr := cnf.DbQ.GetLedBoardByQR(c.Context(), ledBoardQR)
	if ledBoardErr != nil {
		if ledBoardErr.Error() != "sql: no rows in result set" {
			log.Println("CreateController - GetLedBoardByQR error: ", ledBoardErr)
			c.Response().SetStatusCode(400)
			return c.SendString("LedBoard parse error")
		}
	}
	if ledBoard.ID.String() != "00000000-0000-0000-0000-000000000000" {
		log.Println("CreateController - LedBoard already exists: ", ledBoard)
		c.Response().SetStatusCode(400)
		return c.SendString("LedBoard already exists")
	}

	// start transaction
	shouldReturn, err := controllerTransaction(cnf, c, &newController, ledBoardQR)
	if shouldReturn {
		log.Println("CreateController - controllerTransaction error: ", err)
		return c.JSON(newController)
	}

	controllerJson, shouldReturn, err := controllerByID(cnf, c, newController.Id)
	if shouldReturn {
		log.Println("CreateController - New controllerByID error: ", err)
		return c.JSON(newController)
	}



	return c.JSON(controllerJson)
}

func controllerTransaction(cnf *ApiCfg, c *fiber.Ctx, newController *pkg.NewControllerJsonFromFrontend, ledBoardQR sql.NullString) (bool, error) {
	tx, err := cnf.DB.Begin()
	if err != nil {
		log.Println("CreateController - Transaction error: ", err)
		c.Response().SetStatusCode(500)
		return true, c.SendString("Transaction error")
	}
	defer tx.Rollback()
	qtx := cnf.DbQ.WithTx(tx)

	// try to create display
	newDisplayParam, err := validateNewDisplayParams(*newController)
	if err != nil {
		log.Println("CreateController - validateNewDisplayParams error: ", err)
		c.Response().SetStatusCode(400)
		return true, c.SendString("Display parse error")
	}
	newDisplay, err := qtx.CreateDisplayAdapter(c.Context(), newDisplayParam)
	if err != nil {
		log.Println("CreateController - CreateDisplayAdapter error: ", err)
		c.Response().SetStatusCode(400)
		return true, c.SendString("CreateDisplayAdapter error")
	}
	newController.DisplayId = newDisplay.ID

	// try to create miniPcie
	newMiniPcieParam, err := validateNewMiniPcieParams(*newController)
	if err != nil {
		log.Println("CreateController - validateNewMiniPcieParams error: ", err)
		c.Response().SetStatusCode(400)
		return true, c.SendString("pkg.MiniPcie parse error")
	}
	newMiniPcie, err := qtx.CreateMiniPCIeModule(c.Context(), newMiniPcieParam)
	if err != nil {
		log.Println("CreateController - CreateMiniPcieModule error: ", err)
		c.Response().SetStatusCode(400)
		return true, c.SendString("CreateMiniPcieModule error")
	}
	newController.MiniPcieID = newMiniPcie.ID

	// try to create controller
	controllerParam, err := validateNewControllerParams(*newController)
	if err != nil {
		log.Println("CreateController - validateNewControllerParams error: ", err)
		c.Response().SetStatusCode(400)
		return true, c.SendString("Controller parse error")
	}
	newControllerID, err := qtx.CreateController(c.Context(), controllerParam)
	if err != nil {
		log.Println("CreateController - CreateController error: ", err)
		c.Response().SetStatusCode(400)
		return true, c.SendString("CreateController error")
	}
	newController.Id = newControllerID.ID

	// try to create ledBoard
	ledBoardParam, err := validateNevLedBoardParams(ledBoardQR, newController.Id)
	if err != nil {
		log.Println("CreateController - validateNevLedBoardParams error: ", err)
		c.Response().SetStatusCode(400)
		return true, c.SendString("LedBoard parse error")
	}
	newLedBoard, err := qtx.CreateLedBoard(c.Context(), ledBoardParam)
	if err != nil {
		log.Println("CreateController - CreateLedBoard error: ", err)
		c.Response().SetStatusCode(400)
		return true, c.SendString("CreateLedBoard error")
	}
	newController.LedBoardID = newLedBoard.ID

	// try to create encloser
	newEncloserParam, err := validateNewEncloserParams(*newController)
	if err != nil {
		log.Println("CreateController - validateNewEncloserParams error: ", err)
		c.Response().SetStatusCode(400)
		return true, c.SendString("Encloser parse error")
	}
	newEncloser, err := qtx.CreateEncloser(c.Context(), newEncloserParam)
	if err != nil {
		log.Println("CreateController - CreateEncloser error: ", err)
		c.Response().SetStatusCode(400)
		return true, c.SendString("Encloser parse error")
	}
	newController.EncloserID = newEncloser.ID

	commitErr := tx.Commit()
	if commitErr != nil {
		log.Println("CreateController - Commit error: ", commitErr)
	}
	return false, nil
}

func controllerByID(cnf *ApiCfg, c *fiber.Ctx, uuid uuid.UUID) (pkg.DBControllerJsonToFrontend, bool, error) {
	controllerById, err := cnf.DbQ.GetControllerById(c.Context(), uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.Response().SetStatusCode(404)
			return pkg.DBControllerJsonToFrontend{}, true, c.SendString("controller not found")
		}
		c.Response().SetStatusCode(500)
		log.Println("Error getting controller by id from database: ", err)
		return pkg.DBControllerJsonToFrontend{}, true, err
	}

	controllerJson := mapOneControllerToJson(controllerById)
	return controllerJson, false, nil
}

func mapOneControllerToJson(controllerById database.GetControllerByIdRow) pkg.DBControllerJsonToFrontend {
	controllerJson := pkg.DBControllerJsonToFrontend{
		ID:                        controllerById.ID.String(),
		TypeName:                  controllerById.TypeName.String,
		IoModuleSocketAmount:      int(controllerById.IoModuleSocketAmount.Int32),
		ProjectName:               controllerById.ProjectName.String,
		Description:               controllerById.Description.String,
		PcbHwVersion:              controllerById.PcbHwVersion.(string),
		PcbVersionNumber:          int(controllerById.PcbVersionNumber.Int32),
		SerialNumber:              controllerById.SerialNumber.String,
		ManufacturerName:          controllerById.ManufacturerName.String,
		AssemblyDate:              controllerById.AssemblyDate.Time,
		MacAddress:                controllerById.MacAddress.String,
		SimNumber:                 controllerById.SimNumber.String,
		EncloserSerialNumber:      controllerById.EnclosureSerialNumber.String,
		EncloserManufacturerName:  controllerById.EnclosureManufacturer.String,
		MiniPcieModulesName:       controllerById.MiniPcieModulesName.String,
		MiniPcieModuleTypeNumber:  int(controllerById.MiniPcieModuleTypeNumber.Int32),
		MiniPcieSerialNumber:      controllerById.MiniPcieSerialNumber.String,
		M2ModuleName:              controllerById.M2ModuleName.String,
		M2ModuleTypeNumber:        int(controllerById.M2ModuleTypeNumber.Int32),
		LedBoardQrCode:            controllerById.LedBoardQrCode.String,
		DisplayType:               controllerById.DisplayType.String,
		DisplayManufacturerQrCode: controllerById.DisplayManufacturerQrCode.String,
		ArticleNumber:             controllerById.ArticleNumber.String,
		InfoQrCode:                controllerById.InfoQrCode.String,
		Can1Terminated:            controllerById.Can1Terminated.Bool,
		Can2Terminated:            controllerById.Can2Terminated.Bool,
		Can3Terminated:            controllerById.Can3Terminated.Bool,
		Can4Terminated:            controllerById.Can4Terminated.Bool,
		Usb:                       controllerById.Usb.Bool,
		Serial:                    controllerById.Serial.Bool,
		ManufacturerQrCodes:       controllerById.ManufacturerQrCode.String,
		OrderID:                   controllerById.OrderID.UUID,
		CreatedAt:                 controllerById.CreatedAt.Time,
		UpdatedAt:                 controllerById.UpdatedAt.Time,
		IsDeleted:                 controllerById.IsDeleted.Bool,
		SlotPinoutJson:            controllerById.SlotPinoutJson.RawMessage,
	}
	return controllerJson
}

func mapControllersToJson(allController []database.GetAllControllersRow) []pkg.DBControllerJsonToFrontend {
	var controllersJson []pkg.DBControllerJsonToFrontend

	for _, controller := range allController {
		controllersJson = append(controllersJson, pkg.DBControllerJsonToFrontend{
			ID:                        controller.ID.String(),
			TypeName:                  controller.TypeName.String,
			IoModuleSocketAmount:      int(controller.IoModuleSocketAmount.Int32),
			ProjectName:               controller.ProjectName.String,
			Description:               controller.Description.String,
			PcbHwVersion:              controller.PcbHwVersion.(string),
			PcbVersionNumber:          int(controller.PcbVersionNumber.Int32),
			SerialNumber:              controller.SerialNumber.String,
			ManufacturerName:          controller.ManufacturerName.String,
			AssemblyDate:              controller.AssemblyDate.Time,
			MacAddress:                controller.MacAddress.String,
			SimNumber:                 controller.SimNumber.String,
			EncloserSerialNumber:      controller.EnclosureSerialNumber.String,
			EncloserManufacturerName:  controller.EnclosureManufacturer.String,
			MiniPcieModulesName:       controller.MiniPcieModulesName.String,
			MiniPcieModuleTypeNumber:  int(controller.MiniPcieModuleTypeNumber.Int32),
			MiniPcieSerialNumber:      controller.MiniPcieSerialNumber.String,
			M2ModuleName:              controller.M2ModuleName.String,
			M2ModuleTypeNumber:        int(controller.M2ModuleTypeNumber.Int32),
			LedBoardQrCode:            controller.LedBoardQrCode.String,
			DisplayType:               controller.DisplayType.String,
			DisplayManufacturerQrCode: controller.DisplayManufacturerQrCode.String,
			ArticleNumber:             controller.ArticleNumber.String,
			InfoQrCode:                controller.InfoQrCode.String,
			Can1Terminated:            controller.Can1Terminated.Bool,
			Can2Terminated:            controller.Can2Terminated.Bool,
			Can3Terminated:            controller.Can3Terminated.Bool,
			Can4Terminated:            controller.Can4Terminated.Bool,
			Usb:                       controller.Usb.Bool,
			Serial:                    controller.Serial.Bool,
			ManufacturerQrCodes:       controller.ManufacturerQrCode.String,
			OrderID:                   controller.OrderID.UUID,
			CreatedAt:                 controller.CreatedAt.Time,
			UpdatedAt:                 controller.UpdatedAt.Time,
			SlotPinoutJson:            controller.SlotPinoutJson.RawMessage,
			IsDeleted:                 controller.IsDeleted.Bool,
		})
	}
	return controllersJson
}

func validateCanTerminationConfParam(newController pkg.NewControllerJsonFromFrontend) (database.GetTerminationConfigIdByCanTerminatedParams, error) {
	//TODO: add validation
	var canParam database.GetTerminationConfigIdByCanTerminatedParams
	canErr := canParam.Can1Terminated.Scan(newController.Can1Terminated)
	canErr = canParam.Can2Terminated.Scan(newController.Can2Terminated)
	canErr = canParam.Can3Terminated.Scan(newController.Can3Terminated)
	canErr = canParam.Can4Terminated.Scan(newController.Can4Terminated)
	return canParam, canErr
}

func validateNewMiniPcieParams(newController pkg.NewControllerJsonFromFrontend) (database.CreateMiniPCIeModuleParams, error) {
	//TODO: add validation
	var newMiniPcieParams database.CreateMiniPCIeModuleParams
	newMiniPcieErr := newMiniPcieParams.MiniPcieModulesTypeID.Scan(newController.MiniPcieTypeID.String())
	newMiniPcieErr = newMiniPcieParams.SerialNumber.Scan(newController.MiniPcieSerialNumber)
	return newMiniPcieParams, newMiniPcieErr
}

func validateNewControllerParams(newController pkg.NewControllerJsonFromFrontend) (database.CreateControllerParams, error) {
	//TODO: add validation
	var newControllerParams database.CreateControllerParams
	newControllerErr := newControllerParams.ControllerTypesID.Scan(newController.TypesID.String())	
	newControllerErr = newControllerParams.ProjectsID.Scan(newController.ProjectsID.String())
	newControllerErr = newControllerParams.Description.Scan(*newController.Description)
	newControllerErr = newControllerParams.ControllersPcbHwVersionsID.Scan(newController.PcbHwVersionsID.String())
	newControllerErr = newControllerParams.SerialNumber.Scan(newController.SerialNumber)
	newControllerErr = newControllerParams.ManufacturersID.Scan(newController.ManufacturersID.String())
	newControllerErr = newControllerParams.MacAddress.Scan(newController.MacAddress)
	newControllerErr = newControllerParams.SimNumber.Scan(*newController.SimNumber)
	newControllerErr = newControllerParams.MiniPcieModulesID.Scan(newController.MiniPcieID.String()) // pkg.MiniPcie!
	newControllerErr = newControllerParams.M2ModulesTypesID.Scan(newController.M2ModulesTypesID.String())
	newControllerErr = newControllerParams.DisplayAdaptersID.Scan(newController.DisplayId.String()) // Display!
	newControllerErr = newControllerParams.ArticleNumber.Scan(*newController.ArticleNumber)
	newControllerErr = newControllerParams.InfoQrCode.Scan(*newController.InfoQrCode)
	newControllerErr = newControllerParams.CanTerminationConfsID.Scan(newController.CanTerminationID.String()) // pkg.CanTermination!
	newControllerErr = newControllerParams.Usb.Scan(newController.Usb)
	newControllerErr = newControllerParams.Serial.Scan(newController.Serial)
	newControllerErr = newControllerParams.ManufacturerQrCode.Scan(*newController.ManufacturerQrCode)
	newControllerErr = newControllerParams.OrderID.UUID.Scan(newController.OrderID.String())

	return newControllerParams, newControllerErr
}

func validateNevLedBoardParams(ledBoardQR sql.NullString, controllerId uuid.UUID) (database.CreateLedBoardParams, error) {
	//TODO: add validation
	var newLedBoardParams database.CreateLedBoardParams
	newLedBoardErr := newLedBoardParams.ControllersID.Scan(controllerId.String())
	newLedBoardErr = newLedBoardParams.ManufacturerQrCode.Scan(ledBoardQR.String)
	return newLedBoardParams, newLedBoardErr
}

func validateNewDisplayParams(newDisplay pkg.NewControllerJsonFromFrontend) (database.CreateDisplayAdapterParams, error) {
	//TODO: add validation
	var newDisplayParams database.CreateDisplayAdapterParams
	newDisplayErr := newDisplayParams.DisplayAdaptersTypeName.Scan(newDisplay.DisplayTypeName)
	newDisplayErr = newDisplayParams.ManufacturerQrCode.Scan(newDisplay.DisplayManufacturerQrCode)
	return newDisplayParams, newDisplayErr
}

func validateNewEncloserParams(newController pkg.NewControllerJsonFromFrontend) (database.CreateEncloserParams, error) {
	//TODO: add validation
	var newEncloserParams database.CreateEncloserParams
	newEncloserErr := newEncloserParams.ControllerID.Scan(newController.Id.String())
	newEncloserErr = newEncloserParams.SerialNumber.Scan(newController.EncloserSerialNumber)
	newEncloserErr = newEncloserParams.ControllerTypesID.Scan(newController.TypesID.String())
	newEncloserErr = newEncloserParams.ManufacturersID.Scan(newController.EncloserManufacturerID.String())
	return newEncloserParams, newEncloserErr
}
