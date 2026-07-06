//go:build local

package tokenizer

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestCount(t *testing.T) {
	e, f := New()
	assert.FatalOnError(t, f)
	assert.Integer(t, 3, e.Count("Hello world!"))
	assert.Integer(t, 2, e.Count("The fox"))
	assert.Integer(t, 3, e.Count("The fox."))
	assert.Integer(t, 1, e.Count("fox"))
	assert.Integer(t, 2, e.Count("fox."))
	assert.Integer(t, 0, e.Count(""))
}

func TestTokens(t *testing.T) {
	e, f := New()
	assert.FatalOnError(t, f)
	tokens := e.Tokens("Hello world!")
	assert.Count(t, 3, tokens)
	assert.String(t, "Hello", tokens[0])
	assert.String(t, " world", tokens[1])
	assert.String(t, "!", tokens[2])
}
