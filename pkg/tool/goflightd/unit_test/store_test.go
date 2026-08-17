package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"testing"
	"time"
)

func TestStoreRoundTrip(t *testing.T) {
	s := store.New(lite.NewMemory())
	defer s.Close()
	now := time.Now()
	s.MustCreateEvent(
		event.Event{
			Time:      now,
			Process:   "ensembled",
			Subsystem: "com.apple.ensemble",
			Message:   "session started",
		},
	)
	s.MustCreateEvent(
		event.Event{
			Time:    now.Add(-2 * time.Hour),
			Process: "rapportd",
			Message: "old event",
		},
	)
	s.MustCreateSnapshot(
		snapshot.Snapshot{
			Time:  now,
			Kind:  "wireless",
			Key:   "awdl.Channel Sequence",
			Value: "{(44, 80MHz)}",
		},
	)
	events, e := s.EventsByTimeRange(
		now.Add(-1*time.Hour),
		now.Add(time.Hour),
		100,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, len(events))
	assert.String(t, "ensembled", events[0].Process)
	snapshots, f := s.SnapshotsByTimeRange(
		now.Add(-1*time.Hour),
		now.Add(time.Hour),
		100,
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, len(snapshots))
	assert.String(t, "wireless", snapshots[0].Kind)
	last, g := s.LastEventTime()
	assert.FatalOnError(t, g)
	assert.NotNil(t, last)
}

func TestStoreMarks(t *testing.T) {
	s := store.New(lite.NewMemory())
	defer s.Close()
	first := mark.New(time.Now().Add(-time.Minute), "broke", "")
	assert.FatalOnError(t, s.CreateMark(first))
	second := mark.New(time.Now(), "recovered", "mouse cycled")
	assert.FatalOnError(t, s.CreateMark(second))
	marks, e := s.RecentMarks(10)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, len(marks))
	assert.String(t, "recovered", marks[0].Label)
	count, f := s.MarkCount()
	assert.FatalOnError(t, f)
	assert.Integer(t, 2, count)
}
