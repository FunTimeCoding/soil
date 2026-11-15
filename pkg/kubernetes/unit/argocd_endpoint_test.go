package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/kubernetes/unit/argocd_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestApplicationsSendToken(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithAnswer(`{"items":[]}`))
	_, e := argocd_tester.New(t, s.URL).Applications()
	assert.FatalOnError(t, e)
	r := s.Last()
	assert.String(t, "/api/v1/applications", r.URL.Path)
	assert.String(
		t,
		"Bearer alfa-token",
		r.Header.Get(webConstant.Authorization),
	)
}

func TestRefusalCarriesReason(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusForbidden),
		upstream_server.WithAnswer(
			`{"error":"permission denied","code":7,"message":"permission denied"}`,
		),
	)
	_, e := argocd_tester.New(t, s.URL).Applications()
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status: 403: permission denied", e.Error())
}
