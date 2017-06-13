package postgresql

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User model
type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// BeforeCreate an user
func (user *User) BeforeCreate(pScope *gorm.Scope) error {
	user.ID = uuid.NewV4().String()

	return nil
}
