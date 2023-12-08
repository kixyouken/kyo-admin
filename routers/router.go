package routers

import (
	"kyo-admin/controllers"
	"kyo-admin/middlewares"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/views", "./views")

	// r.GET("/", controllers.Index)

	admin := r.Group("admin")
	{
		admin.GET("/captcha", controllers.Captcha)
		admin.POST("/:table/login", controllers.Login)
		admin.GET("/:table/logout", controllers.Logout)
	}

	r.Use(middlewares.Token())

	api := r.Group("api")
	{
		api.GET("/table/:table/menu", controllers.Menu)
		api.GET("/table/:table/get", controllers.Get)
		api.GET("/table/:table/paginate", controllers.Paginate)
		api.GET("/table/:table/count", controllers.Count)

		api.GET("/echarts/:table/line", controllers.Line)

		api.POST("/form/:form", controllers.Save)
		api.PUT("/form/:form/:id", controllers.Update)
		api.DELETE("/table/:table/:id", controllers.Delete)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong~",
		})
	})

	return r
}
