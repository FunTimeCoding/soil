package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
	"time"
)

func TestDeleteQueueWindow(t *testing.T) {
	s := store_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.PushQueue("held", "Frost", constant.QueueTimeout, "first"),
	)
	assert.FatalOnError(
		t,
		s.Store.PushQueue(
			"held",
			"Frost",
			constant.QueueSessionAnnounce,
			"second",
		),
	)
	assert.FatalOnError(
		t,
		s.Store.PushQueue(
			"other",
			"Nora",
			constant.QueueTimeout,
			"other holder",
		),
	)
	before := time.Now().Add(-time.Hour)
	after := time.Now().Add(time.Hour)
	early, e := s.Store.DeleteQueueWindow(
		"Frost",
		before,
		before.Add(time.Minute),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, early)
	removed, f := s.Store.DeleteQueueWindow("Frost", before, after)
	assert.FatalOnError(t, f)
	assert.Integer(t, 2, removed)
	survivor, g := s.Store.DeleteQueueWindow("Nora", before, after)
	assert.FatalOnError(t, g)
	assert.Integer(t, 1, survivor)
}

func TestDeleteNotificationWindow(t *testing.T) {
	s := store_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.SendNotification("held", "Frost", "Dale", "settled"),
	)
	assert.FatalOnError(
		t,
		s.Store.SendNotification("other", "Nora", "Dale", "other holder"),
	)
	before := time.Now().Add(-time.Hour)
	after := time.Now().Add(time.Hour)
	early, e := s.Store.DeleteNotificationWindow(
		"Frost",
		before,
		before.Add(time.Minute),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, early)
	removed, f := s.Store.DeleteNotificationWindow("Frost", before, after)
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, removed)
	survivor, g := s.Store.DeleteNotificationWindow("Nora", before, after)
	assert.FatalOnError(t, g)
	assert.Integer(t, 1, survivor)
}
