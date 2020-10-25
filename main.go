package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/asdine/storm/v3"
	"github.com/gin-gonic/gin"
	"github.com/shelly-tools/core/common"
	"github.com/shelly-tools/core/config"
	endpoint "github.com/shelly-tools/core/endpoints"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Prepare logging environment

	common.LogInstance = log.New()

	common.LogInstance.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	common.LogInstance.SetLevel(log.ErrorLevel)
}

func init() {
	var err error

	// read config file and create a new Config Struct
	fileData, err := ioutil.ReadFile("core_config.yaml")

	if err != nil {
		common.LogInstance.Fatal(err)
	}

	// Generate new config with all defaults
	common.Config, err = config.New(fileData)

	if err != nil {
		common.LogInstance.Println(err)
	}

	fmt.Println("Config loaded", common.Config)

	// Set correct log Level

	var logLevel log.Level
	switch common.Config.Debugging.Logging.LogLevel {
	case "debug":
		logLevel = log.DebugLevel
	case "info":
		logLevel = log.InfoLevel
	case "error":
		logLevel = log.ErrorLevel
	default:
		logLevel = log.DebugLevel
	}

	common.LogInstance.SetLevel(logLevel)

}

func main() {
	var err error
	//prepare Database
	common.DB, err = storm.Open(common.Config.DatabasePath)

	if err != nil {
		common.LogInstance.Errorf("Failed to open database: %s", err)
	} else {
		defer common.DB.Close()
	}

	// Prepare GinMode
	var ginMode string

	switch common.Config.Debugging.Router.Mode {
	case "PROD":
		common.LogInstance.Debugln("Set router Mode to PROD")
		ginMode = gin.ReleaseMode
	case "DEV":
		common.LogInstance.Debugln("Set router Mode to DEV")
		ginMode = gin.TestMode
	}

	gin.SetMode(ginMode)
	gin.Default().AppEngine = common.Config.Debugging.Router.AppEngine

	router := gin.Default()

	router.Static("/assets", "ui/assets")
	router.Static("/"+common.Config.ImageStorePath, common.Config.ImageStorePath)
	router.Static("/webui", "./webui")
	router.LoadHTMLGlob("ui/templates/*")

	app := router.Group("/app")
	apiV1 := router.Group("/api/v1")
	apiV1.Use(CORS)

	endpoint.RegisterAPPEndpoints(app)
	endpoint.RegisterAPIV1Endpoints(apiV1)

	router.Run(common.Config.UI.ListenAdress + ":" + fmt.Sprint(common.Config.UI.ListenPort))
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
