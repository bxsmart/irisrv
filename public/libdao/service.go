package libdao

import (
	"github.com/bxsmart/bxcore/log"
	xormcore "github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"time"
)

const (
	DefaultRdsDriver = "mysql"
)

type RdsServiceImpl struct {
	options *RdSqlOptions
	tables  []interface{}
	Db      *xorm.Engine
}

type RdSqlOptions struct {
	Driver             string
	Address            string
	User               string
	Password           string
	DbName             string
	TablePrefix        string
	MaxOpenConnections int
	MaxIdleConnections int
	ConnMaxLifetime    int
	Debug              bool
}

type XormLogger struct {
	xormcore.ILogger
}

func NewRdsService(options *RdSqlOptions) RdsServiceImpl {
	impl := RdsServiceImpl{}

	impl.options = options

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return options.TablePrefix + defaultTableName
	}

	url := options.User + ":" + options.Password + "@tcp(" + options.Address + ")/" + options.DbName + "?charset=utf8&parseTime=True"
	db, err := xorm.NewEngine(DefaultRdsDriver, url)
	if err != nil {
		log.Fatalf("mysql connection error:%s", err.Error())
	}

	db.DB().SetConnMaxLifetime(time.Duration(options.ConnMaxLifetime) * time.Second)
	db.DB().SetMaxIdleConns(options.MaxIdleConnections)
	db.DB().SetMaxOpenConns(options.MaxOpenConnections)

	db.SetLogger(xorm.NewSimpleLogger(os.Stdout))
	db.Logger().SetLevel(xormcore.LOG_INFO)

	impl.Db = db

	return impl
}
