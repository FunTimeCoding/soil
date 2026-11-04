package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/trace_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
)

func TestAFailureAnsweringACallerCarriesTheCallersTrace(t *testing.T) {
	s := base.New(t)
	c := s.MemoryClient(t)
	s.Execute("ALTER TABLE memory RENAME TO memory_gone")
	_, e := c.GetStatistics(trace_tester.Context(t, nil))
	assert.Error(t, e)
	assert.Count(t, 1, s.Reporter.Events())
	assert.StringContains(
		t,
		"0123456789abcdef0123456789abcdef",
		s.Reporter.Events()[0].Trace,
	)
}
