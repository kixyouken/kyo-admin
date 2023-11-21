package routers

import (
	"kyo-admin/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/views", "./views")

	admin := r.Group("admin")
	{
		admin.GET("/captcha", controllers.Captcha)
		admin.POST("/:table/login", controllers.Find)
		admin.GET("/model/:model", controllers.Model)
		admin.GET("/table/:table", controllers.Table)
		admin.GET("/form/:form", controllers.Form)
	}

	api := r.Group("api")
	{
		api.GET("/table/:table/get", controllers.Get)
		api.GET("/table/:table/paginate", controllers.Paginate)
		api.POST("/form/:form", controllers.Save)
		api.PUT("/form/:form", controllers.Update)
		api.DELETE("/table/:table", controllers.Delete)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong~",
		})
	})

	return r
}
