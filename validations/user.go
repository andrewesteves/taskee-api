package validations

import (
	"github.com/andrewesteves/taskee-api/entities"
	"github.com/andrewesteves/taskee-api/utils"
)

// UserRegister validations
func UserRegister(user entities.User) []string {
	var messages []string

	if user.Name == "" {
		messages = append(messages, "The field name is required")
	}

	if !utils.CheckEmail(user.Email) {
		messages = append(messages, "The field email is invalid")
	}

	if user.Password == "" || len(user.Password) < 8 {
		messages = append(messages, "The field password is required and should have at least 8 characters")
	}

	return messages
}

// UserLogin validations
func UserLogin(user entities.User) []string {
	var messages []string

	if !utils.CheckEmail(user.Email) {
		messages = append(messages, "The field email is invalid")
	}

	if user.Password == "" || len(user.Password) < 8 {
		messages = append(messages, "The field password is required and should have at least 8 characters")
	}

	return messages
}
