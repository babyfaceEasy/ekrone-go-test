package services_test

import (
	"testing"

	"github.com/babyfaceeasy/egroup-test/src/dtos"
	"github.com/babyfaceeasy/egroup-test/src/models"
	"github.com/babyfaceeasy/egroup-test/src/services"
)

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %v expected %v", got, want)
	}
}

// A good library for testing structures, slices
// https://github.com/google/go-cmp
func assertStruct(t testing.TB, got, want dtos.GetLastProjects) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v expected %+v", got, want)
	}
}



type ProjectRepositoryStub struct{}

func (s *ProjectRepositoryStub) GetLastProjects(count int) ([]models.Project, error) {
	projects := []models.Project{
		{
			Name: "atelier7_grabon",
			Description: "",
			ForksCount: 0,		
		},
		{
			Name: "DeathMatch",
			Description: "",
			ForksCount: 0,		
		},
		{
			Name: "project-with-snippets-5f461691683d063d",
			Description: "",
			ForksCount: 0,		
		},
		{
			Name: "devops-demo",
			Description: "",
			ForksCount: 0,		
		},
	}

	return projects, nil
}

func TestGetLastProjects(t *testing.T){
	repo := &ProjectRepositoryStub{}
	projectService := &services.ProjectService{Repo: repo}
	want := dtos.GetLastProjects{
		Names: "atelier7_grabon, DeathMatch, project-with-snippets-5f461691683d063d, devops-demo",
		ForksCount: 0,
	}
	got, err :=  projectService.GetLastProjects(4)
	// assert err is nil
	assertError(t, err, nil)

	// assert got is equal to want
	assertStruct(t, got, want)

}