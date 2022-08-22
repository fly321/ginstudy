package home

import (
	"ginstudy/modules"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
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
		// 获取文件后缀名
		extName := path.Ext(file.Filename)
		allowExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}

		if v, ok := allowExtMap[extName]; !ok || !v {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "不支持的文件类型",
			})
		}

		// 创建图片保存目录
		day := modules.Dates{}.GetDay()
		dir := path.Join("./static/upload/", day)

		err := os.MkdirAll(dir, 0777)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "创建文件夹失败",
			})
		}

		p1 := strconv.FormatInt(modules.Dates{}.GetUnix(), 10) + extName
		filepath := path.Join(dir, p1)
		err1 := ctx.SaveUploadedFile(file, filepath)
		if err1 != nil {
			return
		}

	}
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"file":     file.Filename,
	})
}

func (receiver BaseController) UploadFiles(context *gin.Context) {
	form, _ := context.MultipartForm()
	files := form.File["files"]
	for _, file := range files {
		filepath := path.Join("./static/upload/", file.Filename)
		err := context.SaveUploadedFile(file, filepath)
		if err != nil {
			return
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
