package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DemoRouter(r *gin.Engine) {
	g1 := r.Group("/group1/")
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
