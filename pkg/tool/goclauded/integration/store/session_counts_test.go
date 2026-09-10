package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func TestCountSessionEvents(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	s.EnsureSession("session-2")
	assert.FatalOnError(
		t,
		s.Store.LogEvent(
			"session-1",
			constant.Announce,
			"Ash",
			map[string]string{constant.Topic: "first"},
		),
	)
	assert.FatalOnError(
		t,
		s.Store.LogEvent(
			"session-1",
			constant.Complete,
			"Ash",
			map[string]string{},
		),
	)
	assert.FatalOnError(
		t,
		s.Store.LogEvent(
			"session-2",
			constant.Announce,
			"Blair",
			map[string]string{constant.Topic: "other"},
		),
	)
	count, e := s.Store.CountSessionEvents("session-1")
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, count)
	metadata, f := s.Store.CountSessionEventMetadata("session-1")
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, metadata)
	other, g := s.Store.CountSessionEvents("session-2")
	assert.FatalOnError(t, g)
	assert.Integer(t, 1, other)
	absent, h := s.Store.CountSessionEvents("session-3")
	assert.FatalOnError(t, h)
	assert.Integer(t, 0, absent)
}
