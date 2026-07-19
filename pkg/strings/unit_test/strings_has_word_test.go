package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestHasWord(t *testing.T) {
	assert.True(t, strings.HasWord("alfa bravo", "bravo"))
	assert.False(t, strings.HasWord("alfa bravo", "bro"))
	assert.False(t, strings.HasWord("alfa bravo", "bra"))
	assert.False(t, strings.HasWord("alfa bravo", "charlie"))
}
