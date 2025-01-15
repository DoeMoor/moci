package api

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/google/uuid"
	// "github.com/doemoor/moci/internal/database"
)

func (cnf *ApiConfig) GetAllMiniPCeModulesType(c *fiber.Ctx) error {
	allMiniPcieModules, err := cnf.DbQueries.GetAllMiniPCeModulesType(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		return err
	}

	type miniPcieModulesForJson struct {
		Name string `json:"name"`
		ModuleTypeNumber int `json:"moduleTypeNumber"`
	}

	var miniPcieModulesJson []miniPcieModulesForJson

	for _, miniPcieModule := range allMiniPcieModules {
		miniPcieModulesJson = append(miniPcieModulesJson, miniPcieModulesForJson{
			Name:   miniPcieModule.Name.String,
			ModuleTypeNumber: int(miniPcieModule.ModuleTypeNumber.Int32),
		})
	}

	return c.JSON(miniPcieModulesJson)
}

func (cnf *ApiConfig) GetMiniPCeModuleById(c *fiber.Ctx) error {
	return nil
}

func (cnf *ApiConfig) GetAllMiniPCeModules(c *fiber.Ctx) error {
	return nil
}

func (cnf *ApiConfig) CreateMiniPCeModule(c *fiber.Ctx) error {
	return nil
}

