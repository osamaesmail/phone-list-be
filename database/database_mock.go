package database

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type suite struct {
	DB   *gorm.DB
	Conn *sql.DB
	Mock sqlmock.Sqlmock
}

func MockDb() (*suite, error) {
	s := &suite{}
	var err error

	s.Conn, s.Mock, err = sqlmock.New()
	if err != nil || s.Conn == nil || s.Mock == nil {
		return nil, err
	}

	dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: s.Conn})
	s.DB, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil || s.DB == nil {
		return nil, err
	}

	return s, nil
}
