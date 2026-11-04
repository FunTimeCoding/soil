package subscribed

import (
	chat "github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
)

type Subscribed struct {
	Service    *service.Service
	Watcher    *watcher.Watcher
	Mattermost *chat.Client
	Roster     *goclauded.Client
}
