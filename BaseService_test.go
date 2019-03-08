package creamcore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	creamcore "github.com/vinyguedess/cream-core"
)

func TestHasErrorWithErrors(t *testing.T) {
	service := new(creamcore.BaseService)
	service.AddError("okay", "my boy")

	assert.True(t, service.HasErrors())
}

func TestHasErrorWithoutErrors(t *testing.T) {
	service := new(creamcore.BaseService)
	assert.False(t, service.HasErrors())
}

func TestAddError(t *testing.T) {
	service := new(creamcore.BaseService)
	service.AddError("vish", "maria")

	errors := service.GetErrors()
	assert.Equal(t, errors[0].Field, "vish")
	assert.Equal(t, errors[0].Message, "maria")
}
