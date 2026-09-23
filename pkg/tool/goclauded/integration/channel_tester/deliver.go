package channel_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"testing"
)

func Deliver(
	t *testing.T,
	s *base.Server,
	x context.Context,
	callsign string,
) chan []client.QueueEntry {
	t.Helper()
	result := make(chan []client.QueueEntry, 1)
	go func() {
		response, e := s.RESTClient(t).GetChannelWithResponse(
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
