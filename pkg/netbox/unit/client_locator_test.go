package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/fixture"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientSendsPublicHostThroughLocator(t *testing.T) {
	var h http.Header
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				h = q.Header.Clone()
				w.Header().Set("Content-Type", "application/json")
				_, e := w.Write([]byte(constant.FixtureMacList))
				errors.PanicOnError(e)
			},
		),
	)
	defer s.Close()
	c := netbox.New(
		locator.New("127.0.0.1").
			Port(s.Listener.Addr().(*net.TCPAddr).Port).
			Insecure(),
		"token",
		netbox.WithPublicHost("netbox.example"),
	)
	defer c.Close()
	a, e := net.ParseMAC(constant.FixturePhysicalAddress)
	assert.FatalOnError(t, e)
	_, e = c.PhysicalAddressesByHardware(a)
	assert.FatalOnError(t, e)
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
	_, e = c.PhysicalAddressesByHardware(a)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, testutil.CollectAndCount(h))
}

func TestClientWithoutPublicHostSendsNoForwardedHeaders(t *testing.T) {
	var h http.Header
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				q *http.Request,
			) {
				h = q.Header.Clone()
				w.Header().Set("Content-Type", "application/json")
				_, e := w.Write([]byte(constant.FixtureMacList))
				errors.PanicOnError(e)
			},
		),
	)
	defer s.Close()
	c := netbox.New(
		locator.New("127.0.0.1").
			Port(s.Listener.Addr().(*net.TCPAddr).Port).
			Insecure(),
		"token",
	)
	defer c.Close()
	a, e := net.ParseMAC(constant.FixturePhysicalAddress)
	assert.FatalOnError(t, e)
	_, e = c.PhysicalAddressesByHardware(a)
	assert.FatalOnError(t, e)
	assert.String(t, "", h.Get("X-Forwarded-Host"))
	assert.String(t, "", h.Get("X-Forwarded-Proto"))
}
