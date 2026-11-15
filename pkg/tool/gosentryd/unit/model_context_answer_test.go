package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestUnknownIssueIsAnsweredQuietly(t *testing.T) {
	m := memory.New()
	assert.String(
		t,
		"issue not found: 100",
		daemon_tester.ServedTools(
			t,
			service_tester.NewSource(t),
			m,
		).MustCallToolError(
			constant.GetIssue,
			map[string]any{"identifier": "100"},
		),
	)
	assert.Count(t, 0, m.Events())
}

func TestForbiddenSearchIsAnsweredQuietly(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = forbidden.Format("sentry: %s", "token lacks event:read")
	m := memory.New()
	daemon_tester.ServedTools(t, c, m).MustCallToolError(
		constant.SearchIssues,
		map[string]any{},
	)
	assert.Count(t, 0, m.Events())
}

func TestUnclassifiedFailureIsCaptured(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = errors.New("sentry answered nonsense")
	m := memory.New()
	daemon_tester.ServedTools(t, c, m).MustCallToolError(
		constant.Whoami,
		map[string]any{},
	)
	assert.Count(t, 1, m.Events())
}

func TestOutageIsAnsweredWithoutCapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = outage_tester.RefusedDial()
	m := memory.New()
	s := daemon_tester.ServedTools(t, c, m)

	for range 2 {
		s.MustCallToolError(constant.Whoami, map[string]any{})
	}

	assert.Count(t, 0, m.Events())
}
