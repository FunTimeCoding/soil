package pointer

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestIsSymbol(t *testing.T) {
	assert.True(t, IsSymbol("pkg/web/RecoveryMiddleware"))
	assert.True(t, IsSymbol("pkg/provision/salt.Client"))
	assert.True(t, IsSymbol("pkg/check/memory.LocalLines()"))
	assert.True(t, IsSymbol("pkg/server/Run()"))
	assert.False(t, IsSymbol("pkg/web/recovery_middleware.go"))
	assert.False(t, IsSymbol("doc/ai/spec/naming.md"))
	assert.False(t, IsSymbol("CLAUDE.md"))
	assert.False(t, IsSymbol("pkg/lint"))
}
