package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/doemoor/moci/api"
	"github.com/doemoor/moci/internal/database"
	"github.com/doemoor/moci/internal/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
		os.Exit(1)
	}

	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Println("DB_URL is not set", ":", dbString)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var apiConf = &api.ApiConfig{
		DbQueries: database.New(db),
	}

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	serverReadTimeout := os.Getenv("SERVER_READ_TIMEOUT")
	if serverHost == "" ||
		serverPort == "" ||
		serverReadTimeout == "" {
		log.Printf("SERVER_HOST or SERVER_PORT or SERVER_READ_TIMEOUT are not set:\n"+
			"SERVER_HOST: %s\n"+
			"SERVER_PORT: %s\n"+
			"SERVER_READ_TIMEOUT: %s\n", serverHost, serverPort, serverReadTimeout)
	}

	app := fiber.New(fiber.Config{
		ServerHeader: "Moduline Controller Inventory",
		AppName:      "Moduline Controller Inventory",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:8080, https://localhost:3000, https://localhost:8080, https://localhost:8081, http://localhost:8081",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	app.Static("/", "./app")

	app.Get("/api/controllers", apiConf.GetAllControllers)
	app.Get("/api/controllers/types", apiConf.GetAllControllerTypes)
	app.Get("/api/controllers/:id", apiConf.GetControllerById)

	app.Get("/api/iomodules", apiConf.GetAllIoModules)
	app.Get("/api/iomodules/:id", apiConf.GetIoModuleById)

	app.Get("/api/manufacturers", apiConf.GetAllManufacturers)
	app.Get("/api/manufacturers/:id", apiConf.GetManufacturerById)
	app.Post("/api/manufacturers", apiConf.CreateManufacturer)

	// app.Get("/api/miniPcieModules", apiConf.GetAllMiniPCeModules)
	app.Get("/api/miniPcieModules/types", apiConf.GetAllMiniPCeModulesType)
	// app.Get("/api/miniPcieModules/:id", apiConf.GetMiniPCeModuleById)
	// app.Post("/api/miniPcieModules", apiConf.CreateMiniPCeModule)

	app.Get("/api/m2Modules", apiConf.GetAllm2Modules)
	app.Get("/api/m2Modules/types", apiConf.GetAllm2ModulesType)
	// app.Get("/api/m2Modules/:id", apiConf.GetM2ModuleById)
	app.Post("/api/m2Modules", apiConf.Create2Module)

	utility.ClearTerminal()
	log.Fatal(app.Listen(serverHost + ":" + serverPort))
}
