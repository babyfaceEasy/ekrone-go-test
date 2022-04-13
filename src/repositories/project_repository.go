package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/babyfaceeasy/egroup-test/src/dtos"
	"github.com/babyfaceeasy/egroup-test/src/models"
)

type Repository interface {
	GetLastProjects(int) ([]models.Project, error)
}

type ProjectRepository struct{}

func (pr *ProjectRepository) GetLastProjects(count int) ([]models.Project, error) {
	
	var projects []models.Project

	// connect to gitlab api
	gitlabAPI := os.Getenv("GITLAB_URL")
	url := fmt.Sprintf("%sprojects?order_by=created_at&per_page=%d", gitlabAPI, count)
	
	// build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("New Request: ", err)
		return nil, err
	}

	// create http client.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Do: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	var projectDTOs []dtos.Project
	if err := json.NewDecoder(resp.Body).Decode(&projectDTOs); err != nil {
		log.Println(err)
		return nil, err
	}

	// convert DTOs into projectModels
	for _, projectDTO := range projectDTOs {
		project := projectDTO.ConvertIntoProjectModel()
		projects = append(projects, project)
	}
	fmt.Printf("#%v\n", projects)

	return projects, nil
}
