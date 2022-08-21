package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 中间件
func demoMiddlewareGet(content *gin.Context) {
	fmt.Println("middlewareGet")
	content.Next()
}

func DemoRouter(r *gin.Engine) {
	g1 := r.Group("/group1/", demoMiddlewareGet)
	{
		g1.GET(":uid", func(context *gin.Context) {
			//va := context.Params.ByName("uid")
			va := context.Param("uid")
			context.JSON(http.StatusOK, gin.H{
				"uid": va,
			})
		})
	}
}
