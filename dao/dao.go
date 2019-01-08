package dao

import (
	"byex.io/irisrv/public/libdao"
	"byex.io/irisrv/public/log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PageResult struct {
	Data      []interface{} `json:"data"`
	PageIndex int           `json:"pageNum"`
	PageSize  int           `json:"pageSize"`
	Total     int           `json:"total"`
}

type RdsService struct {
	libdao.RdsServiceImpl
}

func NewDb(options *libdao.RdSqlOptions) *RdsService {
	var s RdsService

	s.RdsServiceImpl = libdao.NewRdsService(options)

	var tables []interface{}

	// TODO 若数据库表已经创建则不用每次都校验
	s.SetTables(tables)

	if err := s.CreateTables(); err != nil {
		log.Fatalf(err.Error())
	}

	return &s
}
