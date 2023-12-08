package utils

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type EchartsFile struct {
	Table string `json:"table"`
	Key   string `json:"key"`
	Type  string `json:"type"`
}

func GetEchartsJson(c *gin.Context, table string) (*EchartsFile, error) {
	echartsFile := EchartsFile{}
	data, err := os.ReadFile("./views/res/kyo/echarts/" + table + ".json")
	if err != nil {
		return nil, nil
	}
	json.Unmarshal(data, &echartsFile)
	return &echartsFile, nil
}
