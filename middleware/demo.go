package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Middlewares struct {
}

// 中间件
func (con Middlewares) InitMd(content *gin.Context) {
	fmt.Println("middlewareGet")
	content.Set("name", "吴亦凡")
	content.Next()
}
