package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "Accept-Language", constant.AcceptLanguage)
	assert.String(t, "User-Agent", constant.UserAgent)
	assert.String(t, "image/x-icon", constant.Icon)
	assert.String(t, "method", constant.FormMethod)
}

func TestAuthorizationConstant(t *testing.T) {
	assert.String(
		t,
		"/.well-known/oauth-protected-resource",
		constant.ProtectedResource,
	)
	assert.String(
		t,
		"authorization_servers",
		constant.AuthorizationServer,
	)
	assert.String(t, "resource", constant.AuthorizationResource)
	assert.String(t, "GET, OPTIONS", constant.ProtectedMethods)
}

func TestLocationConstant(t *testing.T) {
	assert.String(t, "/favicon.ico", constant.LocationFavicon)
	assert.String(t, "/mcp", constant.LocationModelContext)
	assert.String(t, "/shutdown", constant.LocationShutdown)
	assert.String(t, "/status", constant.LocationStatus)
}
