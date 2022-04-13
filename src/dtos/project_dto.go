package dtos

import (
	"strings"

	"github.com/babyfaceeasy/egroup-test/src/models"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ForksCount  int    `json:"forksCount"`
}

func (p *Project) ConvertIntoProjectModel() models.Project {
	return models.Project{
		Name: p.Name,
		Description: p.Description,
		ForksCount: p.ForksCount,
	}
}

type GetLastProjects struct {
	Names      string `json:"names"`
	ForksCount int    `json:"forksCount"`
}

func ConvertIntoGetLastProjects(projects []models.Project) GetLastProjects {
	var projectNames []string
	forksCount := 0

	for _, project := range projects {
		projectNames = append(projectNames, project.Name)
		forksCount += project.ForksCount
	}

	return GetLastProjects{
		Names:      strings.Join(projectNames, ", "),
		ForksCount: forksCount,
	}
}
