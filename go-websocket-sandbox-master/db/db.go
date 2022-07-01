package db

import (
	"websocket-sandbox/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func connect() (*gorm.DB, error) {
	connection := config.GetDataBaseAccess()
	db, err := gorm.Open(
		"mysql",
		connection,
	)
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}

func DB() *gorm.DB {
	if db == nil {
		newDb, err := connect()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)
		//newDb.LogMode(true)
		db = newDb
	}
	return db
}
