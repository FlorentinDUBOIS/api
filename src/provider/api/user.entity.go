package api

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User to use to exchange with the exterior
type User struct {
	ID        *string
	FirstName *string
	LastName  *string
	Email     *string
	Password  *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// EncryptPassword of the user
func (user *User) EncryptPassword() error {
	password, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	passwd := string(password)

	user.Password = &passwd

	return nil
}
