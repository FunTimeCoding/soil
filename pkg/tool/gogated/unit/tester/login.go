package tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func (o *Tester) Login(
	t *testing.T,
	clientIdentifier string,
	redirectLocator string,
	mail string,
	password string,
	source string,
) (int, string) {
	t.Helper()
	getResponse, e := o.client.Get(
		fmt.Sprintf(
			"%s/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=openid&state=state-of-login&code_challenge=%s&code_challenge_method=S256",
			o.server.BaseLocator(),
			url.QueryEscape(clientIdentifier),
			url.QueryEscape(redirectLocator),
			url.QueryEscape(
				computeCodeChallenge(
					"verifier-of-login-at-least-43-characters-long",
				),
			),
		),
	)
	assert.FatalOnError(t, e)
	errors.PanicClose(getResponse.Body)
	var session *http.Cookie

	for _, c := range getResponse.Cookies() {
		if c.Name == constant.CookieName {
			session = c
		}
	}

	if session == nil {
		t.Fatal("no login session cookie returned")
	}

	request, e := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/authorize", o.server.BaseLocator()),
		stringReader(
			url.Values{
				constant.MailField:     {mail},
				constant.PasswordField: {password},
				constant.SourceField:   {source},
			}.Encode(),
		),
	)
	assert.FatalOnError(t, e)
	request.Header.Set(webConstant.ContentType, webConstant.FormEncoded)
	request.AddCookie(session)
	response, e := o.client.Do(request)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(response.Body)
	body, e := io.ReadAll(response.Body)
	assert.FatalOnError(t, e)

	return response.StatusCode, string(body)
}
