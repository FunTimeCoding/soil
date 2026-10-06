package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUnwrapsTheEnvelope(t *testing.T) {
	s, last := upstream_tester.New(
		t,
		http.StatusOK,
		`{"status":"ok","response":{"zones":[]}}`,
	)
	payload, e := newClient(t, s.URL).Get(
		"/zones/records/get?domain=alfa.example&listZone=true",
	)
	assert.FatalOnError(t, e)
	assert.String(t, `{"zones":[]}`, string(payload))
	r := last()
	assert.String(t, "/api/zones/records/get", r.URL.Path)
	assert.String(t, "alfa.example", r.URL.Query().Get("domain"))
	assert.String(t, "true", r.URL.Query().Get("listZone"))
	assert.String(t, "Bearer alfa-token", r.Header.Get(constant.Authorization))
}

func TestErrorEnvelopeIsUnexpected(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusOK,
		`{"status":"error","errorMessage":"No such zone was found: bravo.example"}`,
	)
	_, e := newClient(t, s.URL).Get("/zones/delete?zone=bravo.example")
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"technitium error: No such zone was found: bravo.example",
		e.Error(),
	)
}

func TestInvalidTokenIsUnexpected(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusOK,
		`{"status":"invalid-token","errorMessage":"Invalid token or session expired."}`,
	)
	_, e := newClient(t, s.URL).Get("/zones/list")
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "technitium invalid-token:", e.Error())
}

func TestRefusalCarriesTheErrorMessage(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusBadRequest,
		`{"status":"error","errorMessage":"Parameter 'zone' missing."}`,
	)
	_, e := newClient(t, s.URL).Get("/zones/create")
	assert.True(t, unexpected.Is(e))
	assert.StringContains(
		t,
		"status: 400: Parameter 'zone' missing.",
		e.Error(),
	)
}

func TestUntrustedReachesASelfSignedServer(t *testing.T) {
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
	_, e := newClient(t, s.URL).Get("/zones/list")
	assert.True(t, unreachable.Is(e))
	c := newClient(t, s.URL)
	c.Untrusted()
	_, f := c.Get("/zones/list")
	assert.FatalOnError(t, f)
}
