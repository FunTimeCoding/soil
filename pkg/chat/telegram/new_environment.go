package telegram

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/telegram/store"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	var s *store.Store

	if path := environment.Optional(constant.TelegramDatabaseEnvironment); path != "" {
		s = store.New(lite.New(logger.New(context.Background()), path))
	}

	return New(environment.Required(constant.TelegramTokenEnvironment), s)
}
