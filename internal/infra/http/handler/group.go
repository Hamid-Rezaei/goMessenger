package handler

import (
	"errors"
	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/request"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func (h *Handler) AddGroup(c echo.Context) error {
	userId := userIDFromToken(c)

	var r request.CreateChatRequest

	// Bind Request
	if err := c.Bind(&r); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}
	// Validate Request
	if err := r.Validate(); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}
	_, err := h.userRepo.GetUserByID(c.Request().Context(), r.ReceiverId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Receiver Not Found!")
		}
		return echo.ErrInternalServerError
	}

	check, err := h.chatRepo.GetChat(c.Request().Context(), userId, r.ReceiverId)
	log.Println(check)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.ErrInternalServerError
		}
	}
	if check != nil {
		return c.JSON(http.StatusBadRequest, "You Have Chat With This User!")
	}
	var chat model.Chat
	chat.CreatedAt = time.Now()
	res, err := h.chatRepo.Create(c.Request().Context(), chat)
	if err != nil {
		return echo.ErrInternalServerError
	}

	var people1 model.People
	people1.ChatID = res.ID
	people1.UserID = userId
	var people2 model.People
	people2.ChatID = res.ID
	people2.UserID = r.ReceiverId

	_, err2 := h.peopleRepo.Create(c.Request().Context(), people1)
	if err2 != nil {
		return echo.ErrInternalServerError
	}
	_, err3 := h.peopleRepo.Create(c.Request().Context(), people2)
	if err3 != nil {
		return echo.ErrInternalServerError
	}

	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(201, response.NewChatResponse(res))
}

func (h *Handler) DeleteGroup(c echo.Context) error {

}

func (h *Handler) AddMember(c echo.Context) error {

}

func (h *Handler) DeleteMember(c echo.Context) error {

}
