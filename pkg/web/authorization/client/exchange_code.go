package client

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
	"net/url"
)

func (c *Client) exchangeCode(
	code string,
	verifier string,
	callbackLocator string,
) *tokenResponse {
	form := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {callbackLocator},
		"client_id":     {c.identifier},
		"client_secret": {c.secret},
		"code_verifier": {verifier},
	}
	r, e := http.PostForm(
		join.Empty(c.issuer, "/token"),
		form,
	)
	errors.PanicOnError(e)

	defer errors.PanicClose(r.Body)

	if r.StatusCode != http.StatusOK {
		errors.PanicOnError(
			fmt.Errorf("token exchange failed: %d", r.StatusCode),
		)
	}

	var result tokenResponse
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&result))

	return &result
}
