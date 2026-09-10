package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func TestQueueIsNotDeliveredToTheNextHolderOfTheName(t *testing.T) {
	s := store_tester.New(t)
	addSession(t, s, "first", "Frost", nil, "2026-09-01 10:00:00+00:00")
	assert.FatalOnError(
		t,
		s.Store.PushQueue(
			"first",
			"Frost",
			constant.QueueMessage,
			"for the first Frost",
		),
	)
	addSession(t, s, "second", "Frost", "Frost", "2026-09-08 10:00:00+00:00")
	inherited, e := s.Store.DrainQueue("second", "Frost")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, inherited)
	own, f := s.Store.DrainQueue("first", "Frost")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, own)
	assert.String(t, "for the first Frost", own[0].Body)
}

func TestNotificationIsNotDeliveredToTheNextHolderOfTheName(t *testing.T) {
	s := store_tester.New(t)
	addSession(t, s, "first", "Frost", nil, "2026-09-01 10:00:00+00:00")
	assert.FatalOnError(
		t,
		s.Store.SendNotification("first", "Frost", "Dale", "for the first"),
	)
	addSession(t, s, "second", "Frost", "Frost", "2026-09-08 10:00:00+00:00")
	inherited, e := s.Store.PendingNotifications("second")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, inherited)
	own, f := s.Store.PendingNotifications("first")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, own)
}

func TestQueueStillCarriesTheCallsignForOlderReaders(t *testing.T) {
	s := store_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.PushQueue("holder", "Wren", constant.QueueMessage, "a body"),
	)
	var stored string
	assert.FatalOnError(
		t,
		s.Store.Database().Raw(
			"SELECT callsign FROM queue WHERE body = ?",
			"a body",
		).Scan(&stored).Error,
	)
	assert.String(t, "Wren", stored)
}
