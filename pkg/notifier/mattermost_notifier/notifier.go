package mattermost_notifier

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/mattermost/mattermost/server/public/model"
)

type Notifier struct {
	mattermost face.MattermostClient
	channel    *model.Channel
	prefix     string
}
