package controllers

import (
	"github.com/FlorentinDUBOIS/bouncer/src/services"
	"github.com/gin-gonic/gin"
)

// AuthController structure
type AuthController struct{}

// Register Handler
func (authController *AuthController) Register(pGroup *gin.RouterGroup) {
	pGroup.Use()
	{
		pGroup.POST("/", services.JWTMiddleware.LoginHandler)
	}
}
