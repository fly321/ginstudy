package home

import (
	"ginstudy/modules"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

func (con UserController) IndexSelectAgeGt20(context *gin.Context) {
	userList := []modules.User{}
	modules.DB.Where("age>20").Find(&userList)
	context.JSON(http.StatusOK, userList)
}

func (con UserController) CreateData(context *gin.Context) {
	user := modules.User{
		Age:      33,
		Username: "老吴",
		Email:    "9442@qq.com",
		AddTime:  int(time.Now().Unix()),
	}
	modules.DB.Create(&user)
}

// 修改数据
func (con UserController) UpdateData(context *gin.Context) {
	/*user := modules.User{Id: 3}
	modules.DB.Find(&user)
	user.Username = "吴老菊"
	modules.DB.Save(&user)*/

	user := modules.User{}
	modules.DB.Model(&user).Where("id = ?", 3).Update("username", "aimer")

}
