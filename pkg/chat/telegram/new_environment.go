package telegram

import (
	"github.com/funtimecoding/soil/pkg/chat/telegram/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.TokenEnvironment),
		environment.Optional(constant.DatabaseEnvironment),
	)
}
