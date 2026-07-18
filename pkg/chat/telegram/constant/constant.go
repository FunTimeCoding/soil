package constant

import "github.com/funtimecoding/soil/pkg/console/status/option"

const (
	TokenEnvironment    = "TELEGRAM_TOKEN"
	ChannelEnvironment  = "TELEGRAM_CHANNEL"
	DatabaseEnvironment = "TELEGRAM_DATABASE"
)

var Format = option.Color.Copy()
