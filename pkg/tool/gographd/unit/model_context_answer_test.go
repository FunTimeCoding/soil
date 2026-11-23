package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gographd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestUnknownTeamAnsweredQuietly(t *testing.T) {
	s := base.New(t)
	assert.String(
		t,
		"team not found: Finance",
		s.Tools().MustCallToolError(
			constant.ListChannels,
			map[string]any{"team": "Finance"},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestMissingSignInAnsweredAndCaptured(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = not_configured.Format(microsoftConstant.SignInMissing)
	assert.StringContains(
		t,
		"graph sign-in required: run gograph login",
		s.Tools().MustCallToolError(constant.ListTeams, map[string]any{}),
	)
	assert.Count(t, 1, s.Reporter.Events())
}

func TestUnclassifiedFailureCaptured(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = errors.New("graph answered nonsense")
	s.Tools().MustCallToolError(constant.MyChats, map[string]any{})
	assert.Count(t, 1, s.Reporter.Events())
}

func TestOutageAnsweredWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	c := s.Tools()

	for range 2 {
		c.MustCallToolError(constant.ListTeams, map[string]any{})
	}

	assert.Count(t, 0, s.Reporter.Events())
}
