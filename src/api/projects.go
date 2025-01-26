package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (cnf *ApiCfg) GetALLProjects(c *fiber.Ctx) error {

	allProjects, err := cnf.DbQ.GetAllProjects(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		log.Println("get all projects error:", err)
		return c.SendString("db error")
	}

	type projectJson struct {
		Id   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}

	var result []projectJson

	for _, project := range allProjects {
		result = append(result, projectJson{
			Id:   project.ID,
			Name: project.Name.String,
		})

	}

	return c.JSON(result)

}
