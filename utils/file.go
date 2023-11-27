package utils

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type TableFile struct {
	Table string       `json:"table"`
	Joins []TableJoins `json:"joins"`
}

type TableJoins struct {
	Join    string   `json:"join"`
	Table   string   `json:"table"`
	Foreign string   `json:"foreign"`
	Key     string   `json:"key"`
	Fields  []string `json:"fields"`
}

func GetTableJson(c *gin.Context, table string) (*TableFile, error) {
	tableFile := TableFile{}
	data, err := os.ReadFile("./views/res/json/table/" + table + ".json")
	if err != nil {
		return nil, nil
	}
	json.Unmarshal(data, &tableFile)
	return &tableFile, nil
}
