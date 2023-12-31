package services

import (
	"fmt"
	"kyo-admin/databases"
	"kyo-admin/utils"
	"strconv"
	"strings"
	"time"

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

func (s *sDbServices) Paginate(c *gin.Context, out, columns interface{}, param map[string]string, tableFile *utils.TableFile) error {
	return db.Table(tableFile.Table).
		Scopes(s.Page(c), s.Search(c, param), s.Fields(c, tableFile), s.Joins(c, tableFile), s.Orders(c, tableFile), s.Wheres(c, tableFile)).
		Select(columns).
		Find(out).Error
}

func (s *sDbServices) Find(c *gin.Context, table string, out, columns interface{}, param map[string]interface{}) error {
	return db.Table(table).
		Where(param).
		Select(columns).
		Limit(1).
		Find(out).Error
}

func (s *sDbServices) Count(c *gin.Context, count *int64, param map[string]string, tableFile *utils.TableFile) error {
	return db.Table(tableFile.Table).
		Scopes(s.Search(c, param), s.Joins(c, tableFile), s.Wheres(c, tableFile)).
		Limit(1).
		Count(count).Error
}

func (s *sDbServices) Line(c *gin.Context, out interface{}, echartsFile *utils.EchartsFile) error {
	return db.Table(echartsFile.Table).
		Select("COUNT(*) AS login_count", "DATE(created_at) AS login_date").
		Group("login_date").
		Limit(10).
		Find(out).
		Error
}

func (s *sDbServices) Read(c *gin.Context) {

}

func (s *sDbServices) Create(c *gin.Context) {

}

func (s *sDbServices) Update(c *gin.Context, table string, id, updates interface{}) error {
	return db.Table(table).
		Where(utils.GetKey(c, table)+" = ?", id).
		Updates(updates).Error
}

func (s *sDbServices) Delete(c *gin.Context, table string, id interface{}) error {
	t := time.Now()
	return db.Table(table).
		Where(utils.GetKey(c, table)+" = ?", id).
		Update("deleted_at", t).
		Error
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
				match := strings.ToUpper(slice[2])
				fmt.Println(slice)
				fmt.Println(match)
				switch match {
				case "LIKE":
					v = "%" + v + "%"
					db.Where(slice[0]+"."+slice[1]+" "+match+" ?", v)
				case "BETWEEN":
					between := strings.Split(v, " - ")
					db.Where(slice[0]+"."+slice[1]+" BETWEEN ? AND ?", between[0], between[1])
				default:
					db.Where(slice[0]+"."+slice[1]+" "+match+" ?", v)
				}
			}
		}
		return db
	}
}

func (s *sDbServices) Joins(c *gin.Context, tableFile *utils.TableFile) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tableFile != nil {
			for _, join := range tableFile.Joins {
				db.Joins(strings.ToUpper(join.Join) + " JOIN " + join.Table + " ON " + tableFile.Table + "." + join.Foreign + " = " + join.Table + "." + join.Key)
			}
		}
		return db
	}
}

func (s *sDbServices) Fields(c *gin.Context, tableFile *utils.TableFile) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fields := []string{}
		if tableFile != nil {
			fields = append(fields, tableFile.Table+".*")
			for _, join := range tableFile.Joins {
				for _, v := range join.Fields {
					fields = append(fields, join.Table+"."+v+" AS "+join.Table+"_"+v)
				}
			}
		}
		return db.Select(fields)
	}
}

func (s *sDbServices) Wheres(c *gin.Context, tableFile *utils.TableFile) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tableFile != nil {
			for _, v := range tableFile.Wheres {
				db.Where(v.Field + " " + v.Value)
			}
		}
		return db
	}
}

func (s *sDbServices) Orders(c *gin.Context, tableFile *utils.TableFile) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if tableFile != nil {
			for _, v := range tableFile.Orders {
				db.Order(v.Field + " " + v.Order)
			}
		}
		return db
	}
}
