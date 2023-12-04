package middlewares

import (
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
		c.Next()
	}
}
