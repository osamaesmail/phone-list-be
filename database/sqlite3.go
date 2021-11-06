package database

import (
	"database/sql"
	"phone-list/config"
	"regexp"

	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var sqliteDriverName = "sqlite3_extended"

func openSqlite3DB() (*gorm.DB, error) {
	var err error
	registerSqlite3RegEx()

	conn, err := sql.Open(sqliteDriverName, config.Cfg().Sqlite3Path)
	if err != nil {
		return nil, err
	}

	db, err = gorm.Open(sqlite.Dialector{
		DriverName: sqliteDriverName,
		DSN:        config.Cfg().Sqlite3Path,
		Conn:       conn,
	}, &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func registerSqlite3RegEx() {
	sql.Register(sqliteDriverName,
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				return conn.RegisterFunc("regexp", regexp.MatchString, true)
			},
		})
}
