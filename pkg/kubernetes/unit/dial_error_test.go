package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/kubernetes/dial_error"
	"testing"
)

func TestDialErrorRefusedBecomesUnreachable(t *testing.T) {
	e := dial_error.Translate(
		errors.New(
			"error sending request: Post \"http://localhost:1/exec\": dial tcp [::1]:1: connect: connection refused",
		),
		"localhost:1",
	)
	assert.True(t, unreachable.Is(e))
	assert.String(t, "localhost:1", connection.Classify(e).Host)
}

func TestDialErrorUnknownHostBecomesUnreachable(t *testing.T) {
	e := dial_error.Translate(
		errors.New("dial tcp: lookup alfa.example: no such host"),
		"alfa.example",
	)
	assert.True(t, unreachable.Is(e))
}

func TestDialErrorTimeoutBecomesTimeout(t *testing.T) {
	e := dial_error.Translate(
		errors.New("dial tcp 10.0.0.1:443: i/o timeout"),
		"10.0.0.1",
	)
	assert.True(t, timeout.Is(e))
}

func TestDialErrorLeavesOtherErrorsUntouched(t *testing.T) {
	e := errors.New("command terminated with exit code 1")
	assert.True(t, dial_error.Translate(e, "alfa") == e)
	assert.True(t, dial_error.Translate(nil, "alfa") == nil)
}
