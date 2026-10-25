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
	events, f := s.EventsByTimeRange(at.Add(-time.Hour), at.Add(time.Hour), 100)
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, len(events))
}

func TestStoreNormalizesStoredLocalTimesToUTC(t *testing.T) {
	d := lite.NewMemory()
	store.New(d)
	assert.FatalOnError(
		t,
		d.Exec(
			"INSERT INTO event (time, process) VALUES (?, ?), (?, ?)",
			"2026-07-01 12:00:00.123456+02:00",
			"rapportd",
			"2026-10-25 02:20:59+01:00",
			"bluetoothd",
		).Error,
	)
	store.New(d)
	var stored []string
	assert.FatalOnError(
		t,
		d.Raw("SELECT CAST(time AS TEXT) FROM event ORDER BY time").
			Scan(&stored).Error,
	)
	assert.Strings(
		t,
		[]string{
			"2026-07-01 10:00:00.123456+00:00",
			"2026-10-25 01:20:59+00:00",
		},
		stored,
	)
}
