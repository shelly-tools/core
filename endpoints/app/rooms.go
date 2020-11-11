package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelly-tools/core/common"
	"github.com/shelly-tools/core/models"
)

// GetAllRooms returns all rooms, which are found in the database
func GetAllRooms(c *gin.Context) {
	var rooms []models.Room
	err := common.DB.All(&rooms)
	if err != nil {
		common.LogInstance.Errorf("Failed to get all rooms from the db:", err)
	}

	c.HTML(http.StatusOK, "rooms.html", gin.H{
		"title": "Manage Rooms",
		"data":  rooms,
	})
}

func AddRoom(c *gin.Context) {
	c.HTML(http.StatusOK, "building_create.html", gin.H{
		"title": "Add Building",
	})
}
