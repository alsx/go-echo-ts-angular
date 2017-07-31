package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Versions contains links to endpoint
type Versions struct {
	Links []string
}

// NewVersions initate API versions list
func NewVersions() Versions {
	return Versions{[]string{"v1/"}}
}

// VersionsHandler contains methods to show index of api endpoints
var VersionsHandler = NewVersions()

// List shows index of api endpoints
func (v *Versions) List(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, v, "    ")
}
