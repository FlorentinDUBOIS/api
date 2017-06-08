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
func (*UserService) Find() []postgresql.User {
  return userRepository.Find()
}

// FindOne user
func (*UserService) FindOne(pUuid string) postgresql.User {
  return userRepository.FindByUuid(pUuid)
}

// Save an user
func (*UserService) Save(user postgresql.User) postgresql.User {
  return userRepository.Save(user)
}

// Update an user
func (*UserService) Update(pUuid string, pUser postgresql.User) postgresql.User {
  user := userFactory.Assign(userRepository.FindByUuid(pUuid), pUser)

  return userRepository.Save(user)
}

// Delete an user
func (*UserService) Delete(pUuid string) error {
  return userRepository.Delete(pUuid)
}
