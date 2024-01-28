package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CreateChatRequest struct {
	ReceiverId uint `json:"receiver_id" validate:"required"`
}

func (chr CreateChatRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(chr); err != nil {
		return fmt.Errorf("create chat request validation failed %w", err)
	}

	return nil
}
