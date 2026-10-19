package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/service_tester"
	"strings"
	"testing"
)

func TestPushRejectsAQuestionOverTheLimit(t *testing.T) {
	s := service_tester.New(t)
	_, e := s.Service.PushDecision(
		stringsConstant.LowerAlfa,
		strings.Repeat("x", constant.MaximumTurnLength+1),
		"narrow",
		nil,
		nil,
	)
	assert.True(t, validation.Is(e))
}

func TestPushAcceptsAQuestionAtTheLimit(t *testing.T) {
	s := service_tester.New(t)
	d, e := s.Service.PushDecision(
		stringsConstant.LowerAlfa,
		strings.Repeat("x", constant.MaximumTurnLength),
		"narrow",
		nil,
		nil,
	)
	assert.Nil(t, e)
	assert.NotNil(t, d)
}

func TestPushRejectsAMissingDefaultAction(t *testing.T) {
	s := service_tester.New(t)
	_, e := s.Service.PushDecision(
		stringsConstant.LowerAlfa,
		"which entry rule?",
		"",
		nil,
		nil,
	)
	assert.True(t, validation.Is(e))
}

func TestRejectedPushIsNotStored(t *testing.T) {
	s := service_tester.New(t)
	_, e := s.Service.PushDecision(
		stringsConstant.LowerAlfa,
		"which entry rule?",
		"",
		nil,
		nil,
	)
	assert.Error(t, e)
	assert.Count(t, 0, s.Service.Decisions(stringsConstant.LowerAlfa))
}

func TestTurnRejectsContentOverTheLimit(t *testing.T) {
	s := service_tester.New(t)
	d, f := s.Service.PushDecision(
		stringsConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	assert.FatalOnError(t, f)
	_, e := s.Service.AddTurn(
		d.Identifier,
		constant.AuthorSession,
		strings.Repeat("x", constant.MaximumTurnLength+1),
	)
	assert.True(t, validation.Is(e))
}

func TestSessionTurnsStopAtTheLimit(t *testing.T) {
	s := service_tester.New(t)
	d, f := s.Service.PushDecision(
		stringsConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	assert.FatalOnError(t, f)

	for range constant.MaximumTurnCount {
		_, e := s.Service.AddTurn(
			d.Identifier,
			constant.AuthorSession,
			"narrowing",
		)
		assert.Nil(t, e)
	}

	_, e := s.Service.AddTurn(d.Identifier, constant.AuthorSession, "one more")
	assert.True(t, conflict.Is(e))
}

func TestUserTurnsAreNotCapped(t *testing.T) {
	s := service_tester.New(t)
	d, f := s.Service.PushDecision(
		stringsConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	assert.FatalOnError(t, f)

	for range constant.MaximumTurnCount + 2 {
		_, e := s.Service.AddTurn(d.Identifier, constant.AuthorUser, "go on")
		assert.Nil(t, e)
	}

	assert.Count(t, 5, s.Service.Decisions(stringsConstant.LowerAlfa)[0].Turns)
}
