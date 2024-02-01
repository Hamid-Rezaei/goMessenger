package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CreateGroupRequest struct {
	Name string `json:"name" validate:"required"`
	//Members
}

func (g CreateGroupRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(g); err != nil {
		return fmt.Errorf("create group request validation failed %w", err)
	}

	return nil
}
