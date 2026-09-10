package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func TestMarkClosed(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	assert.True(t, s.GetSession("session-1").ClosedAt == nil)
	assert.FatalOnError(t, s.Store.MarkClosed("session-1", "prompt_input_exit"))
	r := s.GetSession("session-1")
	assert.True(t, r.ClosedAt != nil)
	assert.String(t, "prompt_input_exit", r.ClosedReason)
}

func TestActivityClearsClosedMark(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	assert.FatalOnError(t, s.Store.MarkClosed("session-1", "logout"))
	assert.True(t, s.GetSession("session-1").ClosedAt != nil)
	s.EnsureSession("session-1")
	r := s.GetSession("session-1")
	assert.True(t, r.ClosedAt == nil)
	assert.String(t, "", r.ClosedReason)
}
