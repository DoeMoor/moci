package api

import (
	"log"
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v2")

func (cnf *ApiConfig) GetALLCustomers(c *fiber.Ctx) error {

	allCustomers, err := cnf.DbQueries.GetAllCustomers(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		log.Println("get all customers error:" , err)
		return c.SendString("db error")
	}

	type customerJson struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}

	var result []customerJson

	for _, customer := range allCustomers {
		result = append(result, customerJson{
			Id:   customer.ID,
			Name: customer.Name.String,
		})
	}

	return c.JSON(result)

}