package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getAllStudents)
	e.POST("/students", createStudent)
	e.GET("/students/:id", getStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func getAllStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students\n")
}
func createStudent(c echo.Context) error {
	return c.String(http.StatusOK, "Create students\n")
}
func getStudent(c echo.Context) error {
	id := c.Param("id") // Pega o parametro que foi passado
	message := "GET" + getId(id)
	return c.String(http.StatusOK, message)
}
func updateStudent(c echo.Context) error {
	id := c.Param("id")
	message := "UPDATE" + getId(id)
	return c.String(http.StatusOK, message)
}
func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	message := "DELETE" + getId(id)
	return c.String(http.StatusOK, message)
}

func getId(id string) string {
	return fmt.Sprintf("%s student \n", id)
}
