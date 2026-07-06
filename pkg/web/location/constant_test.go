package location

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "/favicon.ico", Favicon)
	assert.String(t, "/mcp", ModelContext)
	assert.String(t, "/shutdown", Shutdown)
	assert.String(t, "/status", Status)
}
