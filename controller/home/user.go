package home

import (
	"ginstudy/modules"
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

func (con UserController) IndexSelect(context *gin.Context) {
	userList := []modules.User{}
	modules.DB.Find(&userList)
	context.JSON(http.StatusOK, userList)

}
