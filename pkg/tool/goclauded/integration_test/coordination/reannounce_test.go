package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestReannounceAfterRestart(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.CheckLive()
	s.Service.ClearBindings()
	r := a.CheckLive()
	reannounce := clientEntriesByKind(r.Entries, constant.QueueReannounce)
	assert.Count(t, 1, reannounce)
}

func TestReannounceNotSetWithoutPriorBinding(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.CheckLive()
	s.Service.ClearBindings()
	r := a.CheckLive()
	reannounce := clientEntriesByKind(r.Entries, constant.QueueReannounce)
	assert.Count(t, 0, reannounce)
}

func TestAnnounceClearsReannounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "first topic")
	a.CheckLive()
	s.Service.ClearBindings()
	a.Announce(a.Name(), "second topic")
	r := a.CheckLive()
	reannounce := clientEntriesByKind(r.Entries, constant.QueueReannounce)
	assert.Count(t, 0, reannounce)
}

func TestReannounceConsumedByCheck(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.CheckLive()
	s.Service.ClearBindings()
	first := a.CheckLive()
	reannounce := clientEntriesByKind(first.Entries, constant.QueueReannounce)
	assert.Count(t, 1, reannounce)
	second := a.CheckLive()
	reannounce = clientEntriesByKind(second.Entries, constant.QueueReannounce)
	assert.Count(t, 0, reannounce)
}
