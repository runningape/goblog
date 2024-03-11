package bootstrap

import (
	"time"

	"github.com/runningape/goblog/app/models/user"
	"gorm.io/gorm"

	"github.com/runningape/goblog/app/models/article"
	"github.com/runningape/goblog/pkg/model"
)

func SetupDB() {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	migration(db)

}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
