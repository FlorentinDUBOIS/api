package cmd

import (
	"fmt"
	"sync"

	"github.com/FlorentinDUBOIS/bouncer/src/controllers"
	"github.com/FlorentinDUBOIS/bouncer/src/provider/repositories"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	formatter "github.com/x-cray/logrus-prefixed-formatter"
)

var once = new(sync.Once)

// RootCmd launch the aggregator agent
var RootCmd = &cobra.Command{
	Use:   "bouncer",
	Short: "Expose authentication and users",
	Run:   bouncer,
}

// init command and options that may be given
func init() {
	cobra.OnInitialize(configurate)

	// register flags
	RootCmd.Flags().Bool("verbose", false, "Set output to verbose")
	RootCmd.Flags().Uint("port", 8080, "Set port to listen")
	RootCmd.Flags().String("config", "", "Set configuration file")
	RootCmd.Flags().String("postgres-host", "127.0.0.1", "Set ip address of postgres")
	RootCmd.Flags().String("postgres-user", "postgres", "Set user of the database")
	RootCmd.Flags().String("postgres-password", "", "Set password of the database")
	RootCmd.Flags().String("postgres-dbname", "bouncer", "Set the name of the database to connect")
	RootCmd.Flags().String("postgres-sslmode", "disable", "Set the sslmode of the postgres client")

	// bind flags from cobra to viper
	viper.BindPFlags(RootCmd.Flags())
}

// configurate behaviour of the application
func configurate() {
	log.SetFormatter(new(formatter.TextFormatter))

	if viper.GetBool("verbose") {
		gin.SetMode(gin.DebugMode)

		log.SetLevel(log.DebugLevel)
		log.Info("Set log level to debug")
	} else {
		gin.SetMode(gin.ReleaseMode)

		log.SetLevel(log.InfoLevel)
		log.Info("Set log level to info")
	}

	// Allow recuperation from the environnement variable
	viper.SetEnvPrefix("bouncer")
	viper.AutomaticEnv()

	// Allow recuperation from file
	viper.AddConfigPath("/etc/bouncer/")
	viper.AddConfigPath("$HOME/.bouncer")
	viper.AddConfigPath(".")

	// Load config from file
	viper.SetConfigName("config")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debug("No config file found")
		} else {
			log.Panicf("Fatal error in config file: %v \n", err)
		}
	}

	// Load user defined config
	if viper.GetString("config") != "" {
		log.Debugf("Load user configuration file: %s", viper.GetString("config"))
		viper.SetConfigFile(viper.GetString("config"))

		if err := viper.ReadInConfig(); err != nil {
			log.Panicf("Fatal error in config file: %v \n", err)
		}
	}

	log.Debug("Initialize the database access")
	once.Do(repositories.Init)
}

// bouncer main function
func bouncer(cmd *cobra.Command, args []string) {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())

	userController := new(controllers.UserController)

	api := router.Group("/api")

	userController.Register(api.Group("/user"))

	port := fmt.Sprintf(":%d", viper.GetInt("port"))

	log.Infof("Listen and serve on :%d", viper.GetInt("port"))
	router.Run(port)
}
