package routers

import (
	"fmt"
	"ginstudy/controller/home"
	"ginstudy/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DemoRouter(r *gin.Engine) {
	g1 := r.Group("/group1/", middleware.Middlewares{}.InitMd)
	{
		g1.GET("qq/:uid", func(context *gin.Context) {
			//va := context.Params.ByName("uid")
			va := context.Param("uid")
			fmt.Println(context.Get("name"))
			context.JSON(http.StatusOK, gin.H{
				"uid": va,
			})
		})

		g1.GET("cookie/set", func(context *gin.Context) {
			context.SetCookie("name", "value", 3600, "/", "localhost", false, true)
		})

		g1.GET("cookie/get", func(context *gin.Context) {
			name, err := context.Cookie("name")
			if err != nil {
				context.String(http.StatusOK, "err: %s", err.Error())
			} else {
				context.String(http.StatusOK, "name: %s", name)
			}
		})

		g1.POST("doUpload", home.BaseController{}.UploadFile)
		g1.POST("doUploads", home.BaseController{}.UploadFiles)

	}
}
