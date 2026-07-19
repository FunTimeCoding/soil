package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
	"time"
)

func TestAnnounceClearsPendingTimeout(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Announce("session-1", r.Callsign, "working", "")
	s.Check("session-1")
	s.Store.Advance(2 * time.Hour)
	s.Service.RunTimeoutSweep()
	s.Announce("session-1", r.Callsign, "back again", "")
	r = s.Check("session-1")
	timeouts := entriesByKind(r.Entries, constant.QueueTimeout)
	assert.Count(t, 0, timeouts)
}
