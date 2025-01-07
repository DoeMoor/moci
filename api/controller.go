package api

import (
	"log"
	// "time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	// "github.com/doemoor/moci/internal/database"
)

func (cnf *ApiConfig) GetAllControllers(c *fiber.Ctx) error {
	allController, err := cnf.DbQueries.GetAllControllers(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		log.Println(err)
		return err
	}

	return c.JSON(allController)
}

func (cnf *ApiConfig) GetControllerById(c *fiber.Ctx) error {

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

	controllerById, err := cnf.DbQueries.GetControllerById(c.Context(), uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.Response().SetStatusCode(404)
			return c.SendString("controller not found")
		}
		c.Response().SetStatusCode(500)
		log.Println("Error getting controller by id from database: ",err)
		return err
	}
	return c.JSON(controllerById)
}