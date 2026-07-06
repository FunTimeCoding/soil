//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestMoment(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Moment("session-1", r.Callsign, "figured out the root cause")
	events := s.Store.Events(event_query.New().Kind(constant.Moment).SetLimit(10))
	assert.Count(t, 1, events)
	assert.String(
		t,
		"figured out the root cause",
		events[0].Metadata[constant.Line],
	)
}
