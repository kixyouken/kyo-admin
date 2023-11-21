package controllers

import (
	"kyo-admin/services"
	"kyo-admin/utils"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	param := utils.GetParam(c)

	table := c.Param("table")
	result := []map[string]interface{}{}
	services.DbServices.Get(c, table, &result, "*", param)
	utils.Dataful(c, result)
}

func Paginate(c *gin.Context) {
	param := utils.GetParam(c)

	table := c.Param("table")
	var count int64
	services.DbServices.Count(c, table, &count, param)
	result := []map[string]interface{}{}
	services.DbServices.Paginate(c, table, &result, "*", param)
	utils.Paginate(c, result, count)
}

func Save(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
