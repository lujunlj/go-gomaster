package v1

import (
	"gomaster/common"
	"gomaster/config"
	"gomaster/config/result"
	"gomaster/service"

	"github.com/gin-gonic/gin"
)

//上传单一文件
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		common.Log().Error(config.FILE_UPLOAD_FAIL)
		return
	}
	//上传
	err = service.Upload(c, file)
	if err != nil {
		result.FailWithMessage(config.FILE_UPLOAD_FAIL, c)
		return
	}
	//入库
	result.OkWithMessage(config.FILE_UPLOAD_SUCCESS, c)
}

//上传多个文件
func UploadFileS(c *gin.Context) {
	form, _ := c.MultipartForm()
	//file, err := c.FormFile("file")
	//if err != nil {
	//	common.Log().Error(config.FILE_UPLOAD_FAIL)
	//	return
	//}
	files := form.File["file"]
	for _, file := range files {
		common.Log().Info(file.Filename)
		//上传
		err := service.Upload(c, file)
		if err != nil {
			result.FailWithMessage(config.FILE_UPLOAD_FAIL, c)
			return
		}
	}
	result.OkWithMessage(config.FILE_UPLOAD_SUCCESS, c)
}

//下载文件
func DownloadFile(c *gin.Context) {
	//id := c.GetString("id")
	filename := c.Param("filename")
	service.Download(filename, c)
}
