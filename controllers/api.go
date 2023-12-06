package controllers

import (
	"kyo-admin/services"
	"kyo-admin/utils"

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

// func Find(c *gin.Context) {
// 	param := utils.GetForm(c)
// 	if !utils.VerifyCaptcha(param["id"].(string), param["vercode"].(string)) {
// 		utils.Errorful(c, "验证码错误")
// 		return
// 	}

// 	table := c.Param("table")
// 	result := map[string]interface{}{}
// 	delete(param, "id")
// 	delete(param, "vercode")
// 	services.DbServices.Find(c, table, &result, "*", param)
// 	if len(result) == 0 {
// 		utils.Errorful(c, "账号或密码错误")
// 		return
// 	}
// 	utils.Dataful(c, result)
// }

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
