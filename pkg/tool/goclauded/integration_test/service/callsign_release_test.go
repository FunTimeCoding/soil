//go:build local

package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"sort"
	"testing"
	"time"
)

func TestCallsignReleasedAfterThreeDays(t *testing.T) {
	s := service_tester.New(t)
	r := s.Register("session-1")
	assert.True(t, r.Callsign != "")
	callsign := r.Callsign
	s.Store.Advance(73 * time.Hour)
	s.Service.RunTimeoutSweep()
	e := s.Store.GetSession("session-1")
	assert.True(t, e.Callsign == nil)
	assert.String(t, callsign, e.Name)
}

func TestCallsignNotReleasedBeforeThreeDays(t *testing.T) {
	s := service_tester.New(t)
	r := s.Register("session-1")
	assert.True(t, r.Callsign != "")
	s.Store.Advance(71 * time.Hour)
	s.Service.RunTimeoutSweep()
	e := s.Store.GetSession("session-1")
	assert.True(t, e.Callsign != nil)
}

func TestCallsignRecycledAfterSweepRelease(t *testing.T) {
	s := service_tester.New(t)
	expected := []string{
		"Ash", "Blair", "Cedar", "Dale", "Ellis",
		"Frost", "Glen", "Harbor", "Jade", "Kent",
	}
	var callsigns []string

	for i := 0; ; i++ {
		r := s.Register(fmt.Sprintf("fill-%d", i))

		if r.Callsign == "" {
			break
		}

		callsigns = append(callsigns, r.Callsign)
	}

	sort.Strings(callsigns)
	assert.Integer(t, len(expected), len(callsigns))

	for i, v := range expected {
		assert.String(t, v, callsigns[i])
	}

	s.Store.Advance(73 * time.Hour)
	s.Service.RunTimeoutSweep()
	r := s.Register("after-sweep")
	assert.True(t, r.Callsign != "")
}
