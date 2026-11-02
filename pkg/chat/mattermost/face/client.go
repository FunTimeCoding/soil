package face

import (
	"context"
	"github.com/mattermost/mattermost/server/public/model"
)

type Client interface {
	ChannelClient
	UserClient
	FileClient
	DialogClient
	PostingClient
	PostReadingClient
	Close()
	Configuration() (map[string]string, error)
	Host() string
	MustConfiguration() map[string]string
	MustRefreshSocket()
	Reachable(x context.Context) error
	RefreshSocket(x context.Context) error
	WebSocket() *model.WebSocketClient
}
