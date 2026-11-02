package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/docker/constant"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestTagsReadsThePage(t *testing.T) {
	s, last := upstream_tester.New(
		t,
		http.StatusOK,
		`{"results":[{"name":"1.27-alpine"}]}`,
	)
	tags, e := newHubClient(t, s.URL).Tags("library/alfa")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, tags)
	assert.String(t, "1.27-alpine", tags[0].Name)
	r := last()
	assert.String(t, "/v2/repositories/library/alfa/tags", r.URL.Path)
	assert.String(t, "100", r.URL.Query().Get(constant.PageSizeParameter))
	assert.String(t, "application/json", r.Header.Get(webConstant.Accept))
}

func TestMissingImageIsNotFound(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusNotFound,
		`{"message":"object not found"}`,
	)
	_, e := newHubClient(t, s.URL).Tags("library/bravo")
	assert.True(t, not_found.Is(e))
}

func TestAnOutageIsUnavailable(t *testing.T) {
	for _, status := range []int{
		http.StatusTooManyRequests,
		http.StatusServiceUnavailable,
	} {
		s, _ := upstream_tester.New(t, status, `{"message":"slow down"}`)
		_, e := newHubClient(t, s.URL).Tags("library/charlie")
		assert.True(t, unavailable.Is(e))
	}
}

func TestARefusedTokenIsForbidden(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusForbidden, `{"message":"denied"}`)
	_, e := newHubClient(t, s.URL).Tags("library/delta")
	assert.True(t, forbidden.Is(e))
}

func TestARequestWeWroteBadlyStaysUnexpected(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusBadRequest, `{"message":"bad"}`)
	_, e := newHubClient(t, s.URL).Tags("library/echo")
	assert.Error(t, e)
	assert.String(t, "", classify.Class(e))
}

func TestReachableReadsTheProbeImage(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"name":"alpine"}`)
	assert.FatalOnError(
		t,
		newHubClient(t, s.URL).Reachable(context.Background()),
	)
	assert.String(t, "/v2/repositories/library/alpine", last().URL.Path)
}
