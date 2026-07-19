package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/authorization/constant"
	"testing"
)

func TestAuthorizationConstant(t *testing.T) {
	assert.String(
		t,
		"/.well-known/oauth-protected-resource",
		constant.ProtectedResource,
	)
	assert.String(t, "authorization_servers", constant.AuthorizationServer)
	assert.String(t, "resource", constant.Resource)
	assert.String(t, "GET, OPTIONS", constant.ProtectedMethods)
}
