package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/unit/stalled"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"io"
	"net"
	"net/url"
	"os"
	"syscall"
	"testing"
)

func TestConnectionRefusedIsUnreachable(t *testing.T) {
	f := connection.Classify(
		&url.Error{
			Op:  "Get",
			URL: "http://alfa.example:8080/api/items?api_key=secret",
			Err: &net.OpError{
				Op:  "dial",
				Net: "tcp",
				Err: os.NewSyscallError("connect", syscall.ECONNREFUSED),
			},
		},
	)
	assert.True(t, unreachable.Is(f))
	assert.String(t, "alfa.example:8080: connection refused", f.Error())
	assert.String(t, "/api/items", f.Path)
	assert.Strings(
		t,
		[]string{"unreachable", "alfa.example:8080"},
		f.Fingerprint(),
	)
}

func TestConnectionUnknownHostIsUnreachable(t *testing.T) {
	f := connection.Classify(
		&url.Error{
			Op:  "Get",
			URL: "https://bravo.example/",
			Err: &net.DNSError{Err: "no such host", Name: "bravo.example"},
		},
	)
	assert.True(t, unreachable.Is(f))
	assert.String(t, "bravo.example: unknown host", f.Error())
}

func TestConnectionDeadlineIsTimeout(t *testing.T) {
	f := connection.Classify(
		&url.Error{
			Op:  "Get",
			URL: "http://charlie.example:8594/api/processes",
			Err: &net.OpError{Op: "dial", Net: "tcp", Err: stalled.New()},
		},
	)
	assert.True(t, timeout.Is(f))
	assert.String(t, "charlie.example:8594: timed out", f.Error())
}

func TestConnectionBrokenAnswerIsDropped(t *testing.T) {
	f := connection.Classify(
		&url.Error{
			Op:  "Get",
			URL: "https://delta.example/repos",
			Err: io.ErrUnexpectedEOF,
		},
	)
	assert.True(t, dropped.Is(f))
	assert.String(t, "delta.example: connection dropped", f.Error())
}

func TestConnectionLeavesOtherErrorsAlone(t *testing.T) {
	assert.True(t, connection.Classify(errors.New("bad request")) == nil)
	assert.True(t, connection.Classify(nil) == nil)
}
