package graph_tester

import (
	"github.com/funtimecoding/soil/pkg/microsoft/graph_server"
	"github.com/funtimecoding/soil/pkg/microsoft/mock_token_store"
	"golang.org/x/oauth2"
	"net/http"
)

func (g *Tester) LoginRefusal(code string) (*mock_token_store.Store, error) {
	g.t.Helper()
	s := graph_server.New(g.t).
		Route(DevicePath(), http.StatusOK, DeviceCode(900)).
		Route(TokenPath(), http.StatusBadRequest, Refusal(code, "AADSTS70000"))
	store := mock_token_store.New(g.t, nil)

	return store, g.Client(s, store).Login(
		g.t.Context(),
		func(*oauth2.DeviceAuthResponse) {},
	)
}
