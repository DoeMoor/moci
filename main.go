package main

import (
	// "database/sql"
	"log"
	"os"
	"os/exec"
	"runtime"

	// "github.com/doemoor/moci/api"
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
    log.Println("SERVER_HOST or SERVER_PORT or SERVER_READ_TIMEOUT are not set", serverHost, serverPort, serverReadTimeout)
  }


	// db, err := sql.Open("postgres", dbString)
	// if err != nil {
	// 	log.Println(err)
	// 	os.Exit(1)
	// }

	app := fiber.New(fiber.Config{
    ServerHeader: "Moduline Controller Inventory",
	})

  app.Static("/", "./app")

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

	clearTerminal()
	log.Fatal(app.Listen(serverHost + ":" + serverPort))
}

func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
