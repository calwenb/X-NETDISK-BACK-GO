package util

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"netdisk-back/app/config"
	"time"
)

var db *gorm.DB
var (
	dsn = config.Get("mysql.dataSourceName")
)

func init() {
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB, err := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	db = DB

}
func GetDBConn() (*gorm.DB, error) {
	if db != nil {
		return db, errors.New("数据库连接失败")
	}
	return db, nil
}
