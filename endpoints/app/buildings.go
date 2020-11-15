package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shelly-tools/core/common"
	"github.com/shelly-tools/core/models"
)

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
		"type":  "create",
	})
}

// DeleteBuilding shows a template for deleting a building
func DeleteBuilding(c *gin.Context) {

	var building models.Building
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	err := common.DB.One("ID", i, &building)
	if err != nil {
		common.LogInstance.Errorf("Failed to get all buildings from the db:", err)
	}

	c.HTML(http.StatusOK, "building_delete.html", gin.H{
		"title": "Delete Building",
		"data":  building,
	})
}

// DeleteBuilding shows a template for deleting a building
func RemoveBuilding(c *gin.Context) {

	var building models.Building

	if err := c.ShouldBind(&building); err == nil {
		err := common.DB.DeleteStruct(&building)
		if err != nil {
			common.LogInstance.Errorf("Failed to store room instance in database: %s", err)
		}
		c.Redirect(http.StatusFound, "/app/buildings")

	} else {
		c.String(http.StatusBadRequest, "JSON Structur for a room is wrong")
	}
}

// InsertBuilding inserts a building into the database
func InsertBuilding(c *gin.Context) {

	var building models.Building

	if err := c.ShouldBind(&building); err == nil {

		file, _ := c.FormFile("picture")
		if file != nil {
			// image is available. Save it
			err = c.SaveUploadedFile(file, common.Config.ImageStorePath)

			if err != nil {
				common.LogInstance.Errorf("Failed to store image. %v", err)
			}

			building.PicturePath = file.Filename
		}
		// save data to the database
		err := common.DB.Save(&building)
		if err != nil {
			common.LogInstance.Errorf("Failed to store room instance in database: %s", err)
		}
		c.HTML(http.StatusOK, "building_create.html", gin.H{
			"title": "Delete Building",
			"type":  "created",
			"data":  building,
		})

	} else {
		c.String(http.StatusBadRequest, "JSON Structure for a Building is wrong")
	}
}
