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
	contact.DELETE("/:contact_id", h.DeleteContact)

	guestUsers := v1.Group("/users")
	guestUsers.POST("/register", h.SignUp)
	guestUsers.POST("/login", h.Login)

	user := v1.Group("/users", middleware.JWT(utils.GetSigningKey()))
	user.GET("/:id", h.CurrentUser)
	user.PATCH("/:id", h.UpdateUser)
	user.DELETE("/:id", h.DeleteUser)
	user.GET("", h.SearchUser)

	chat := v1.Group("/chats", middleware.JWT(utils.GetSigningKey()))
	chat.GET("/:chat_id", h.GetChat)
	chat.DELETE("/:chat_id", h.DeleteChat)
	chat.GET("", h.GetChatsList)
	chat.POST("", h.AddChat)
	chat.DELETE("/:chat_id/messages/:message_id", h.DeleteMessage)
	chat.POST("/:chat_id/messages", h.AddMessage)
	chat.GET("/:chat_id/messages/new", h.GetChatNewMessages)

	group := v1.Group("/groups", middleware.JWT(utils.GetSigningKey()))
	group.POST("", h.AddGroup)
	group.DELETE("/:group_id", h.DeleteGroup)
	group.PATCH("/:group_id", h.AddMember)
	group.DELETE("/:group_id/:user_id", h.DeleteMember)

}
