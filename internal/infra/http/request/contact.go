package request

type ContactAddRequest struct {
	ContactId   uint   `json:"contact_id" validate:"required"`
	ContactName string `json:"contact_name" validate:"required"`
}
