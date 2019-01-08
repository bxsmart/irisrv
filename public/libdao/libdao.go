package libdao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Table interface {
	TableName() string
}

type RdsService interface {
	// create tables
	SetTables(tables []interface{})
	CreateTables() error

	// base functions
	Add(item interface{}) error
	Del(item interface{}) error
	Take(item interface{}) error
	First(item interface{}) error
	Last(item interface{}) error
	Save(item interface{}) error
	FindAll(item interface{}) error
	FindForUpdate(item interface{}, tx *gorm.DB) error
}

// find for update
func (s *RdsServiceImpl) ForUpdate(tx *gorm.DB) *gorm.DB {
	return tx.Set("gorm:query_option", " FOR UPDATE ")
}

// add single item
func (s *RdsServiceImpl) Add(item interface{}) error {
	return s.Db.Create(item).Error
}

// del single item
func (s *RdsServiceImpl) Del(item interface{}) error {
	return s.Db.Delete(item).Error
}

// select one item order asc
func (s *RdsServiceImpl) Take(item interface{}) error {
	return s.Db.Take(item).Error
}

// select first item order by primary key asc
func (s *RdsServiceImpl) First(item interface{}) error {
	return s.Db.First(item).Error
}

// select the last item order by primary key asc
func (s *RdsServiceImpl) Last(item interface{}) error {
	return s.Db.Last(item).Error
}

// update single item
func (s *RdsServiceImpl) Save(item interface{}) error {
	return s.Db.Save(item).Error
}

// find all items in table where primary key > 0
func (s *RdsServiceImpl) FindAll(item interface{}) error {
	table := item.(Table)
	return s.Db.Table(table.TableName()).Find(item, s.Db.Where("id > ", 0)).Error
}

func (s *RdsServiceImpl) SetTables(tables []interface{}) {
	s.tables = tables
}

func (s *RdsServiceImpl) CreateTables() error {
	for _, t := range s.tables {
		if ok := s.Db.HasTable(t); !ok {
			if err := s.Db.CreateTable(t).Error; err != nil {
				return fmt.Errorf("create mysql table error:%s", err.Error())
			}
		}
	}

	// auto migrate to keep schema update to date
	// AutoMigrate will ONLY create tables, missing columns and missing indexes,
	// and WON'T change existing column's type or delete unused columns to protect your data
	return s.Db.AutoMigrate(s.tables...).Error
}
