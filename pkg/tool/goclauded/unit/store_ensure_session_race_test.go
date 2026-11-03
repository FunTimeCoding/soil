package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"sync"
	"testing"
)

func TestConcurrentEnsuresShareOneSessionAndPickDistinctCallsigns(t *testing.T) {
	for range 50 {
		s := store_tester.New(t)
		shared := make([]string, 8)
		distinct := make([]string, 8)
		var g sync.WaitGroup

		for i := range 8 {
			g.Go(
				func() {
					r, e := s.Store.EnsureSession("shared")
					assert.FatalOnError(t, e)
					shared[i] = r.Callsign
				},
			)
			g.Go(
				func() {
					r, e := s.Store.EnsureSession(fmt.Sprintf("distinct-%d", i))
					assert.FatalOnError(t, e)
					distinct[i] = r.Callsign
				},
			)
		}

		g.Wait()
		assert.Count(t, 9, s.ListSessions())
		seen := map[string]bool{shared[0]: true}

		for _, callsign := range shared {
			assert.String(t, shared[0], callsign)
		}

		for _, callsign := range distinct {
			assert.False(t, seen[callsign])
			seen[callsign] = true
		}
	}
}

func TestDiscoveryRacingEnsuresCreatesOneSession(t *testing.T) {
	for range 50 {
		s := store_tester.New(t)
		var g sync.WaitGroup

		for range 8 {
			g.Go(
				func() {
					_, e := s.Store.EnsureSession("shared")
					assert.FatalOnError(t, e)
				},
			)
			g.Go(func() { s.Store.CreateDiscovered("shared") })
		}

		g.Wait()
		all, e := s.Store.AllSessions(0, 0)
		assert.FatalOnError(t, e)
		assert.Count(t, 1, all)
	}
}
