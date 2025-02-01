package api

import (
	"log"
	"fmt"

	// "github.com/doemoor/moci/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	// "github.com/doemoor/moci/internal/database"
)

func (cnf *ApiCfg) GetAllIoModules(c *fiber.Ctx) error {
	allIoModules, err := cnf.DbQ.GetAllIoModules(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		return err
	}

	return c.JSON(allIoModules)
}

func (cnf *ApiCfg) GetIoModuleById(c *fiber.Ctx) error {
	if c.Params("id") == "" {
		c.Response().SetStatusCode(404)
		log.Println(c.OriginalURL(), " id is empty")
		return c.SendString("GetIoModuleById id is empty")
	}

	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		c.Response().SetStatusCode(404)
		log.Println("error parsing uuid: ", c.OriginalURL(), "\n", "  error: ", err)
		return c.SendString("GetIoModuleById wrong uuid")
	}

	ioModuleById, err := cnf.DbQ.GetIoModuleById(c.Context(), uuid)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.Response().SetStatusCode(404)
			return c.SendString("io module not found")
		}
		c.Response().SetStatusCode(500)
		log.Println("Error getting io module by id from database: ", err)
		return c.SendString("500")
	}

	return c.JSON(ioModuleById)
}

func (cnf *ApiCfg) CreateIoModule(c *fiber.Ctx) error {
	return nil
}

func (cnf *ApiCfg) GetAllIoModuleTypes(c *fiber.Ctx) error {

	allIoModuleTypes, err := cnf.DbQ.GetAllIoModuleTypes(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		log.Println("Error getting all io module types from database: ", err)
		return c.SendString("Error getting all io module types from database")
	}

	type IoModuleTypesJson struct {
		ID               uuid.UUID `json:"id"`
		ModuleTypeName   string    `json:"moduleTypeName"`
		ModuleTypeNumber int       `json:"moduleTypeNumber"`
		FullName         string    `json:"fullName"`
	}

	var allIoModuleTypesJson []IoModuleTypesJson
	for _, ioModuleType := range allIoModuleTypes {
		allIoModuleTypesJson = append(allIoModuleTypesJson, IoModuleTypesJson{
			ID:               ioModuleType.ID,
			ModuleTypeName:   ioModuleType.ModuleTypeName,
			ModuleTypeNumber: int(ioModuleType.ModuleTypeNumber.Int32),
			FullName:         fmt.Sprint(ioModuleType.ModuleTypeName, " ", ioModuleType.ModuleTypeNumber.Int32),
		})
	}
	return c.JSON(allIoModuleTypesJson)
}

func (cnf *ApiCfg) CreateIoModuleType(c *fiber.Ctx) error {
	return nil
}
