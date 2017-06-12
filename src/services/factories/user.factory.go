package factories

import (
	"github.com/sirupsen/logrus"
	"github.com/FlorentinDUBOIS/api/src/provider/postgresql"
)

// UserFactory structure
type UserFactory struct{}

// Assign user to another one
func (*UserFactory) Assign(pDest *postgresql.User, pSrc *postgresql.User) *postgresql.User {
	logrus.WithField("dest", pDest).WithField("src", pSrc).Debug("UserFactory; Start assignment")

	if pDest != nil && pSrc != nil {
		if pSrc.ID != "" {
			pDest.ID = pSrc.ID
		}

		if pSrc.FirstName != "" {
			pDest.FirstName = pSrc.FirstName
		}

		if pSrc.LastName != "" {
			pDest.LastName = pSrc.LastName
		}

		if pSrc.Email != "" {
			pDest.Email = pSrc.Email
		}

		if pSrc.Password != "" {
			pDest.Password = pSrc.Password
		}

		if pSrc.CreatedAt.String() != "" {
			pDest.CreatedAt = pSrc.CreatedAt
		}

		if pSrc.UpdatedAt.String() != "" {
			pDest.UpdatedAt = pSrc.UpdatedAt
		}

		pDest.DeletedAt = pSrc.DeletedAt
	}

	logrus.WithField("dest", pDest).WithField("src", pSrc).Debug("UserFactory; Assignment done")
	return pDest
}
