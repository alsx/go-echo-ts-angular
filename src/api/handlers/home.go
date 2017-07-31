package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Index contains links to endpoint
type Index struct {
	Links []string
}

// NewIndex initate API Index list
func NewIndex() Index {
	return Index{[]string{"signin/", "login/", "user/", "fb-login/", "fb-callback/"}}
}

// IndexHandler contains methods to show index of api endpoints
var IndexHandler = NewIndex()

// List shows index of api endpoints
func (v *Index) List(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, v, "    ")
}
