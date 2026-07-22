package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"testing"
)

func TestJoinKeyValue(t *testing.T) {
	assert.String(t, "a=b", key_value.Equals("a", "b"))
	assert.String(t, "a b", key_value.Space("a", "b"))
	assert.String(t, "ab", key_value.Empty("a", "b"))
	assert.String(t, "a:b", key_value.Colon("a", "b"))
	assert.String(t, "a/b", key_value.Slash("a", "b"))
	assert.String(t, "a@b", key_value.At("a", "b"))
}
