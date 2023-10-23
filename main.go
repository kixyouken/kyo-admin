package main

import (
	"kyo-admin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := routers.GetRouter()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong~",
		})
	})
	// 监听并在 0.0.0.0:9090 上启动服务
	r.Run(":9090")
}
