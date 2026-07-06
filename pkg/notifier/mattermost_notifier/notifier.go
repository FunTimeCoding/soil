package mattermost_notifier

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/mattermost/mattermost/server/public/model"
)

type Notifier struct {
	mattermost *mattermost.Client
	channel    *model.Channel
	prefix     string
}
