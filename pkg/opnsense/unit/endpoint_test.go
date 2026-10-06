package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	opnsense "github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSendsKeyAndQuery(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `[{"action":"pass"}]`)
	var out []map[string]string
	assert.FatalOnError(
		t,
		newClient(t, s.URL, false).Get(
			opnsense.LogRead,
			map[string]string{"limit": "5"},
			&out,
		),
	)
	assert.String(t, "pass", out[0]["action"])
	r := last()
	assert.String(t, "/api/diagnostics/firewall/log", r.URL.Path)
	assert.String(t, "5", r.URL.Query().Get("limit"))
	key, secret, okay := r.BasicAuth()
	assert.True(t, okay)
	assert.String(t, "alfa-key", key)
	assert.String(t, "bravo-secret", secret)
	assert.String(t, "application/json", r.Header.Get(constant.Accept))
}

func TestPostSendsNotation(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"result":"saved"}`)
	var out map[string]string
	assert.FatalOnError(
		t,
		newClient(t, s.URL, false).Post(
			opnsense.HostAdd,
			map[string]string{"host": "charlie"},
			&out,
		),
	)
	assert.String(t, "saved", out["result"])
	r := last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "application/json", r.Header.Get(constant.ContentType))
	b, e := io.ReadAll(r.Body)
	assert.FatalOnError(t, e)
	assert.String(t, `{"host":"charlie"}`, string(b))
}

func TestMissingEndpointIsNotFound(t *testing.T) {
	s, _ := upstream_tester.New(t, http.StatusNotFound, "")
	var out map[string]any
	e := newClient(t, s.URL, false).Get("delta/echo", nil, &out)
	assert.True(t, not_found.Is(e))
}

func TestRefusalCarriesTheMessage(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusUnauthorized,
		`{"status":401,"message":"Authentication Failed"}`,
	)
	var out map[string]any
	e := newClient(t, s.URL, false).Get("delta/echo", nil, &out)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "status: 401: Authentication Failed", e.Error())
}

func TestUntrustedReachesASelfSignedFirewall(t *testing.T) {
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
	assert.FatalOnError(t, newClient(t, s.URL, true).Get("foxtrot", nil, &out))
	assert.String(t, "ok", out["status"])
	e := newClient(t, s.URL, false).Get("foxtrot", nil, &out)
	assert.True(t, unreachable.Is(e))
	assert.StringContains(t, "untrusted certificate", e.Error())
}
