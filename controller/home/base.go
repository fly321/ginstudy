package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (receiver BaseController) Success(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    []string{},
	})
}
