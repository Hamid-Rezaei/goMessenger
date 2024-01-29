package handler

import (
	"errors"
	"log"
	"net/http"
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

	people, err2 := h.peopleRepo.Get(c.Request().Context(), userId, uint(id))
	if err2 != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Chat Not Found!")
		}
		return echo.ErrInternalServerError
	}
	if people == nil {
		return c.JSON(http.StatusNotFound, "Chat Not Found!")
	} else {
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
		err := h.chatRepo.Delete(c.Request().Context(), uint(chatId))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNoContent, "Chat Not Found!")
			} else {
				return echo.ErrInternalServerError
			}
		}
		return c.JSON(http.StatusOK, "Chat Deleted!")
	}
}
