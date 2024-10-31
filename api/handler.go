package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/adrian-feijo-fagundes/my-api-golang/db"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}
	return c.JSON(http.StatusOK, students)
}
func (api *API) createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := api.DB.AddStudent(student); err != nil { // FORMA "simplificada de tratar um erro"
		return c.String(http.StatusInternalServerError, "Error to create students\n")
	}

	return c.String(http.StatusOK, "Create students\n")
}
func (api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) // Pega o parametro que foi passado
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ID")
	}
	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student id not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}
	return c.JSON(http.StatusOK, student)
}
func (api *API) updateStudent(c echo.Context) error {

	return c.String(http.StatusOK, "")
}
func (api *API) deleteStudent(c echo.Context) error {

	return c.String(http.StatusOK, "")
}
