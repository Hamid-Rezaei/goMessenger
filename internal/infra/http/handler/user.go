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
	"strconv"
)

func (h *Handler) SignUp(c echo.Context) error {
	var r request.UserRegisterRequest

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

	var user model.User

	hash, err := user.HashPassword(r.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	//imagePath, err := user.UploadImage(r.Image)
	//if err != nil {
	//	return err
	//}
	//user.Image = imagePath
	user.Username = r.Username
	user.Firstname = r.Firstname
	user.Lastname = r.Lastname
	user.Bio = r.Bio
	user.Phone = r.Phone

	id, err := h.userRepo.Create(c.Request().Context(), user)

	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, id)
}

func (h *Handler) Login(c echo.Context) error {
	var r request.UserLoginRequest

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

	username := r.Username
	phone := r.Phone
	user, err := h.userRepo.GetByUsernamePhone(c.Request().Context(), username, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if user == nil {
		return c.JSON(http.StatusForbidden, "Access Forbidden!")
	}
	if !user.CheckPassword(r.Password) {
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

func (h *Handler) UpdateUser(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	var r request.UserUpdateRequest

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

	userID := userIDFromToken(c)
	if userID != uint(id) {
		return c.JSON(http.StatusForbidden, "Access Forbidden!")
	}

	user, err := h.userRepo.GetUserByID(c.Request().Context(), uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "User Not Found!")
		}

		return echo.ErrInternalServerError
	}

	hash, err := user.HashPassword(r.Password)
	if err != nil {
		return err
	}

	u := model.User{
		Firstname: r.Firstname,
		Lastname:  r.Lastname,
		Username:  r.Username,
		Password:  hash,
		Phone:     r.Phone,
		Bio:       r.Bio,
	}

	if err := h.userRepo.Update(c.Request().Context(), &u, uint(id)); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "User Was Updated Successfully.")

}

func (h *Handler) DeleteUser(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("%v\n", err)
		return echo.ErrBadRequest
	}

	userID := userIDFromToken(c)
	if userID != uint(id) {
		return c.JSON(http.StatusForbidden, "Access Forbidden!")
	}

	if err := h.userRepo.Delete(c.Request().Context(), uint(id)); err != nil {
		log.Printf("%v\n", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, "User Was Deleted Successfully.")
}

func (h *Handler) SearchUser(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	// TODO: here must check jwt expiration

	user, err := h.userRepo.SearchUser(c.Request().Context(), keyword)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNoContent, "User Not Found!")
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, response.NewUserSearchResponse(user))
}

func userIDFromToken(c echo.Context) uint {
	id, ok := c.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}
