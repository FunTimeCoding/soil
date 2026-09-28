package tester

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"net/url"
	"testing"
)

func (o *Tester) ExchangeCode(
	t *testing.T,
	clientIdentifier string,
	clientSecret string,
	a *AuthorizeResult,
) *TokenResult {
	t.Helper()
	form := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {a.Code},
		"redirect_uri":  {a.RedirectLocator},
		"client_id":     {clientIdentifier},
		"client_secret": {clientSecret},
		"code_verifier": {a.CodeVerifier},
	}
	r, e := o.client.PostForm(
		fmt.Sprintf("%s/token", o.server.BaseLocator()),
		form,
	)
	assert.FatalOnError(t, e)

	defer errors.PanicClose(r.Body)

	if r.StatusCode != 200 {
		t.Fatalf(
			"token exchange failed: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	var result TokenResult
	assert.FatalOnError(t, json.NewDecoder(r.Body).Decode(&result))

	return &result
}
