package graph_tester

import (
	"github.com/funtimecoding/soil/pkg/microsoft/face"
	"github.com/funtimecoding/soil/pkg/microsoft/graph"
	"github.com/funtimecoding/soil/pkg/microsoft/graph_server"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
)

func (g *Tester) Client(
	s *graph_server.Server,
	store face.TokenStore,
) *graph.Client {
	g.t.Helper()
	host, port := upstream_server.HostPort(g.t, s.Address())
	result := graph.New(
		locator.New(host).Port(port).Insecure(),
		locator.New(host).Port(port).Insecure(),
		"tenant",
		"client",
		store,
	)
	g.t.Cleanup(result.Close)

	return result
}
