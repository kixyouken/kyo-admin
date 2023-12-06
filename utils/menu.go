package utils

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type MenuInfo struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	Jump     string `json:"jump"`
	ParentID string `json:"parent_id"`
	List     []*MenuInfo
}

func GetMenu(c *gin.Context) (*MenuInfo, error) {
	menu := MenuInfo{}
	data, err := os.ReadFile("./views/res/kyo/menu.json")
	if err != nil {
		return nil, nil
	}
	json.Unmarshal(data, &menu)
	return &menu, nil
}

// BuildTree 处理树状结构
//
//	@param c
//	@param menuInfo
//	@param parentID
//	@return []map
func BuildTree(c *gin.Context, menuInfo []map[string]interface{}, parentID int) []map[string]interface{} {
	menu, _ := GetMenu(c)
	var tree []map[string]interface{}
	for i := range menuInfo {
		// fmt.Printf("类型：%T ", menuInfo[i]["parent_id"])
		// fmt.Printf("类型：%T ", menuInfo[i]["id"])
		menuInfo[i]["jump"] = menuInfo[i][menu.Jump]
		if int(menuInfo[i][menu.ParentID].(int64)) == parentID {
			list := BuildTree(c, menuInfo, int(menuInfo[i][menu.Key].(uint64)))
			if len(list) > 0 {
				menuInfo[i]["list"] = list
			}
			tree = append(tree, menuInfo[i])
		}
	}
	return tree
}
