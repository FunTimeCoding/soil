package flow

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/integration/tester"
	"net/http"
	"net/url"
	"testing"
)

func TestDiscoveryAdvertisesEndSession(t *testing.T) {
	o := tester.New(t)
	d := o.Discovery(t)

	if d.EndSessionEndpoint == "" {
		t.Fatal("expected end_session_endpoint in discovery")
	}
}

func TestLogoutRevokesSessionAcrossRelyingParties(t *testing.T) {
	o := tester.New(t)
	r := o.Register(t, []string{"http://localhost/callback"})
	first := o.Authorize(t, r.ClientIdentifier, "http://localhost/callback")
	o.Logout(t, first.AuthenticationCookie)
	response := o.ResumeAuthorize(
		t,
		r.ClientIdentifier,
		"http://localhost/callback",
		first.AuthenticationCookie,
		"",
	)
	errors.PanicClose(response.Body)
	assert.Integer(t, http.StatusOK, response.StatusCode)
	assert.String(t, "", response.Header.Get("Location"))
}

func TestLogoutGetDoesNotRevoke(t *testing.T) {
	o := tester.New(t)
	r := o.Register(t, []string{"http://localhost/callback"})
	first := o.Authorize(t, r.ClientIdentifier, "http://localhost/callback")
	request, e := http.NewRequest(
		http.MethodGet,
		join.Empty(o.BaseLocator(), constant.LogoutPath),
		nil,
	)
	assert.FatalOnError(t, e)
	request.AddCookie(first.AuthenticationCookie)
	confirm, e := o.Do(request)
	assert.FatalOnError(t, e)
	errors.PanicClose(confirm.Body)
	assert.Integer(t, http.StatusOK, confirm.StatusCode)
	response := o.ResumeAuthorize(
		t,
		r.ClientIdentifier,
		"http://localhost/callback",
		first.AuthenticationCookie,
		"",
	)
	errors.PanicClose(response.Body)
	location := response.Header.Get("Location")

	if location == "" {
		t.Fatal("session should survive the confirmation page")
	}

	redirect, e := url.Parse(location)
	assert.FatalOnError(t, e)

	if redirect.Query().Get(constant.ResponseCode) == "" {
		t.Fatalf("expected code in redirect: %s", location)
	}
}
