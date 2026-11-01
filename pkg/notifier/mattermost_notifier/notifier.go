package mattermost_notifier

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	"github.com/mattermost/mattermost/server/public/model"
)

type Notifier struct {
	mattermost face.Client
	channel    *model.Channel
	prefix     string
}
