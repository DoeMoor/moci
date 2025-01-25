package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (cnf *ApiConfig) GetAllMiniPCeModulesType(c *fiber.Ctx) error {
	allMiniPcieModules, err := cnf.DbQ.GetAllMiniPCIeModulesType(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		return err
	}

	type miniPcieModulesForJson struct {
		Id               string `json:"id"`
		Name             string `json:"name"`
		ModuleTypeNumber int    `json:"moduleTypeNumber"`
		FullName         string `json:"fullName"`
	}

	var miniPcieModulesJson []miniPcieModulesForJson

	for _, miniPcieModule := range allMiniPcieModules {
		miniPcieModulesJson = append(miniPcieModulesJson, miniPcieModulesForJson{
			Id:               miniPcieModule.ID.String(),
			Name:             miniPcieModule.Name.String,
			ModuleTypeNumber: int(miniPcieModule.ModuleTypeNumber.Int32),
			FullName:         fmt.Sprint(miniPcieModule.Name.String, " ", miniPcieModule.ModuleTypeNumber.Int32),
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
