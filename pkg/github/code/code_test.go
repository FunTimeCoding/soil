package code

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/google/go-github/v89/github"
	"testing"
)

func TestCode(t *testing.T) {
	r := New(
		&github.CodeResult{
			SHA:  new(upper.Alfa),
			Name: new(upper.Bravo),
			Path: new(upper.Charlie),
		},
	)
	r.Raw = nil
	assert.Any(t, &Code{Hash: "Alfa", Name: "Bravo", Path: "Charlie"}, r)
}
