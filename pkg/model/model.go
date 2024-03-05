package model

import (
	"github.com/runningape/goblog/pkg/logger"
	"gorm.io/gorm"

	// GORM 的 MySQL 数据库驱动导入
	"gorm.io/driver/mysql"
	gormlogger "gorm.io/gorm/logger"
)

// DB gorm.DB 对象
var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:dyhuangZz223@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})
	logger.LogError(err)

	return DB
}
