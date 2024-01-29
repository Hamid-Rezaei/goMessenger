package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type AddMessageRequest struct {
	Content string `json:"content" validate:"required"`
}

func (mr AddMessageRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(mr); err != nil {
		return fmt.Errorf("new message request validation failed %w", err)
	}

	return nil
}
