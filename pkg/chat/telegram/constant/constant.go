package constant

import "github.com/funtimecoding/soil/pkg/console/status/option"

const (
	TokenEnvironment    = "TELEGRAM_TOKEN"
	ChannelEnvironment  = "TELEGRAM_CHANNEL"
	DatabaseEnvironment = "TELEGRAM_DATABASE"

	UserBucket    = "user"
	ChannelBucket = "channel"
)

var Format = option.Color.Copy()
