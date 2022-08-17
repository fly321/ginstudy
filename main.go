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
			"message": "pong1",
		})
	})

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1",
			"success": true,
		})
	})

	r.GET("/json2", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"message": "pong2",
			"success": true,
		})
	})

	r.Run(":8080") // listen and serve on
}
