package service_tester

import (
	chat "github.com/funtimecoding/soil/pkg/chat/face"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"time"
)

func (o *Tester) daemonOption() *option.Mattermost {
	result := option.New()
	result.LitePath = o.Path
	result.Mattermost = func() chat.MattermostClient { return o.client }
	result.Connector = func() goclauded.Client { return o.connector }

	if o.monitor {
		result.Monitor.Enabled = true
		result.Monitor.Interval = time.Hour
	}

	return result
}
