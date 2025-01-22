package api

import (
	"database/sql"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (cnf *ApiConfig) GetAllManufacturers(c *fiber.Ctx) error {
	allManufacturers, err := cnf.DbQueries.GetAllManufacturers(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		return err
	}

	var manufacturersJSON []struct {
		ID   string `json:"id"`
		Name string    `json:"name"`
	}

	for _, manufacturer := range allManufacturers {
		manufacturersJSON = append(manufacturersJSON, struct {
			ID   string `json:"id"`
			Name string    `json:"name"`
		}{
			ID:   manufacturer.ID.String(),
			Name: manufacturer.Name.String,
		})
	}

	return c.JSON(manufacturersJSON)
}

func (cnf *ApiConfig) GetManufacturerById(c *fiber.Ctx) error {

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

	manufacturerById, err := cnf.DbQueries.GetManufacturerById(c.Context(), uuid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.Response().SetStatusCode(404)
			return c.SendString("manufacturer not found")
		}
		c.Response().SetStatusCode(500)
		log.Println("Error getting manufacturer by id from database: ", err)
		return err		
	}
	return c.JSON(manufacturerById)
}

func (cnf *ApiConfig) CreateManufacturer(c *fiber.Ctx) error {

	if c.Get("Content-Type") != "application/json" {
		c.Response().SetStatusCode(400)	
		return c.SendString("Content-Type must be application/json")
	}

	var body struct {
		Name string `json:"name"`
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	manufacturer, err := cnf.DbQueries.CreateManufacturer(c.Context(), sql.NullString{
		String: body.Name,
		Valid:  true,})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.Response().SetStatusCode(400)
			return c.SendString("manufacturer already exists")
		}
		c.Response().SetStatusCode(500)
		return err
	}

	return c.JSON(manufacturer)
}
