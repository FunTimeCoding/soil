//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestSetListening(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	assert.FatalOnError(t, s.Store.SetListening(r.Callsign, true))
	e := s.GetSession("session-1")
	assert.True(t, e.Listening)
	assert.FatalOnError(t, s.Store.SetListening(r.Callsign, false))
	e = s.GetSession("session-1")
	assert.False(t, e.Listening)
}
