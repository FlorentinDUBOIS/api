package repositories

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var database *gorm.DB

// InitDatabase function
func InitDatabase() {
	host := viper.GetString("postgres-host")
	user := viper.GetString("postgres-user")
	password := viper.GetString("postgres-password")
	sslmode := viper.GetString("postgres-sslmode")
	dbname := viper.GetString("postgres-dbname")

	uri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", host, user, dbname, sslmode, password)

	db, error := gorm.Open("postgres", uri)

	if error != nil {
		log.Panic(error)
	}

	db.LogMode(viper.GetBool("verbose"))

	database = db
}
