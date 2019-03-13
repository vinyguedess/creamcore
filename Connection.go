package creamcore

import (
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Imports SQLite parser
)

// ConnectionAuth holds and treats database connection
type ConnectionAuth struct {
	Driver string
	Host   string
	Port   int64
	Name   string
	User   string
	Pass   string
}

// GetDriver returns connection driver
func (auth *ConnectionAuth) GetDriver() string {
	if auth.Driver == "sqlite" {
		return "sqlite3"
	}

	return auth.Driver
}

// GetConnString gets connection string formatted for GORM
func (auth *ConnectionAuth) GetConnString() string {
	if auth.Driver == "sqlite" {
		return auth.Name
	} else if auth.Driver == "mysql" {
		return auth.User + ":" + auth.Pass + "@/" + auth.Name + "?charset=utf8&parseTime=True&loc=Local"
	}

	return ""
}

var conn *gorm.DB

func makeConnection() {
	port, _ := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 0)
	auth := ConnectionAuth{
		Driver: os.Getenv("DB_DRIVER"),
		Host:   os.Getenv("DB_HOST"),
		Port:   port,
		Name:   os.Getenv("DB_NAME"),
		User:   os.Getenv("DB_USER"),
		Pass:   os.Getenv("DB_PASS"),
	}

	db, err := gorm.Open(auth.GetDriver(), auth.GetConnString())
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

// DestroyConnection is responsible of destroying connection to database
func DestroyConnection() {
	if conn != nil {
		conn.Close()
	}
	conn = nil
}
