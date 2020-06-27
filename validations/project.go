package validations

import (
	"github.com/andrewesteves/taskee-api/entities"
)

// ProjectStore validations
func ProjectStore(project entities.Project) []string {
	var messages []string

	if project.Description == "" {
		messages = append(messages, "The field description is required")
	}

	return messages
}
