package repositories

import (
	"github.com/FlorentinDUBOIS/bouncer/provider/postgresql"
	_ "github.com/jinzhu/gorm/dialects/postgres" // import postgresql driver needed by gorm
)

// UserRepository structure
type UserRepository struct{}

// Find users
func (*UserRepository) Find() []*postgresql.User {
	var users []*postgresql.User

	database.Find(&users)

	return users
}

// FindByUUID an user
func (*UserRepository) FindByUUID(pUUID string) *postgresql.User {
	user := new(postgresql.User)

	database.Where("id = ?", pUUID).Find(user)

	return user
}

// FindByEmail an user
func (*UserRepository) FindByEmail(pEmail string) *postgresql.User {
	user := new(postgresql.User)

	database.Where("email = ?", pEmail).Find(user)

	return user
}

// Save an user
func (*UserRepository) Save(pUser *postgresql.User) (*postgresql.User, error) {
	tx := database.Begin()

	if database.NewRecord(pUser) {
		if error := tx.Create(pUser).Error; error != nil {
			tx.Rollback()

			return nil, error
		}
	} else {
		if error := tx.Save(pUser).Error; error != nil {
			tx.Rollback()

			return nil, error
		}
	}

	tx.Commit()

	return pUser, nil
}

// Delete an user
func (*UserRepository) Delete(pUUID string) error {
	return database.Where("id = ?", pUUID).Delete(postgresql.User{}).Error
}
