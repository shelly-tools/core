package endpoints

import (
	"github.com/gin-gonic/gin"
	API "github.com/shelly-tools/core/endpoints/api"
	APP "github.com/shelly-tools/core/endpoints/app"
)

// RegisterAPIV1Endpoints registers all api endpoints to the api router group
func RegisterAPIV1Endpoints(api *gin.RouterGroup) {
	api.POST("/rooms/create")

	api.GET("/rooms/get/all", API.GetAllRooms)
}

// RegisterAPPEndpoints registers all app endpoints to the app router group
func RegisterAPPEndpoints(app *gin.RouterGroup) {
	app.GET("/", APP.Root)

	app.GET("/rooms", APP.Rooms)
}
