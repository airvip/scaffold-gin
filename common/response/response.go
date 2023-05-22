package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RECODE_OK        = 1
	RECODE_FAIL      = 0
	RECODE_DBERR     = 4001
	RECODE_NODATA    = 4002
	RECODE_UNKNOWERR = 9999
)

var recodeText = map[int]string{
	RECODE_OK:        "成功",
	RECODE_DBERR:     "数据库错误",
	RECODE_NODATA:    "无数据",
	RECODE_UNKNOWERR: "未知错误",
}

func RecodeText(code int) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Success(c *gin.Context, data gin.H) {
	Response(c, http.StatusOK, RECODE_OK, data, RecodeText(RECODE_OK))
}

func Fail(c *gin.Context, msg string) {
	Response(c, http.StatusOK, RECODE_FAIL, nil, msg)
}

func SuccessStruct(c *gin.Context, data interface{}){
	c.JSON(http.StatusOK, gin.H{
		"code": RECODE_OK,
		"msg":  RecodeText(RECODE_OK),
		"data": data,
	})
}
