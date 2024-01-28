package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) AddChat(c echo.Context) error {
	user_id := userIDFromToken(c)

	var r request.ChatAddRequest

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

	user, err := h.userRepo.GetUserByID(c.Request().Context(), r.ReceiverId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Receiver Not Found!")
		}
		return echo.ErrInternalServerError
	}

	check, err := h.chatRepo.GetChat(c.Request().Context(), user_id, r.ReceiverId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.ErrInternalServerError
		}
	}
	if check != nil {
		return c.JSON(http.StatusNoContent, "You Have Chat With This User!")
	}

	var chat model.Chat
	chat.People = []uint{user_id, r.ReceiverId}
	chat.CreatedAt = time.Now()

	res, err := h.chatRepo.Create(c.Request().Context(), chat)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(201, res)
}

func (h *Handler) GetChatsList(c echo.Context) error {
	user_id := userIDFromToken(c)
	chats, err := h.chatRepo.GetChatList(c.Request().Context(), uint(user_id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "No Chat Not Found!")
		}
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, chats)
}

func (h *Handler) GetChat(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("chat_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	user_id := userIDFromToken(c)
	chat, err := h.chatRepo.GetChatById(c.Request().Context(), uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Chat Not Found!")
		}
		return echo.ErrInternalServerError
	}

	if slices.Contains(chat.People, user_id){
		messages, err := h.messageRepo.GetMessagesOfAChat(c.Request().Context(), uint(id))
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.ErrInternalServerError
			}
		}
		return c.JSON(http.StatusOK, {chat : chat, messages : messages})
	}
	else{
		return c.JSON(http.StatusNotFound, "Chat Not Found!")
	} 
}

func (h *Handler) DeleteChat(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("chat_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	user_id := userIDFromToken(c)

	check, err := h.chatRepo.GetChatById(c.Request().Context(), chat_id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "Chat Not Found!")
		}
		else{
			return echo.ErrInternalServerError
		}
	}
	if(check == nil){
		return c.JSON(http.StatusNoContent, "Chat Not Found!")
	}

	if slices.Contains(chat.People, user_id){
		err := h.chatRepo.Delete(c.Request().Context(), chat_id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNoContent, "Chat Not Found!")
			}
			else{
				return echo.ErrInternalServerError
			}
		}
		return err
	}
	else{
		return c.JSON(http.StatusNotFound, "Chat Not Found!")
	} 
}
