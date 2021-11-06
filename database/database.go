package database

import (
	"log"

	"phone-list/config"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	if config.Cfg().DBDriver == "sqlite3" {
		db, err = openSqlite3DB()
		if err != nil {
			log.Fatalf("database.init err: %v", err)
		}
		return
	}

	log.Fatalf("DB_DRIVER %s not available", config.Cfg().DBDriver)
}

func DB() *gorm.DB {
	return db
}
