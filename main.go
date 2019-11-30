package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"nc_student.com/v1/route"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	route.All(e)

	log.Println(e.Start(":8080"))
}
