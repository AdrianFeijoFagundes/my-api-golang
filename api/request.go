package api

import "fmt"

type StudentRequest struct {
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active *bool  `json:"active"` // force input true/false
}

func errParamRequired(param, typ string) error {
	return fmt.Errorf("param '%s' of type '%s' is required", param, typ)
}

func (s *StudentRequest) Validate() error {
	if s.Name == "" {
		return errParamRequired("name", "string")
	}
	if s.CPF == "" {
		return errParamRequired("cpf", "string")

	}
	if s.Email == "" {
		return errParamRequired("email", "string")
	}
	if s.Age == 0 {
		return errParamRequired("age", "int")
	}
	if s.Active == nil {
		return errParamRequired("active", "boolean")
	}
	return nil
}
