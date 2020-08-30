package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{"user_id":1,"name":"example 1"}`
)

func TestCreateUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	w := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, w)

	if assert.NoError(t, CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, userJSON+"\n", w.Body.String())
	}
}

func TestUsers(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	w := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, w)

	if assert.NoError(t, Users(c)) {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "["+userJSON+"]\n", w.Body.String())
	}
}
