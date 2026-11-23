package graph_tester

import (
	"github.com/funtimecoding/soil/pkg/microsoft/graph_server"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
)

func (g *Tester) Paged() *graph_server.Server {
	g.t.Helper()
	s := graph_server.New(g.t)
	s.Route(
		"/v1.0/items",
		http.StatusOK,
		join.Empty(
			`{"value":[1,2],"@odata.nextLink":"`,
			s.Address(),
			`/v1.0/next?skip=2"}`,
		),
	)
	s.Route("/v1.0/next", http.StatusOK, `{"value":[3,4]}`)

	return s
}
