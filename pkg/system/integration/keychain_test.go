package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/system/integration/keychain_tester"
	"github.com/funtimecoding/soil/pkg/system/keychain"
	"strings"
	"testing"
)

func TestKeychainSaveKeepsSecretOffArguments(t *testing.T) {
	k := keychain_tester.New(t, "")
	assert.FatalOnError(t, keychain.Save("gographd", "client", "s3cret"))
	assert.String(t, "-i\n", k.Arguments())
	assert.String(
		t,
		"add-generic-password -U -s \"gographd\" -a \"client\" -w \"s3cret\"\n",
		k.Input(),
	)
}

func TestKeychainSaveRefusesQuote(t *testing.T) {
	k := keychain_tester.New(t, "")
	e := keychain.Save("gographd", "client", "s3\"cret")
	assert.True(t, validation.Is(e))
	assert.String(t, "", k.Arguments())
}

func TestKeychainSaveRefusesLineSecurityCuts(t *testing.T) {
	k := keychain_tester.New(t, "")
	e := keychain.Save("gographd", "client", strings.Repeat("a", 4042))
	assert.True(t, validation.Is(e))
	assert.String(t, "", k.Arguments())
}

func TestKeychainSaveTakesLongestLineSecurityReads(t *testing.T) {
	k := keychain_tester.New(t, "")
	assert.FatalOnError(
		t,
		keychain.Save("gographd", "client", strings.Repeat("a", 4041)),
	)
	assert.String(t, "-i\n", k.Arguments())
}

func TestKeychainFindReadsSecret(t *testing.T) {
	k := keychain_tester.New(t, "")
	secret, found, e := keychain.Find("gographd", "client")
	assert.FatalOnError(t, e)
	assert.True(t, found)
	assert.String(t, "stored-secret", secret)
	assert.String(
		t,
		"find-generic-password\n-s\ngographd\n-a\nclient\n-w\n",
		k.Arguments(),
	)
}

func TestKeychainFindMissingNotFound(t *testing.T) {
	keychain_tester.New(t, "44")
	secret, found, e := keychain.Find("gographd", "client")
	assert.FatalOnError(t, e)
	assert.False(t, found)
	assert.String(t, "", secret)
}

func TestKeychainFindFailure(t *testing.T) {
	keychain_tester.New(t, "71")
	_, found, e := keychain.Find("gographd", "client")
	assert.Error(t, e)
	assert.False(t, found)
}
