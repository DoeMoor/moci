package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (cnf *ApiCfg) GetALLCustomers(c *fiber.Ctx) error {

	allCustomers, err := cnf.DbQ.GetAllCustomers(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		log.Println("get all customers error:", err)
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
