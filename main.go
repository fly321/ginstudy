package main

import (
	"encoding/xml"
	"fmt"
	"ginstudy/controller/home"
	"ginstudy/routers"
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

	// 设置静态文件服务
	r.Static("/static", "./static")

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

	r.GET("/html2", home.UserController{}.IndexView)
	r.GET("/baseSuccess", home.BaseController{}.Success)

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

	r.GET("/get1", func(context *gin.Context) {
		name := context.Query("name")
		age := context.DefaultQuery("age", "20")
		context.String(http.StatusOK, "name: %s, age: %s", name, age)
	})

	r.POST("/post1", func(context *gin.Context) {
		name := context.PostForm("name")
		age := context.DefaultPostForm("age", "20")
		context.String(http.StatusOK, "name: %s, age: %s", name, age)
	})

	r.GET("/get2", func(context *gin.Context) {
		student := &Student{}
		if err := context.ShouldBindQuery(student); err != nil {
			//context.String(http.StatusOK, "err: %s", err.Error())
			context.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			//context.String(http.StatusOK, "name: %s, age: %d", student.Name, student.Age)
			context.JSON(http.StatusOK, student)
		}
	})

	r.POST("/xml1", func(context *gin.Context) {
		article := &Article1{}
		b, _ := context.GetRawData()
		fmt.Println(string(b))
		if err := xml.Unmarshal(b, &article); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, article)
		}
	})

	r.POST("/xml2", func(context *gin.Context) {
		article := &Article1{}
		if err := context.ShouldBindXML(article); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, article)
		}
	})

	r.GET("/get3/:uid", func(context *gin.Context) {
		//va := context.Params.ByName("uid")
		va := context.Param("uid")
		context.JSON(http.StatusOK, gin.H{
			"uid": va,
		})
	})

	routers.DemoRouter(r)

	err := r.Run(":8080")
	if err != nil {
		return
	} // listen and serve on
}

type Student struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}

type Article1 struct {
	Title string `json:"title" xml:"title"`
	Desc  string `json:"desc" xml:"desc"`
}
