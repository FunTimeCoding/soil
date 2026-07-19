package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
	"testing"
)

func TestIsSymbol(t *testing.T) {
	assert.True(t, pointer.IsSymbol("pkg/web/RecoveryMiddleware"))
	assert.True(t, pointer.IsSymbol("pkg/provision/salt.Client"))
	assert.True(t, pointer.IsSymbol("pkg/check/memory.LocalLines()"))
	assert.True(t, pointer.IsSymbol("pkg/server/Run()"))
	assert.False(t, pointer.IsSymbol("pkg/web/recovery_middleware.go"))
	assert.False(t, pointer.IsSymbol("doc/ai/spec/naming.md"))
	assert.False(t, pointer.IsSymbol("CLAUDE.md"))
	assert.False(t, pointer.IsSymbol("pkg/lint"))
}
