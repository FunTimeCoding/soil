package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gographd/unit/base"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestChatMessageToolAnswersSentMessage(t *testing.T) {
	s := base.New(t)
	s.Source.AddChat(chat.New("19:alpha@thread.v2", "group"))
	s.Source.AddChatMember(
		"19:alpha@thread.v2",
		member.New("Juliett", "user-9"),
	)
	assert.StringContains(
		t,
		`"text": "@Juliett on it"`,
		s.Tools().MustCallTool(
			constant.SendChatMessage,
			map[string]any{
				"chat":    "19:alpha@thread.v2",
				"text":    "@Juliett on it",
				"mention": []string{"Juliett"},
			},
		),
	)
	assert.Count(t, 1, s.Source.Mentioned)
}

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
