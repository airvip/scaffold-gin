package pub

import (
	"encoding/json"
	"fmt"
	"regexp"
	"scaffold-gin/common/global"
	"scaffold-gin/common/response"
	"scaffold-gin/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SendAliCode struct {
	Code int `json:"code,string"`
}

// SmsAliCode
// @Summary 获取验证码
// @Schemes
// @Description 获取验证码
// @Tags 公用方法
// @Accept json
// @Produce json
// @Param mobile query string true "手机号"
// @Success 200 {string} json "{"code":200,"msg":"","data":""}"
// @Router /sms-code-ali [get]
func SmsAliCode(c *gin.Context) {
	mobile := c.Query("mobile")
	if mobile == "" {
		response.Fail(c, "mobile is required")
		return
	}

	matched, err := regexp.MatchString("^1(3|4|5|6|8|9)[0-9]{9}$", mobile)
	if err != nil {
		response.Fail(c, "mobile regexp is failed,error:"+err.Error())
		return
	}
	if !matched {
		response.Fail(c, "Incorrect mobile phone format ")
		return
	}

	alcode := &SendAliCode{Code: util.RandNum()}
	b, err := json.Marshal(&alcode)
	if err != nil {
		response.Fail(c, "randnum marshal failed,error:"+err.Error())
		return
	}
	params := string(b)
	ssr, err := global.SendAliSms(mobile, "滇医通", "SMS_202568112", params)
	if err != nil {
		response.Fail(c, "send sms failed,error:"+err.Error())
		return
	}
	if *ssr.Body.Code == "OK" {
		response.Success(c, gin.H{})
		return
	}
	response.Fail(c, *ssr.Body.Message)
}

// SmsTxCode
// @Summary 获取验证码
// @Schemes
// @Description 获取验证码
// @Tags 公用方法
// @Accept json
// @Produce json
// @Param mobile query string true "手机号"
// @Success 200 {string} json "{"code":200,"msg":"","data":""}"
// @Router /sms-code-tx [get]
func SmsTxCode(c *gin.Context) {
	mobile := c.Query("mobile")
	if mobile == "" {
		response.Fail(c, "mobile is required")
		return
	}

	matched, err := regexp.MatchString("^1(3|4|5|6|8|9)[0-9]{9}$", mobile)
	if err != nil {
		response.Fail(c, "mobile regexp is failed,error:"+err.Error())
		return
	}
	if !matched {
		response.Fail(c, "Incorrect mobile phone format ")
		return
	}

	code := strconv.Itoa(util.RandNum())
	ssr, err := global.SendTxSms("滇医通", "618488", []string{code}, []string{mobile})
	if err != nil {
		response.Fail(c, "send sms failed,error:"+err.Error())
		return
	}

	flag := false
	for _, v := range ssr.Response.SendStatusSet {
		global.ZAPLOGGER.Info(fmt.Sprintf("[%s] %s |#| %s", *v.PhoneNumber, *v.Message, *v.Code))
		if "+86" + mobile == *v.PhoneNumber && *v.Code == "Ok" {
			flag = true
			goto Loop
		}
	}

Loop:
	if flag {
		response.Success(c, gin.H{})
		return
	}
	response.Fail(c, "发送失败")

}