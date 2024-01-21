package response

import (
	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"github.com/Hamid-Rezaei/goMessenger/internal/utils"
)

type UserResponse struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Bio       string `json:"bio"`
	//Image     image.Image `json:"file,omitempty"`
	Token string `json:"token"`
}

func NewUserResponse(u *model.User) *UserResponse {
	r := new(UserResponse)
	r.Firstname = u.Firstname
	r.Lastname = u.Lastname
	r.Username = u.Username
	r.Phone = u.Phone
	r.Bio = u.Bio
	//r.Image, _ = u.RetrieveImage(u.Image)
	r.Token = utils.GenerateJWT(u.ID)
	return r
}
