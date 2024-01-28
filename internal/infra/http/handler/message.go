package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) DeleteMessage(c echo.Context) error {
	chat_id, err := strconv.ParseUint(c.Param("chat_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	message_id, err := strconv.ParseUint(c.Param("message_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	user_id := userIDFromToken(c)

	check, err := h.messageRepo.GetMessage(c.Request().Context(), chat_id, message_id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Message Not Found!")
		}
		else{
			return echo.ErrInternalServerError
		}
	}
	if(check == nil){
		return c.JSON(http.StatusNotFound, "Message Not Found!")
	}
	if(check.SenderId == user_id){
		err := h.messageRepo.Delete(c.Request().Context(), chat_id, message_id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNotFound, "Message Not Found!")
			}
			else{
				return echo.ErrInternalServerError
			}
		}
		return err
	}
	else{
		return c.JSON(http.StatusNotFound, "Message Not Found!")
	}

	
}
