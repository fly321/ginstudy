package main

import "github.com/gin-gonic/gin"

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

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

	r.GET("/json3", func(context *gin.Context) {
		context.JSON(200, Article{
			Title:   "GO语言",
			Desc:    "GO语言是一种简单的面向对象的编程语言",
			Content: "GO语言是一种简单的面向对象的编程语言",
		})
	})

	r.Run(":8080") // listen and serve on
}
