package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/mock_closer"
	"github.com/funtimecoding/soil/pkg/log/log_tester"
	"net"
	"os"
	"testing"
)

func TestLogRecloseIgnoresAlreadyClosed(t *testing.T) {
	b := log_tester.Capture(t)
	file :=mock_closer.New(os.ErrClosed)
	listener := mock_closer.New(net.ErrClosed)
	errors.LogReclose(file)
	errors.LogReclose(listener)
	assert.Integer(t, 1, file.Closed())
	assert.Integer(t, 1, listener.Closed())
	assert.String(t, "", b.String())
}

func TestLogRecloseLogsOtherErrors(t *testing.T) {
	b := log_tester.Capture(t)
	errors.LogReclose(mock_closer.New(os.ErrPermission))
	assert.StringContains(t, "permission denied", b.String())
}
