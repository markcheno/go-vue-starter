package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// sqlite db driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB abstraction
type DB struct {
	*gorm.DB
}

// NewPostgresDB - postgres database
func NewPostgresDB(dataSourceName string) *DB {

	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	//db.LogMode(true)

	return &DB{db}
}

// NewSqliteDB - sqlite database
func NewSqliteDB(databaseName string) *DB {

	db, err := gorm.Open("sqlite3", databaseName)
	if err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	//db.LogMode(true)

	return &DB{db}
}
