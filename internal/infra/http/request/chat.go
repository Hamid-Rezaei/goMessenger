package request

type ChatAddRequest struct {
	ReceiverId uint `json:"receivername" validate:"required"`
}
