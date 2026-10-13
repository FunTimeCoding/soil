package mattermost

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/user_map"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

type Client struct {
	context     context.Context
	host        string
	insecure    bool
	token       string
	teamName    string
	channelName string
	keepAlive   time.Duration
	client      *model.Client4
	team        *model.Team
	channel     *model.Channel
	webSocket   *model.WebSocketClient
	stopPing    chan struct{}
	meCache     *model.User
	user        *user_map.Map
	verbose     bool
	pool        *pool.Pool
}
