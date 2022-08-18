package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func UnixToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	// Creates a gin router with default middleware:
	r := gin.Default()
	// 自定义模板函数 要写在加载模板之前
	r.SetFuncMap(template.FuncMap{
		"UnixToDate": UnixToDate,
	})

	// 配置模板目录 ** 代表目录
	r.LoadHTMLGlob("templates/**/*")

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
		context.JSON(http.StatusOK, Article{
			Title:   "GO语言",
			Desc:    "GO语言是一种简单的面向对象的编程语言",
			Content: "GO语言是一种简单的面向对象的编程语言",
		})
	})

	r.GET("/html1", func(context *gin.Context) {
		context.HTML(http.StatusOK, "home.1", gin.H{
			"name": "Main website",
		})
	})

	r.GET("/html2", func(context *gin.Context) {
		context.HTML(http.StatusOK, "home.index", gin.H{
			"name": "Main website",
		})
	})

	r.GET("/html3", func(context *gin.Context) {
		context.HTML(http.StatusOK, "home.index3", gin.H{
			"title":    "Main website",
			"msg":      "hello world",
			"score":    100,
			"bool":     true,
			"hobby":    []string{"篮球", "足球", "游泳"},
			"unixtime": 1660837281,
			"newsList": []interface{}{
				&Article{
					"Go1",
					"GO语言是一种简单的面向对象的编程语言",
					"GO语言是一种简单的面向对象的编程语言",
				},
				&Article{
					"Go2",
					"GO语言是一种简单的面向对象的编程语言",
					"GO语言是一种简单的面向对象的编程语言",
				},
			},
		})
	})

	r.Run(":8080") // listen and serve on
}
