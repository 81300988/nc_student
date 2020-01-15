package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	db "nc_student.com/v1/db"
	"nc_student.com/v1/model"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetAllStudents(c echo.Context) error {
	students, err := db.GetAllStudent()
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, students)
}

func GetAllStudentsWithKeyword(c echo.Context) error {
	var student model.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	students, err2 := db.FindNameStartsWith(&student)
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, students)
}

func GetOneStudent(c echo.Context) error {
	id := c.QueryParam("id")
	student, err := db.GetStudentbyID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, student)
}
