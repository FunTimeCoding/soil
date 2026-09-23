package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/service_tester"
	"strings"
	"testing"
)

func TestPushRejectsAQuestionOverTheLimit(t *testing.T) {
	s := service_tester.New(t)
	_, e := s.Service.PushDecision(
		stringConstant.LowerAlfa,
		strings.Repeat("x", constant.MaximumTurnLength+1),
		"narrow",
		nil,
		nil,
	)
	assert.Error(t, e)
}

func TestPushAcceptsAQuestionAtTheLimit(t *testing.T) {
	s := service_tester.New(t)
	d, e := s.Service.PushDecision(
		stringConstant.LowerAlfa,
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
		stringConstant.LowerAlfa,
		"which entry rule?",
		"",
		nil,
		nil,
	)
	assert.Error(t, e)
}

func TestRejectedPushIsNotStored(t *testing.T) {
	s := service_tester.New(t)
	_, e := s.Service.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"",
		nil,
		nil,
	)
	assert.Error(t, e)
	assert.Count(t, 0, s.Service.Decisions(stringConstant.LowerAlfa))
}

func TestTurnRejectsContentOverTheLimit(t *testing.T) {
	s := service_tester.New(t)
	d, f := s.Service.PushDecision(
		stringConstant.LowerAlfa,
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
	assert.Error(t, e)
}

func TestSessionTurnsStopAtTheLimit(t *testing.T) {
	s := service_tester.New(t)
	d, f := s.Service.PushDecision(
		stringConstant.LowerAlfa,
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
	assert.Error(t, e)
}

func TestUserTurnsAreNotCapped(t *testing.T) {
	s := service_tester.New(t)
	d, f := s.Service.PushDecision(
		stringConstant.LowerAlfa,
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

	assert.Count(t, 5, s.Service.Decisions(stringConstant.LowerAlfa)[0].Turns)
}
