package handler

import (
	"errors"
	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/request"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
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
	group.Members = members
	group.Owner = owner
	g, err := h.groupRepo.Create(c.Request().Context(), group)
	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, g.ID)
}

func (h *Handler) DeleteGroup(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	userID := userIDFromToken(c)
	g, err := h.groupRepo.GetGroupByOwnerID(c.Request().Context(), userID, uint(id))

	if g == nil {
		log.Printf("%s", "User does not have permission!")
		return c.JSON(http.StatusForbidden, "User does not have permission or not found!")
	}
	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	if err := h.groupRepo.Delete(c.Request().Context(), uint(int(id))); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "Group Was Deleted Successfully.")
}

//func (h *Handler) AddMember(c echo.Context) error {
//
//}
//
//func (h *Handler) DeleteMember(c echo.Context) error {
//
//}
