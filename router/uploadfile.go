package router

import (
	v1 "gomaster/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUploadFileRouter(Router *gin.RouterGroup) {
	router := Router.Group("/upload")
	{
		//单一文件 上传
		router.POST("", v1.UploadFile)
		//多个文件 批量上传
		router.POST("many", v1.UploadFileS)
		//查看文件
		router.GET(":filename", v1.LookFile)
	}
}
