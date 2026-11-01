package goclaude

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"time"
)

func AwaitCallsign(
	c face.Client,
	session string,
	since time.Time,
	interval time.Duration,
) string {
	for failures := 0; failures < constant.ChannelCallsignAttempts; {
		r, found, e := c.GetChannelCallsign(
			context.Background(),
			&client.GetChannelCallsignParams{Session: session, Since: since},
		)

		if e == nil && found {
			return r.Callsign
		}

		if e == nil {
			continue
		}

		failures++
		time.Sleep(interval)
	}

	return ""
}
