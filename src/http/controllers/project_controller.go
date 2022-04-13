package controllers

import (
	"fmt"

	"github.com/babyfaceeasy/egroup-test/src/repositories"
	"github.com/babyfaceeasy/egroup-test/src/services"
	"github.com/gofiber/fiber/v2"
)


func GetLastProjects(c *fiber.Ctx) error  {
	repo := &repositories.ProjectRepository{}
	projectService := &services.ProjectService{Repo: repo}
	projects, err := projectService.GetLastProjects(4)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": "An error occurred, please try again later.",
		})
	}

	fmt.Printf("#%+v\n", projects)

	return c.JSON(fiber.Map{
		"error": false,
		"data": projects,
	})
}