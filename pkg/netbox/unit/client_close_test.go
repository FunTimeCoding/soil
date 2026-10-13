package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"net"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestClientCloseDropsIdleConnections(t *testing.T) {
	var opened atomic.Int64
	s := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.Header().Set("Content-Type", "application/json")
				_, e := w.Write([]byte(constant.FixtureMacList))
				errors.PanicOnError(e)
			},
		),
	)
	s.Config.ConnState = func(
		_ net.Conn,
		c http.ConnState,
	) {
		if c == http.StateNew {
			opened.Add(1)
		}
	}
	s.Start()
	defer s.Close()
	c := netbox.New("netbox.example.org", "token", netbox.WithServer(s.URL))
	defer c.Close()
	a, e := net.ParseMAC(constant.FixturePhysicalAddress)
	assert.FatalOnError(t, e)
	_, e = c.PhysicalAddressesByHardware(a)
	assert.FatalOnError(t, e)
	c.Close()
	_, e = c.PhysicalAddressesByHardware(a)
	assert.FatalOnError(t, e)
	assert.Integer(t, int64(2), opened.Load())
}
