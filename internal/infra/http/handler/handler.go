package handler

import "github.com/Hamid-Rezaei/goMessenger/internal/infra/repository"

type Handler struct {
	userRepo repository.UserRepo,
	contactRepo repository.ContactRepo
}

func NewHandler(ur repository.UserRepo, cr repository.ContactRepo) *Handler {
	return &Handler{
		userRepo: ur,
		contactRepo: cr
	}
}
