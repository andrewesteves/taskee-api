package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Task entity
type Task struct {
	gorm.Model
	ProjectID   uint
	Project     Project    `json:"project"`
	Description string     `json:"description"`
	CompletedAt *time.Time `json:"completedAt"`
}
