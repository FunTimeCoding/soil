package constant

import "github.com/funtimecoding/soil/pkg/console/status/option"

const (
	TelegramTokenEnvironment    = "TELEGRAM_TOKEN"
	TelegramChannelEnvironment  = "TELEGRAM_CHANNEL"
	TelegramDatabaseEnvironment = "TELEGRAM_DATABASE"
)

var TelegramFormat = option.Color.Copy()
