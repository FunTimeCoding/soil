package flow

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/integration/tester"
	"testing"
)

func TestDiscoveryReturnsIssuer(t *testing.T) {
	o := tester.New(t)
	d := o.Discovery(t)
	assert.String(t, "http://localhost", d.Issuer)
}

func TestDiscoveryIncludesEndpoints(t *testing.T) {
	o := tester.New(t)
	d := o.Discovery(t)
	assert.StringContains(t, "/authorize", d.AuthorizationEndpoint)
	assert.StringContains(t, "/token", d.TokenEndpoint)
	assert.StringContains(t, "/jwks", d.SigningKeysLocator)
	assert.StringContains(t, "/register", d.RegistrationEndpoint)
}

func TestJWKSReturnsKey(t *testing.T) {
	o := tester.New(t)
	keys := o.SigningKeys(t)
	assert.Integer(t, 1, len(keys.Keys))
	assert.String(t, "RSA", keys.Keys[0].KeyType)
	assert.String(t, "sig", keys.Keys[0].Use)
	assert.String(t, "RS256", keys.Keys[0].Algorithm)
	assert.StringContains(t, "-", keys.Keys[0].KeyIdentifier)
}
