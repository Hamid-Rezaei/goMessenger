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
)

func (h *Handler) AddGroup(c echo.Context) error {
	userId := userIDFromToken(c)

	var r request.CreateGroupRequest

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

	owner, err := h.userRepo.GetUserByID(c.Request().Context(), userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "owner Not Found!")
		}
		return echo.ErrInternalServerError
	}

	members, err := h.userRepo.GetUsersByID(c.Request().Context(), r.Members)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Member Not Found!")
		}
		return echo.ErrInternalServerError
	}

	var group model.Group
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
