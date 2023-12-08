package controllers

import (
	"fmt"
	"kyo-admin/services"
	"kyo-admin/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// Menu 获取 menu 所有数据
//
//	@param c
func Menu(c *gin.Context) {
	table := c.Param("table")
	result := []map[string]interface{}{}
	services.DbServices.Get(c, table, &result, "*", nil)
	menu := utils.BuildTree(c, result, 0)
	utils.Dataful(c, menu)
}

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
	tableFile, _ := utils.GetTableJson(c, table)
	var count int64
	services.DbServices.Count(c, &count, param, tableFile)
	result := []map[string]interface{}{}
	services.DbServices.Paginate(c, &result, "*", param, tableFile)
	utils.Paginate(c, result, count)
}

func Count(c *gin.Context) {
	param := utils.GetParam(c)
	table := c.Param("table")

	tableFile, _ := utils.GetTableJson(c, table)
	fmt.Println(tableFile)
	var count int64

	services.DbServices.Count(c, &count, param, tableFile)
	result := map[string]interface{}{}
	for _, v := range tableFile.Result.Count {
		result[v] = count
	}
	utils.Dataful(c, result)
}

func Line(c *gin.Context) {
	table := c.Param("table")
	echartsFile, _ := utils.GetEchartsJson(c, table)
	result := []map[string]interface{}{}
	services.DbServices.Line(c, &result, echartsFile)
	xAxis := []string{}
	yAxis := []int64{}
	for _, v := range result {
		t := v["login_date"].(time.Time).Format("2006-01-02")
		xAxis = append(xAxis, t)
		yAxis = append(yAxis, v["login_count"].(int64))
	}
	res := map[string]interface{}{
		"count": yAxis,
		"date":  xAxis,
	}
	utils.Dataful(c, res)
}

func Save(c *gin.Context) {

}

func Update(c *gin.Context) {
	param := utils.GetForm(c)
	table := c.Param("form")
	id := c.Param("id")
	services.DbServices.Update(c, table, id, param)
	utils.Successful(c, "更新成功")
}

func Delete(c *gin.Context) {
	table := c.Param("table")
	id := c.Param("id")
	services.DbServices.Delete(c, table, id)
	utils.Successful(c, "删除成功")
}
