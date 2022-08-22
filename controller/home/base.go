package home

import (
	gin "github.com/gin-gonic/gin"
	"net/http"
	"path"
)

type BaseController struct {
}

func (receiver BaseController) Success(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    []string{},
	})
}

func (con BaseController) UploadFile(ctx *gin.Context) {
	username := ctx.PostForm("username")
	file, err := ctx.FormFile("file")
	if err == nil {
		filepath := path.Join("./static/upload/", file.Filename)
		ctx.SaveUploadedFile(file, filepath)

	}
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"file":     file.Filename,
	})
}
