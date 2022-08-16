package main

import "github.com/gin-gonic/gin"

func main() {
	// Creates a gin router with default middleware:
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		/*c.JSON(200, gin.H{
			"message": "pong",
		})*/

		c.String(200, "值：%v", "你好golang")
	})

	r.GET("/jsonDemo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // listen and serve on
}
