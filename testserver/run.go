package testserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RunGinServer(port int64) {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"port":    port,
		})
	})
	r.Run(fmt.Sprintf(":%d", port))
}
