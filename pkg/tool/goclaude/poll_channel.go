package goclaude

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func PollChannel(
	c face.Client,
	s *channel.Server,
	callsign string,
	failures int,
) int {
	stalled := map[string]string{
		constant.ChannelKindMeta: constant.ChannelStalledKind,
	}
	r, e := c.GetChannel(
		context.Background(),
		&client.GetChannelParams{Callsign: callsign},
	)

	if e != nil {
		failures++

		if failures == constant.ChannelStallThreshold {
			s.Push(
				fmt.Sprintf(constant.ChannelStalledMessage, failures),
				stalled,
			)
		}

		return failures
	}

	if failures >= constant.ChannelStallThreshold {
		s.Push(constant.ChannelResumedMessage, stalled)
	}

	for _, entry := range r.Entries {
		s.Push(
			entry.Body,
			map[string]string{constant.ChannelKindMeta: entry.Kind},
		)
	}

	return 0
}
