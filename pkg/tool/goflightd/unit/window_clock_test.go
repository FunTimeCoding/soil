package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/base"
	"testing"
	"time"
)

func TestEventWindowRunsOnServiceClock(t *testing.T) {
	s := base.NewWith(t, windowNow)
	assert.FatalOnError(
		t,
		store.New(s.Mapper).CreateEvent(
			*event.New(
				windowNow().Add(-30*time.Minute),
				"airportd",
				"com.apple.wifi",
				"scan",
				"clocked",
			),
		),
	)
	r, e := s.Client.GetEventsWithResponse(
		t.Context(),
		&client.GetEventsParams{},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *r.JSON200)
	assert.String(t, "clocked", (*r.JSON200)[0].Message)
}
