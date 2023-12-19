package main

import (
	"GoPlayground/ginproto"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		var req ginproto.Request
		if err := c.Bind(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": req.GetName()})
	})

	r.Run()
}
