package request

type ChatAddRequest struct {
	ReceiverId uint `json:"receiver_name" validate:"required"`
}
