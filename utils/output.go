package utils

import "github.com/gin-gonic/gin"

func Successful(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":   0,
		"status": 200,
		"msg":    "success",
	})
}

func Errorful(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":   0,
		"status": 400,
		"msg":    "error",
	})
}

func Dataful(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code":   0,
		"status": 200,
		"data":   data,
		"msg":    "success",
	})
}

func Paginate(c *gin.Context, data, count interface{}) {
	c.JSON(200, gin.H{
		"code":   0,
		"status": 200,
		"data":   data,
		"count":  count,
		"msg":    "success",
	})
}
