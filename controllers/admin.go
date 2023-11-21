package controllers

import (
	"kyo-admin/utils"

	"github.com/gin-gonic/gin"
)

func Model(c *gin.Context) {

}

func Table(c *gin.Context) {

}

func Form(c *gin.Context) {

}

func Captcha(c *gin.Context) {
	captcha, _ := utils.GetCaptcha(c)
	utils.Dataful(c, captcha)
}
