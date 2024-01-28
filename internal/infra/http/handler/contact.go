package handler

import (
	"errors"
	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"log"
	"net/http"
	"strconv"

	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) GetContactsList(c echo.Context) error {
	log.Printf("lll\n")
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	userId := userIDFromToken(c)
	if userId != uint(id) {
		return c.JSON(http.StatusForbidden, "Access Forbidden!")
	}
	contacts, err := h.contactRepo.GetList(c.Request().Context(), uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "No Contact Found!")
		}
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, contacts)
}

func (h *Handler) AddContact(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	userId := userIDFromToken(c)
	if userId != uint(id) {
		return c.JSON(http.StatusForbidden, "Access Forbidden!")
	}

	var r request.ContactAddRequest

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

	_, err = h.userRepo.GetUserByID(c.Request().Context(), r.ContactId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "Contact Not Found!")
		}
		return echo.ErrInternalServerError
	}

	check, err := h.contactRepo.GetById(c.Request().Context(), userId, r.ContactId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

		} else {
			return echo.ErrInternalServerError
		}
	}
	if check != nil {
		return c.JSON(http.StatusNoContent, "You Have This Contact!")
	}

	var contact model.Contact
	contact.UserId = userId
	contact.ContactId = r.ContactId
	contact.ContactName = r.ContactName

	res, err := h.contactRepo.Create(c.Request().Context(), contact)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(201, res)
}

func (h *Handler) DeleteContact(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	contactId, err := strconv.ParseUint(c.Param("contact_id"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	userId := userIDFromToken(c)
	if userId != uint(id) {
		return c.JSON(http.StatusForbidden, "Access Forbidden!")
	}
	log.Print(userId)
	check, err := h.contactRepo.GetById(c.Request().Context(), userId, uint(contactId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "Contact Not Found!")
		} else {
			return echo.ErrInternalServerError
		}
	}
	if check == nil {
		return c.JSON(http.StatusNoContent, "Contact Not Found!")
	}

	err2 := h.contactRepo.Delete(c.Request().Context(), check.ID)
	if err2 != nil {
		if errors.Is(err2, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "Contact Not Found!")
		} else {
			return echo.ErrInternalServerError
		}
	}
	return c.JSON(http.StatusOK, "Contact Deleted!")
}
