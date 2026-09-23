package channel_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel"
	"github.com/funtimecoding/soil/pkg/identity"
)

func NewChannelServer() *channel.Server {
	return channel.New(
		identity.New("goclaude", "channel test", "goclaude [command]"),
		"0.0.0",
		"",
	)
}
