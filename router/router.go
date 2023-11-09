package router

import "github.com/gin-gonic/gin"

func Router(c *gin.Context) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    200,
		})
	})
	return r
}
