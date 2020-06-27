package validations

import (
	"testing"

	"github.com/andrewesteves/taskee-api/entities"
)

func TestUserFields(t *testing.T) {
	user := entities.User{
		Name:     "Andrew Esteves",
		Email:    "andrewesteves@github.com",
		Password: "12345678",
	}

	actual := len(UserRegister(user))
	expected := 0

	if actual != expected {
		t.Errorf("actual: %d, expected: %d", actual, expected)
	}
}

func TestUserWithoutPassword(t *testing.T) {
	user := entities.User{
		Name:  "Andrew Esteves",
		Email: "andrewesteves@github.com",
	}

	actual := len(UserRegister(user))
	expected := 1

	if actual != expected {
		t.Errorf("actual: %d, expected: %d", actual, expected)
	}
}
