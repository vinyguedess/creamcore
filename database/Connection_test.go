package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jinzhu/gorm"

	"github.com/tembici/temapi-stations-list/database"
)

type UserTest struct {
	gorm.Model
	Name string `json:"name"`
}

func (m *UserTest) TableName() string {
	return "users"
}

func TestConnectingToDatabase(t *testing.T) {
	conn := database.GetDB()
	conn.CreateTable(&UserTest{})

	assert.True(t, conn.HasTable("users"))
}
