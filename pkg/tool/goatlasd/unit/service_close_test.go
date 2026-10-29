package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	"testing"
)

func TestCloseClosesEveryPeer(t *testing.T) {
	v := service_tester.New(t)
	v.Service.Close()
	assert.Integer(t, 1, v.Netbox.Closed)
	assert.Integer(t, 1, v.Process.Closed)
	assert.Integer(t, 1, v.Lease.Closed)
}
