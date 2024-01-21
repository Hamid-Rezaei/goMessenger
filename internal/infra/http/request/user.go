package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserRegisterRequest struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	//Image     *multipart.FileHeader `json:"-"`
	Bio string `json:"bio"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

func (ur UserRegisterRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(ur); err != nil {
		return fmt.Errorf("register request validation failed %w", err)
	}

	return nil
}

func (ul UserLoginRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(ul); err != nil {
		return fmt.Errorf("user login request validation failed %w", err)
	}

	return nil
}
