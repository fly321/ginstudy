package routers

import (
	"ginstudy/controller/home"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	r.GET("/user/findAll", home.UserController{}.IndexSelect)
}
