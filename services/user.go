package services

import "github.com/FlorentinDUBOIS/bouncer/provider/api"

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
	if _, err := userRepository.Save(user); err != nil {
		return nil, err
	}

	return apiUserFactory.FromPostgres(user), nil
}

// Delete an user
func (*UserService) Delete(pUUID string) error {
	return userRepository.Delete(pUUID)
}
