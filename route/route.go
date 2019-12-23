package route

import (
	"github.com/labstack/echo/v4"
	"nc_student.com/v1/handler"
)

func All(e *echo.Echo) {
	Private(e)
	Staff(e)
	Public(e)
}

func Private(e *echo.Echo) {

}

func Staff(e *echo.Echo) {
	g := e.Group("api/student/v1/staff")
	g.POST("/student", handler.AddStudent)
	g.PUT("/student", handler.UpdateStudent)
	g.DELETE("/student", handler.DeleteOneStudent)
}

func Public(e *echo.Echo) {
	g := e.Group("api/student/v1/public")
	g.GET("/health", handler.HealthCheck)
	g.GET("/students", handler.GetAllStudents)
	g.GET("/student", handler.GetOneStudent)
	g.PATCH("/studentWithKeywords", handler.GetAllStudentsWithKeyword)
}
