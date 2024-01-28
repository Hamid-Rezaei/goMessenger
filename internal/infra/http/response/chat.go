package response

import "github.com/Hamid-Rezaei/goMessenger/internal/domain/model"

type ChatResponse struct {
	Chat     *model.Chat      `json:"chat"`
	Messages *[]model.Message `json:"messages"`
}
