package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ContactAddRequest struct {
	ContactId   uint   `json:"contact_id" validate:"required"`
	ContactName string `json:"contact_name" validate:"required"`
}

func (chr ContactAddRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(chr); err != nil {
		return fmt.Errorf("create chat request validation failed %w", err)
	}

	return nil
}
