package unit

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
	"github.com/funtimecoding/soil/pkg/microsoft/graph_server"
	"github.com/funtimecoding/soil/pkg/microsoft/graph_tester"
	"github.com/funtimecoding/soil/pkg/microsoft/mock_token_store"
	"net/http"
	"testing"
)

func TestRequestCarriesStoredAccessToken(t *testing.T) {
	s := graph_server.New(t).Route(
		constant.ProfilePath,
		http.StatusOK,
		graph_tester.Profile(),
	)
	c := graph_tester.New(t).Client(
		s,
		mock_token_store.New(t, graph_tester.Valid()),
	)
	assert.FatalOnError(t, c.Reachable(t.Context()))
	x := s.Exchanges(constant.ProfilePath)
	assert.Count(t, 1, x)
	assert.String(t, "Bearer stored-access", x[0].Authorization)
	assert.Count(t, 0, s.Exchanges(graph_tester.TokenPath()))
}

func TestExpiredTokenRefreshedAndSaved(t *testing.T) {
	s := graph_server.New(t).
		Route(constant.ProfilePath, http.StatusOK, graph_tester.Profile()).
		Route(graph_tester.TokenPath(), http.StatusOK, graph_tester.Refreshed())
	store := mock_token_store.New(t, graph_tester.Expired())
	assert.FatalOnError(
		t,
		graph_tester.New(t).Client(s, store).Reachable(t.Context()),
	)
	refresh := s.Exchanges(graph_tester.TokenPath())
	assert.Count(t, 1, refresh)
	assert.String(t, "refresh_token", refresh[0].Form.Get("grant_type"))
	assert.String(t, "stored-refresh", refresh[0].Form.Get("refresh_token"))
	assert.String(t, "client", refresh[0].Form.Get("client_id"))
	assert.String(
		t,
		"Bearer fresh-access",
		s.Exchanges(constant.ProfilePath)[0].Authorization,
	)
	saved := store.Saved()
	assert.Count(t, 1, saved)
	assert.String(t, "fresh-refresh", saved[0].RefreshToken)
}

func TestMissingTokenNotConfigured(t *testing.T) {
	s := graph_server.New(t).Route(
		constant.ProfilePath,
		http.StatusOK,
		graph_tester.Profile(),
	)
	e := graph_tester.New(t).Client(s, mock_token_store.New(t, nil)).Reachable(
		t.Context(),
	)
	assert.True(t, not_configured.Is(e))
	assert.String(t, "graph sign-in required: run gograph login", e.Error())
	assert.Count(t, 0, s.Exchanges(constant.ProfilePath))
}

func TestExpiredGrantNotConfigured(t *testing.T) {
	e := graph_tester.New(t).RefreshRefusal(
		constant.ExpiredGrant,
		"AADSTS700082",
	)
	assert.True(t, not_configured.Is(e))
	assert.String(
		t,
		"graph sign-in expired: run gograph login: invalid_grant AADSTS700082",
		e.Error(),
	)
}

func TestRequiredInteractionNotConfigured(t *testing.T) {
	e := graph_tester.New(t).RefreshRefusal(
		constant.InteractionRequired,
		"AADSTS50076",
	)
	assert.True(t, not_configured.Is(e))
	assert.String(
		t,
		"graph sign-in expired: run gograph login: interaction_required AADSTS50076",
		e.Error(),
	)
}

func TestInvalidClientNotConfigured(t *testing.T) {
	e := graph_tester.New(t).RefreshRefusal(
		constant.InvalidClient,
		"AADSTS7000215",
	)
	assert.True(t, not_configured.Is(e))
	assert.String(
		t,
		"graph app registration refused the sign-in: invalid_client AADSTS7000215",
		e.Error(),
	)
}

func TestUnauthorizedClientNotConfigured(t *testing.T) {
	e := graph_tester.New(t).RefreshRefusal(
		constant.UnauthorizedClient,
		"AADSTS700016",
	)
	assert.True(t, not_configured.Is(e))
	assert.String(
		t,
		"graph app registration refused the sign-in: unauthorized_client AADSTS700016",
		e.Error(),
	)
}

func TestFailedSaveKeepsFreshToken(t *testing.T) {
	s := graph_server.New(t).
		Route(constant.ProfilePath, http.StatusOK, graph_tester.Profile()).
		Route(graph_tester.TokenPath(), http.StatusOK, graph_tester.Refreshed())
	store := mock_token_store.New(t, graph_tester.Expired())
	store.SaveFailure = errors.New("keychain locked")
	c := graph_tester.New(t).Client(s, store)
	e := c.Reachable(t.Context())
	assert.String(t, "save graph token: keychain locked", fmt.Sprint(e))
	assert.FatalOnError(t, c.Reachable(t.Context()))
	assert.Count(t, 1, s.Exchanges(graph_tester.TokenPath()))
	x := s.Exchanges(constant.ProfilePath)
	assert.Count(t, 1, x)
	assert.String(t, "Bearer fresh-access", x[0].Authorization)
}

func TestRefusedTokenRenewedOnce(t *testing.T) {
	s := graph_server.New(t).
		Route(
			constant.ProfilePath,
			http.StatusUnauthorized,
			`{"error":{"message":"expired"}}`,
		).
		Route(constant.ProfilePath, http.StatusOK, graph_tester.Profile()).
		Route(graph_tester.TokenPath(), http.StatusOK, graph_tester.Refreshed())
	assert.FatalOnError(
		t,
		graph_tester.New(t).Client(
			s,
			mock_token_store.New(t, graph_tester.Valid()),
		).Reachable(t.Context()),
	)
	assert.Count(t, 1, s.Exchanges(graph_tester.TokenPath()))
	x := s.Exchanges(constant.ProfilePath)
	assert.Count(t, 2, x)
	assert.String(t, "Bearer fresh-access", x[1].Authorization)
}

func TestMissingResourceNotFound(t *testing.T) {
	e := graph_tester.New(t).Client(
		graph_server.New(t),
		mock_token_store.New(t, graph_tester.Valid()),
	).Reachable(t.Context())
	assert.True(t, not_found.Is(e))
}
