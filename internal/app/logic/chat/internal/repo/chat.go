package repo

import (
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/adapter"
	"github.com/cd-home/Hissssss/internal/app/logic/chat/internal/model"
)

type ChatRepo struct {
}

func NewChatRepo() adapter.ChatRepo {
	return &ChatRepo{}
}

func (c *ChatRepo) CreateAllMessage(msg *model.AllMessage) error {
	return nil
}
