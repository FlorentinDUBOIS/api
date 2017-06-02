package controllers

import (
	"net/http"

	"gitlab.com/FlorentinDUBOIS/api/services"

	"github.com/gin-gonic/gin"
)

// FindUsers gin.HandlerFunc
func FindUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, services.FindUsers())
	}
}

// CreateUser gin.HandlerFunc
func CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotImplemented)
	}
}

// FindUser gin.HandlerFunc
func FindUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotImplemented)
	}
}

// UpdateUser gin.HandlerFunc
func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotImplemented)
	}
}

// DeleteUser gin.HandlerFunc
func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotImplemented)
	}
}
