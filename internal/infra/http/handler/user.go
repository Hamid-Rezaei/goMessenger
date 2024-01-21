package handler

import (
	"github.com/Hamid-Rezaei/goMessenger/internal/domain/model"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/request"
	"github.com/Hamid-Rezaei/goMessenger/internal/infra/http/response"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (h *Handler) SignUp(c echo.Context) error {
	var req request.UserRegisterRequest

	// Bind Request
	if err := c.Bind(&req); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}
	// Validate Request
	if err := req.Validate(); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	var user model.User

	hash, err := user.HashPassword(req.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	//imagePath, err := user.UploadImage(req.Image)
	//if err != nil {
	//	return err
	//}
	//user.Image = imagePath
	user.Username = req.Username
	user.Firstname = req.Firstname
	user.Lastname = req.Lastname
	user.Bio = req.Bio
	user.Phone = req.Phone

	id, err := h.userRepo.Create(c.Request().Context(), user)

	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, id)
}

func (h *Handler) Login(c echo.Context) error {
	var req request.UserLoginRequest

	// Bind Request
	if err := c.Bind(&req); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}
	// Validate Request
	if err := req.Validate(); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	username := req.Username
	phone := req.Phone
	user, err := h.userRepo.GetByUsernamePhone(c.Request().Context(), username, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if user == nil {
		return c.JSON(http.StatusForbidden, "Access Forbidden!")
	}
	if !user.CheckPassword(req.Password) {
		return c.JSON(http.StatusForbidden, "Access Forbidden!")
	}

	return c.JSON(http.StatusOK, response.NewUserResponse(user))
}

func (h *Handler) CurrentUser(c echo.Context) error {
	user, err := h.userRepo.GetUserByID(c.Request().Context(), userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, "User Does not exist!")
	}
	return c.JSON(http.StatusOK, response.NewUserResponse(user))
}

func userIDFromToken(c echo.Context) uint {
	id, ok := c.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}
