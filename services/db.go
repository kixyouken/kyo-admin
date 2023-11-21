package services

import (
	"fmt"
	"kyo-admin/databases"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sDbServices struct{}

var DbServices = sDbServices{}

var db = databases.InitMysql()

func (s *sDbServices) Get(c *gin.Context, table string, out, columns interface{}, param map[string]string) error {
	return db.Table(table).
		Scopes(s.Search(c, param)).
		Select(columns).
		Find(out).Error
}

func (s *sDbServices) Paginate(c *gin.Context, table string, out, columns interface{}, param map[string]string) error {
	return db.Table(table).
		Scopes(s.Page(c), s.Search(c, param)).
		Select(columns).
		Find(out).Error
}

func (s *sDbServices) Count(c *gin.Context, table string, count *int64, param map[string]string) error {
	return db.Table(table).
		Scopes(s.Search(c, param)).
		Count(count).Error
}

func (s *sDbServices) Read(c *gin.Context) {

}

func (s *sDbServices) Create(c *gin.Context) {

}

func (s *sDbServices) Update(c *gin.Context) {

}

func (s *sDbServices) Delete(c *gin.Context) {

}

func (s *sDbServices) Page(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("limit"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *sDbServices) Search(c *gin.Context, param map[string]string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for k, v := range param {
			if v != "" {
				slice := strings.Split(k, ".")
				switch slice[1] {
				case "like":
					v = "%" + v + "%"
				}
				where := fmt.Sprintf("%s %s \"%s\"", slice[0], slice[1], v)
				db.Where(where)
			}
		}
		return db
	}
}
