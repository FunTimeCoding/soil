package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
)

func TestAnsweredDecisionsStartUnpulsed(t *testing.T) {
	s := store_tester.New(t)
	s.Answered("first?")
	s.Answered("second?")
	assert.Count(t, 2, s.Store.UnpulsedAnswers(constant.LowerAlfa))
}

func TestPulsingABatchEmptiesIt(t *testing.T) {
	s := store_tester.New(t)
	s.Answered("first?")
	s.Answered("second?")
	s.Store.MarkPulsed(constant.LowerAlfa)
	assert.Count(t, 0, s.Store.UnpulsedAnswers(constant.LowerAlfa))
}

func TestANewAnswerAfterPulsingTriggersAgain(t *testing.T) {
	s := store_tester.New(t)
	s.Answered("first?")
	s.Store.MarkPulsed(constant.LowerAlfa)
	s.Answered("second?")
	assert.Count(t, 1, s.Store.UnpulsedAnswers(constant.LowerAlfa))
}

func TestPollingAlsoStopsTheTrigger(t *testing.T) {
	s := store_tester.New(t)
	s.Answered("first?")
	s.Store.MarkSeen(constant.LowerAlfa)
	assert.Count(t, 0, s.Store.UnpulsedAnswers(constant.LowerAlfa))
}

func TestUnansweredDecisionsNeverTrigger(t *testing.T) {
	s := store_tester.New(t)
	s.Store.PushDecision(constant.LowerAlfa, "unanswered?", "narrow", nil, nil)
	assert.Count(t, 0, s.Store.UnpulsedAnswers(constant.LowerAlfa))
}
