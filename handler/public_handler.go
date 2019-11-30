package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	database "nc_student.com/v1/db"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func TestDb(c echo.Context) error {
	database.Test()
	return c.String(http.StatusOK, "OK")
}
