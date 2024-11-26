package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/shohan-pherones/blood-bank-management.git/database"
	"github.com/shohan-pherones/blood-bank-management.git/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()

	database.ConnectToMongoDB()

	app.Get("/api/v1/health", func(c *fiber.Ctx) error {
		return utils.SendResponse(c, fiber.StatusOK, "Server is running healthy!", nil)
	})

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		log.Println("Shutting down server...")

		if err := app.Shutdown(); err != nil {
			log.Fatalf("Error shutting down server: %v", err)
		}
	}()

	log.Fatal(app.Listen(":4000"))
}
