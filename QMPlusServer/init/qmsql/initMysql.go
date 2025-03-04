package qmsql

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gin-vue-admin/config"
)

var DEFAULTDB *gorm.DB

// 初始化数据库并产生数据库全局变量
func InitMysql(admin config.Admin) *gorm.DB {
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		log.Printf("DEFAULTDB数据库启动异常%S", err)
	} else {
		DEFAULTDB = db
		DEFAULTDB.DB().SetMaxIdleConns(10)
		DEFAULTDB.DB().SetMaxOpenConns(100)
	}
	DEFAULTDB.LogMode(true)
	return DEFAULTDB
}
