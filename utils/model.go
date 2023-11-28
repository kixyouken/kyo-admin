package utils

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type ModelFile struct {
	Table string `json:"table"`
	Key   string `json:"key"`
}

func GetModel(c *gin.Context, model string) (*ModelFile, error) {
	modelFile := ModelFile{}
	data, err := os.ReadFile("./views/res/kyo/model/" + model + ".json")
	if err != nil {
		return nil, nil
	}
	json.Unmarshal(data, &modelFile)
	return &modelFile, nil
}

func GetKey(c *gin.Context, model string) string {
	modelFile, _ := GetModel(c, model)
	if modelFile.Key == "" {
		modelFile.Key = "id"
	}

	return modelFile.Key
}
