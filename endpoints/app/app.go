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

func GetAllBuildings(c *gin.Context) {

	var buildings []models.Building
	err := common.DB.All(&buildings)
	if err != nil {
		common.LogInstance.Errorf("Failed to get all buildings from the db:", err)
	}

	c.HTML(http.StatusOK, "buildings.html", gin.H{
		"title": "Manage Buildings",
		"data":  buildings,
	})
}

func AddBuilding(c *gin.Context) {

	c.HTML(http.StatusOK, "building_create.html", gin.H{
		"title": "Add Building",
	})
}

// InsertBuilding inserts a building into the database
func InsertBuilding(c *gin.Context) {

	var building models.Building

	if err := c.ShouldBind(&building); err == nil {
		err := common.DB.Save(&building)
		if err != nil {
			common.LogInstance.Errorf("Failed to store room instance in database: %s", err)
		}
		c.Redirect(http.StatusFound, "/app/buildings")

	} else {
		c.String(http.StatusBadRequest, "JSON Structur for a room is wrong")
	}
}
