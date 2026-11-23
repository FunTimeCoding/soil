package graph_tester

import (
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
	"github.com/funtimecoding/soil/pkg/microsoft/graph_server"
	"github.com/funtimecoding/soil/pkg/microsoft/mock_token_store"
	"net/http"
)

func (g *Tester) RefreshRefusal(
	code string,
	directory string,
) error {
	g.t.Helper()
	s := graph_server.New(g.t).
		Route(constant.ProfilePath, http.StatusOK, Profile()).
		Route(TokenPath(), http.StatusBadRequest, Refusal(code, directory))

	return g.Client(s, mock_token_store.New(g.t, Expired())).Reachable(
		g.t.Context(),
	)
}
