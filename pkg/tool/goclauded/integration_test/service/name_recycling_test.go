//go:build local

package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestNameRecycledAfterRelease(t *testing.T) {
	s := service_tester.New(t)
	var callsigns []string

	for i := 0; ; i++ {
		r := s.Register(fmt.Sprintf("fill-%d", i))

		if r.Callsign == "" {
			break
		}

		callsigns = append(callsigns, r.Callsign)
	}

	released := callsigns[0]
	s.Release("fill-0", released)
	r := s.Register("after-release")
	assert.String(t, released, r.Callsign)
}

func TestReleasedSessionKeepsName(t *testing.T) {
	s := service_tester.New(t)
	r := s.Register("session-1")
	callsign := r.Callsign
	s.Release("session-1", callsign)
	e := s.Store.GetSession("session-1")
	assert.True(t, e != nil)
	assert.String(t, callsign, e.Name)
	assert.True(t, e.Callsign == nil)
}
