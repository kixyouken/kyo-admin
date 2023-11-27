package utils

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type TableFile struct {
	Table  string        `json:"table"`
	Wheres []TableWheres `json:"wheres"`
	Joins  []TableJoins  `json:"joins"`
	Orders []TableOrders `json:"orders"`
}

type TableJoins struct {
	Join    string   `json:"join"`
	Table   string   `json:"table"`
	Foreign string   `json:"foreign"`
	Key     string   `json:"key"`
	Fields  []string `json:"fields"`
}

type TableOrders struct {
	Field string `json:"field"`
	Order string `json:"Order"`
}

type TableWheres struct {
	Field string `json:"field"`
	Value string `json:"value"`
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
