package tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"testing"
)

func (o *Tester) ResumeAuthorize(
	t *testing.T,
	clientIdentifier string,
	redirectLocator string,
	authenticationCookie *http.Cookie,
	extra string,
) *http.Response {
	t.Helper()
	challenge := computeCodeChallenge(
		fmt.Sprintf("%s%s", uuid.New(), uuid.New()),
	)
	locator := fmt.Sprintf(
		"%s/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=openid+offline&state=%s&code_challenge=%s&code_challenge_method=S256%s",
		o.server.BaseLocator(),
		url.QueryEscape(clientIdentifier),
		url.QueryEscape(redirectLocator),
		url.QueryEscape(uuid.New().String()),
		url.QueryEscape(challenge),
		extra,
	)
	request, e := http.NewRequest(http.MethodGet, locator, nil)
	assert.FatalOnError(t, e)

	if authenticationCookie != nil {
		request.AddCookie(authenticationCookie)
	}

	response, e := o.client.Do(request)
	assert.FatalOnError(t, e)

	return response
}
