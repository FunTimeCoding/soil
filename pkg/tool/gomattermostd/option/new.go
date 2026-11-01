package option

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
)

func New() *Mattermost {
	return &Mattermost{
		Mattermost: func() face.Client {
			return mattermost.NewEnvironment()
		},
		Connector:      func() goclauded.Client { return nil },
		Monitor:        option.New(),
		DebounceWindow: constant.DebounceWindow,
		PurgeWindow:    constant.PurgeWindow,
		PurgeInterval:  constant.PurgeInterval,
	}
}
