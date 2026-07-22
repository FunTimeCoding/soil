package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	authorization "github.com/funtimecoding/soil/pkg/web/authorization/constant"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/location"
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
		authorization.ProtectedResource,
	)
	assert.String(t, "authorization_servers", authorization.AuthorizationServer)
	assert.String(t, "resource", authorization.Resource)
	assert.String(t, "GET, OPTIONS", authorization.ProtectedMethods)
}

func TestLocationConstant(t *testing.T) {
	assert.String(t, "/favicon.ico", location.Favicon)
	assert.String(t, "/mcp", location.ModelContext)
	assert.String(t, "/shutdown", location.Shutdown)
	assert.String(t, "/status", location.Status)
}
