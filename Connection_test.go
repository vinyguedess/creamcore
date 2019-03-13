package creamcore_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/vinyguedess/creamcore"
)

type UserTest struct {
	gorm.Model
	Name string `json:"name"`
}

func (m *UserTest) TableName() string {
	return "users"
}

func TestConnectingToDatabase(t *testing.T) {
	godotenv.Load()

	conn := creamcore.GetConnection()
	conn.DropTableIfExists(&UserTest{})
	conn.CreateTable(&UserTest{})

	assert.True(t, conn.HasTable("users"))
}

func TestConnectingToDatabaseErrorIfWrongAuthentication(t *testing.T) {
	creamcore.DestroyConnection()

	os.Setenv("DB_DRIVER", "mysql")
	os.Setenv("DB_HOST", "127.0.0.2")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "unknow_database")

	assert.Panics(t, func() {
		creamcore.GetConnection()
	})
}

func TestConnectionAuthGetDriverSqlite(t *testing.T) {
	auth := creamcore.ConnectionAuth{
		Driver: "sqlite",
	}

	assert.Equal(t, auth.GetDriver(), "sqlite3")
}

func TestConnectionAuthGetDriverOthers(t *testing.T) {
	auth := creamcore.ConnectionAuth{
		Driver: "mysql",
	}

	assert.Equal(t, auth.GetDriver(), "mysql")
}

func TestConnectionAuthGetConnStringSqlite(t *testing.T) {
	auth := creamcore.ConnectionAuth{
		Driver: "sqlite",
		Name:   "anydb.sqlite",
	}

	assert.Equal(t, auth.GetConnString(), "anydb.sqlite")
}

func TestConnectionAuthGetConnStringMySQL(t *testing.T) {
	auth := creamcore.ConnectionAuth{
		Driver: "mysql",
		Host:   "localhost",
		Port:   3306,
		Name:   "anydb",
		User:   "root",
		Pass:   "123321",
	}

	assert.Equal(t, auth.GetConnString(), "root:123321@/anydb?charset=utf8&parseTime=True&loc=Local")
}

func TestConnectionAuthGetConnStringOthers(t *testing.T) {
	auth := creamcore.ConnectionAuth{
		Driver: "unknown",
	}

	assert.Equal(t, auth.GetConnString(), "")
}
