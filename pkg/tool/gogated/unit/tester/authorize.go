package tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	gogated "github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/types/result"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"testing"
)

func (o *Tester) Authorize(
	t *testing.T,
	clientIdentifier string,
	redirectLocator string,
) *result.Authorize {
	t.Helper()
	verifier := fmt.Sprintf("%s%s", uuid.New(), uuid.New())
	challenge := computeCodeChallenge(verifier)
	state := uuid.New().String()
	authorizeLocator := fmt.Sprintf(
		"%s/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=openid+offline&state=%s&code_challenge=%s&code_challenge_method=S256",
		o.server.BaseLocator(),
		url.QueryEscape(clientIdentifier),
		url.QueryEscape(redirectLocator),
		url.QueryEscape(state),
		url.QueryEscape(challenge),
	)
	getResponse, e := o.client.Get(authorizeLocator)
	assert.FatalOnError(t, e)
	errors.PanicClose(getResponse.Body)
	var sessionCookie *http.Cookie

	for _, c := range getResponse.Cookies() {
		if c.Name == "gogated_login" {
			sessionCookie = c
		}
	}

	if sessionCookie == nil {
		t.Fatal("no login session cookie returned")
	}

	postForm := url.Values{
		"mail":     {gogated.FixtureMail},
		"password": {gogated.FixturePassword},
	}
	postRequest, e := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/authorize", o.server.BaseLocator()),
		nil,
	)
	assert.FatalOnError(t, e)
	postRequest.Header.Set(constant.ContentType, constant.FormEncoded)
	postRequest.Body = http.NoBody
	postRequest.URL.RawQuery = ""
	formEncoded := postForm.Encode()
	postRequest, e = http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/authorize", o.server.BaseLocator()),
		stringReader(formEncoded),
	)
	assert.FatalOnError(t, e)
	postRequest.Header.Set(constant.ContentType, constant.FormEncoded)
	postRequest.AddCookie(sessionCookie)
	postResponse, e := o.client.Do(postRequest)
	assert.FatalOnError(t, e)
	errors.PanicClose(postResponse.Body)
	location := postResponse.Header.Get("Location")

	if location == "" {
		t.Fatalf("no redirect location, status: %d", postResponse.StatusCode)
	}

	parsed, e := url.Parse(location)
	assert.FatalOnError(t, e)
	code := parsed.Query().Get(gogated.ResponseCode)

	if code == "" {
		t.Fatalf("no code in redirect: %s", location)
	}

	assert.String(t, state, parsed.Query().Get("state"))
	var authenticationCookie *http.Cookie

	for _, c := range postResponse.Cookies() {
		if c.Name == gogated.AuthenticationCookieName {
			authenticationCookie = c
		}
	}

	return result.NewAuthorize(
		code,
		verifier,
		redirectLocator,
		authenticationCookie,
	)
}
