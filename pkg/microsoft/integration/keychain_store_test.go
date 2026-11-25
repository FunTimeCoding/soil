package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/microsoft/keychain_store"
	"github.com/funtimecoding/soil/pkg/system/integration/keychain_tester"
	"golang.org/x/oauth2"
	"testing"
)

func TestKeychainStoreSavesOnlyRefreshToken(t *testing.T) {
	k := keychain_tester.New(t, "")
	assert.FatalOnError(
		t,
		keychain_store.New("charlie", "client").SaveToken(
			&oauth2.Token{AccessToken: "alfa", RefreshToken: "bravo"},
		),
	)
	assert.String(
		t,
		"add-generic-password -U -s \"charlie\" -a \"client\" -w \"bravo\"\n",
		k.Input(),
	)
}

func TestKeychainStoreRefusesTokenWithoutRefresh(t *testing.T) {
	k := keychain_tester.New(t, "")
	e := keychain_store.New("charlie", "client").SaveToken(
		&oauth2.Token{AccessToken: "alfa"},
	)
	assert.True(t, validation.Is(e))
	assert.String(t, "", k.Arguments())
}

func TestKeychainStoreLoadsRefreshTokenForRenewal(t *testing.T) {
	keychain_tester.New(t, "")
	token, found, e := keychain_store.New("charlie", "client").FindToken()
	assert.FatalOnError(t, e)
	assert.True(t, found)
	assert.String(t, "stored-secret", token.RefreshToken)
	assert.False(t, token.Valid())
}
