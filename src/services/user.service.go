package services

import (
	"gitlab.com/FlorentinDUBOIS/api/src/provider/postgresql"
	"gitlab.com/FlorentinDUBOIS/api/src/provider/repositories"
	"gitlab.com/FlorentinDUBOIS/api/src/services/factories"
)

var userRepository = repositories.UserRepository{}
var userFactory = factories.UserFactory{}

// UserService structure
type UserService struct{}

// Find users
func (*UserService) Find() []*postgresql.User {
	return userRepository.Find()
}

// FindOne user
func (*UserService) FindOne(pUUID string) *postgresql.User {
	return userRepository.FindByUUID(pUUID)
}

// Save an user
func (*UserService) Save(pUser *postgresql.User) (*postgresql.User, error) {
	return userRepository.Save(pUser)
}

// Update an user
func (*UserService) Update(pUUID string, pUser *postgresql.User) (*postgresql.User, error) {
	user := userFactory.Assign(userRepository.FindByUUID(pUUID), pUser)

	if error := user.EncryptPassword(); error != nil {
		return nil, error
	}

	return userRepository.Save(user)
}

// Delete an user
func (*UserService) Delete(pUUID string) error {
	return userRepository.Delete(pUUID)
}
