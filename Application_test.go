package creamcore_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vinyguedess/creamcore"
)

func TestSimpleHandler(t *testing.T) {
	app := creamcore.NewApplication("Simple Handler Test")
	app.Register("/", func(request *http.Request) (int, string) {
		return http.StatusOK, "Testing"
	}, "GET")

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	app.Router.ServeHTTP(recorder, request)

	assert.Equal(t, recorder.Body.String(), "Testing")
	assert.Equal(t, recorder.Code, http.StatusOK)
}
