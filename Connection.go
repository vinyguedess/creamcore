package creamcore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Imports SQLite parser
)

var conn *gorm.DB

func makeConnection() {
	db, err := gorm.Open("sqlite3", "creamcore.sqlite3")
	if err != nil {
		panic(err)
	}

	conn = db
}

// GetConnection responsible for generate a connection if necessary and return
// its object
func GetConnection() *gorm.DB {
	if conn == nil {
		makeConnection()
	}

	return conn
}
