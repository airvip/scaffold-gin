package service

import (
	"fmt"
	"log"
	"scaffold-gin/common/def"
	"scaffold-gin/common/global"
	"scaffold-gin/common/response"
	"scaffold-gin/internal/dto/req"
	"scaffold-gin/internal/model"
	"scaffold-gin/internal/validate"
	"scaffold-gin/util"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// AddUser
// @Summary 用户注册
// @Schemes
// @Description 用户注册
// @Tags 用户相关
// @Accept json
// @Produce json
// @Param userinfo body req.AddUserDto true "用户信息"
// @Success 200 {string} json "{"code":200,"msg":"","data":""}"
// @Router /user-register [post]
func AddUser(c *gin.Context) {

	var json req.AddUserDto
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Fail(c, validate.Translate(err))
		return
	}

	d := model.GetUserModel()
	uuidv4 := uuid.NewV4()
	userBase := model.UserBasic{
		Identity: uuidv4.String(),
		Password: util.Md5String(json.Password),
		Nickname: json.Nickname,
		Email:    json.Email,
		Mobile:   json.Mobile,
	}
	if err := d.Create(&userBase).Error; err != nil {
		msg := fmt.Sprintf("create user faield error:%s", err.Error())
		response.Fail(c, msg)
		return
	}
	response.Success(c, gin.H{})
}

// LoginUser
// @Summary 用户登录
// @Schemes
// @Description 获取用户token
// @Tags 用户相关
// @Accept json
// @Produce json
// @Param userinfo body req.LoginUserDto true "用户信息"
// @Success 200 {string} json "{"code":200,"msg":"","data":""}"
// @Router /user-login [post]
func LoginUser(c *gin.Context) {

	var json req.LoginUserDto
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Fail(c, validate.Translate(err))
		return
	}

	user := new(model.UserBasic)
	// var user models.UserBasic
	err :=  model.GetUserModel().Where("mobile=? AND password=?", json.Mobile, util.Md5String(json.Password)).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Fail(c, "user not exist")
			return
		}
		response.Fail(c, "query failed, err:"+err.Error())
		return
	}
	token, err := util.ReleaseToken(user.Identity, user.Nickname)
	if err != nil {
		response.Fail(c, "release token failed, err:"+err.Error())
		return
	}
	response.Success(c, gin.H{"token": token})
}

// GetUserDetail
// @Summary 获取用户详情
// @Schemes
// @Description 获取用户详情
// @Tags 用户相关
// @Accept json
// @Produce json
// @Param nickname query string false "昵称"
// @Success 200 {object}  req.UserBase
// @Router /v1/user-detail [get]
// @Security ApiKeyAuth
func GetUserDetail(c *gin.Context) {

	nickname := c.Query("nickname")
	if nickname == "" {
		response.Fail(c, "nickname is required")
		return
	}

	user := new(req.UserBase)
	// var user models.UserBasic
	// err := global.DB.Omit("password").Where("nickname=?", nickname).First(&user).Error
	err := model.GetUserModel().Omit("password").Where("nickname=?", nickname).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Fail(c, "user not exist")
			return
		}
		response.Fail(c, "query failed, err:"+err.Error())
		return
	}

	/* b, _ := json.Marshal(user)
	    po := new(req.UserBase)
		json.Unmarshal(b,po)
	*/

	response.SuccessStruct(c, user)
}

// GetUserList
// @Summary 获取用户列表
// @Schemes
// @Description 获取用户列表
// @Tags 用户相关
// @Accept json
// @Produce json
// @Param page_num query int false "请输入页码"
// @Param page_size query int false "请输入每页数量"
// @Success 200 {object} req.UserBase
// @Router /v1/user-list [get]
// @Security ApiKeyAuth
func GetUserList(c *gin.Context) {
	page_size, err := strconv.Atoi(c.DefaultQuery("page_size", def.PAGE_SIZE))
	if err != nil {
		global.ZAPLOGGER.Sugar().Info("get user list strconv page_size error:", err)
		msg := fmt.Sprintf("get user list strconv page_size error:%s", err.Error())
		response.Fail(c, msg)
		return
	}

	page_num, err := strconv.Atoi(c.DefaultQuery("page_num", def.PAGE_NUM))
	if err != nil {
		global.ZAPLOGGER.Sugar().Info("get user list strconv page_num error:", err)
		msg := fmt.Sprintf("get user list strconv page_num error:%s", err.Error())
		response.Fail(c, msg)
		return
	}
	page_num = (page_num - 1) * page_size
	nickname := c.Query("nickname")

	var count int64
	// list := make([]*models.UserBasic, 0)
	list := make([]*req.UserBase, 0)
	d := model.GetUserList(nickname)
	err = d.Count(&count).Offset(page_num).Limit(page_size).Find(&list).Error
	if err != nil {
		log.Println("get user list error:", err)
		msg := fmt.Sprintf("get user list error:%s", err.Error())
		response.Fail(c, msg)
		return
	}
	// c.String(http.StatusOK, "Get User List")
	// utils.Success(c, map[string]interface{}{"list":  list,"count": count,}, "ok")
	response.Success(c, gin.H{"list": list, "count": count})
}

// UpdateUser
// @Summary 用户更新
// @Schemes
// @Description 用户更新
// @Tags 用户相关
// @Accept json
// @Produce json
// @Param userinfo body req.UpdateUserDto true "用户信息"
// @Success 200 {string} json "{"code":200,"msg":"","data":""}"
// @Router /v1/user-update [put]
// @Security ApiKeyAuth
func UpdateUser(c *gin.Context){
	var json req.UpdateUserDto
	if err := c.ShouldBindJSON(&json); err != nil {
		response.Fail(c, validate.Translate(err))
		return
	}

	d := model.GetUserModel()
	user := new(model.UserBasic)
	err := d.Where("email=?", json.Email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Fail(c, "user not exist")
			return
		}
		response.Fail(c, "query failed, err:"+err.Error())
		return
	}

	if json.Password != "" {
		user.Password = util.Md5String(json.Password)
	}
	if json.Nickname != "" {
		user.Nickname = json.Nickname
	}
	if json.Mobile != "" {
		user.Mobile = json.Mobile
	}
	log.Printf("%v", json)
	log.Printf("%v", user)

	if err := d.Save(&user).Error; err != nil {
		msg := fmt.Sprintf("update user faield error: %s", err.Error())
		response.Fail(c, msg)
		return
	}
	response.Success(c, gin.H{})

}