package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/FlorentinDUBOIS/api/src/controllers"
)

var userController = controllers.UserController{}

func main() {
	router := gin.New()

	logrus.SetOutput(os.Stdout)
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())

	api := router.Group("/api")

	userController.Register(api.Group("/user"))

	if port := os.Getenv("PORT"); port == "" {
		logrus.WithField("PORT", 8080).Info("No environment variable PORT, default port is in used")
	} else {
		logrus.WithField("PORT", port).Info("Set port from environment variable PORT")
	}

	logrus.Info("Listen and Serve")
	router.Run()
}
