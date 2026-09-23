package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
)

func TestPushedDecisionCarriesItsFrames(t *testing.T) {
	s := store_tester.New(t)
	s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		[]string{"handoff", "counting"},
	)
	assert.Count(t, 2, s.Store.Decisions(stringConstant.LowerAlfa)[0].Frames)
}

func TestTwoDecisionsShareOneFrameRow(t *testing.T) {
	s := store_tester.New(t)
	s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"first?",
		"narrow",
		nil,
		[]string{"handoff"},
	)
	s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"second?",
		"narrow",
		nil,
		[]string{"handoff"},
	)
	assert.Count(t, 1, s.Store.Frames(stringConstant.LowerAlfa))
}

func TestFramesAreScopedToTheirSession(t *testing.T) {
	s := store_tester.New(t)
	s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"mine?",
		"narrow",
		nil,
		[]string{"handoff"},
	)
	s.Store.PushDecision(
		stringConstant.LowerBravo,
		"theirs?",
		"narrow",
		nil,
		[]string{"handoff"},
	)
	assert.Count(t, 1, s.Store.Frames(stringConstant.LowerAlfa))
}

func TestOpenDecisionsExcludeAnswered(t *testing.T) {
	s := store_tester.New(t)
	first := s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"first?",
		"narrow",
		nil,
		[]string{"handoff"},
	)
	s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"second?",
		"narrow",
		nil,
		[]string{"handoff"},
	)
	frames := s.Store.Frames(stringConstant.LowerAlfa)
	assert.Count(t, 2, s.Store.OpenDecisions(frames[0].Identifier))
	s.Store.Answer(
		first.Identifier,
		constant.AnswerKindOther,
		constant.AnswerChannelQueue,
		"neither",
	)
	assert.Count(t, 1, s.Store.OpenDecisions(frames[0].Identifier))
}

func TestOpenDecisionsExcludeClearedByDefault(t *testing.T) {
	s := store_tester.New(t)
	d := s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"first?",
		"narrow",
		nil,
		[]string{"handoff"},
	)
	frames := s.Store.Frames(stringConstant.LowerAlfa)
	s.Store.Clear(d.Identifier, "went narrow", constant.ResolutionDefault)
	assert.Count(t, 0, s.Store.OpenDecisions(frames[0].Identifier))
}
