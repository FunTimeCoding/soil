package unit_test

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"testing"
)

func TestAmbiguousClassifiesByType(t *testing.T) {
	e := ambiguous.Format("expected 1 group, got %d", 3)
	assert.String(t, "expected 1 group, got 3", e.Error())
	assert.True(t, ambiguous.Is(e))
	assert.False(t, ambiguous.Is(errors.New("expected 1 group, got 3")))
	assert.False(t, not_found.Is(e))
	assert.False(t, ambiguous.Is(nil))
}
