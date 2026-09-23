package goclaude

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/generative/model_context/channel"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/command_context"
	"time"
)

func runChannel(
	c *command_context.Context,
	s *channel.Server,
	interval time.Duration,
	v *recovery.Recovery,
) {
	s.Resolve(awaitCallsign(c, time.Now(), interval))
	callsign := s.Attach()
	failures := 0

	for {
		okay := false
		v.Run(
			func() {
				failures = pollChannel(c, s, callsign, failures)
				okay = failures == 0
			},
		)

		if !okay {
			time.Sleep(interval)
		}
	}
}
