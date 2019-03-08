package creamcore_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	creamcore "github.com/vinyguedess/cream-core"
)

type UserTest struct {
	gorm.Model
	Name string `json:"name"`
}

func (m *UserTest) TableName() string {
	return "users"
}

func TestConnectingToDatabase(t *testing.T) {
	conn := creamcore.GetConnection()
	conn.CreateTable(&UserTest{})

	assert.True(t, conn.HasTable("users"))
}
