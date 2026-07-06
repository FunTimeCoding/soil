//go:build local

package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestReleasePreservesSession(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	assert.FatalOnError(t, s.Store.ReleaseByCallsign(r.Callsign))
	e := s.GetSession("session-1")
	assert.True(t, e != nil)
	assert.String(t, r.Callsign, e.Name)
	assert.True(t, e.Callsign == nil)
}

func TestReleaseFreesCallsign(t *testing.T) {
	s := store_tester.New(t)
	var callsigns []string

	for i := 0; ; i++ {
		r := s.EnsureSession(fmt.Sprintf("fill-%d", i))

		if r.Callsign == "" {
			break
		}

		callsigns = append(callsigns, r.Callsign)
	}

	released := callsigns[0]
	assert.FatalOnError(t, s.Store.ReleaseByCallsign(released))
	r := s.EnsureSession("after-release")
	assert.String(t, released, r.Callsign)
}
