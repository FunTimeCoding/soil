package channel_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"testing"
	"time"
)

func Callsign(
	t *testing.T,
	s *base.Server,
	x context.Context,
	since time.Time,
) chan string {
	t.Helper()
	result := make(chan string, 1)
	go func() {
		response, e := s.RESTClient(t).GetChannelCallsignWithResponse(
			x,
			&client.GetChannelCallsignParams{
				Session: "session-1",
				Since:   since,
			},
		)

		if e != nil || response.JSON200 == nil {
			result <- ""

			return
		}

		result <- response.JSON200.Callsign
	}()

	return result
}
