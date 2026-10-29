package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	generated "github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/server"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/service_tester"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestTheServiceKeepsUsersAndGroups(t *testing.T) {
	v := service_tester.Shared(t, service_tester.NewSource(), memory.New())
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

func TestAQuietRefusalIsNotCapturedByTheGuard(t *testing.T) {
	m := memory.New()
	v := service_tester.Shared(t, service_tester.NewSource(), m)
	_, e := v.User("ghost")
	assert.True(t, not_found.Is(e))
	f := v.DeleteUser("ghost")
	assert.True(t, not_found.Is(f))
	assert.Count(t, 0, m.Events())
}

func TestAnOutageIsAnsweredAsUnavailableWithoutACapture(t *testing.T) {
	c := service_tester.NewSource()
	c.Failure = &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}
	m := memory.New()
	s := server.New(service_tester.Shared(t, c, m), m)

	for range 2 {
		r, e := s.GetUser(
			context.Background(),
			generated.GetUserRequestObject{},
		)
		assert.FatalOnError(t, e)
		_, away := r.(generated.GetUser503JSONResponse)
		assert.True(t, away)
	}

	_, g := service_tester.Shared(t, c, memory.New()).Users()
	assert.True(t, unavailable.Is(g))
	assert.Count(t, 0, m.Events())
}

func TestARefusedServiceAccountIsAnsweredAsForbidden(t *testing.T) {
	c := service_tester.NewSource()
	c.Failure = forbidden.Format(
		"directory cn=admin: %s",
		"invalid credentials",
	)
	m := memory.New()
	r, e := server.New(service_tester.Shared(t, c, m), m).GetGroup(
		context.Background(),
		generated.GetGroupRequestObject{},
	)
	assert.FatalOnError(t, e)
	_, refused := r.(generated.GetGroup403JSONResponse)
	assert.True(t, refused)
	assert.Count(t, 0, m.Events())
}

func TestWaitAnswersOnceTheDirectoryAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		service_tester.Shared(t, service_tester.NewSource(), memory.New()).
			Wait(t.Context()),
	)
}
