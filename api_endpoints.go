package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/asdine/storm/v3"
	"github.com/gin-gonic/gin"
	"github.com/shelly-tools/core/models"
)

func prepareAPIV1(api *gin.RouterGroup) {

	api.POST("/rooms/create", func(c *gin.Context) {
		db, err := storm.Open(GlobalConfig.DatabasePath)
		if err != nil {
			fmt.Println("Error", err)
		}
		defer db.Close()

		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println("Error", err)
		}
		filePath := GlobalConfig.ImageStorePath + file.Filename
		fmt.Println("Store to", filePath)
		err = c.SaveUploadedFile(file, filePath)

		if err != nil {
			log.Println(err)
		}

		roomName := c.PostForm("roomName")
		roomInstance := models.Room{
			RoomPicture: "/" + filePath,
			RoomName:    roomName,
		}

		err = db.Save(&roomInstance)
		if err != nil {
			fmt.Println("Error", err)
		}
	})

	api.POST("/buildings/create", func(c *gin.Context) {
		db, err := storm.Open(GlobalConfig.DatabasePath)
		if err != nil {
			fmt.Println("Error", err)
		}
		defer db.Close()

		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println("Error", err)
		}
		filePath := GlobalConfig.ImageStorePath + file.Filename
		fmt.Println("Store to", filePath)
		err = c.SaveUploadedFile(file, filePath)

		if err != nil {
			log.Println(err)
		}

		buildingName := c.PostForm("buildingName")
		buildingInstance := models.Building{
			BuildingPicture: "/" + filePath,
			BuildingName:    buildingName,
		}

		err = db.Save(&buildingInstance)
		if err != nil {
			fmt.Println("Error", err)
		}
	})
	api.GET("/buildings", func(c *gin.Context) {
		db, err := storm.Open(GlobalConfig.DatabasePath)
		if err != nil {
			fmt.Println("Error", err)
		}
		defer db.Close()

		var buildings []models.Building

		err = db.All(&buildings)
		if err != nil {
			fmt.Println("Error", err)
		}
		c.JSON(http.StatusOK, buildings)
	})
	api.GET("/buildings/get/:id", func(c *gin.Context) {
		id := c.Param("id")
		i, _ := strconv.Atoi(id)
		db, err := storm.Open(GlobalConfig.DatabasePath)
		if err != nil {
			fmt.Println("Error", err)
		}
		defer db.Close()

		var building models.Building
		err = db.One("IDBuilding", i, &building)
		fmt.Println(building)
		if err != nil {
			fmt.Println("Error", err)
		}
		c.JSON(http.StatusOK, building)
	})
}
