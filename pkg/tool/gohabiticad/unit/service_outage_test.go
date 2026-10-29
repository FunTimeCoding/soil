package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestAnOutageMarksTheGuardDownWithoutAnEarlyCapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}
	m := memory.New()
	v := service_tester.Shared(t, c, m)
	_, e := v.UserStats()
	assert.True(t, unavailable.Is(e))
	_, f := v.Tags()
	assert.True(t, unavailable.Is(f))
	assert.Count(t, 0, m.Events())
}
