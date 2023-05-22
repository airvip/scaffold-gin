package middleware

import (
	"net/http"
	"scaffold-gin/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 authorization header
		s := c.GetHeader("Authorization")
		if s == "" || !strings.HasPrefix(s, "Bearer ") {
			
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "未携带token, 无权限访问",
			})
			c.Abort()
			return
		}

		token := s[7:]
		// log.Println("get token", token)
		// c2, err := common.ParseToken(token)
		_, err := util.ParseToken(token)
		if err != nil {

			if err == util.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": 402,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 403,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		// 验证通过之后获取claim中的数据
		/*uid := c2.Id
		log.Println("user identity is:",c2.StandardClaims.Id)
		user := new(models.UserBasic)
		models.DB.First(&user, uid)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 402,
				"msg":  "权限异常",
			})
			c.Abort()
			return
		}
		// 用户存在写入 Context
		// log.Println("user info", *user)
		c.Set("user", user)*/
		c.Next()
	}
}

