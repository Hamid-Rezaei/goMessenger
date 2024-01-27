package handler

import "github.com/Hamid-Rezaei/goMessenger/internal/infra/repository"

type Handler struct {
	userRepo repository.UserRepo,
	contactRepo repository.ContactRepo,
	chatRepo repository.ChatRepo,
	messageRepo repository.MessageRepo
}

func NewHandler(ur repository.UserRepo, cr repository.ContactRepo, chr repository.ChatRepo, mr repository.MessageRepo) *Handler {
	return &Handler{
		userRepo: ur,
		contactRepo: cr,
		chatRepo: chr,
		messageRepo: mr
	}
}
