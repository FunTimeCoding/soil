package telegram

import (
	"github.com/funtimecoding/soil/pkg/bolt"
	"github.com/funtimecoding/soil/pkg/chat/telegram/database/channel"
	"github.com/funtimecoding/soil/pkg/chat/telegram/database/user"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	client   *tgbotapi.BotAPI
	Verbose  bool
	database *bolt.Client
	channels []*channel.Channel
	users    []*user.User
}
