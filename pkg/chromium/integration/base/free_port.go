package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"net"
	"testing"
)

func freePort(t *testing.T) int {
	t.Helper()
	l, e := net.Listen(constant.Transmission, "localhost:0")
	assert.FatalOnError(t, e)
	port := l.Addr().(*net.TCPAddr).Port
	assert.FatalOnError(t, l.Close())

	return port
}
