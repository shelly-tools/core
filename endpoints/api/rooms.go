package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shelly-tools/core/common"
	"github.com/shelly-tools/core/models"
)

// GetAllRooms returns all rooms found in the database specified in the config
func GetAllRooms(c *gin.Context) {
	var rooms []models.Room

	err := common.DB.All(&rooms)

	if err != nil {
		common.LogInstance.Errorf("Failed to get all rooms from the db:", err)
	}

	c.JSON(http.StatusOK, rooms)
}

// InsertOneRoom inserts a room into the database
func InsertOneRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindBodyWith(&room, binding.JSON); err == nil {
		c.String(http.StatusOK, "Body loaded", room.Name)

		err := common.DB.Save(&room)
		if err != nil {
			common.LogInstance.Errorf("Failed to store room instance in database: %s", err)
		}

	} else {
		c.String(http.StatusBadRequest, "JSON Structur for a room is wrong")
	}
}
