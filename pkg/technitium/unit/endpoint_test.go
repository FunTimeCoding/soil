package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"github.com/funtimecoding/soil/pkg/technitium/unit/technitium_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetUnwrapsEnvelope(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"status":"ok","response":{"zones":[]}}`))
	payload, e := technitium_tester.New(t, s.URL).Get(
		"/zones/records/get?domain=alfa.example&listZone=true",
	)
	assert.FatalOnError(t, e)
	assert.String(t, `{"zones":[]}`, string(payload))
	r := s.Last()
	assert.String(t, "/api/zones/records/get", r.URL.Path)
	assert.String(t, "alfa.example", r.URL.Query().Get("domain"))
	assert.String(t, "true", r.URL.Query().Get("listZone"))
	assert.String(
		t,
		"Bearer alfa-token",
		r.Header.Get(webConstant.Authorization),
	)
}

func TestPostSendsFormInBody(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"status":"ok","response":{}}`))
	_, e := technitium_tester.New(t, s.URL).Post(
		constant.SettingsSetPath,
		url.Values{constant.KeysParameter: {"alfa|c2VjcmV0|hmac-sha256"}},
	)
	assert.FatalOnError(t, e)
	r := s.Last()
	assert.String(t, http.MethodPost, r.Method)
	assert.String(t, "/api/settings/set", r.URL.Path)
	assert.String(t, "", r.URL.RawQuery)
	assert.String(
		t,
		"application/x-www-form-urlencoded",
		r.Header.Get(webConstant.ContentType),
	)
	assert.String(
		t,
		"Bearer alfa-token",
		r.Header.Get(webConstant.Authorization),
	)
	assert.FatalOnError(t, r.ParseForm())
	assert.String(
		t,
		"alfa|c2VjcmV0|hmac-sha256",
		r.PostForm.Get(constant.KeysParameter),
	)
}

func TestPostErrorEnvelopeIsUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"status":"error","errorMessage":"Invalid TSIG key name."}`,
		))
	_, e := technitium_tester.New(t, s.URL).Post(
		constant.ZoneOptionsSetPath,
		url.Values{},
	)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"technitium error: Invalid TSIG key name.",
		e.Error(),
	)
}

func TestErrorEnvelopeIsUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"status":"error","errorMessage":"No such zone was found: bravo.example"}`,
		))
	_, e := technitium_tester.New(t, s.URL).Get(
		"/zones/delete?zone=bravo.example",
	)
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"technitium error: No such zone was found: bravo.example",
		e.Error(),
	)
}

func TestInvalidTokenIsUnexpected(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"status":"invalid-token","errorMessage":"Invalid token or session expired."}`,
		))
	_, e := technitium_tester.New(t, s.URL).Get("/zones/list")
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "technitium invalid-token:", e.Error())
}

func TestRefusalCarriesErrorMessage(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(
			`{"status":"error","errorMessage":"Parameter 'zone' missing."}`,
		),
	)
	_, e := technitium_tester.New(t, s.URL).Get("/zones/create")
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"status: 400: Parameter 'zone' missing.",
		e.Error(),
	)
}

func TestUntrustedReachesSelfSignedServer(t *testing.T) {
	s := httptest.NewTLSServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				_, e := w.Write([]byte(`{"status":"ok","response":{}}`))
				errors.PanicOnError(e)
			},
		),
	)
	t.Cleanup(s.Close)
	_, e := technitium_tester.New(t, s.URL).Get("/zones/list")
	assert.True(t, unreachable.Is(e))
	c := technitium_tester.New(t, s.URL)
	defer c.Close()
	c.Untrusted()
	_, f := c.Get("/zones/list")
	assert.FatalOnError(t, f)
}
