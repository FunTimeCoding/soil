package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/docker/constant"
	"github.com/funtimecoding/soil/pkg/docker/unit/hub_tester"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestTagsReadsPage(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"results":[{"name":"1.27-alpine"}]}`))
	tags, e := hub_tester.New(t, s.URL).Tags("library/alfa")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, tags)
	assert.String(t, "1.27-alpine", tags[0].Name)
	r := s.Last()
	assert.String(t, "/v2/repositories/library/alfa/tags", r.URL.Path)
	assert.String(t, "100", r.URL.Query().Get(constant.PageSizeParameter))
	assert.String(t, "application/json", r.Header.Get(webConstant.Accept))
}

func TestMissingImageIsNotFound(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusNotFound),
		upstream_server.WithAnswer(`{"message":"object not found"}`),
	)
	_, e := hub_tester.New(t, s.URL).Tags("library/bravo")
	assert.True(t, not_found.Is(e))
}

func TestOutageIsUnavailable(t *testing.T) {
	for _, status := range []int{
		http.StatusTooManyRequests,
		http.StatusServiceUnavailable,
	} {
		s := upstream_server.New(
			t,
			upstream_server.WithStatus(status),
			upstream_server.WithAnswer(`{"message":"slow down"}`),
		)
		_, e := hub_tester.New(t, s.URL).Tags("library/charlie")
		assert.True(t, unavailable.Is(e))
	}
}

func TestRefusedTokenIsForbidden(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithAnswer(`{"message":"denied"}`),
	)
	_, e := hub_tester.New(t, s.URL).Tags("library/delta")
	assert.True(t, forbidden.Is(e))
}

func TestRequestWeWroteBadlyStaysUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(`{"message":"bad"}`),
	)
	_, e := hub_tester.New(t, s.URL).Tags("library/echo")
	assert.Error(t, e)
	assert.String(t, "", classify.Class(e))
}

func TestReachableReadsProbeImage(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{"name":"alpine"}`))
	assert.FatalOnError(
		t,
		hub_tester.New(t, s.URL).Reachable(context.Background()),
	)
	assert.String(t, "/v2/repositories/library/alpine", s.Last().URL.Path)
}
