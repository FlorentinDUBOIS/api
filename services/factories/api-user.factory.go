package factories

import (
	"github.com/FlorentinDUBOIS/bouncer/provider/api"
	"github.com/FlorentinDUBOIS/bouncer/provider/postgresql"
)

// APIUserFactory structure
type APIUserFactory struct{}

// FromPostgres user
func (*APIUserFactory) FromPostgres(pUser *postgresql.User) *api.User {
	var user *api.User

	if pUser != nil {
		user = new(api.User)

		user.ID = &pUser.ID
		user.FirstName = &pUser.FirstName
		user.LastName = &pUser.LastName
		user.Email = &pUser.Email
		user.Password = &pUser.Password
		user.CreatedAt = &pUser.CreatedAt
		user.UpdatedAt = &pUser.UpdatedAt
		user.DeletedAt = pUser.DeletedAt
	}

	return user
}

// FromPostgresArray users
func (factory *APIUserFactory) FromPostgresArray(pUsers []*postgresql.User) []*api.User {
	users := []*api.User{}

	for _, user := range pUsers {
		users = append(users, factory.FromPostgres(user))
	}

	return users
}
