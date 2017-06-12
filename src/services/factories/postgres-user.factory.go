package factories

import (
	"github.com/FlorentinDUBOIS/bouncer/src/provider/api"
	"github.com/FlorentinDUBOIS/bouncer/src/provider/postgresql"
)

// PostgresUserFactory structure
type PostgresUserFactory struct{}

// FromAPI user
func (*PostgresUserFactory) FromAPI(pUser *api.User) *postgresql.User {
	var user *postgresql.User

	if pUser != nil {
		user = new(postgresql.User)

		if pUser.ID != nil {
			user.ID = *pUser.ID
		}

		if pUser.FirstName != nil {
			user.FirstName = *pUser.FirstName
		}

		if pUser.LastName != nil {
			user.LastName = *pUser.LastName
		}

		if pUser.Email != nil {
			user.Email = *pUser.Email
		}

		if pUser.Password != nil {
			user.Password = *pUser.Password
		}

		if pUser.CreatedAt != nil {
			user.CreatedAt = *pUser.CreatedAt
		}

		if pUser.UpdatedAt != nil {
			user.UpdatedAt = *pUser.UpdatedAt
		}

		user.DeletedAt = pUser.DeletedAt
	}

	return user
}

// Compose users
func (*PostgresUserFactory) Compose(pDest *postgresql.User, pSrc *api.User) *postgresql.User {
	if pDest != nil && pSrc != nil {
		if pSrc.ID != nil {
			pDest.ID = *pSrc.ID
		}

		if pSrc.FirstName != nil {
			pDest.FirstName = *pSrc.FirstName
		}

		if pSrc.LastName != nil {
			pDest.LastName = *pSrc.LastName
		}

		if pSrc.Email != nil {
			pDest.Email = *pSrc.Email
		}

		if pSrc.Password != nil {
			pDest.Password = *pSrc.Password
		}

		if pSrc.CreatedAt != nil {
			pDest.CreatedAt = *pSrc.CreatedAt
		}

		if pSrc.UpdatedAt != nil {
			pDest.UpdatedAt = *pSrc.UpdatedAt
		}

		pDest.DeletedAt = pSrc.DeletedAt
	}

	return pDest
}
