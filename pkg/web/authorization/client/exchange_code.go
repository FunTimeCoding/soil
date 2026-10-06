package client

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
	"net/url"
)

func (c *Client) exchangeCode(
	code string,
	verifier string,
	callbackLocator string,
) (*tokenResponse, error) {
	form := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {callbackLocator},
		"client_id":     {c.identifier},
		"client_secret": {c.secret},
		"code_verifier": {verifier},
	}
	q := request.Absolute(join.Empty(c.issuer, "/token")).WithBody(
		constant.FormEncoded,
		[]byte(form.Encode()),
	)
	q.Method = http.MethodPost
	var result tokenResponse

	if e := c.requester.Notation(q, &result); e != nil {
		return nil, e
	}

	return &result, nil
}
