package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

func (con UserController) IndexView(context *gin.Context) {
	context.HTML(http.StatusOK, "home.index", gin.H{
		"name": "Main website",
	})
}
