package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/delivery_tester"
	"testing"
)

func TestReannounceAfterRestart(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "building things")
	a.CheckLive()
	s.Restart()
	r := a.CheckLive()
	reannounce := delivery_tester.ClientKindEntries(
		r.Entries,
		constant.QueueReannounce,
	)
	assert.Count(t, 1, reannounce)
}

func TestReannounceNotSetWithoutPriorBinding(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.CheckLive()
	s.Restart()
	r := a.CheckLive()
	reannounce := delivery_tester.ClientKindEntries(
		r.Entries,
		constant.QueueReannounce,
	)
	assert.Count(t, 0, reannounce)
}

func TestAnnounceClearsReannounce(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "first topic")
	a.CheckLive()
	s.Restart()
	a.Announce(a.Name(), "second topic")
	r := a.CheckLive()
	reannounce := delivery_tester.ClientKindEntries(
		r.Entries,
		constant.QueueReannounce,
	)
	assert.Count(t, 0, reannounce)
}

func TestReannounceConsumedByCheck(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "building things")
	a.CheckLive()
	s.Restart()
	first := a.CheckLive()
	reannounce := delivery_tester.ClientKindEntries(
		first.Entries,
		constant.QueueReannounce,
	)
	assert.Count(t, 1, reannounce)
	second := a.CheckLive()
	reannounce = delivery_tester.ClientKindEntries(
		second.Entries,
		constant.QueueReannounce,
	)
	assert.Count(t, 0, reannounce)
}
