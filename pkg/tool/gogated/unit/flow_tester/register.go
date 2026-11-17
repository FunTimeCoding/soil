package flow_tester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/types/register"
)

func (o *Tester) Register(redirectLocators []string) *register.Response {
	o.t.Helper()
	body, e := json.Marshal(
		map[string]any{
			"redirect_uris": redirectLocators,
			"grant_types": []string{
				"authorization_code",
				"refresh_token",
			},
			"response_types":             []string{"code"},
			"scope":                      "openid offline",
			"token_endpoint_auth_method": "client_secret_post",
		},
	)
	assert.FatalOnError(o.t, e)
	r, e := o.client.Post(
		fmt.Sprintf("%s/register", o.server.BaseLocator()),
		"application/json",
		bytes.NewReader(body),
	)
	assert.FatalOnError(o.t, e)

	defer errors.PanicClose(r.Body)
	assert.Integer(o.t, 201, r.StatusCode)
	var result register.Response
	assert.FatalOnError(o.t, json.NewDecoder(r.Body).Decode(&result))

	return &result
}
