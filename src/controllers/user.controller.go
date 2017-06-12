package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/FlorentinDUBOIS/api/src/provider/postgresql"
	"github.com/FlorentinDUBOIS/api/src/services"
)

var userService = services.UserService{}
var userController = UserController{}

// UserController structure
type UserController struct{}

// Register routes
func (*UserController) Register(pRouter *gin.RouterGroup) {
	pRouter.Use()
	{
		pRouter.GET("/", userController.Find)
		pRouter.POST("/", userController.Create)
		pRouter.GET("/:uuid", userController.FindOne)
		pRouter.PUT("/:uuid", userController.Update)
		pRouter.DELETE("/:uuid", userController.Delete)
	}
}

// Find users
func (*UserController) Find(pContext *gin.Context) {
	logrus.Info("Handle search of users")
	pContext.JSON(http.StatusOK, userService.Find())
}

// Create an user
func (*UserController) Create(pContext *gin.Context) {
	user := &postgresql.User{}

	if error := pContext.BindJSON(user); error != nil {
		logrus.Error(error)
		pContext.JSON(http.StatusBadRequest, error)

		return
	}

	logrus.WithField("user", user).Info("Handle creation of user")

	if savedUser, error := userService.Save(user); error == nil {
		logrus.WithField("user", savedUser).Debug("Create new user")
		pContext.JSON(http.StatusOK, user)
	} else {
		logrus.Error(error)
		pContext.AbortWithError(http.StatusInternalServerError, error)
	}
}

// FindOne user
func (*UserController) FindOne(pContext *gin.Context) {
	uuid := pContext.Param("uuid")

	logrus.WithField("uuid", uuid).Info("Handle search of user")
	pContext.JSON(http.StatusOK, userService.FindOne(uuid))
}

// Update user
func (*UserController) Update(pContext *gin.Context) {
	user := &postgresql.User{}
	uuid := pContext.Param("uuid")

	if error := pContext.BindJSON(user); error != nil {
		logrus.Error(error)
		pContext.AbortWithError(http.StatusBadRequest, error)

		return
	}

	logrus.WithField("uuid", uuid).WithField("user", user).Info("Handle update of user")
	if updatedUser, error := userService.Update(uuid, user); error == nil {
		logrus.WithField("user", updatedUser).Debug("Updated user")
		pContext.JSON(http.StatusOK, updatedUser)
	} else {
		logrus.Error(error)
		pContext.AbortWithError(http.StatusInternalServerError, error)
	}
}

// Delete use
func (*UserController) Delete(pContext *gin.Context) {
	uuid := pContext.Param("uuid")

	logrus.WithField("uuid", uuid).Info("Delete user")
	if error := userService.Delete(uuid); error != nil {
		logrus.Error(error)
		pContext.AbortWithError(http.StatusInternalServerError, error)

		return
	}

	pContext.AbortWithStatus(http.StatusNoContent)
}
