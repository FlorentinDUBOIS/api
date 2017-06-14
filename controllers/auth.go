package controllers

import (
	"github.com/FlorentinDUBOIS/bouncer/services"
	"github.com/FlorentinDUBOIS/gin-jwt-forwarder/handlers"
	"github.com/gin-gonic/gin"
)

// AuthController structure
type AuthController struct{}

// Register Handler
func (authController *AuthController) Register(pGroup *gin.RouterGroup) {
	pGroup.Use()
	{
		pGroup.POST("/", services.JWTMiddleware.LoginHandler)

		pGroup.GET("/validate", handlers.Ok)
	}
}
