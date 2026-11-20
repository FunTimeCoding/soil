package channel_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
)

func (o *Tester) Deliver(
	s *base.Server,
	x context.Context,
	callsign string,
) chan []client.QueueEntry {
	o.t.Helper()
	result := make(chan []client.QueueEntry, 1)
	go func() {
		response, e := s.RESTClient().GetChannelWithResponse(
			x,
			&client.GetChannelParams{Callsign: callsign},
		)

		if e != nil || response.JSON200 == nil {
			result <- nil

			return
		}

		result <- response.JSON200.Entries
	}()

	return result
}
