package bootstrap

import (
	"time"

	"github.com/runningape/goblog/app/models/user"
	"gorm.io/gorm"

	"github.com/runningape/goblog/app/models/article"
	"github.com/runningape/goblog/pkg/config"
	"github.com/runningape/goblog/pkg/model"
)

func SetupDB() {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	migration(db)

}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
