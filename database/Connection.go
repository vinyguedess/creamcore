package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var conn *gorm.DB

func makeConnection() {
	db, err := gorm.Open("sqlite3", "creamcore.sqlite3")
	if err != nil {
		panic(err)
	}

	conn = db
}

func GetConnection() *gorm.DB {
	if conn == nil {
		makeConnection()
	}

	return conn
}
