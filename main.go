package main

import (
	"log"
	"os"

	"github.com/doemoor/moci/internal/utility"
  "github.com/doemoor/moci/api"
	"github.com/gofiber/fiber/v2"
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
		log.Println("DB_URL is not set",":", dbString)
		os.Exit(1)
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
    AppName: "Moduline Controller Inventory",
	})

  app.Static("/", "./app")

  app.Get("/api/controllers", api.GetAllControllers)
  app.Get("/api/controllers/:id", api.GetControllerById)
  
  app.Get("/api/iomodules", api.GetAllIoModules)
  app.Get("/api/iomodules/:id", api.GetIoModuleById)
  
  utility.ClearTerminal()
	log.Fatal(app.Listen(serverHost + ":" + serverPort))
}


