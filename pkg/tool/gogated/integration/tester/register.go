package tester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/types/register"
	"testing"
)

func (o *Tester) Register(
	t *testing.T,
	redirectLocators []string,
) *register.Response {
	t.Helper()
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
	assert.FatalOnError(t, e)
	r, e := o.client.Post(
		fmt.Sprintf("%s/register", o.server.BaseLocator()),
		"application/json",
		bytes.NewReader(body),
	)
	assert.FatalOnError(t, e)

	defer errors.PanicClose(r.Body)
	assert.Integer(t, 201, r.StatusCode)
	var result register.Response
	assert.FatalOnError(t, json.NewDecoder(r.Body).Decode(&result))

	return &result
}
