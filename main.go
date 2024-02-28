package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	prometheus := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)
	router.Use(prometheus.Instrument())
	router.GET("/foo", fooGet)
	router.Run()
}

func fooGet(c *gin.Context) {

	delay, _ := c.GetQuery("delay")

	if delay != "" {
		duration, _ := time.ParseDuration(delay)
		fmt.Print("Sleeping for ", duration, " seconds")
		time.Sleep(duration)
	} else {
		delay = "0"
	}
	c.JSON(http.StatusOK, gin.H{"message": "foo", "delay": delay})
}
