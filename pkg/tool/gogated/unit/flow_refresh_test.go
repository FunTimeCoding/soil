package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/flow_tester"
	"testing"
)

func TestRefreshTokenGrantsNewTokens(t *testing.T) {
	o := flow_tester.New(t)
	reg := o.Register([]string{"http://localhost/callback"})
	auth := o.Authorize(reg.ClientIdentifier, "http://localhost/callback")
	tokens := o.ExchangeCode(reg.ClientIdentifier, reg.ClientSecret, auth)
	refreshed := o.RefreshToken(
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
