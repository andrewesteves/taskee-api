package validations

import (
	"testing"

	"github.com/andrewesteves/taskee-api/entities"
)

func TestTaskFields(t *testing.T) {
	task := entities.Task{
		Description: "Awesome task",
		ProjectID:   1,
	}

	actual := len(TaskStore(task))
	expected := 0

	if actual != expected {
		t.Errorf("actual: %d, expected: %d", actual, expected)
	}
}

func TestTaskWithoutProject(t *testing.T) {
	task := entities.Task{
		Description: "Awesome task",
	}

	actual := len(TaskStore(task))
	expected := 1

	if actual != expected {
		t.Errorf("actual: %d, expected: %d", actual, expected)
	}
}
