package services

import (
	"fmt"
	"kyo-admin/databases"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sDbServices struct{}

var DbServices = sDbServices{}

var db = databases.InitMysql()

func (s *sDbServices) Get(c *gin.Context, table string, out, columns interface{}) error {
	return db.Table(table).
		Select(columns).
		Find(out).Error
}

func (s *sDbServices) Paginate(c *gin.Context, table string, out, columns interface{}) error {
	return db.Table(table).
		Select(columns).
		Scopes(s.Page(c)).
		Find(out).Error
}

func (s *sDbServices) Count(c *gin.Context, table string, count *int64) error {
	return db.Table(table).
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
		fmt.Println(c.Query("page"))
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}

		fmt.Println(c.Query("limit"))
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
