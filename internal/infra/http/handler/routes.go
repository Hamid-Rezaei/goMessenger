package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	//user := v1.Group("/user", middleware.JWT(utils.GetSigningKey()))
	//user.GET("", h.CurrentUser)

}
