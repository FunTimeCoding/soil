package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/service_tester"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestUnknownIssueIsAnsweredQuietly(t *testing.T) {
	m := memory.New()
	assert.String(
		t,
		"issue not found: 100",
		servedTools(t, service_tester.NewSource(t), m).MustCallToolError(
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
	servedTools(t, c, m).MustCallToolError(
		constant.SearchIssues,
		map[string]any{},
	)
	assert.Count(t, 0, m.Events())
}

func TestUnclassifiedFailureIsCaptured(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = errors.New("sentry answered nonsense")
	m := memory.New()
	servedTools(t, c, m).MustCallToolError(constant.Whoami, map[string]any{})
	assert.Count(t, 1, m.Events())
}

func TestOutageIsAnsweredWithoutCapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}
	m := memory.New()
	s := servedTools(t, c, m)

	for range 2 {
		s.MustCallToolError(constant.Whoami, map[string]any{})
	}

	assert.Count(t, 0, m.Events())
}
