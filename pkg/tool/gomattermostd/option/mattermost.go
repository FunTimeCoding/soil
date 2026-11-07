package option

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
	webOption "github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Mattermost struct {
	webOption.Web
	LitePath         string
	PostgresLocator  string
	Mattermost       func() chat.MattermostClient
	Connector        func() goclauded.Client
	Monitor          *option.Monitor
	DebounceWindow   time.Duration
	PurgeWindow      time.Duration
	PurgeInterval    time.Duration
	FlushInterval    time.Duration
	ReconnectMinimum time.Duration
	ReconnectMaximum time.Duration
}
