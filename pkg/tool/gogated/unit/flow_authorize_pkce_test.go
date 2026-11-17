package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/flow_tester"
	"testing"
)

func TestFullAuthorizationCodeFlow(t *testing.T) {
	o := flow_tester.New(t)
	r := o.Register([]string{"http://localhost/callback"})
	authorization := o.Authorize(
		r.ClientIdentifier,
		"http://localhost/callback",
	)

	if authorization.Code == "" {
		t.Fatal("expected authorization code")
	}

	tokens := o.ExchangeCode(r.ClientIdentifier, r.ClientSecret, authorization)

	if tokens.AccessToken == "" {
		t.Fatal("expected access token")
	}

	if tokens.RefreshToken == "" {
		t.Fatal("expected refresh token")
	}

	if tokens.IdentityToken == "" {
		t.Fatal("expected identifier token")
	}

	assert.String(t, "bearer", tokens.TokenType)
	assert.StringContains(t, "openid", tokens.Scope)
}
