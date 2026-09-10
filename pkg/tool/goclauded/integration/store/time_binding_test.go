package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
	"time"
)

func TestQueueWindowIgnoresTheBoundLocation(t *testing.T) {
	s := store_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.PushQueue("", "Nobody", constant.QueueTimeout, "now"),
	)
	east, e := time.LoadLocation("Europe/Berlin")
	assert.FatalOnError(t, e)
	now := time.Now()
	removed, f := s.Store.DeleteQueueWindow(
		"Nobody",
		now.Add(-time.Hour).In(east),
		now.Add(time.Hour).In(east),
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, removed)
}

func TestQueueWindowExcludesRowOutsideTheWindow(t *testing.T) {
	s := store_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.PushQueue("", "Nobody", constant.QueueTimeout, "now"),
	)
	utc, e := time.LoadLocation("UTC")
	assert.FatalOnError(t, e)
	now := time.Now()
	removed, f := s.Store.DeleteQueueWindow(
		"Nobody",
		now.Add(time.Hour).In(utc),
		now.Add(2*time.Hour).In(utc),
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 0, removed)
}
