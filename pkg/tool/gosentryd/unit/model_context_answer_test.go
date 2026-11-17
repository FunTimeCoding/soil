package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestUnknownIssueIsAnsweredQuietly(t *testing.T) {
	s := base.New(t)
	assert.String(
		t,
		"issue not found: 100",
		s.Tools().MustCallToolError(
			constant.GetIssue,
			map[string]any{"identifier": "100"},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestForbiddenSearchIsAnsweredQuietly(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = forbidden.Format("sentry: %s", "token lacks event:read")
	s.Tools().MustCallToolError(constant.SearchIssues, map[string]any{})
	assert.Count(t, 0, s.Reporter.Events())
}

func TestUnclassifiedFailureIsCaptured(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = errors.New("sentry answered nonsense")
	s.Tools().MustCallToolError(constant.Whoami, map[string]any{})
	assert.Count(t, 1, s.Reporter.Events())
}

func TestOutageIsAnsweredWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	c := s.Tools()

	for range 2 {
		c.MustCallToolError(constant.Whoami, map[string]any{})
	}

	assert.Count(t, 0, s.Reporter.Events())
}
