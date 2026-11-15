package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/integers"
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/alertmanager_tester"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/silence_tester"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestAlertmanagerLinkFollowsConnection(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(constant.FixtureSilences),
		upstream_server.WithObject(),
	)
	_, port := upstream_server.HostPort(t, s.URL)
	assert.String(
		t,
		join.Empty(
			"http://127.0.0.1:",
			integers.ToString(port),
			"/#/silences/Alfa",
		),
		silence_tester.Link(t, alertmanager_tester.New(t, s.URL)),
	)
}

func TestAlertmanagerLinkUsesLinkEndpoint(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(constant.FixtureSilences),
		upstream_server.WithObject(),
	)
	assert.String(
		t,
		"https://alertmanager.example/#/silences/Alfa",
		silence_tester.Link(
			t,
			alertmanager_tester.New(
				t,
				s.URL,
				alertmanager.WithLink("alertmanager.example", 0, true),
			),
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
