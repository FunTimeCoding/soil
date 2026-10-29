package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generated "github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/server"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestCloseClosesTheHabiticaClient(t *testing.T) {
	c := service_tester.NewSource(t)
	service_tester.Service(t, c).Close()
	assert.Integer(t, 1, c.Closed)
}

func TestWaitAnswersOnceHabiticaAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		service_tester.Service(t, service_tester.NewSource(t)).Wait(t.Context()),
	)
}

func TestEquipAnswersTheEquippedGear(t *testing.T) {
	v := service_tester.Service(t, service_tester.NewSource(t))
	g, e := v.Equip("weapon_warrior_1")
	assert.FatalOnError(t, e)
	assert.String(t, "weapon_warrior_1", g.Equipped["weapon"])
}

func TestAnOutageIsAnsweredAsUnavailableWithoutACapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}
	m := memory.New()
	s := server.New(service_tester.Service(t, c), m)

	for range 2 {
		r, e := s.GetStats(
			context.Background(),
			generated.GetStatsRequestObject{},
		)
		assert.FatalOnError(t, e)
		_, away := r.(generated.GetStats503JSONResponse)
		assert.True(t, away)
	}

	assert.Count(t, 0, m.Events())
}

func TestAnExpiredTokenIsAnsweredAsForbiddenWithoutACapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = forbidden.Format("habitica: %s", "token expired")
	m := memory.New()
	r, e := server.New(service_tester.Service(t, c), m).GetTasks(
		context.Background(),
		generated.GetTasksRequestObject{},
	)
	assert.FatalOnError(t, e)
	_, refused := r.(generated.GetTasks403JSONResponse)
	assert.True(t, refused)
	assert.Count(t, 0, m.Events())
}

func TestAConflictingCronIsAnsweredAsConflict(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = conflict.Format("habitica: %s", "cron already run")
	m := memory.New()
	r, e := server.New(service_tester.Service(t, c), m).RunCron(
		context.Background(),
		generated.RunCronRequestObject{},
	)
	assert.FatalOnError(t, e)
	_, refused := r.(generated.RunCron409JSONResponse)
	assert.True(t, refused)
	assert.Count(t, 0, m.Events())
}
