package utils

import (
	"github.com/gin-gonic/gin"
)

func GetParam(c *gin.Context) map[string]string {
	// 获取所有查询参数
	queryParams := c.Request.URL.Query()
	// 创建一个 map 用于存储参数
	params := make(map[string]string)

	// 遍历查询参数并存储到 map 中
	for key, values := range queryParams {
		switch key {
		case "page", "limit":
		default:
			// 如果有多个值，只取第一个值
			params[key] = values[0]
		}
	}

	return params
}
