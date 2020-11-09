package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelly-tools/core/common"
	"github.com/shelly-tools/core/models"
)

func Root(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

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
