package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/runningape/goblog/logger"
)

var DB *sql.DB

func Initialize() {
	initDB()
	createTables()
}

func initDB() {
	var err error

	config := mysql.Config{
		User:                 "root",
		Passwd:               "dyhuangZz223",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	logger.LogError(err)

}

func createTables() {
	createArticlesSQL := `
		CREATE TABLE IF NOT EXISTS articles(
			id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
			title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
			body longtext COLLATE utf8mb4_unicode_ci
		);
	`
	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
