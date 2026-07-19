package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/option"
	"testing"
)

func TestWhitespace(t *testing.T) {
	c := option.Compact()
	assert.False(t, c.NewlineAtEnd)
	assert.Integer(t, 0, c.AllowedBlankLines)
}
