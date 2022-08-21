package routers

import (
	"fmt"
	"ginstudy/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DemoRouter(r *gin.Engine) {
	g1 := r.Group("/group1/", middleware.Middlewares{}.InitMd)
	{
		g1.GET(":uid", func(context *gin.Context) {
			//va := context.Params.ByName("uid")
			va := context.Param("uid")
			fmt.Println(context.Get("name"))
			context.JSON(http.StatusOK, gin.H{
				"uid": va,
			})
		})
	}
}
