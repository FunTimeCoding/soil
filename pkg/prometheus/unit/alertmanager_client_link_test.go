package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net"
	"testing"
)

func TestAlertmanagerLinkFollowsConnection(t *testing.T) {
	s := silenceServer()
	defer s.Close()
	assert.String(
		t,
		join.Empty(
			"http://127.0.0.1:",
			integers.ToString(s.Listener.Addr().(*net.TCPAddr).Port),
			"/#/silences/Alfa",
		),
		silenceLink(t, s),
	)
}

func TestAlertmanagerLinkUsesLinkEndpoint(t *testing.T) {
	s := silenceServer()
	defer s.Close()
	assert.String(
		t,
		"https://alertmanager.example/#/silences/Alfa",
		silenceLink(
			t,
			s,
			alertmanager.WithLink("alertmanager.example", 0, true),
		),
	)
}

func TestPrometheusGraphLinkFollowsConnection(t *testing.T) {
	p := prometheus.New("127.0.0.1", 9090, false, "", "")
	defer p.Close()
	assert.String(
		t,
		"http://127.0.0.1:9090/graph?g0.expr=up",
		p.GraphLink(constant.Up),
	)
}

func TestPrometheusGraphLinkUsesLinkEndpoint(t *testing.T) {
	p := prometheus.New(
		"127.0.0.1",
		9090,
		false,
		"",
		"",
		prometheus.WithLink("prometheus.example", 0, true),
	)
	defer p.Close()
	assert.String(
		t,
		"https://prometheus.example/graph?g0.expr=up",
		p.GraphLink(constant.Up),
	)
}
