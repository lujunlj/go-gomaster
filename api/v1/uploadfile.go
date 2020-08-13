package v1

import (
	"fmt"
	"gomaster/common"
	"gomaster/config"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	uploadpath = common.Conf().Upload.Path
)

func upload(c *gin.Context, file *multipart.FileHeader) error {
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
	err := c.SaveUploadedFile(file, uploadpath+file.Filename)
	if err != nil {
		common.Log().Error(config.FILE_UPLOAD_FAIL)
		return err
	}
	return nil
}

//上传单一文件
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		common.Log().Error(config.FILE_UPLOAD_FAIL)
		return
	}
	//上传
	err = upload(c, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": config.FILE_UPLOAD_FAIL,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": config.FILE_UPLOAD_SUCCESS,
	})
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
		err := upload(c, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": config.FILE_UPLOAD_FAIL,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": config.FILE_UPLOAD_SUCCESS,
	})
}

//查看文件
func LookFile(c *gin.Context) {
	//id := c.GetString("id")
	filename := c.GetString("filename")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", filename))
	c.File(uploadpath + filename)
}
