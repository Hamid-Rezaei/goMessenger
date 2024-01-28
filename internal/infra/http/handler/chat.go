package handler

import (
	"errors"
	"log"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/request"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) AddChat(c echo.Context) error {
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
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.ErrInternalServerError
		}
	}
	if check != nil {
		return c.JSON(http.StatusNoContent, "You Have Chat With This User!")
	}

	var chat model.Chat
	chat.People = []uint{userId, r.ReceiverId}
	chat.CreatedAt = time.Now()

	res, err := h.chatRepo.Create(c.Request().Context(), chat)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(201, response.NewChatResponse(res))
}

func (h *Handler) GetChatsList(c echo.Context) error {
	userId := userIDFromToken(c)
	chats, err := h.chatRepo.GetChatList(c.Request().Context(), userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "No Chat Not Found!")
		}
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, response.NewChatsResponse(chats))
}

func (h *Handler) GetChat(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("chat_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	userId := userIDFromToken(c)
	chat, err := h.chatRepo.GetChatById(c.Request().Context(), uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Chat Not Found!")
		}
		return echo.ErrInternalServerError
	}

	if slices.Contains(chat.People, userId) {
		messages, err := h.messageRepo.GetMessagesOfAChat(c.Request().Context(), uint(id))
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.ErrInternalServerError
			}
		}
		return c.JSON(http.StatusOK, response.ChatWithMessageResponse{
			Chat:     chat,
			Messages: messages,
		})
	} else {
		return c.JSON(http.StatusNotFound, "Chat Not Found!")
	}
}

func (h *Handler) DeleteChat(c echo.Context) error {
	chatId, err := strconv.ParseUint(c.Param("chat_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	userId := userIDFromToken(c)

	chat, err := h.chatRepo.GetChatById(c.Request().Context(), uint(chatId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "Chat Not Found!")
		} else {
			return echo.ErrInternalServerError
		}
	}
	if chat == nil {
		return c.JSON(http.StatusNoContent, "Chat Not Found!")
	}

	if slices.Contains(chat.People, userId) {
		err := h.chatRepo.Delete(c.Request().Context(), uint(chatId))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNoContent, "Chat Not Found!")
			} else {
				return echo.ErrInternalServerError
			}
		}
		return c.JSON(http.StatusOK, "Chat Deleted!")
	} else {
		return c.JSON(http.StatusNotFound, "Chat Not Found!")
	}
}
