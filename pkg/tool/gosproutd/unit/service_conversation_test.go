package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/service_tester"
	"testing"
)

func TestClearWithHeardRecordsTheConversationAnswer(t *testing.T) {
	s := service_tester.New(t)
	d, e := s.Service.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	assert.FatalOnError(t, e)
	s.Service.Clear(
		d.Identifier,
		"going narrow",
		"keep it narrow, said in passing",
	)
	r := s.Service.Decisions(stringConstant.LowerAlfa)[0]
	assert.String(t, "answered", string(r.State))
	assert.String(t, "keep it narrow, said in passing", r.Answer)
}

func TestConversationAnswerIsMarkedAsSuch(t *testing.T) {
	s := service_tester.New(t)
	d, e := s.Service.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	assert.FatalOnError(t, e)
	s.Service.Clear(
		d.Identifier,
		"going narrow",
		"keep it narrow, said in passing",
	)
	assert.String(
		t,
		"conversation",
		string(s.Service.Decisions(stringConstant.LowerAlfa)[0].AnswerChannel),
	)
}

func TestConversationAnswerResolvesAsAnswerNotDefault(t *testing.T) {
	s := service_tester.New(t)
	d, e := s.Service.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	assert.FatalOnError(t, e)
	s.Service.Clear(
		d.Identifier,
		"going narrow",
		"keep it narrow, said in passing",
	)
	assert.String(
		t,
		"answer",
		string(s.Service.Decisions(stringConstant.LowerAlfa)[0].Resolution),
	)
}

func TestHeardNeverOverwritesAQueueAnswer(t *testing.T) {
	s := service_tester.New(t)
	d, e := s.Service.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	assert.FatalOnError(t, e)
	s.Service.Answer(d.Identifier, constant.AnswerKindChoice, "wide")
	s.Service.Clear(
		d.Identifier,
		"going narrow",
		"keep it narrow, said in passing",
	)
	r := s.Service.Decisions(stringConstant.LowerAlfa)[0]
	assert.String(t, "wide", r.Answer)
	assert.String(t, "queue", string(r.AnswerChannel))
}

func TestClearWithoutHeardStillTakesTheDefault(t *testing.T) {
	s := service_tester.New(t)
	d, e := s.Service.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	assert.FatalOnError(t, e)
	s.Service.Clear(d.Identifier, "went narrow", "")
	r := s.Service.Decisions(stringConstant.LowerAlfa)[0]
	assert.String(t, "default", string(r.Resolution))
	assert.String(t, "open", string(r.State))
}
