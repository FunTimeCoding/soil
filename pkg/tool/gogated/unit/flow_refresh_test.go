package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/tester"
	"testing"
)

func TestRefreshTokenGrantsNewTokens(t *testing.T) {
	o := tester.New(t)
	reg := o.Register(t, []string{"http://localhost/callback"})
	auth := o.Authorize(t, reg.ClientIdentifier, "http://localhost/callback")
	tokens := o.ExchangeCode(t, reg.ClientIdentifier, reg.ClientSecret, auth)
	refreshed := o.RefreshToken(
		t,
		reg.ClientIdentifier,
		reg.ClientSecret,
		tokens.RefreshToken,
	)

	if refreshed.AccessToken == "" {
		t.Fatal("expected access token")
	}

	if refreshed.RefreshToken == "" {
		t.Fatal("expected refresh token")
	}

	if tokens.AccessToken == refreshed.AccessToken {
		t.Fatal("expected new access token after refresh")
	}

	if tokens.RefreshToken == refreshed.RefreshToken {
		t.Fatal("expected new refresh token after rotation")
	}
}
