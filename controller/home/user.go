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
	var userList []modules.User
	modules.DB.Find(&userList)
	context.JSON(http.StatusOK, userList)
}

func (con UserController) IndexSelectAgeGt20(context *gin.Context) {
	var userList []modules.User
	modules.DB.Where("age>20").Find(&userList)
	context.JSON(http.StatusOK, userList)
}

func (con UserController) CreateData(*gin.Context) {
	user := modules.User{
		Age:      33,
		Username: "老吴",
		Email:    "9442@qq.com",
		AddTime:  int(time.Now().Unix()),
	}
	modules.DB.Create(&user)
}

// UpdateData 修改数据
//
//goland:noinspection ALL
func (con UserController) UpdateData(*gin.Context) {
	/*user := modules.User{Id: 3}
	modules.DB.Find(&user)
	user.Username = "吴老菊"
	modules.DB.Save(&user)*/

	user := modules.User{}
	modules.DB.Model(&user).Where("id = ?", 3).Update("username", "aimer")
}

// DeleteData 删除
func (con UserController) DeleteData(context *gin.Context) {
	/*user := modules.User{Id: 3}
	modules.DB.Delete(&user)*/

	user := modules.User{}
	modules.DB.Where("id = ?", 5).Delete(&user)

	context.String(200, "删除用户")
}
