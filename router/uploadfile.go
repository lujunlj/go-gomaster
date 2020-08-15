package router

import (
	v1 "gomaster/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUploadFileRouter(Router *gin.RouterGroup) {
	router := Router.Group("/file")
	{
		//单一文件 上传
		router.POST("/upload", v1.UploadFile)
		//多个文件 批量上传
		router.POST("/uploads", v1.UploadFileS)
		//下载文件
		router.GET("/:filename", v1.DownloadFile)
	}
}
