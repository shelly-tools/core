package api

import (
	"fmt"
	"log"

	"github.com/asdine/storm/v3"
	"github.com/gin-gonic/gin"
	"github.com/shelly-tools/core/common"
	"github.com/shelly-tools/core/models"
)

// GetAllRooms returns all rooms found in the database specified in the config
func GetAllRooms(c *gin.Context) {
	db, err := storm.Open(common.Config.DatabasePath)
	if err != nil {
		fmt.Println("Error", err)
	}
	defer db.Close()

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("Error", err)
	}
	filePath := common.Config.ImageStorePath + file.Filename
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
}
