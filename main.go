package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/babyfaceeasy/egroup-test/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	app := fiber.New()

	// logger
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":  "Olakunle Odegbaro",
			"email": "oodegbaro@gmail.com",
		})
	})

	routes.ProjectRoutes(app.Group("api/projects"))

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // catch os signals
		<-sigint

		// received an interrupt signal, shutdown server
		if err := app.Shutdown(); err != nil {
			log.Printf("Oops... server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// start server
	if err := app.Listen(os.Getenv("APP_URL")); err != nil {
		log.Printf("Oops... server is not running! Reason: %v", err)
	}
	<-idleConnsClosed
}
