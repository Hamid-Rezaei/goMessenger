package request

type ContactAddRequest struct {
	ContactId   uint   `json:"contactid" validate:"required"`
	ContactName string `json:"contactname" validate:"required"`
}
