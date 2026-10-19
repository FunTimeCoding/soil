package outage

import (
	"github.com/funtimecoding/soil/pkg/peer/constant"
	"time"
)

func (o *OutageError) ErrorContext() (string, map[string]any) {
	return constant.ContextKey, map[string]any{
		constant.HostKey:     o.Peer,
		constant.StateKey:    constant.DownState,
		constant.SinceKey:    o.Since.Format(time.DateTime),
		constant.DurationKey: o.Duration.String(),
	}
}
