package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
	"time"
)

func TestAnsweredAtTakesTheInjectedClock(t *testing.T) {
	s := store_tester.New(t)
	d := s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	s.Advance(time.Hour)
	s.Store.Answer(
		d.Identifier,
		constant.AnswerKindChoice,
		constant.AnswerChannelQueue,
		"narrow",
	)
	answered := s.Store.Decisions(stringConstant.LowerAlfa)[0].AnsweredAt
	assert.True(t, answered != nil)
	assert.True(t, !answered.After(s.Clock()()))
	assert.True(t, answered.Equal(s.Clock()()))
}

func TestDismissTakesTheInjectedClock(t *testing.T) {
	s := store_tester.New(t)
	d := s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	s.Advance(90 * time.Minute)
	s.Store.Dismiss(d.Identifier, constant.StatePostponed)
	answered := s.Store.Decisions(stringConstant.LowerAlfa)[0].AnsweredAt
	assert.True(t, answered != nil)
	assert.True(t, answered.Equal(s.Clock()()))
}
