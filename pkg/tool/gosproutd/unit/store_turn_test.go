package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"testing"
)

func TestTurnsKeepTheirOrder(t *testing.T) {
	s := store_tester.New(t)
	d := s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	s.Store.AddTurn(d.Identifier, constant.AuthorSession, "narrowing it")
	s.Store.AddTurn(d.Identifier, constant.AuthorUser, "go on")
	turns := s.Store.Decisions(stringConstant.LowerAlfa)[0].Turns
	assert.Count(t, 2, turns)
	assert.String(t, "narrowing it", turns[0].Content)
	assert.String(t, "go on", turns[1].Content)
}

func TestTurnCountCountsOneAuthor(t *testing.T) {
	s := store_tester.New(t)
	d := s.Store.PushDecision(
		stringConstant.LowerAlfa,
		"which entry rule?",
		"narrow",
		nil,
		nil,
	)
	s.Store.AddTurn(d.Identifier, constant.AuthorSession, "narrowing it")
	s.Store.AddTurn(d.Identifier, constant.AuthorUser, "go on")
	s.Store.AddTurn(d.Identifier, constant.AuthorSession, "understood")
	assert.Integer(
		t,
		int64(2),
		s.Store.TurnCount(d.Identifier, constant.AuthorSession),
	)
}
