package base

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/unit/service_tester"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
	o ...service_tester.Option,
) *Server {
	t.Helper()

	return &Server{
		Tester:   service_tester.NewFixture(t, c, o...),
		Registry: prometheus.NewRegistry(),
		t:        t,
	}
}
