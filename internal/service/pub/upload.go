package pub

import (
	"fmt"
	"path"
	"scaffold-gin/common/config"
	"scaffold-gin/common/global"
	"scaffold-gin/common/response"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// UploadObj
// @Summary 上传文件
// @Schemes
// @Description 上传文件
// @Tags 文件上传
// @Accept json
// @Produce json
// @Param fileObj formData file true "文件"
// @Success 200 {string} json "{"code":200,"msg":"","data":""}"
// @Router /v1/upload [post]
// @Security ApiKeyAuth
func UploadObj(c *gin.Context) {
	file, _ := c.FormFile("fileObj")
	fileHandle, err := file.Open()
    if err != nil {
        response.Fail(c, "can't open fialeHandle, error:" + err.Error())
		return
    }
    defer fileHandle.Close()

    fileExt := path.Ext(file.Filename)  // 获取文件后缀名
	uuidv4 := uuid.NewV4()
	if(err != nil){
		response.Fail(c, "uuid generate failed, error:"+err.Error() )
		return
	}
    flieName := uuidv4.String() // 使用uuid作为文件名
    tTime := time.Now().Format("200601") // 以年月为文件目录进行分类
	ossFilePath := fmt.Sprintf("%s/%s%s", tTime, flieName, fileExt) // 年月/文件名.扩展名（注意不要再定义的目录前面加/）
	err = global.OSS.PutObject(ossFilePath, fileHandle)
	if(err != nil){
		response.Fail(c, "up file to oss failed, error:"+err.Error() )
		return
	}
	allPath := fmt.Sprintf("%s/%s",config.Conf.OSS.BasePath,ossFilePath)
	response.SuccessStruct(c,gin.H{"url":allPath})
}