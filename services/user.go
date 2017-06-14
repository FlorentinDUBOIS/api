package services

import (
	"errors"

	"github.com/FlorentinDUBOIS/bouncer/provider/api"
	"github.com/badoux/checkmail"
)

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
	if pUser == nil || pUser.FirstName == nil || pUser.LastName == nil || pUser.Email == nil || pUser.Password == nil {
		return nil, errors.New("You must set a FirstName, LastName, Email and Password to your user in order to be created")
	}

	if err := pUser.EncryptPassword(); err != nil {
		return nil, err
	}

	if err := checkmail.ValidateFormat(*pUser.Email); err != nil {
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
	if pUser == nil {
		return nil, errors.New("You must give at least one user field to update")
	}

	if pUser.Password != nil {
		if err := pUser.EncryptPassword(); err != nil {
			return nil, err
		}
	}

	if pUser.Email != nil {
		if err := checkmail.ValidateFormat(*pUser.Email); err != nil {
			return nil, err
		}
	}

	user := pgUserFactory.Compose(userRepository.FindByUUID(pUUID), pUser)
	if _, err := userRepository.Save(user); err != nil {
		return nil, err
	}

	return apiUserFactory.FromPostgres(user), nil
}

// Delete an user
func (*UserService) Delete(pUUID string) error {
	return userRepository.Delete(pUUID)
}
