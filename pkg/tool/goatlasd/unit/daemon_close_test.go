package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/base"
	"testing"
)

func TestCloseClosesEveryPeer(t *testing.T) {
	s := base.New(t)
	s.Daemon().Close()
	assert.Integer(t, 1, s.Netbox.Closed)
	assert.Integer(t, 1, s.Process.Closed)
	assert.Integer(t, 1, s.Lease.Closed)
}
