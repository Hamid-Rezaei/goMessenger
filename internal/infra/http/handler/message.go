package handler

import (
	"errors"
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
		return err
	} else {
		return c.JSON(http.StatusNotFound, "Message Not Found!")
	}

}
