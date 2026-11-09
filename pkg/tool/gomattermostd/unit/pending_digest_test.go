package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/digest/event"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/service/pending_digest"
	"testing"
	"time"
)

func TestPendingRetainDropsSubscriptionsThatAreGone(t *testing.T) {
	p := pending_digest.New(0, time.Minute)
	now := time.Now()
	e := event.New(constant.MessageEvent, "Foxtrot", "first", now)
	p.Keep(1, "alfa", []*event.Event{e}, now)
	p.Keep(2, "alfa", []*event.Event{e}, now)
	p.Keep(3, "bravo", []*event.Event{e}, now)
	p.Retain("alfa", []uint{1})
	assert.Count(t, 1, p.Take(1))
	assert.Count(t, 0, p.Take(2))
	assert.Count(t, 1, p.Take(3))
	assert.Count(t, 0, p.Roots(now))
}

func TestPendingDigestRetriesWithDoublingDelay(t *testing.T) {
	p := pending_digest.New(30*time.Second, 10*time.Minute)
	start := time.Date(2026, 11, 4, 12, 0, 0, 0, time.UTC)
	e := event.New(constant.MessageEvent, "Foxtrot", "first", start)
	p.Keep(1, "alfa", []*event.Event{e}, start)
	assert.Count(t, 0, p.Roots(start.Add(29*time.Second)))
	assert.Count(t, 1, p.Roots(start.Add(30*time.Second)))
	retry := start.Add(30 * time.Second)
	p.Keep(1, "alfa", p.Take(1), retry)
	assert.Count(t, 0, p.Roots(retry.Add(59*time.Second)))
	assert.Count(t, 1, p.Roots(retry.Add(time.Minute)))
}

func TestPendingDigestDelayStopsAtLimit(t *testing.T) {
	p := pending_digest.New(30*time.Second, 10*time.Minute)
	now := time.Date(2026, 11, 4, 12, 0, 0, 0, time.UTC)
	e := event.New(constant.MessageEvent, "Foxtrot", "first", now)

	for range 12 {
		p.Keep(1, "alfa", []*event.Event{e}, now)
	}

	assert.Count(t, 0, p.Roots(now.Add(10*time.Minute-time.Second)))
	assert.Count(t, 1, p.Roots(now.Add(10*time.Minute)))
}

func TestSuccessfulRetryResetsDelay(t *testing.T) {
	p := pending_digest.New(30*time.Second, 10*time.Minute)
	now := time.Date(2026, 11, 4, 12, 0, 0, 0, time.UTC)
	e := event.New(constant.MessageEvent, "Foxtrot", "first", now)
	p.Keep(1, "alfa", []*event.Event{e}, now)
	p.Keep(1, "alfa", p.Take(1), now)
	p.Take(1)
	p.Done(1)
	p.Keep(1, "alfa", []*event.Event{e}, now)
	assert.Count(t, 1, p.Roots(now.Add(30*time.Second)))
}
