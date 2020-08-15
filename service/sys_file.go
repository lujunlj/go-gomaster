package service

import (
	"fmt"
	"gomaster/common"
	"gomaster/config"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

var uploadpath string

func Get_uploadpath() string {
	if uploadpath == "" {
		uploadpath = common.Conf().Upload.Path
	}
	return uploadpath
}

func Upload(c *gin.Context, file *multipart.FileHeader) error {
	//open, err := file.Open()
	//if err != nil {
	//	common.Log().Error("打开上传文件出错")
	//	return
	//}
	//defer open.Close()
	//out, err := os.Create("./upload/" + file.Filename)
	//if err != nil {
	//	common.Log().Error("文件创建失败!")
	//	return
	//}
	//defer out.Close()
	//io.Copy(out, open)
	err := c.SaveUploadedFile(file, Get_uploadpath()+file.Filename)
	if err != nil {
		common.Log().Error(config.FILE_UPLOAD_FAIL)
		return err
	}
	return nil
}

func Download(filename string, c *gin.Context) {
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", filename))
	c.File(Get_uploadpath() + filename)
}
