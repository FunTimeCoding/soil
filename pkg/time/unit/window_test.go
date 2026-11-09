package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/window"
	"testing"
	"time"
)

func TestWholeWindowHasNoEnds(t *testing.T) {
	w := window.Whole()
	_, started := w.Start()
	_, ended := w.End()
	assert.False(t, started)
	assert.False(t, ended)
	early := time.Date(1990, time.March, 4, 0, 0, 0, 0, time.UTC)
	assert.True(t, w.Contains(early))
}

func TestWindowAnswersItsEnds(t *testing.T) {
	w := window.New(
		time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2030, time.January, 31, 0, 0, 0, 0, time.UTC),
	)
	start, started := w.Start()
	end, ended := w.End()
	assert.True(t, started)
	assert.True(t, ended)
	assert.String(t, "2030-01-01T00:00:00Z", start.Format(time.RFC3339))
	assert.String(t, "2030-01-31T00:00:00Z", end.Format(time.RFC3339))
}

func TestWindowWithOneEndLeavesOtherOpen(t *testing.T) {
	w := window.Whole().WithEnd(
		time.Date(2030, time.January, 31, 0, 0, 0, 0, time.UTC),
	)
	_, started := w.Start()
	_, ended := w.End()
	assert.False(t, started)
	assert.True(t, ended)
}

func TestWindowContainsBothEnds(t *testing.T) {
	start := time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2030, time.January, 31, 0, 0, 0, 0, time.UTC)
	w := window.New(start, end)
	assert.True(t, w.Contains(start))
	assert.True(t, w.Contains(end))
	assert.False(t, w.Contains(start.Add(-time.Second)))
	assert.False(t, w.Contains(end.Add(time.Second)))
}

func TestBoundWindowEndsAtNowAndReachesBackSpan(t *testing.T) {
	w := window.Whole().Bound(
		time.Date(2030, time.January, 31, 12, 0, 0, 0, time.UTC),
		time.Hour,
	)
	start, _ := w.Start()
	end, _ := w.End()
	assert.String(t, "2030-01-31T11:00:00Z", start.Format(time.RFC3339))
	assert.String(t, "2030-01-31T12:00:00Z", end.Format(time.RFC3339))
}

func TestBoundWindowReachesBackFromGivenEnd(t *testing.T) {
	w := window.Whole().WithEnd(
		time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC),
	).Bound(time.Date(2030, time.January, 31, 12, 0, 0, 0, time.UTC), time.Hour)
	start, _ := w.Start()
	end, _ := w.End()
	assert.String(t, "2029-12-31T23:00:00Z", start.Format(time.RFC3339))
	assert.String(t, "2030-01-01T00:00:00Z", end.Format(time.RFC3339))
}

func TestBoundWindowKeepsGivenStart(t *testing.T) {
	w := window.Whole().WithStart(
		time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC),
	).Bound(time.Date(2030, time.January, 31, 12, 0, 0, 0, time.UTC), time.Hour)
	start, _ := w.Start()
	end, _ := w.End()
	assert.String(t, "2030-01-01T00:00:00Z", start.Format(time.RFC3339))
	assert.String(t, "2030-01-31T12:00:00Z", end.Format(time.RFC3339))
}

func TestWindowOpenStartContainsEverythingBeforeEnd(t *testing.T) {
	w := window.Whole().WithEnd(
		time.Date(2030, time.January, 31, 0, 0, 0, 0, time.UTC),
	)
	early := time.Date(1990, time.March, 4, 0, 0, 0, 0, time.UTC)
	late := time.Date(2030, time.February, 1, 0, 0, 0, 0, time.UTC)
	assert.True(t, w.Contains(early))
	assert.False(t, w.Contains(late))
}
