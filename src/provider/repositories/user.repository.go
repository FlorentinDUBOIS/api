package repositories

import (
  "os"

  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres" // import postgresql driver needed by gorm
  "github.com/sirupsen/logrus"
  "gitlab.com/FlorentinDUBOIS/api/src/provider/postgresql"
)

var database *gorm.DB

func init() {
  db, error := gorm.Open("postgres", os.Getenv("POSTGRESQL_URI"))

  if error != nil {
    logrus.Panic(error)
  }

  if gin.Mode() == gin.DebugMode {
    db.LogMode(true)
  } else {
    db.LogMode(false)
  }

  database = db
}

// UserRepository structure
type UserRepository struct{}

// Find users
func (userRepository *UserRepository) Find() []postgresql.User {
  users := []postgresql.User{}

  logrus.Info("Retrieve information from users")
  database.Find(&users)

  logrus.Debug("Users retrieved from users", users)
  return users
}

// FindByUuid an user
func (userRepository *UserRepository) FindByUuid(pUuid string) postgresql.User {
  user := postgresql.User{}

  logrus.WithField("uuid", pUuid).Info("Retrieve an user from users")
  database.Where("id = ?", pUuid).Find(&user)

  logrus.WithField("user", user).Debug("User retrieved from user")
  return user
}

// Save an user
func (userRepository *UserRepository) Save(pUser postgresql.User) postgresql.User {
  tx := database.Begin()

  logrus.Info("Persist user in database")
  if database.NewRecord(pUser) {

    logrus.WithField("user", pUser).Debug("Create new user")
    if error := tx.Create(&pUser).Error; error != nil {
      logrus.Error(error)
      tx.Rollback()
    } else {
      tx.Commit()
    }
  } else {

    logrus.WithField("user", pUser).Debug("Update user")
    if error := tx.Save(&pUser).Error; error != nil {
      logrus.Error(error)
      tx.Rollback()
    } else {
      tx.Commit()
    }
  }

  logrus.WithField("user", pUser).Debug("Persited user in database")
  return pUser
}

func (*UserRepository) Delete(pUuid string) error {
  return database.Where("id = ?", pUuid).Delete(&postgresql.User{}).Error
}
