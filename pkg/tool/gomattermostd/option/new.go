package option

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	peerConstant "github.com/funtimecoding/soil/pkg/peer/constant"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
)

func New() *Mattermost {
	return &Mattermost{
		Mattermost: func() chat.MattermostClient {
			return mattermost.NewEnvironment()
		},
		Connector:        func() goclauded.Client { return nil },
		Monitor:          option.New(),
		DebounceWindow:   constant.DebounceWindow,
		PurgeWindow:      constant.PurgeWindow,
		PurgeInterval:    constant.PurgeInterval,
		FlushInterval:    constant.FlushTick,
		ReconnectMinimum: peerConstant.ProbeMinimum,
		ReconnectMaximum: peerConstant.ProbeMaximum,
	}
}
