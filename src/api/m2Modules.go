package api

import (
	"fmt"
	"log"

	"database/sql"

	"strings"

	"github.com/doemoor/moci/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type m2ModulesTypeForJson struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	ModuleTypeNumber int    `json:"moduleTypeNumber"`
	FullName         string `json:"fullName"`
}

func (cnf *ApiCfg) GetAllm2ModulesType(c *fiber.Ctx) error {
	DBm2ModulesTypes, err := cnf.DbQ.GetAllM2ModulesTypes(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		log.Println(err)
		return err
	}

	var m2ModulesTypeJson []m2ModulesTypeForJson

	for _, m2ModuleType := range DBm2ModulesTypes {
		m2ModulesTypeJson = append(m2ModulesTypeJson, m2ModulesTypeForJson{
			Id:               m2ModuleType.ID.String(),
			Name:             m2ModuleType.Name.String,
			ModuleTypeNumber: int(m2ModuleType.ModuleTypeNumber.Int32),
			FullName:         fmt.Sprint(m2ModuleType.Name.String, " ", m2ModuleType.ModuleTypeNumber.Int32),
		})
	}

	return c.JSON(m2ModulesTypeJson)
}

func (cnf *ApiCfg) GetAllm2Modules(c *fiber.Ctx) error {

	allM2Modules, err := cnf.DbQ.GetAllM2Modules(c.Context())
	if err != nil {
		c.Response().SetStatusCode(500)
		return err
	}
	type m2ModulesForJson struct {
		Id               uuid.UUID `json:"id"`
		Name             string    `json:"name"`
		ModuleTypeNumber int       `json:"moduleTypeNumber"`
	}

	var m2ModulesJson []m2ModulesForJson

	for _, m2Module := range allM2Modules {
		m2ModulesJson = append(m2ModulesJson, m2ModulesForJson{
			Id:               m2Module.ID,
			Name:             m2Module.Name.String,
			ModuleTypeNumber: int(m2Module.ModuleTypeNumber.Int32),
		})
	}

	return c.JSON(m2ModulesJson)
}

func (cnf *ApiCfg) Create2Module(c *fiber.Ctx) error {

	if c.Get("Content-Type") != "application/json" {
		c.Response().SetStatusCode(400)
		return c.SendString("Content-Type must be application/json")
	}

	type newM2Module struct {
		Name             string `json:"name"`
		ModuleTypeNumber int    `json:"moduleTypeNumber"`
	}

	var body newM2Module

	if err := c.BodyParser(&body); err != nil {
		log.Println(err)
		log.Println(string(c.Request().Body()))
		return c.SendString(err.Error())
	}

	param := database.CreateM2ModuleParams{
		Name: sql.NullString{
			String: body.Name,
			Valid:  true,
		},
		ModuleTypeNumber: sql.NullInt32{
			Int32: int32(body.ModuleTypeNumber),
			Valid: true,
		},
	}

	resultOfInsert, err := cnf.DbQ.CreateM2Module(c.Context(), param)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.Response().SetStatusCode(400)
			log.Println("Error creating m2 module: ", err)
			return c.SendString("m2 module already exists")
		}
		c.Response().SetStatusCode(500)
		log.Println("Error creating m2 module: ", err)
		return c.SendString("internal server error")
	}

	type m2ModuleForJson struct {
		Id               uuid.UUID `json:"id"`
		Name             string    `json:"name"`
		ModuleTypeNumber int       `json:"moduleTypeNumber"`
	}

	resultOfInsertForJson := m2ModuleForJson{
		Id:               resultOfInsert.ID,
		Name:             resultOfInsert.Name.String,
		ModuleTypeNumber: int(resultOfInsert.ModuleTypeNumber.Int32),
	}

	return c.JSON(resultOfInsertForJson)
}
