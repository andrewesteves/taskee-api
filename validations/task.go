package validations

import (
	"github.com/andrewesteves/taskee-api/entities"
)

// TaskStore validations
func TaskStore(task entities.Task) []string {
	var messages []string

	if task.Description == "" {
		messages = append(messages, "The field description is required")
	}

	if task.ProjectID < 1 {
		messages = append(messages, "The related project is required")
	}

	return messages
}
