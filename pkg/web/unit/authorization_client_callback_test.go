package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"net/http"
	"testing"
)

func TestRefusedExchangeAnswersWithTheProviderReason(t *testing.T) {
	s := newFakeIssuer(t)
	r := memory.New()
	w := signInAndReturn(t, newAuthorizationClient(s.URL, r), "expired")
	assert.Integer(t, http.StatusBadGateway, w.Code)
	assert.StringContains(t, "token exchange failed", w.Body.String())
	assert.StringContains(t, "Code expired", w.Body.String())
	assert.Count(t, 1, r.Events())
}

func TestFailedDiscoveryIsRetriedNextLogin(t *testing.T) {
	s := newFakeIssuer(t)
	r := memory.New()
	c := newAuthorizationClient(s.URL, r)
	first := signInAndReturn(t, c, "golf")
	assert.Integer(t, http.StatusBadGateway, first.Code)
	assert.StringContains(
		t,
		"identity provider discovery failed",
		first.Body.String(),
	)
	second := signInAndReturn(t, c, "golf")
	assert.Integer(t, http.StatusBadGateway, second.Code)
	assert.Integer(t, 2, s.discovery)
}
