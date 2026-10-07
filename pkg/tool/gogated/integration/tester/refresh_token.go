package tester

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/types/result"
	"net/url"
	"testing"
)

func (o *Tester) RefreshToken(
	t *testing.T,
	clientIdentifier string,
	clientSecret string,
	refreshToken string,
) *result.Token {
	t.Helper()
	form := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
		"client_id":     {clientIdentifier},
		"client_secret": {clientSecret},
	}
	r, e := o.client.PostForm(
		fmt.Sprintf("%s/token", o.server.BaseLocator()),
		form,
	)
	assert.FatalOnError(t, e)

	defer errors.PanicClose(r.Body)

	if r.StatusCode != 200 {
		t.Fatalf("refresh failed: %d", r.StatusCode)
	}

	var result result.Token
	assert.FatalOnError(t, json.NewDecoder(r.Body).Decode(&result))

	return &result
}
