package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grandcat/zeroconf"
	"github.com/shelly-tools/core/models"
)

// DiscoverShellys returns all buildings found in the database specified in the config
func DiscoverShellys(c *gin.Context) {
	log.Println("Shelly mDNS Discovery in progress..")
	var devices []models.Device
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)

	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			fmt.Println("New Result")
			ipstring := ""
			if len(entry.AddrIPv4) > 0 {
				ipstring = entry.AddrIPv4[0].String()
			}
			hostname := strings.Split(entry.HostName, ".")

			device := models.Device{
				IDDevice:   1,
				DeviceName: hostname[0],
				DeviceIP:   ipstring,
			}

			devices = append(devices, device)

			fmt.Println(hostname, ipstring)
		}
		fmt.Println("")
		log.Println("mDNS Discovery finished.")
		return
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(20))
	defer cancel()

	err = resolver.Browse(ctx, "_http._tcp", "local", entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()
	// Wait some additional time to see debug messages on go routine shutdown.
	time.Sleep(1 * time.Second)

	c.JSON(http.StatusOK, devices)
}
