//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestSummarize(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("session-1")
	e := s.Service.Summarize(
		"session-1",
		"Cedar",
		"Fixed the auth race condition",
	)
	assert.FatalOnError(t, e)
	body := s.Store.SummaryBySession("session-1")
	assert.String(t, "Fixed the auth race condition", body)
	events := s.Store.Events(event_query.New().Kind(constant.Summarize).SetLimit(10))
	assert.Count(t, 1, events)
}

func TestSummarizeReplace(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("session-1")
	assert.FatalOnError(
		t,
		s.Service.Summarize("session-1", "Cedar", "first summary"),
	)
	assert.FatalOnError(
		t,
		s.Service.Summarize("session-1", "Cedar", "revised summary"),
	)
	body := s.Store.SummaryBySession("session-1")
	assert.String(t, "revised summary", body)
}
