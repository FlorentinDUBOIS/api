package services

import (
	"github.com/FlorentinDUBOIS/bouncer/src/provider/api"
	"github.com/FlorentinDUBOIS/bouncer/src/provider/repositories"
	"github.com/FlorentinDUBOIS/bouncer/src/services/factories"
)

var userRepository = new(repositories.UserRepository)
var apiUserFactory = new(factories.APIUserFactory)
var pgUserFactory = new(factories.PostgresUserFactory)

// UserService structure
type UserService struct{}

// Find users
func (*UserService) Find() []*api.User {
	return apiUserFactory.FromPostgresArray(userRepository.Find())
}

// FindOne user
func (*UserService) FindOne(pUUID string) *api.User {
	return apiUserFactory.FromPostgres(userRepository.FindByUUID(pUUID))
}

// Save an user
func (*UserService) Save(pUser *api.User) (*api.User, error) {
	if err := pUser.EncryptPassword(); err != nil {
		return nil, err
	}

	user, err := userRepository.Save(pgUserFactory.FromAPI(pUser))
	if err != nil {
		return nil, err
	}

	return apiUserFactory.FromPostgres(user), nil
}

// Update an user
func (*UserService) Update(pUUID string, pUser *api.User) (*api.User, error) {
	if pUser.Password != nil {
		if err := pUser.EncryptPassword(); err != nil {
			return nil, err
		}
	}

	user := pgUserFactory.Compose(userRepository.FindByUUID(pUUID), pUser)

	_, err := userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return apiUserFactory.FromPostgres(user), nil
}

// Delete an user
func (*UserService) Delete(pUUID string) error {
	return userRepository.Delete(pUUID)
}
