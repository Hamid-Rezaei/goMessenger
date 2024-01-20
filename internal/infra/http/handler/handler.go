package handler

import "github.com/Hamid-Rezaei/goMessenger/internal/infra/repository"

type Handler struct {
	userRepo repository.UserRepo
}

func NewHandler(ur repository.UserRepo) *Handler {
	return &Handler{
		userRepo: ur,
	}
}
