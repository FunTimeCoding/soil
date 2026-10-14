package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/mock_closer"
	"net"
	"os"
	"testing"
)

func TestLogRecloseIgnoresAlreadyClosed(t *testing.T) {
	b := captureLog(t)
	file := mock_closer.New(os.ErrClosed)
	listener := mock_closer.New(net.ErrClosed)
	errors.LogReclose(file)
	errors.LogReclose(listener)
	assert.Integer(t, 1, file.Closed())
	assert.Integer(t, 1, listener.Closed())
	assert.String(t, "", b.String())
}

func TestLogRecloseLogsOtherErrors(t *testing.T) {
	b := captureLog(t)
	errors.LogReclose(mock_closer.New(os.ErrPermission))
	assert.StringContains(t, "permission denied", b.String())
}
