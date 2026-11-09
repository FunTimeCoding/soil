package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/mock_client"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/face"
	"net/http"
	"testing"
)

func TestUtilizationClientReadsWithBearer(t *testing.T) {
	var bearer string
	var beta string
	c := utilizationUpstream(
		t,
		func(w http.ResponseWriter, q *http.Request) {
			bearer = q.Header.Get("Authorization")
			beta = q.Header.Get(constant.AnthropicBetaHeader)
			_, e := w.Write([]byte(fixture.Read("claude", "utilization.json")))
			errors.PanicOnError(e)
		},
	)
	r, e := c.Read(t.Context(), "alfa-token")
	assert.FatalOnError(t, e)
	assert.Integer(t, 22, r.WeeklyPercent)
	assert.String(t, "Bearer alfa-token", bearer)
	assert.String(t, "oauth-2025-04-20", beta)
}

func TestOverloadedUtilizationAnswerIsUnavailable(t *testing.T) {
	c := utilizationUpstream(
		t,
		func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusServiceUnavailable)
		},
	)
	_, e := c.Read(t.Context(), "alfa-token")
	assert.True(t, unavailable.Is(e))
}

func TestRefusedCredentialIsNoOutage(t *testing.T) {
	c := utilizationUpstream(
		t,
		func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		},
	)
	_, e := c.Read(t.Context(), "expired")
	assert.NotNil(t, e)
	assert.False(t, unavailable.Is(e))
	assert.FatalOnError(t, c.Reachable(t.Context()))
}

func TestAnyAnswerMakesUtilizationHostReachable(t *testing.T) {
	c := utilizationUpstream(
		t,
		func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		},
	)
	assert.FatalOnError(t, c.Reachable(t.Context()))
}

func TestGoneUtilizationHostIsUnreachable(t *testing.T) {
	c := utilization.New(closedLocator(t))
	defer c.Close()
	assert.True(t, unreachable.Is(c.Reachable(t.Context())))
}

func TestUtilizationClientAndMockHoldFace(t *testing.T) {
	c := utilization.New(closedLocator(t))
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	assert.Count(t, 2, []face.AnthropicClient{c, m})
}
