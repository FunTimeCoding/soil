package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/service_tester"
	"testing"
)

func TestAQuietRefusalIsNotCapturedByTheGuard(t *testing.T) {
	m := memory.New()
	v := service_tester.Shared(t, service_tester.NewSource(t), m)
	_, e := v.Issue("100")
	assert.True(t, not_found.Is(e))
	f := v.DeleteIssue("200")
	assert.True(t, not_found.Is(f))
	assert.Count(t, 0, m.Events())
}
