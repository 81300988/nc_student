package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	db "nc_student.com/v1/db"
	"nc_student.com/v1/model"
)

func AddStudent(c echo.Context) error {
	var student model.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.AddOneStudent(&student)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateStudent(c echo.Context) error {
	var student model.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.UpdateOneStudent(&student)
	if err != nil {
		log.Printf("update error :%v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)

}

func DeleteOneStudent(c echo.Context) error {
	var student model.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.DeleteOneStudent(&student)
	if err != nil {
		log.Printf("update error :%v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}
