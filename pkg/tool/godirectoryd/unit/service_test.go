package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestServiceKeepsUsersAndGroups(t *testing.T) {
	v := service_tester.New(t).Service()
	u, e := v.CreateUser("alfa", "Alfa", "Person", "alfa@example.test", "")
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", u.Account)
	g, f := v.CreateGroup("operators")
	assert.FatalOnError(t, f)
	assert.String(t, "operators", g.Name)
	m, h := v.AddMember("operators", "alfa")
	assert.FatalOnError(t, h)
	assert.Count(t, 1, m.Member)
	_, i := v.AddMember("operators", "alfa")
	assert.True(t, conflict.Is(i))
}

func TestQuietRefusalIsNotCapturedByGuard(t *testing.T) {
	s := service_tester.New(t)
	v := s.Service()
	_, e := v.User("ghost")
	assert.True(t, not_found.Is(e))
	f := v.DeleteUser("ghost")
	assert.True(t, not_found.Is(f))
	assert.Count(t, 0, s.Reporter.Events())
}

func TestOutageIsAnsweredAsUnavailableByService(t *testing.T) {
	s := service_tester.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	_, e := s.Service().Users()
	assert.True(t, unavailable.Is(e))
	assert.Count(t, 0, s.Reporter.Events())
}
