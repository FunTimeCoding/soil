package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"net"
	"testing"
)

func TestClaimedPortListensOnLoopbackOnly(t *testing.T) {
	p, l := system.ClaimPort()
	defer func() { assert.FatalOnError(t, l.Close()) }()
	a := l.Addr().(*net.TCPAddr)
	assert.String(t, "127.0.0.1", a.IP.String())
	assert.Integer(t, p, a.Port)
}
