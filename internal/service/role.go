package service

import (
	"encoding/json"
	"fmt"
	"log"
	"scaffold-gin/internal/def"
	"scaffold-gin/common/global"
	"scaffold-gin/common/response"
	"scaffold-gin/internal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetRoleList
// @Summary 获取角色列表
// @Schemes
// @Description 获取角色列表
// @Tags 角色相关
// @Accept json
// @Produce json
// @Param keyword query string false "关键词"
// @Param rule_identity query string false "规则"
// @Param page_num query int false "请输入页码"
// @Param page_size query int false "请输入每页数量"
// @Success 200 {object} req.RoleBasic
// @Router /v1/role-list [get]
// @Security ApiKeyAuth
func GetRoleList(c *gin.Context) {
	page_size, err := strconv.Atoi(c.DefaultQuery("page_size", def.PAGE_SIZE))
	if err != nil {
		log.Println("get rule list strconv page_size error:", err)
		response.Fail(c, "get rule list strconv page_size failed, err:"+err.Error())
		return
	}

	page_num, err := strconv.Atoi(c.DefaultQuery("page_num", def.PAGE_NUM))
	if err != nil {
		log.Println("get rule list strconv page_num error:", err)
		global.ZAPLOGGER.Sugar().Errorf("get rule list strconv page_num error: %s", err)
		response.Fail(c, "get rule list strconv page_num failed, err:"+err.Error())
		return
	}
	page_num = (page_num - 1) * page_size
	keyword := c.Query("keyword")
	rule_identity := c.Query("rule_identity")

	var count int64
	list := make([]*model.RoleBasic, 0)
	d := model.GetRoleList(keyword, rule_identity)
	err = d.Count(&count).Offset(page_num).Limit(page_size).Find(&list).Error
	if err != nil {
		response.Fail(c, "get rule list failed, err:"+err.Error())
		return
	}
	// c.String(http.StatusOK, "Get User List")
	response.Success(c, gin.H{"list": list, "count": count})
}

// GetRoleDetail
// @Summary 获取角色详情
// @Schemes
// @Description 获取角色详情
// @Tags 角色相关
// @Accept json
// @Produce json
// @Param identity query string false "唯一标识"
// @Success 200 {object} req.RoleBasic
// @Router /v1/role-detail [get]
// @Security ApiKeyAuth
func GetRoleDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		response.Fail(c, "identity is required")
		return
	}

	data := new(model.RoleBasic)

	cacheKey := fmt.Sprintf(def.GetCacheString(def.USER_DETAIL), identity)
	s, err := global.REDIS.Get(cacheKey).Result()
	if err == nil {
		err = json.Unmarshal([]byte(s), &data)
		if err != nil {
			global.ZAPLOGGER.Sugar().Errorf(cacheKey+" unmarshal failed, err: %s", err)
		}
		response.Success(c, gin.H{"item": data})
		return
	}

	err = global.DB.Preload("RoleRule").Preload("RoleRule.RuleBasic").Where("identity=?", identity).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Fail(c, "query no data")
			return
		}
		response.Fail(c, "query failed, err:"+err.Error())
		return
	}
	b, _ := json.Marshal(data)
	err = global.REDIS.Set(cacheKey, b, 10 * time.Second).Err()
	if err != nil {
		global.ZAPLOGGER.Sugar().Errorf(cacheKey+"set in redis failed, err:%s", err)
	}

	response.Success(c, gin.H{"item": data})
}