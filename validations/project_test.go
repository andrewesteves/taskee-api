package validations

import (
	"testing"

	"github.com/andrewesteves/taskee-api/entities"
)

func TestProjectFields(t *testing.T) {
	project := entities.Project{
		Description: "Awesome project",
	}

	actual := len(ProjectStore(project))
	expected := 0

	if actual != expected {
		t.Errorf("actual: %d, expected: %d", actual, expected)
	}
}

func TestProjectWithoutDescription(t *testing.T) {
	project := entities.Project{}

	actual := len(ProjectStore(project))
	expected := 1

	if actual != expected {
		t.Errorf("actual: %d, expected: %d", actual, expected)
	}
}
