package routes

import (
	"github.com/babyfaceeasy/egroup-test/src/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(router fiber.Router) {
	router.Get("/", controllers.GetLastProjects)
}