package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"testing"
)

func TestInScopeEmpty(t *testing.T) {
	o := option.New("", false)
	assert.True(t, lint.InScope(o, "doc/ai/spec/naming.md"))
}

func TestInScopeFile(t *testing.T) {
	o := option.New("", false)
	o.Scopes = []string{"doc/ai/spec/naming.md"}
	assert.True(t, lint.InScope(o, "doc/ai/spec/naming.md"))
	assert.False(t, lint.InScope(o, "doc/ai/spec/testing.md"))
}

func TestInScopeDirectory(t *testing.T) {
	o := option.New("", false)
	o.Scopes = []string{"doc/ai"}
	assert.True(t, lint.InScope(o, "doc/ai/spec/naming.md"))
	assert.False(t, lint.InScope(o, "doc/aircraft/naming.md"))
	assert.False(t, lint.InScope(o, "pkg/lint/lint.go"))
}

func TestInScopeSeveral(t *testing.T) {
	o := option.New("", false)
	o.Scopes = []string{"doc/ai", "pkg/lint"}
	assert.True(t, lint.InScope(o, "pkg/lint/lint.go"))
	assert.False(t, lint.InScope(o, "pkg/system/remove.go"))
}
