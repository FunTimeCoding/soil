package channel_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"time"
)

func (o *Tester) Callsign(
	s *base.Server,
	x context.Context,
	since time.Time,
) chan string {
	o.t.Helper()
	result := make(chan string, 1)
	go func() {
		response, e := s.RESTClient().GetChannelCallsignWithResponse(
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
