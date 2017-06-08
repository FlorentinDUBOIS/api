package factories

import (
  "github.com/sirupsen/logrus"
  "gitlab.com/FlorentinDUBOIS/api/src/provider/postgresql"
)

// UserFactory structure
type UserFactory struct{}

// Assign user to another one
func (*UserFactory) Assign(pDest postgresql.User, pSrc postgresql.User) postgresql.User {
  logrus.WithField("dest", pDest).WithField("src", pSrc).Debug("UserFactory; Start assignment")

  if len(pSrc.ID) > 0 {
    pDest.ID = pSrc.ID
  }

  if len(pSrc.FirstName) > 0 {
    pDest.FirstName = pSrc.FirstName
  }

  if len(pSrc.LastName) > 0 {
    pDest.FirstName = pSrc.FirstName
  }

  if len(pSrc.Email) > 0 {
    pDest.Email = pSrc.Email
  }

  if len(pSrc.Password) > 0 {
    pDest.Password = pSrc.Password

    if error := pDest.EncryptPassword(); error != nil {
      logrus.Error(error)
    }
  }

  pDest.DeletedAt = pSrc.DeletedAt

  logrus.WithField("dest", pDest).WithField("src", pSrc).Debug("UserFactory; Assignment done")
  return pDest
}
