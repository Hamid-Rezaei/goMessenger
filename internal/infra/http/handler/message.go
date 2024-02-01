package handler

import (
	"errors"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/request"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) DeleteMessage(c echo.Context) error {
	chatId, err := strconv.ParseUint(c.Param("chat_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	messageId, err := strconv.ParseUint(c.Param("message_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	userId := userIDFromToken(c)

	check, err := h.messageRepo.GetMessage(c.Request().Context(), uint(chatId), uint(messageId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Message Not Found!")
		} else {
			return echo.ErrInternalServerError
		}
	}
	if check == nil {
		return c.JSON(http.StatusNotFound, "Message Not Found!")
	}
	if check.SenderId == userId {
		err := h.messageRepo.Delete(c.Request().Context(), uint(chatId), uint(messageId))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNotFound, "Message Not Found!")
			} else {
				return echo.ErrInternalServerError
			}
		}
		return c.JSON(http.StatusOK, "Message Deleted")
	} else {
		return c.JSON(http.StatusNotFound, "Message Not Found!")
	}

}

func (h *Handler) AddMessage(c echo.Context) error {
	chatId, err := strconv.ParseUint(c.Param("chat_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	userId := userIDFromToken(c)

	var r request.AddMessageRequest

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
	chat, err := h.chatRepo.GetChatById(c.Request().Context(), uint(chatId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusBadRequest, "Chat Not Found!")
		} else {
			return echo.ErrInternalServerError
		}
	}
	if chat == nil {
		return c.JSON(http.StatusBadRequest, "Chat Not Found!")
	}

	people, err2 := h.peopleRepo.Get(c.Request().Context(), userId, uint(chat.ID))
	if err2 != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusBadRequest, "Chat Not Found!")
		}
		return echo.ErrInternalServerError
	}
	if people == nil {
		return c.JSON(http.StatusNotFound, "Chat Not Found!")
	} else {
		temp, _ := h.peopleRepo.GetChatUsers(c.Request().Context(), chat.ID)
		for _, element := range temp {
			if element != userId {
				message, err := h.messageRepo.AddMessage(c.Request().Context(), uint(chatId), r.Content, userId, element)
				if err != nil {
					log.Println(people)
					return echo.ErrInternalServerError
				}
				h.peopleRepo.AddNewMessages(c.Request().Context(), chat.ID, userId)
				return c.JSON(http.StatusOK, message)
			}
		}
		return c.JSON(http.StatusNotFound, "Chat Not Found!")
	}
}
