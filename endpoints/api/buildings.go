package api

import (
	"net/http"
	"strconv"

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

// GetAllBuildings returns all buildings found in the database specified in the config
func GetBuilding(c *gin.Context) {
	var building models.Building
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	err := common.DB.One("IDBuilding", i, &building)
	if err != nil {
		common.LogInstance.Errorf("Failed to get all buildings from the db:", err)
	}

	c.JSON(http.StatusOK, building)
}

// InsertOneBuidling inserts a building into the database
func InsertOneBuilding(c *gin.Context) {
	var msg struct {
		Status string `json:"status"`
	}
	var building models.Building
	if err := c.ShouldBindBodyWith(&building, binding.JSON); err == nil {
		msg.Status = "ok"
		c.JSON(http.StatusOK, msg)

		err := common.DB.Save(&building)
		if err != nil {
			common.LogInstance.Errorf("Failed to store building instance in database: %s", err)
		}
	} else {
		c.String(http.StatusBadRequest, "JSON structure for building is wrong")
	}
}

func DeleteOneBuilding(c *gin.Context) {
	var msg struct {
		Status string `json:"status"`
	}
	var building models.Building
	if err := c.ShouldBindBodyWith(&building, binding.JSON); err == nil {
		msg.Status = "ok"
		c.JSON(http.StatusOK, msg)

		err := common.DB.DeleteStruct(&building)
		if err != nil {
			common.LogInstance.Errorf("Failed to store building instance in database: %s", err)
		}
	} else {
		c.String(http.StatusBadRequest, "JSON structure for building is wrong")
	}
}
