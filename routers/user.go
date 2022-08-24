package routers

import (
	"ginstudy/controller/home"
	"ginstudy/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	r1 := r.Group("/user/", middleware.Middlewares{}.InitMd)
	{
		r1.GET("findAll", home.UserController{}.IndexSelect)
		r1.GET("findAllGtAge20", home.UserController{}.IndexSelectAgeGt20)
		r1.GET("createData", home.UserController{}.CreateData)
		r1.GET("updateData", home.UserController{}.UpdateData)
		r1.GET("deleteData", home.UserController{}.DeleteData)
	}
}
