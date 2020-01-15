package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"nc_student.com/v1/middlewareCustom"
	"nc_student.com/v1/route"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	// e.Use(middlewareCustom.SimpleLogger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: middlewareCustom.Logger,
	}))
	route.All(e)

	log.Println(e.Start(":9090"))
}
