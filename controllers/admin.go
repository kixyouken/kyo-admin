package controllers

import (
	"kyo-admin/services"
	"kyo-admin/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Model(c *gin.Context) {

}

func Table(c *gin.Context) {

}

func Form(c *gin.Context) {

}

func Login(c *gin.Context) {
	param := utils.GetForm(c)
	if !utils.VerifyCaptcha(param["id"].(string), param["vercode"].(string)) {
		utils.Errorful(c, "验证码错误")
		return
	}

	table := c.Param("table")
	result := map[string]interface{}{}
	delete(param, "id")
	delete(param, "vercode")
	services.DbServices.Find(c, table, &result, "*", param)
	if len(result) == 0 {
		utils.Errorful(c, "账号或密码错误")
		return
	}
	key := utils.GetKey(c, table)
	access_token := uuid.New()
	services.DbServices.Update(c, table, result[key], map[string]interface{}{
		"access_token": access_token,
	})
	result["access_token"] = access_token
	utils.Dataful(c, result)
}

func Logout(c *gin.Context) {
	table := c.Param("table")
	utils.Successful(c, table)
}

func Captcha(c *gin.Context) {
	captcha, _ := utils.GetCaptcha(c)
	utils.Dataful(c, captcha)
}
