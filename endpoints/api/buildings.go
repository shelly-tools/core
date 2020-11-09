package api

import (
	"encoding/base64"
	"net/http"
	"os"
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

// GetBuilding returns all buildings found in the database specified in the config
func GetBuilding(c *gin.Context) {
	var building models.Building
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	err := common.DB.One("ID", i, &building)
	if err != nil {
		common.LogInstance.Errorf("Failed to get all buildings from the db:", err)
	}

	c.JSON(http.StatusOK, building)
}

// InsertOneBuilding inserts a building into the database
func InsertOneBuilding(c *gin.Context) {
	var msg struct {
		Status string `json:"status"`
	}

	var building models.Building

	if err := c.ShouldBindBodyWith(&building, binding.JSON); err == nil {
		msg.Status = "ok"
		c.JSON(http.StatusOK, msg)

		decodedString, err := base64.StdEncoding.DecodeString(building.PictureData)
		if err != nil {
			panic(err)
		}

		building.PicturePath = common.Config.ImageStorePath + building.PicturePath

		f, err := os.Create(building.PicturePath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err := f.Write(decodedString); err != nil {
			panic(err)
		}

		if err := f.Sync(); err != nil {
			panic(err)
		}

		building.PictureData = ""
		err = common.DB.Save(&building)
		if err != nil {
			common.LogInstance.Errorf("Failed to store building instance in database: %s", err)
		}
	} else {
		c.String(http.StatusBadRequest, "JSON structure for building is wrong")
	}
}

// DeleteOneBuilding deletes a selected building over the ID
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
