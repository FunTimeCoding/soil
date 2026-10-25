package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"testing"
	"time"
	_ "time/tzdata"
)

func TestStoreRangeSpansTheFallBack(t *testing.T) {
	berlin, e := time.LoadLocation("Europe/Berlin")
	assert.FatalOnError(t, e)
	at := time.Date(2026, 10, 25, 0, 30, 0, 0, time.UTC).In(berlin)
	s := store.New(lite.NewMemory())
	defer s.Close()
	s.MustCreateEvent(event.Event{Time: at, Process: "ensembled"})
	events, f := s.EventsByTimeRange(
		at.Add(-time.Hour),
		at.Add(time.Hour),
		100,
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, len(events))
}
