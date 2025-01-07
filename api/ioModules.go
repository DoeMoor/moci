package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	// "github.com/doemoor/moci/internal/database"
)

func (cnf *ApiConfig) GetAllIoModules(c *fiber.Ctx) error {
	allIoModules, err := cnf.DbQueries.GetAllIoModules(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		return err
	}

	return c.JSON(allIoModules)
}

func (cnf *ApiConfig) GetIoModuleById(c *fiber.Ctx) error {
	if c.Params("id") == "" {
		c.Response().SetStatusCode(404)
		log.Println(c.OriginalURL()," id is empty")
		return c.SendString("id is empty")
	}

	uuid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		c.Response().SetStatusCode(404)
		log.Println("error parsing uuid: ",c.OriginalURL(),"\n","  error: ",err)
		return c.SendString("wrong uuid")
	}

	ioModuleById, err := cnf.DbQueries.GetIoModuleById(c.Context(), uuid)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.Response().SetStatusCode(404)
			return c.SendString("io module not found")
		}
		c.Response().SetStatusCode(500)
		log.Println("Error getting io module by id from database: ",err)
		return c.SendString("500")
	}

	return c.JSON(ioModuleById)
}