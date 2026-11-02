package goclaude

import (
	"context"
	"fmt"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
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
		generativeConstant.ChannelKindMeta: generativeConstant.ChannelStalledKind,
	}
	r, e := c.GetChannel(
		context.Background(),
		&client.GetChannelParams{Callsign: callsign},
	)

	if e != nil {
		failures++

		if failures == generativeConstant.ChannelStallThreshold {
			s.Push(
				fmt.Sprintf(generativeConstant.ChannelStalledMessage, failures),
				stalled,
			)
		}

		return failures
	}

	if failures >= generativeConstant.ChannelStallThreshold {
		s.Push(generativeConstant.ChannelResumedMessage, stalled)
	}

	for _, entry := range r.Entries {
		s.Push(
			entry.Body,
			map[string]string{generativeConstant.ChannelKindMeta: entry.Kind},
		)
	}

	return 0
}
