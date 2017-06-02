package postgresql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

// User model
type User struct {
	gorm.Model

	UID       string `gorm:"type:UUID;primary_key"`
	FirstName string `gorm:"type:varchar(64)"`
	LastName  string `gorm:"type:varchar(64)"`
	Email     string `gorm:"type:varchar(256);index;unique"`
	Password  string `gorm:"type:varchar(256)"`
}

var database *gorm.DB

func init() {
	db, error := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=api sslmode=disable password=")

	if error != nil {
		logrus.Panic(error)
	}

	database = db
}

// FindUsers from the database
func (*User) FindUsers() []User {
	users := []User{}

	database.Find(&users)

	return users
}
