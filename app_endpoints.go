package main

import (
	"fmt"
	"net/http"

	"github.com/asdine/storm/v3"
	"github.com/gin-gonic/gin"
	"github.com/shelly-tools/core/models"
)

func prepareApp(app *gin.RouterGroup) {
	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	app.GET("/rooms", func(c *gin.Context) {
		db, err := storm.Open(GlobalConfig.DatabasePath)
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
	})

	app.GET("/rooms/manage/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "rooms_create.html", gin.H{
			"title": "Main website",
		})
	})
}
