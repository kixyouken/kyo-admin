package middlewares

import (
	"kyo-admin/services"
	"kyo-admin/utils"

	"github.com/gin-gonic/gin"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token := c.Query("access_token")
		if access_token == "" {
			utils.Logoutful(c, "access_token 不存在，请重新登录~")
			c.Abort()
			return
		}

		result := map[string]interface{}{}
		services.DbServices.Find(c, "admins", &result, "*", map[string]interface{}{"access_token": access_token})
		if result == nil {
			utils.Logoutful(c, "access_token 错误，请重新登录~")
			c.Abort()
			return
		}
		c.Next()
	}
}
