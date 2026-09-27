package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func TestEventsAfterWalksForward(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")

	for _, topic := range []string{"first", "second", "third"} {
		assert.FatalOnError(
			t,
			s.Store.LogEvent(
				"session-1",
				constant.Announce,
				"Ash",
				map[string]string{constant.Topic: topic},
			),
		)
	}

	events, e := s.Store.EventsAfter(0, nil, 10)
	assert.FatalOnError(t, e)
	assert.Count(t, 3, events)
	assert.String(t, "first", events[0].Metadata[constant.Topic])
	assert.String(t, "third", events[2].Metadata[constant.Topic])
}

func TestEventsAfterRespectsCursor(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")

	for _, topic := range []string{"first", "second", "third"} {
		assert.FatalOnError(
			t,
			s.Store.LogEvent(
				"session-1",
				constant.Announce,
				"Ash",
				map[string]string{constant.Topic: topic},
			),
		)
	}

	all, e := s.Store.EventsAfter(0, nil, 10)
	assert.FatalOnError(t, e)
	events, e := s.Store.EventsAfter(all[0].Identifier, nil, 10)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, events)
	assert.String(t, "second", events[0].Metadata[constant.Topic])
}

func TestEventsAfterFiltersKinds(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	assert.FatalOnError(
		t,
		s.Store.LogEvent(
			"session-1",
			constant.Announce,
			"Ash",
			map[string]string{constant.Topic: "announced"},
		),
	)
	assert.FatalOnError(
		t,
		s.Store.LogEvent(
			"session-1",
			constant.Label,
			"Ash",
			map[string]string{constant.Key: "character"},
		),
	)
	events, e := s.Store.EventsAfter(0, []string{constant.Label}, 10)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, events)
	assert.String(t, "character", events[0].Metadata[constant.Key])
}

func TestEventsAfterHonoursLimit(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")

	for _, topic := range []string{"first", "second", "third"} {
		assert.FatalOnError(
			t,
			s.Store.LogEvent(
				"session-1",
				constant.Announce,
				"Ash",
				map[string]string{constant.Topic: topic},
			),
		)
	}

	events, e := s.Store.EventsAfter(0, nil, 2)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, events)
	assert.String(t, "first", events[0].Metadata[constant.Topic])
}

func TestEventsAfterEmptyBeyondEnd(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	assert.FatalOnError(
		t,
		s.Store.LogEvent(
			"session-1",
			constant.Announce,
			"Ash",
			map[string]string{constant.Topic: "only"},
		),
	)
	events, e := s.Store.EventsAfter(9000, nil, 10)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, events)
}
