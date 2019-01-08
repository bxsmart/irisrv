package libdao

import (
	"byex.io/irisrv/public/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

const (
	DefaultRdsDriver = "mysql"
)

type RdsServiceImpl struct {
	options *RdSqlOptions
	tables  []interface{}
	Db      *gorm.DB
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

type LogWriter struct {
	gorm.Logger
}

func (writer *LogWriter) Println(values ...interface{}) {
	log.Println(gorm.LogFormatter(values...)...)
}

func NewRdsService(options *RdSqlOptions) RdsServiceImpl {
	impl := RdsServiceImpl{}

	impl.options = options

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return options.TablePrefix + defaultTableName
	}

	if len(options.Driver) <= 0 {
		options.Driver = DefaultRdsDriver
	}

	url := options.User + ":" + options.Password + "@tcp(" + options.Address + ")/" + options.DbName + "?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(options.Driver, url)
	if err != nil {
		log.Fatalf("mysql connection error:%s", err.Error())
	}

	db.DB().SetConnMaxLifetime(time.Duration(options.ConnMaxLifetime) * time.Second)
	db.DB().SetMaxIdleConns(options.MaxIdleConnections)
	db.DB().SetMaxOpenConns(options.MaxOpenConnections)

	db.SetLogger(&gorm.Logger{LogWriter: &LogWriter{}})

	db.LogMode(options.Debug)

	impl.Db = db

	return impl
}
