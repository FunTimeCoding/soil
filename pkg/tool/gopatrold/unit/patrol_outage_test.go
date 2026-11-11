package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/patrol/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/unit/daemon_tester"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestHostOutageStaysQuietAndStopsAsking(t *testing.T) {
	c := mock_client.New()
	c.Failure = &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}
	r := memory.New()
	d := daemon_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": c},
		metric.New(),
		r,
	)

	for range 3 {
		poll(t, d)
	}

	assert.Count(t, 0, r.Events())
	assert.Integer(t, 1, c.Calls)
}

func TestHostFaultIsCapturedEachRound(t *testing.T) {
	c := mock_client.New()
	c.Failure = errors.New("sudo: a password is required")
	r := memory.New()
	d := daemon_tester.New(
		t,
		map[string]*mock_client.Client{"alfa": c},
		metric.New(),
		r,
	)
	poll(t, d)
	poll(t, d)
	assert.Count(t, 2, r.Events())
}
