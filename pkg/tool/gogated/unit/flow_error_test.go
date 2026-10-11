package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/tester"
	"net/url"
	"testing"
)

func TestInvalidCodeFails(t *testing.T) {
	o := tester.New(t)
	reg := o.Register(t, []string{"http://localhost/callback"})
	form := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {"invalid-code"},
		"redirect_uri":  {"http://localhost/callback"},
		"client_id":     {reg.ClientIdentifier},
		"client_secret": {reg.ClientSecret},
		"code_verifier": {"anything"},
	}
	status := o.ExchangeCodeRaw(t, form)

	if status == 200 {
		t.Fatal("expected token exchange to fail with invalid code")
	}
}

func TestReplayedCodeFails(t *testing.T) {
	o := tester.New(t)
	reg := o.Register(t, []string{"http://localhost/callback"})
	auth := o.Authorize(t, reg.ClientIdentifier, "http://localhost/callback")
	_ = o.ExchangeCode(t, reg.ClientIdentifier, reg.ClientSecret, auth)
	form := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {auth.Code},
		"redirect_uri":  {auth.RedirectLocator},
		"client_id":     {reg.ClientIdentifier},
		"client_secret": {reg.ClientSecret},
		"code_verifier": {auth.CodeVerifier},
	}
	status := o.ExchangeCodeRaw(t, form)

	if status == 200 {
		t.Fatal("expected replayed code to fail")
	}
}

func TestRegistrationRequiresRedirectURIs(t *testing.T) {
	o := tester.New(t)
	status := o.RegisterRaw(
		t,
		map[string]any{
			"grant_types":    []string{"authorization_code"},
			"response_types": []string{"code"},
		},
	)
	assert.Integer(t, 400, status)
}
