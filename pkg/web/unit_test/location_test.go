package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/location"
	"testing"
)

func TestLocationConstant(t *testing.T) {
	assert.String(t, "/favicon.ico", location.Favicon)
	assert.String(t, "/mcp", location.ModelContext)
	assert.String(t, "/shutdown", location.Shutdown)
	assert.String(t, "/status", location.Status)
}
