package main

import (
	"fmt"
	"log"

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
}
