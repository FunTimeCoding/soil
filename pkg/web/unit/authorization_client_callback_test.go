package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/web/unit/issuer_server"
	"net/http"
	"testing"
)

func TestRefusedExchangeAnswersWithProviderReason(t *testing.T) {
	s := issuer_server.New(t)
	r := memory.New()
	owned := s.Client(r)
	defer owned.Close()
	w := issuer_server.SignIn(t, owned, "expired")
	assert.Integer(t, http.StatusBadGateway, w.Code)
	assert.StringContains(t, "token exchange failed", w.Body.String())
	assert.StringContains(t, "Code expired", w.Body.String())
	assert.Count(t, 1, r.Events())
}

func TestFailedDiscoveryIsRetriedNextLogin(t *testing.T) {
	s := issuer_server.New(t)
	r := memory.New()
	c := s.Client(r)
	defer c.Close()
	first := issuer_server.SignIn(t, c, "golf")
	assert.Integer(t, http.StatusBadGateway, first.Code)
	assert.StringContains(
		t,
		"identity provider discovery failed",
		first.Body.String(),
	)
	second := issuer_server.SignIn(t, c, "golf")
	assert.Integer(t, http.StatusBadGateway, second.Code)
	assert.Integer(t, 2, s.Discoveries())
}
