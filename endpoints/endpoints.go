package endpoints

import (
	"github.com/gin-gonic/gin"
	API "github.com/shelly-tools/core/endpoints/api"
	APP "github.com/shelly-tools/core/endpoints/app"
)

// RegisterAPIV1Endpoints registers all api endpoints to the api router group
func RegisterAPIV1Endpoints(api *gin.RouterGroup) {
	// endpoints for rooms
	api.POST("/rooms/create", API.InsertOneRoom)
	api.GET("/rooms/get/all", API.GetAllRooms)

	// endpoints for buildings
	api.POST("/buildings/create", API.InsertOneBuilding)
	api.POST("/buildings/delete", API.DeleteOneBuilding)
	api.GET("/buildings/get/all", API.GetAllBuildings)
	api.GET("/buildings/get/id/:id", API.GetBuilding)

	// endpoints for devices
	api.GET("/devices/discover", API.DiscoverShellys)
}

// RegisterAPPEndpoints registers all app endpoints to the app router group
func RegisterAPPEndpoints(app *gin.RouterGroup) {
	app.GET("/", APP.Root)

	app.GET("/buildings", APP.GetAllBuildings)
	app.GET("/buildings/create", APP.AddBuilding)
	app.POST("/buildings/insert", APP.InsertBuilding)
	app.GET("/buildings/delete/:id", APP.DeleteBuilding)
	app.POST("/buildings/remove/", APP.RemoveBuilding)

	app.GET("/rooms", APP.GetAllRooms)
}
