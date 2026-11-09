package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
)

func TestAnsweredDecisionsStartUnpulsed(t *testing.T) {
	s := store_tester.New(t)
	s.Answered("first?")
	s.Answered("second?")
	assert.Count(t, 2, s.Store.UnpulsedAnswers(stringsConstant.LowerAlfa))
}

func TestPulsingBatchEmptiesIt(t *testing.T) {
	s := store_tester.New(t)
	s.Answered("first?")
	s.Answered("second?")
	s.Store.MarkPulsed(stringsConstant.LowerAlfa)
	assert.Count(t, 0, s.Store.UnpulsedAnswers(stringsConstant.LowerAlfa))
}

func TestNewAnswerAfterPulsingTriggersAgain(t *testing.T) {
	s := store_tester.New(t)
	s.Answered("first?")
	s.Store.MarkPulsed(stringsConstant.LowerAlfa)
	s.Answered("second?")
	assert.Count(t, 1, s.Store.UnpulsedAnswers(stringsConstant.LowerAlfa))
}

func TestPollingAlsoStopsTrigger(t *testing.T) {
	s := store_tester.New(t)
	s.Answered("first?")
	s.Store.MarkSeen(stringsConstant.LowerAlfa)
	assert.Count(t, 0, s.Store.UnpulsedAnswers(stringsConstant.LowerAlfa))
}

func TestUnansweredDecisionsNeverTrigger(t *testing.T) {
	s := store_tester.New(t)
	s.Store.PushDecision(
		stringsConstant.LowerAlfa,
		"unanswered?",
		"narrow",
		nil,
		nil,
	)
	assert.Count(t, 0, s.Store.UnpulsedAnswers(stringsConstant.LowerAlfa))
}
