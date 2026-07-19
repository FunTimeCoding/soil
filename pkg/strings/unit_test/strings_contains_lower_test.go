package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestContainsLower(t *testing.T) {
	assert.True(t, strings.ContainsLower("TestFixture", "fixture"))
	assert.False(t, strings.ContainsLower("TestFixture", "other"))
}
