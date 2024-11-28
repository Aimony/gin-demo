package mysql

import (
	"fmt"
	"gin-demo/dao"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

const (
	UserName     = "root"
	PassWord     = "123456"
	Host         = "127.0.0.1"
	Port         = 3306
	Database     = "gotest"
	MaxLifetime  = 60 * time.Second
	MaxIdletime  = 30 * time.Second
	MaxOpenconns = 6
	MaxIdleconns = 2
	Dialect      = "mysql"
)

type DataSouce struct {
	db *gorm.DB
}

func NewDataSource() *DataSouce {
	var db *gorm.DB

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Asia%%2FShanghai",
		UserName, PassWord, Host, Port, Database)
	db, err := gorm.Open(Dialect, dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	db.DB().SetConnMaxLifetime(MaxLifetime)
	db.DB().SetConnMaxIdleTime(MaxIdletime)
	db.DB().SetMaxOpenConns(MaxOpenconns)
	db.DB().SetMaxOpenConns(MaxIdleconns)

	return &DataSouce{db: db}
}

// 操作book 表
func (d *DataSouce) BookDao() dao.BookDao {
	return &dao.BookDaoImpl{
		DB: d.db,
	}
}
