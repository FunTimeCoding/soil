package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/tester"
	"net/http"
	"net/url"
	"testing"
)

func TestSingleSignOnSkipsLoginForm(t *testing.T) {
	o := tester.New(t)
	r := o.Register(t, []string{"http://localhost/callback"})
	first := o.Authorize(t, r.ClientIdentifier, "http://localhost/callback")

	if first.AuthenticationCookie == nil {
		t.Fatal("expected authentication cookie after login")
	}

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
		t.Fatalf(
			"expected redirect without login form, status: %d",
			response.StatusCode,
		)
	}

	redirect, e := url.Parse(location)
	assert.FatalOnError(t, e)

	if redirect.Query().Get(constant.ResponseCode) == "" {
		t.Fatalf("expected code in redirect: %s", location)
	}
}

func TestSingleSignOnWithoutCookieRendersForm(t *testing.T) {
	o := tester.New(t)
	r := o.Register(t, []string{"http://localhost/callback"})
	response := o.ResumeAuthorize(
		t,
		r.ClientIdentifier,
		"http://localhost/callback",
		nil,
		"",
	)
	errors.PanicClose(response.Body)
	assert.Integer(t, http.StatusOK, response.StatusCode)
	assert.String(t, "", response.Header.Get("Location"))
}

func TestPromptLoginForcesReauthentication(t *testing.T) {
	o := tester.New(t)
	r := o.Register(t, []string{"http://localhost/callback"})
	first := o.Authorize(t, r.ClientIdentifier, "http://localhost/callback")
	response := o.ResumeAuthorize(
		t,
		r.ClientIdentifier,
		"http://localhost/callback",
		first.AuthenticationCookie,
		"&prompt=login",
	)
	errors.PanicClose(response.Body)
	assert.Integer(t, http.StatusOK, response.StatusCode)
	assert.String(t, "", response.Header.Get("Location"))
}
