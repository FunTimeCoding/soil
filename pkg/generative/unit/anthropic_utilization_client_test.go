package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/mock_client"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/face"
	"github.com/funtimecoding/soil/pkg/generative/unit/utilization_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestUtilizationClientReadsWithBearer(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(fixture.Read("claude", "utilization.json")),
	)
	r, e := utilization_tester.New(t, s.URL).Read(t.Context(), "alfa-token")
	assert.FatalOnError(t, e)
	assert.Integer(t, 22, r.WeeklyPercent)
	h := s.Last().Header
	assert.String(t, "Bearer alfa-token", h.Get(webConstant.Authorization))
	assert.String(t, "oauth-2025-04-20", h.Get(constant.AnthropicBetaHeader))
}

func TestOverloadedUtilizationAnswerIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
	)
	_, e := utilization_tester.New(t, s.URL).Read(t.Context(), "alfa-token")
	assert.True(t, unavailable.Is(e))
}

func TestRefusedCredentialIsNoOutage(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
	)
	c := utilization_tester.New(t, s.URL)
	_, e := c.Read(t.Context(), "expired")
	assert.NotNil(t, e)
	assert.False(t, unavailable.Is(e))
	assert.FatalOnError(t, c.Reachable(t.Context()))
}

func TestAnyAnswerMakesUtilizationHostReachable(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithStatus(http.StatusNotFound))
	assert.FatalOnError(
		t,
		utilization_tester.New(t, s.URL).Reachable(t.Context()),
	)
}

func TestGoneUtilizationHostIsUnreachable(t *testing.T) {
	c := utilization.New(locator.NewOrigin(outage_tester.RefusedAddress(t)))
	defer c.Close()
	assert.True(t, unreachable.Is(c.Reachable(t.Context())))
}

func TestUtilizationClientAndMockHoldFace(t *testing.T) {
	c := utilization.New(locator.NewOrigin(outage_tester.RefusedAddress(t)))
	defer c.Close()
	m := mock_client.New()
	defer m.Close()
	assert.Count(t, 2, []face.AnthropicClient{c, m})
}
