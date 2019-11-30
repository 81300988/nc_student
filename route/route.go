package route

import (
	"github.com/labstack/echo/v4"
	"nc_student.com/v1/handler"
)

func All(e *echo.Echo) {
	Private(e)
	Saff(e)
	Public(e)
}

func Private(e *echo.Echo) {

}

func Saff(e *echo.Echo) {

}

func Public(e *echo.Echo) {
	g := e.Group("api/student/v1/public")
	g.GET("/health", handler.HealthCheck)
	g.GET("/test", handler.TestDb)
}
