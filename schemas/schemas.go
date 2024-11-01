package schemas

import (
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}
type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"registration"`
}
