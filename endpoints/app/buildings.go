package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shelly-tools/core/common"
	"github.com/shelly-tools/core/models"
)

// GetAllBuildings shows a template with all buildings listed
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

// AddBuilding shows a template for adding a building
func AddBuilding(c *gin.Context) {
	c.HTML(http.StatusOK, "building_create.html", gin.H{
		"title": "Add Building",
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

// RemoveBuilding shows a template for deleting a building
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
		if err != nil {
			fmt.Println("Error", err)
		}
		err = common.DB.Save(&building)
		if err != nil {
			common.LogInstance.Errorf("Failed to store building instance in database: %s", err)
		}
		c.Redirect(http.StatusFound, "/app/buildings")

	} else {
		c.String(http.StatusBadRequest, "ShouldBind for building failed")
	}
}

// EditBuilding shows a template for editing a building
func EditBuilding(c *gin.Context) {

	var building models.Building
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	err := common.DB.One("ID", i, &building)
	if err != nil {
		common.LogInstance.Errorf("Failed to get all buildings from the db:", err)
	}

	c.HTML(http.StatusOK, "building_edit.html", gin.H{
		"title": "Edit Building",
		"data":  building,
	})
}

// UpdateBuilding stores the modified building to our database
func UpdateBuilding(c *gin.Context) {

	var building models.Building

	if err := c.ShouldBind(&building); err == nil {
		err := common.DB.Update(&building)
		if err != nil {
			common.LogInstance.Errorf("Failed to update building instance in database: %s", err)
		}
		c.Redirect(http.StatusFound, "/app/buildings")

	} else {
		c.String(http.StatusBadRequest, "JSON Structure for building was wrong")
	}
}
