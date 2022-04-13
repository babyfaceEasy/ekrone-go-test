package services

import (
	"github.com/babyfaceeasy/egroup-test/src/dtos"
	"github.com/babyfaceeasy/egroup-test/src/repositories"
)

type ProjectService struct{
	Repo repositories.Repository
}

func (ps *ProjectService) GetLastProjects(count int) (dtos.GetLastProjects, error) {
	/*
	projectRepository := &repositories.ProjectRepository{}
	lastProjects, err := projectRepository.GetLastProjects(count)
	*/
	lastProjects, err := ps.Repo.GetLastProjects(count)
	if err != nil {
		return dtos.GetLastProjects{}, err
	}

	getLastProjects := dtos.ConvertIntoGetLastProjects(lastProjects)
	return getLastProjects, nil
}
