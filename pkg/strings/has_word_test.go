package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestHasWord(t *testing.T) {
	assert.True(t, HasWord("alfa bravo", "bravo"))
	assert.False(t, HasWord("alfa bravo", "bro"))
	assert.False(t, HasWord("alfa bravo", "bra"))
	assert.False(t, HasWord("alfa bravo", "charlie"))
}
