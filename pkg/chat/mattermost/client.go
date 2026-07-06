package mattermost

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/user_map"
	"github.com/mattermost/mattermost/server/public/model"
)

type Client struct {
	context     context.Context
	host        string
	token       string
	teamName    string
	channelName string
	client      *model.Client4
	team        *model.Team
	channel     *model.Channel
	webSocket   *model.WebSocketClient
	meCache     *model.User
	user        *user_map.Map
}
