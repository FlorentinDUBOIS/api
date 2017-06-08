package postgresql

import (
  "time"

  "github.com/jinzhu/gorm"
  uuid "github.com/satori/go.uuid"
  "github.com/sirupsen/logrus"
  "golang.org/x/crypto/bcrypt"
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

// EncryptPassword of the user
func (user *User) EncryptPassword() error {
  password, error := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

  if error != nil {
    logrus.Error(error)
    return error
  }

  user.Password = string(password)

  return nil
}

// BeforeCreate an user
func (user *User) BeforeCreate(pScope *gorm.Scope) error {
  error := user.EncryptPassword()

  if error == nil {
    id := uuid.NewV4().String()

    logrus.WithField("uuid", id).Info("Set identifier")
    user.ID = id
  }

  return error
}
