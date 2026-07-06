//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
	"time"
)

func TestWaitMessageNotListening(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	messages, e := s.Store.WaitMessage(r.Callsign, time.Millisecond)
	assert.FatalOnError(t, e)
	assert.True(t, messages == nil)
}

func TestWaitMessageWithPending(t *testing.T) {
	s := store_tester.New(t)
	r1 := s.EnsureSession("session-1")
	r2 := s.EnsureSession("session-2")
	s.SetListening(r1.Callsign, true)
	s.SendMessage(r2.Callsign, r1.Callsign, "wake up")
	messages, e := s.Store.WaitMessage(r1.Callsign, time.Second)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, messages)
	assert.String(t, "wake up", messages[0].Body)
}

func TestWaitMessageTimeout(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	s.SetListening(r.Callsign, true)
	messages, e := s.Store.WaitMessage(r.Callsign, 10*time.Millisecond)
	assert.FatalOnError(t, e)
	assert.True(t, messages == nil)
}
