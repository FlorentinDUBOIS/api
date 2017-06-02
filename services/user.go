package services

import (
	"gitlab.com/FlorentinDUBOIS/api/provider/postgresql"
)

// FindUsers (all)
func FindUsers() []postgresql.User {
	user := postgresql.User{}

	return user.FindUsers()
}
