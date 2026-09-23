package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
)

func TestPushedDecisionStartsOpen(t *testing.T) {
	s := store_tester.New(t)
	s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"take the narrow one",
		nil,
		nil,
	)
	assert.String(
		t,
		"open",
		string(s.Store.Decisions(stringConstant.LowerAlfa)[0].State),
	)
}

func TestPushedDecisionKeepsItsDefaultAction(t *testing.T) {
	s := store_tester.New(t)
	s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"take the narrow one",
		nil,
		nil,
	)
	assert.String(
		t,
		"take the narrow one",
		s.Store.Decisions(stringConstant.LowerAlfa)[0].DefaultAction,
	)
}

func TestChoicesKeepPushOrder(t *testing.T) {
	s := store_tester.New(t)
	s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"take the narrow one",
		[]string{"narrow", "wide", "unfiltered"},
		nil,
	)
	choices := s.Store.Decisions(stringConstant.LowerAlfa)[0].Choices
	assert.Count(t, 3, choices)
	assert.String(t, "narrow", choices[0].Label)
	assert.String(t, "unfiltered", choices[2].Label)
}

func TestDecisionsAreScopedToTheirSession(t *testing.T) {
	s := store_tester.New(t)
	s.Store.PushDecision(stringConstant.LowerAlfa, "mine?", "proceed", nil, nil)
	s.Store.PushDecision(
		stringConstant.LowerBravo,
		"theirs?",
		"proceed",
		nil,
		nil,
	)
	assert.Count(t, 1, s.Store.Decisions(stringConstant.LowerAlfa))
}

func TestAnswerMarksTheDecisionAnswered(t *testing.T) {
	s := store_tester.New(t)
	d := s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	s.Store.Answer(
		d.Identifier,
		constant.AnswerKindChoice,
		constant.AnswerChannelQueue,
		"narrow",
	)
	assert.String(
		t,
		"answered",
		string(s.Store.Decisions(stringConstant.LowerAlfa)[0].State),
	)
}

func TestAnswerRecordsKindAndValue(t *testing.T) {
	s := store_tester.New(t)
	d := s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	s.Store.Answer(
		d.Identifier,
		constant.AnswerKindConstraint,
		constant.AnswerChannelQueue,
		"you rule, stay cheap",
	)
	answered := s.Store.Decisions(stringConstant.LowerAlfa)[0]
	assert.String(t, "constraint", string(answered.AnswerKind))
	assert.String(t, "you rule, stay cheap", answered.Answer)
}

func TestDismissKeepsTheGivenState(t *testing.T) {
	s := store_tester.New(t)
	d := s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	s.Store.Dismiss(d.Identifier, constant.StateIrrelevant)
	assert.String(
		t,
		"irrelevant",
		string(s.Store.Decisions(stringConstant.LowerAlfa)[0].State),
	)
}

func TestBumpAccumulates(t *testing.T) {
	s := store_tester.New(t)
	d := s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	s.Store.Bump(d.Identifier)
	s.Store.Bump(d.Identifier)
	assert.Integer(t, 2, s.Store.Decisions(stringConstant.LowerAlfa)[0].Bumped)
}

func TestClearRecordsTheUnderstanding(t *testing.T) {
	s := store_tester.New(t)
	d := s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	s.Store.Clear(
		d.Identifier,
		"narrow, unfiltered entry",
		constant.ResolutionDefault,
	)
	cleared := s.Store.Decisions(stringConstant.LowerAlfa)[0]
	assert.String(t, "narrow, unfiltered entry", cleared.ClearLine)
	assert.String(t, "default", string(cleared.Resolution))
}
