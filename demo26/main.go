package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type SignUpRequest struct {
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required"`
	PasswordConformation string `json:"password_conformation" validate:"required,eqfield=password"`
}

func main() {
	var (
		v             *validator.Validate
		signUpRequest SignUpRequest
		err           error
	)

	v = validator.New()

	signUpRequest = SignUpRequest{
		Name:                 "",
		Email:                "",
		Password:             "",
		PasswordConformation: "",
	}

	if err = v.Struct(signUpRequest); err != nil {
		fmt.Printf("%v\n", err)
	}

}
