package option

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/face"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/monitor/option"
	"time"
)

type Mattermost struct {
	Address         string
	ServiceTokens   []string
	LitePath        string
	PostgresLocator string
	Mattermost      func() face.Client
	Connector       func() goclauded.Client
	Monitor         *option.Monitor
	DebounceWindow  time.Duration
	PurgeWindow     time.Duration
	PurgeInterval   time.Duration
}
