package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shelly-tools/core/common"
	"github.com/shelly-tools/core/models"
)

// GetAllBuildings returns all buildings found in the database specified in the config
func GetAllBuildings(c *gin.Context) {
	var buildings []models.Building

	err := common.DB.All(&buildings)

	if err != nil {
		common.LogInstance.Errorf("Failed to get all buildings from the db:", err)
	}

	c.JSON(http.StatusOK, buildings)
}

// InsertOneBuidling inserts a building into the database
func InsertOneBuilding(c *gin.Context) {
	fmt.Println("hier")
	var building models.Building
	if err := c.ShouldBindBodyWith(&building, binding.JSON); err == nil {
		c.String(http.StatusOK, "Body loaded", building.BuildingName)

		err := common.DB.Save(&building)
		if err != nil {
			common.LogInstance.Errorf("Failed to store building instance in database: %s", err)
		}

	} else {
		c.String(http.StatusBadRequest, "JSON structire for building is wrong")
	}
}
