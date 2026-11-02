package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/service_tester"
	"testing"
)

func TestTheirQueryRunningTooLongIsQuiet(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	clients["alfa"].Failure = timeout.Format(
		"canceling statement due to statement timeout",
	)
	_, e := s.Query(t.Context(), "alfa", "SELECT pg_sleep(60)")
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "query ran too long", e.Error())
}

func TestOurQueryRunningTooLongIsCaptured(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	clients["alfa"].Failure = timeout.Format(
		"canceling statement due to statement timeout",
	)
	_, e := s.TableSizes(t.Context(), "alfa", "")
	assert.True(t, timeout.Is(e))
	assert.True(t, classify.Reportable(e))
}

func TestAnUnknownInstanceStaysNotFoundOnACatalogue(t *testing.T) {
	s, _ := service_tester.New(t, "alfa")
	_, e := s.ListTables(t.Context(), "zulu", "")
	assert.True(t, not_found.Is(e))
}

func TestACallerWhoLeftIsNotCaptured(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	clients["alfa"].Failure = errors.Format(
		"query: %s",
		"connection closed mid-query",
	)
	x, cancel := context.WithCancel(t.Context())
	cancel()
	_, e := s.Query(x, "alfa", "SELECT 1")
	assert.False(t, classify.Reportable(e))
	_, f := s.ListSchemas(x, "alfa")
	assert.False(t, classify.Reportable(f))
}
