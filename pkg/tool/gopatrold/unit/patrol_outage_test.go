package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/unit/base"
	"net"
	"os"
	"syscall"
	"testing"
)

func TestHostOutageStaysQuietAndStopsAsking(t *testing.T) {
	s := base.New(t)
	s.Alfa().Failure = &net.OpError{
		Op:  "dial",
		Net: "tcp",
		Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
	}

	for range 3 {
		s.Poll()
	}

	assert.Count(t, 0, s.Reporter.Events())
	assert.Integer(t, 1, s.Alfa().Calls)
}

func TestHostFaultIsCapturedEachRound(t *testing.T) {
	s := base.New(t)
	s.Alfa().Failure = errors.New("sudo: a password is required")
	s.Poll()
	s.Poll()
	assert.Count(t, 2, s.Reporter.Events())
}
