package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelly-tools/core/config"
	log "github.com/sirupsen/logrus"
)

var (
	GlobalConfig *config.Config
	LogInstance  *log.Logger
)

func init() {
	// Prepare logging environment

	LogInstance = log.New()

	LogInstance.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	LogInstance.SetLevel(log.ErrorLevel)
}

func init() {
	var err error

	// read config file and create a new Config Struct
	fileData, err := ioutil.ReadFile("core_config.yaml")

	if err != nil {
		LogInstance.Fatal(err)
	}

	// Generate new config with all defaults
	GlobalConfig, err = config.New(fileData)

	if err != nil {
		LogInstance.Println(err)
	}

	fmt.Println("Config loaded", GlobalConfig)

	// Set correct log Level

	var logLevel log.Level
	switch GlobalConfig.Debugging.Logging.LogLevel {
	case "debug":
		logLevel = log.DebugLevel
	case "info":
		logLevel = log.InfoLevel
	case "error":
		logLevel = log.ErrorLevel
	default:
		logLevel = log.DebugLevel
	}

	LogInstance.SetLevel(logLevel)
}

func main() {
	// Prepare GinMode
	var ginMode string

	switch GlobalConfig.Debugging.Router.Mode {
	case "PROD":
		LogInstance.Debugln("Set router Mode to PROD")
		ginMode = gin.ReleaseMode
	case "DEV":
		LogInstance.Debugln("Set router Mode to DEV")
		ginMode = gin.TestMode
	}

	gin.SetMode(ginMode)
	gin.Default().AppEngine = GlobalConfig.Debugging.Router.AppEngine

	router := gin.Default()

	router.Static("/assets", "ui/assets")
	router.Static("/"+GlobalConfig.ImageStorePath, GlobalConfig.ImageStorePath)
	router.LoadHTMLGlob("ui/templates/*")

	app := router.Group("/app")
	apiV1 := router.Group("/api/v1")
	apiV1.Use(CORS)

	prepareAPIV1(apiV1)
	prepareApp(app)

	router.Run(GlobalConfig.UI.ListenAdress + ":" + fmt.Sprint(GlobalConfig.UI.ListenPort))
}

// CORS Middleware
func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
