package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/tester"
	"testing"
)

func TestFullAuthorizationCodeFlow(t *testing.T) {
	o := tester.New(t)
	r := o.Register(t, []string{"http://localhost/callback"})
	authorization := o.Authorize(
		t,
		r.ClientIdentifier,
		"http://localhost/callback",
	)

	if authorization.Code == "" {
		t.Fatal("expected authorization code")
	}

	tokens := o.ExchangeCode(
		t,
		r.ClientIdentifier,
		r.ClientSecret,
		authorization,
	)

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
