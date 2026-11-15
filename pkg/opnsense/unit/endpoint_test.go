package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/unit/opnsense_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSendsKeyAndQuery(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`[{"action":"pass"}]`),
	)
	var out []map[string]string
	assert.FatalOnError(
		t,
		opnsense_tester.Basic(t, s.URL, false).Get(
			constant.LogRead,
			map[string]string{"limit": "5"},
			&out,
		),
	)
	assert.String(t, "pass", out[0]["action"])
	r := s.Last()
	assert.String(t, "/api/diagnostics/firewall/log", r.URL.Path)
	assert.String(t, "5", r.URL.Query().Get("limit"))
	key, secret, okay := r.BasicAuth()
	assert.True(t, okay)
	assert.String(t, "alfa-key", key)
	assert.String(t, "bravo-secret", secret)
	assert.String(t, "application/json", r.Header.Get(webConstant.Accept))
}

func TestPostSendsNotation(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"result":"saved"}`),
	)
	var out map[string]string
	assert.FatalOnError(
		t,
		opnsense_tester.Basic(t, s.URL, false).Post(
			constant.HostAdd,
			map[string]string{"host": "charlie"},
			&out,
		),
	)
	assert.String(t, "saved", out["result"])
	r := s.Last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "application/json", r.Header.Get(webConstant.ContentType))
	b, e := io.ReadAll(r.Body)
	assert.FatalOnError(t, e)
	assert.String(t, `{"host":"charlie"}`, string(b))
}

func TestMissingEndpointIsNotFound(t *testing.T) {
	s := upstream_server.New(t, upstream_server.WithStatus(http.StatusNotFound))
	var out map[string]any
	e := opnsense_tester.Basic(t, s.URL, false).Get("delta/echo", nil, &out)
	assert.True(t, not_found.Is(e))
}

func TestRefusalCarriesMessage(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
		upstream_server.WithAnswer(
			`{"status":401,"message":"Authentication Failed"}`,
		),
	)
	var out map[string]any
	e := opnsense_tester.Basic(t, s.URL, false).Get("delta/echo", nil, &out)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status: 401: Authentication Failed", e.Error())
}

func TestUntrustedReachesSelfSignedFirewall(t *testing.T) {
	s := httptest.NewTLSServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				_, e := w.Write([]byte(`{"status":"ok"}`))
				errors.PanicOnError(e)
			},
		),
	)
	t.Cleanup(s.Close)
	var out map[string]string
	assert.FatalOnError(
		t,
		opnsense_tester.Basic(t, s.URL, true).Get("foxtrot", nil, &out),
	)
	assert.String(t, "ok", out["status"])
	e := opnsense_tester.Basic(t, s.URL, false).Get("foxtrot", nil, &out)
	assert.True(t, unreachable.Is(e))
	assert.StringContains(t, "untrusted certificate", e.Error())
}
