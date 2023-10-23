package routers

import (
	"kyo-admin/controllers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/views", "./views")

	api := r.Group("api")
	{
		api.GET("/table/:table/get", controllers.Get)
		api.GET("/table/:table/paginate", controllers.Paginate)
		api.POST("/form/:form", controllers.Save)
		api.PUT("/form/:form", controllers.Update)
		api.DELETE("/table/:table", controllers.Delete)
	}

	return r
}
