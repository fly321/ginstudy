package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Middlewares struct {
}

// 中间件
func (con Middlewares) InitMd(content *gin.Context) {
	fmt.Println("middlewareGet")
	content.Set("name", "fly321")

	// 必须使用copy对象 协程
	cpc := content.Copy()

	// 记录日志
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Deno ! in middleware" + cpc.Request.URL.Path)
	}()
	content.Next()
}
