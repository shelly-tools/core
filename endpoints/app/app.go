package app

import (
	"fmt"
	"net/http"

	"github.com/asdine/storm/v3"
	"github.com/gin-gonic/gin"
	"github.com/shelly-tools/core/common"
	"github.com/shelly-tools/core/models"
)

func Root(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

func Rooms(c *gin.Context) {
	db, err := storm.Open(common.Config.DatabasePath)
	if err != nil {
		fmt.Println("Error", err)
	}
	defer db.Close()

	var rooms []models.Room

	err = db.All(&rooms)
	if err != nil {
		fmt.Println("Error", err)
	}
	c.HTML(http.StatusOK, "rooms.html", gin.H{
		"title": "Main website",
		"data":  rooms,
	})
}
