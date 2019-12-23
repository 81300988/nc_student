package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	db "nc_student.com/v1/db"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetAllStudents(c echo.Context) error {
	students, err := db.GetAllStudent()
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, students)
}

func GetAllStudentsWithKeyword(c echo.Context) error {
	var student db.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	students, err2 := db.FindNameStartsWith(&student)
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, students)
}

func GetOneStudent(c echo.Context) error {
	id := c.QueryParam("id")
	student, err := db.GetOneStudent(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, student)
}
