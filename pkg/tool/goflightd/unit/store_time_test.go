package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/base"
	"testing"
	"time"
	_ "time/tzdata"
)

func TestStoreRangeSpansTheFallBack(t *testing.T) {
	berlin, e := time.LoadLocation("Europe/Berlin")
	assert.FatalOnError(t, e)
	at := time.Date(2026, 10, 25, 0, 30, 0, 0, time.UTC).In(berlin)
	s := base.New(t)
	assert.FatalOnError(
		t,
		s.Store.CreateEvent(event.Event{Time: at, Process: "ensembled"}),
	)
	events, f := s.Service.Events(at.Add(-time.Hour), at.Add(time.Hour), 100)
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, len(events))
}
