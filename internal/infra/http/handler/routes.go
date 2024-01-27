package handler

import (
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/router/middleware"
	"github.com/Hamid-Rezaei/goMessenger/internal/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	contact := v1.Group("/users/:user_id/contacts", middleware.JWT(utils.GetSigningKey()))
	contact.GET("", h.GetContactsList)
	contact.POST("", h.AddContact)
	contact.DELETE("/:contact_id")

	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	user := v1.Group("/users", middleware.JWT(utils.GetSigningKey()))
	user.GET("/:id", h.CurrentUser)
	user.PATCH("/:id", h.UpdateUser)
	user.DELETE("/:id", h.DeleteUser)
	user.GET("", h.SearchUser)
}
