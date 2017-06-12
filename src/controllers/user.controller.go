package controllers

import (
	"net/http"

	"github.com/FlorentinDUBOIS/bouncer/src/provider/api"
	"github.com/FlorentinDUBOIS/bouncer/src/services"
	"github.com/gin-gonic/gin"
)

var userService = new(services.UserService)

// UserController structure
type UserController struct{}

// Register routes
func (userController *UserController) Register(pRouter *gin.RouterGroup) {
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
	pContext.JSON(http.StatusOK, userService.Find())
}

// Create an user
func (*UserController) Create(pContext *gin.Context) {
	user := new(api.User)

	if err := pContext.BindJSON(user); err != nil {
		pContext.JSON(http.StatusBadRequest, err)

		return
	}

	if userSaved, err := userService.Save(user); err == nil {
		pContext.JSON(http.StatusOK, userSaved)
	} else {
		pContext.AbortWithError(http.StatusInternalServerError, err)
	}
}

// FindOne user
func (*UserController) FindOne(pContext *gin.Context) {
	uuid := pContext.Param("uuid")

	pContext.JSON(http.StatusOK, userService.FindOne(uuid))
}

// Update user
func (*UserController) Update(pContext *gin.Context) {
	user := new(api.User)
	uuid := pContext.Param("uuid")

	if err := pContext.BindJSON(user); err != nil {
		pContext.AbortWithError(http.StatusBadRequest, err)

		return
	}

	if userUpdated, err := userService.Update(uuid, user); err == nil {
		pContext.JSON(http.StatusOK, userUpdated)
	} else {
		pContext.AbortWithError(http.StatusInternalServerError, err)
	}
}

// Delete user
func (*UserController) Delete(pContext *gin.Context) {
	uuid := pContext.Param("uuid")

	if err := userService.Delete(uuid); err != nil {
		pContext.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	pContext.AbortWithStatus(http.StatusNoContent)
}
