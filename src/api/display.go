package api

import (
	// "log"
	"github.com/gofiber/fiber/v2"
)

func (cnf *ApiCfg) GetAllDisplayTypes(c *fiber.Ctx) error {

	allDisplay, err := cnf.DbQ.GetAllDisplayTypes(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		return err
	}

	var result []struct {
		Name string `json:"name"`
	}

	for _, display := range allDisplay {
		result = append(result, struct {
			Name string `json:"name"`
		}{
			Name: display,
		})
	}

	return c.JSON(result)
}
