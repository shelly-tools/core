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

<<<<<<< HEAD
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
=======
	r := mux.NewRouter()
	staticDir := "/ui/assets/"
	// Create the route
	r.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/v1/devices", DeviceIndex)
	r.HandleFunc("/", DashboardIndex)

	// Rooms
	r.HandleFunc("/ui/manage/rooms", RoomList)
	r.HandleFunc("/ui/manage/rooms/add", RoomAdd)
	r.HandleFunc("/ui/manage/rooms/edit/{id}", RoomEdit)
	r.HandleFunc("/ui/manage/rooms/delete/{id}", RoomDelete)
	// Buildings
	r.HandleFunc("/ui/manage/buildings", BuildingList)
	r.HandleFunc("/ui/manage/buildings/add", BuildingAdd)
	r.HandleFunc("/ui/manage/buildings/insert", BuildingInsert)
	r.HandleFunc("/ui/manage/buildings/edit/{id}", BuildingEdit)
	r.HandleFunc("/ui/manage/buildings/update", BuildingUpdate)
	r.HandleFunc("/ui/manage/buildings/delete/{id}", BuildingDelete)
	r.HandleFunc("/ui/manage/buildings/remove/{id}", BuildingRemove)

	// Devices
	r.HandleFunc("/ui/manage/devices/discovered", DiscoveredDevices)

	//r.HandleFunc("/api/v1/device/{id}", ApiDeviceHandler)

	go func() {
		fmt.Println("CoIoT Listener started")
		log.Fatal(http.ListenAndServe(":8000", r))
	}()

	mux := coiot.NewServeMux()
	mux.Handle("/cit/s", coiot.FuncHandler(CoIoTHandler))
	log.Fatal(coiot.ListenAndServe("udp", "224.0.1.187:5683", mux))
}
func DeviceIndex(w http.ResponseWriter, r *http.Request) {
>>>>>>> 39b25c89a9324b9298b1e2e7d1a61467a3dfbaaa

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
