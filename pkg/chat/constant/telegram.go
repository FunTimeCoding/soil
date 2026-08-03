package constant

import "github.com/funtimecoding/soil/pkg/console/constant"

const (
	TelegramTokenEnvironment    = "TELEGRAM_TOKEN"
	TelegramChannelEnvironment  = "TELEGRAM_CHANNEL"
	TelegramDatabaseEnvironment = "TELEGRAM_DATABASE"
)

var TelegramFormat = constant.ColorFormat.Copy()
