package home

import (
	"github.com/gin-gonic/gin"
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

func (receiver BaseController) UploadFile(ctx *gin.Context) {
	username := ctx.PostForm("username")
	file, err := ctx.FormFile("file")
	if err == nil {
		filepath := path.Join("./static/upload/", file.Filename)
		err := ctx.SaveUploadedFile(file, filepath)
		if err != nil {
			return
		}

	}
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"file":     file.Filename,
	})
}
