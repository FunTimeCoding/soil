package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/fixture"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"net"
	"testing"
)

func TestClientSendsPublicHostThroughLocator(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(constant.FixtureMacList),
		upstream_server.WithObject(),
	)
	host, port := upstream_server.HostPort(t, s.URL)
	c := netbox.New(
		locator.New(host).Port(port).Insecure(),
		"token",
		netbox.WithPublicHost("netbox.example"),
	)
	defer c.Close()
	a, e := net.ParseMAC(constant.FixturePhysicalAddress)
	assert.FatalOnError(t, e)
	_, e = c.HardwarePhysicalAddresses(a)
	assert.FatalOnError(t, e)
	h := s.Last().Header
	assert.String(t, "netbox.example", h.Get("X-Forwarded-Host"))
	assert.String(t, "https", h.Get("X-Forwarded-Proto"))
}

func TestClientObservesRequestDuration(t *testing.T) {
	h := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{Name: "alfa_seconds"},
		[]string{"host", "code", "method"},
	)
	s := fixture.NewServer(fixture.NewRecorder())
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
		netbox.WithRequestDuration(h),
	)
	defer c.Close()
	a, e := net.ParseMAC(constant.FixturePhysicalAddress)
	assert.FatalOnError(t, e)
	_, e = c.HardwarePhysicalAddresses(a)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, testutil.CollectAndCount(h))
}

func TestClientWithoutPublicHostSendsNoForwardedHeaders(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(constant.FixtureMacList),
		upstream_server.WithObject(),
	)
	host, port := upstream_server.HostPort(t, s.URL)
	c := netbox.New(locator.New(host).Port(port).Insecure(), "token")
	defer c.Close()
	a, e := net.ParseMAC(constant.FixturePhysicalAddress)
	assert.FatalOnError(t, e)
	_, e = c.HardwarePhysicalAddresses(a)
	assert.FatalOnError(t, e)
	h := s.Last().Header
	assert.String(t, "", h.Get("X-Forwarded-Host"))
	assert.String(t, "", h.Get("X-Forwarded-Proto"))
}
